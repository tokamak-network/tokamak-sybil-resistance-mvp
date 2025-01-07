package til

import (
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/log"

	ethCommon "github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/iden3/go-iden3-crypto/babyjub"
)

func newBatchData(batchNum int) common.BatchData {
	return common.BatchData{
		Batch: common.Batch{
			BatchNum: common.BatchNum(batchNum),

			AccountRoot: big.NewInt(0),
			VouchRoot:   big.NewInt(0),
			ScoreRoot:   big.NewInt(0),
			ExitRoot:    big.NewInt(0),
		},
	}
}

func newBlock(blockNum int64) common.BlockData {
	return common.BlockData{
		Block: common.Block{
			Num: blockNum,
		},
		Rollup: common.RollupData{
			L1UserTxs: []common.L1Tx{},
		},
	}
}

type contextExtra struct {
	openToForge     int64
	toForgeL1TxsNum int64
	nonces          map[common.AccountIdx]common.Nonce
	idx             int
	idxByTxID       map[common.TxID]common.AccountIdx
}

// Context contains the data of the test
type Context struct {
	instructions          []Instruction
	accountNames          []string
	Accounts              map[string]*Account // Name -> *Account
	AccountsByIdx         map[int]*Account
	LastRegisteredTokenID common.TokenID
	l1CreatedAccounts     map[string]*Account // (Name, TokenID) -> *Account

	// rollupConstMaxL1UserTx Maximum L1-user transactions allowed to be
	// queued in a batch
	rollupConstMaxL1UserTx int

	chainID      uint64
	idx          int
	currBlock    common.BlockData
	currBatch    common.BatchData
	currBatchNum int
	Queues       [][]L1Tx
	ToForgeNum   int
	openToForge  int
	blockNum     int64

	extra contextExtra
}

// NewContext returns a new Context
func NewContext(chainID uint64, rollupConstMaxL1UserTx int) *Context {
	currBatchNum := 1 // The protocol defines the first batchNum to be 1
	return &Context{
		Accounts:          make(map[string]*Account),
		l1CreatedAccounts: make(map[string]*Account),
		AccountsByIdx:     make(map[int]*Account),

		rollupConstMaxL1UserTx: rollupConstMaxL1UserTx,
		chainID:                chainID,
		idx:                    common.UserThreshold,
		// We use some placeholder values for StateRoot and ExitTree
		// because these values will never be nil
		currBlock:    newBlock(2),
		currBatch:    newBatchData(currBatchNum),
		currBatchNum: currBatchNum,
		// start with 2 queues, one for toForge, and the other for openToForge
		Queues:      make([][]L1Tx, 2),
		ToForgeNum:  0,
		openToForge: 1,
		//nolint:gomnd
		blockNum: 2, // rollup genesis blockNum
		extra: contextExtra{
			openToForge:     0,
			toForgeL1TxsNum: 0,
			nonces:          make(map[common.AccountIdx]common.Nonce),
			idx:             common.UserThreshold,
			idxByTxID:       make(map[common.TxID]common.AccountIdx),
		},
	}
}

// Account contains the data related to the account for a specific TokenID of a User
type Account struct {
	Name     string
	Idx      common.AccountIdx
	Addr     ethCommon.Address
	BJJ      *babyjub.PrivateKey
	Balance  *big.Int
	Nonce    common.Nonce
	BatchNum int
}

// L1Tx is the data structure used internally for transaction test generation,
// which contains a common.L1Tx data plus some intermediate data for the
// transaction generation.
type L1Tx struct {
	lineNum     int
	fromIdxName string
	toIdxName   string
	L1Tx        common.L1Tx
}

// GenerateBlocks returns an array of BlockData for a given set made of a
// string. It uses the users (keys & nonces) of the Context.
func (tc *Context) GenerateBlocks(set string) ([]common.BlockData, error) {
	parser := newParser(strings.NewReader(set))
	parsedSet, err := parser.parse()
	if err != nil {
		return nil, common.Wrap(err)
	}
	if parsedSet.typ != SetTypeBlockchain {
		return nil,
			common.Wrap(fmt.Errorf("expected set type: %s, found: %s",
				SetTypeBlockchain, parsedSet.typ))
	}

	tc.instructions = parsedSet.instructions
	tc.accountNames = parsedSet.users

	return tc.generateBlocks()
}

// GenerateBlocksFromInstructions returns an array of BlockData for a given set
// made of instructions. It uses the users (keys & nonces) of the Context.
func (tc *Context) GenerateBlocksFromInstructions(set []Instruction) ([]common.BlockData, error) {
	accountNames := []string{}
	addedNames := make(map[string]bool)
	for _, inst := range set {
		// TODO: sybil doesn't have `inst.From != ""` condition but without it
		// empty string gets added to the accountNames
		if _, ok := addedNames[inst.From]; !ok && inst.From != "" {
			// If the name wasn't already added
			accountNames = append(accountNames, inst.From)
			addedNames[inst.From] = true
		}
	}
	tc.accountNames = accountNames
	tc.instructions = set
	return tc.generateBlocks()
}

func (tc *Context) generateBlocks() ([]common.BlockData, error) {
	tc.generateKeys(tc.accountNames)

	var blocks []common.BlockData
	for _, inst := range tc.instructions {
		switch inst.Typ {
		case common.TxTypeCreateAccountDeposit:
			// tx source: L1UserTx
			tx := common.L1Tx{
				FromEthAddr:   tc.Accounts[inst.From].Addr,
				FromBJJ:       tc.Accounts[inst.From].BJJ.Public().Compress(),
				Amount:        big.NewInt(0),
				DepositAmount: inst.DepositAmount,
				Type:          inst.Typ,
			}
			testTx := L1Tx{
				lineNum:     inst.LineNum,
				fromIdxName: inst.From,
				toIdxName:   inst.To,
				L1Tx:        tx,
			}
			if err := tc.addToL1UserQueue(testTx); err != nil {
				return nil, common.Wrap(err)
			}
		case common.TxTypeDeposit: // tx source: L1UserTx
			if err := tc.checkIfAccountExists(inst.From, inst); err != nil {
				log.Error(err)
				return nil, common.Wrap(fmt.Errorf("line %d: %s", inst.LineNum, err.Error()))
			}
			tx := common.L1Tx{
				DepositAmount: inst.DepositAmount,
				Type:          inst.Typ,
			}
			testTx := L1Tx{
				lineNum:     inst.LineNum,
				fromIdxName: inst.From,
				L1Tx:        tx,
			}
			if err := tc.addToL1UserQueue(testTx); err != nil {
				return nil, common.Wrap(err)
			}
		case common.TxTypeCreateVouch:
			tx := common.L1Tx{
				Amount:        big.NewInt(0),
				DepositAmount: big.NewInt(0),
				Type:          common.TxTypeCreateVouch,
			}
			testTx := L1Tx{
				lineNum:     inst.LineNum,
				fromIdxName: inst.From,
				toIdxName:   inst.To,
				L1Tx:        tx,
			}
			if err := tc.addToL1UserQueue(testTx); err != nil {
				return nil, common.Wrap(err)
			}
		case common.TxTypeDeleteVouch:
			tx := common.L1Tx{
				Amount:        big.NewInt(0),
				DepositAmount: big.NewInt(0),
				Type:          common.TxTypeDeleteVouch,
			}
			testTx := L1Tx{
				lineNum:     inst.LineNum,
				fromIdxName: inst.From,
				toIdxName:   inst.To,
				L1Tx:        tx,
			}
			if err := tc.addToL1UserQueue(testTx); err != nil {
				return nil, common.Wrap(err)
			}
		case common.TxTypeForceExit: // tx source: L1UserTx
			tx := common.L1Tx{
				ToIdx:         common.AccountIdx(1), // as is an Exit
				Amount:        inst.Amount,
				DepositAmount: big.NewInt(0),
				Type:          common.TxTypeForceExit,
			}
			testTx := L1Tx{
				lineNum:     inst.LineNum,
				fromIdxName: inst.From,
				toIdxName:   inst.To,
				L1Tx:        tx,
			}
			if err := tc.addToL1UserQueue(testTx); err != nil {
				return nil, common.Wrap(err)
			}
		case TypeNewBatchL1:
			// for each L1UserTx of the Queues[ToForgeNum], accumulate Txs into map
			for _, tx := range tc.Queues[tc.ToForgeNum] {
				tc.l1CreatedAccounts[tx.fromIdxName] = tc.Accounts[tx.fromIdxName]
				tc.AccountsByIdx[tc.idx] = tc.Accounts[tx.fromIdxName]
				tc.idx++
			}
			tc.currBatch.L1Batch = true
			if err := tc.setCurrBatch(); err != nil {
				return nil, common.Wrap(err)
			}
			toForgeL1TxsNum := int64(tc.openToForge)
			tc.currBatch.Batch.ForgeL1TxsNum = &toForgeL1TxsNum
			// advance batch
			tc.ToForgeNum++
			if tc.ToForgeNum == tc.openToForge {
				tc.openToForge++
				newQueue := []L1Tx{}
				tc.Queues = append(tc.Queues, newQueue)
			}
		case TypeNewBlock:
			blocks = append(blocks, tc.currBlock)
			tc.blockNum++
			tc.currBlock = newBlock(tc.blockNum)
		default:
			return nil, common.Wrap(fmt.Errorf("line %d: Unexpected type: %s", inst.LineNum, inst.Typ))
		}
	}

	return blocks, nil
}

// setCurrBatch sets the Idxs to the transactions of the tc.currBatch
func (tc *Context) setCurrBatch() error {
	tc.currBatch.Batch.LastIdx = int64(tc.idx - 1) // `-1` because tc.idx is the next available idx
	tc.currBlock.Rollup.Batches = append(tc.currBlock.Rollup.Batches, tc.currBatch)
	tc.currBatchNum++
	tc.currBatch = newBatchData(tc.currBatchNum)
	return nil
}

// addToL1UserQueue adds the L1UserTx into the queue that is open and has space
func (tc *Context) addToL1UserQueue(tx L1Tx) error {
	if len(tc.Queues[tc.openToForge]) >= tc.rollupConstMaxL1UserTx {
		// if current OpenToForge queue reached its Max, move into a
		// new queue
		tc.openToForge++
		newQueue := []L1Tx{}
		tc.Queues = append(tc.Queues, newQueue)
	}
	// Fill L1UserTx specific parameters
	tx.L1Tx.UserOrigin = true
	toForgeL1TxsNum := int64(tc.openToForge)
	tx.L1Tx.ToForgeL1TxsNum = &toForgeL1TxsNum
	tx.L1Tx.EthBlockNum = tc.blockNum
	tx.L1Tx.Position = len(tc.Queues[tc.openToForge])

	// When an L1UserTx is generated, all idxs must be available (except when idx == 0 or idx == 1)
	if tx.L1Tx.Type != common.TxTypeCreateAccountDeposit {
		tx.L1Tx.FromIdx = tc.Accounts[tx.fromIdxName].Idx
	}
	tx.L1Tx.FromEthAddr = tc.Accounts[tx.fromIdxName].Addr
	tx.L1Tx.FromBJJ = tc.Accounts[tx.fromIdxName].BJJ.Public().Compress()
	if tx.toIdxName == "" {
		tx.L1Tx.ToIdx = common.AccountIdx(0)
	} else {
		account, ok := tc.Accounts[tx.toIdxName]
		if !ok {
			return common.Wrap(fmt.Errorf("line %d: Transfer to User: %s, "+
				"while account not created yet", tx.lineNum, tx.toIdxName))
		}
		tx.L1Tx.ToIdx = account.Idx
	}
	if tx.L1Tx.Type == common.TxTypeForceExit {
		tx.L1Tx.ToIdx = common.AccountIdx(1)
	}
	nTx, err := common.NewL1Tx(&tx.L1Tx)
	if err != nil {
		return common.Wrap(fmt.Errorf("line %d: %s", tx.lineNum, err.Error()))
	}
	tx.L1Tx = *nTx

	tc.Queues[tc.openToForge] = append(tc.Queues[tc.openToForge], tx)
	tc.currBlock.Rollup.L1UserTxs = append(tc.currBlock.Rollup.L1UserTxs, tx.L1Tx)

	return nil
}

func (tc *Context) checkIfAccountExists(tf string, inst Instruction) error {
	if tc.Accounts[tf] == nil {
		return common.Wrap(fmt.Errorf("%s at User: %s, while account not created yet",
			inst.Typ, tf))
	}
	return nil
}

// RestartNonces sets all the Users.Accounts.Nonces to 0
func (tc *Context) RestartNonces() {
	for name := range tc.Accounts {
		tc.Accounts[name].Nonce = common.Nonce(0)
	}
}

// generateKeys generates BabyJubJub & Address keys for the given list of user
// names in a deterministic way. This means, that for the same given
// 'userNames' in a certain order, the keys will be always the same.
func (tc *Context) generateKeys(userNames []string) {
	for i := 1; i < len(userNames)+1; i++ {
		if _, ok := tc.Accounts[userNames[i-1]]; ok {
			// account already created
			continue
		}

		u := NewUser(i, userNames[i-1])
		tc.Accounts[userNames[i-1]] = &u
	}
}

// NewUser creates a User deriving its keys at the path keyDerivationIndex
func NewUser(keyDerivationIndex int, name string) Account {
	// babyjubjub key
	var sk babyjub.PrivateKey
	var iBytes [8]byte
	binary.LittleEndian.PutUint64(iBytes[:], uint64(keyDerivationIndex))
	copy(sk[:], iBytes[:]) // only for testing

	// eth address
	var key ecdsa.PrivateKey
	key.D = big.NewInt(int64(keyDerivationIndex)) // only for testing
	key.PublicKey.X, key.PublicKey.Y = ethCrypto.S256().ScalarBaseMult(key.D.Bytes())
	key.Curve = ethCrypto.S256()
	addr := ethCrypto.PubkeyToAddress(key.PublicKey)

	// Idx
	idx := common.AccountIdx(255 + keyDerivationIndex)

	// Balance
	balance := big.NewInt(0)

	// Nonce
	nonce := common.Nonce(0)

	return Account{
		Name:     name,
		Idx:      idx,
		Addr:     addr,
		BJJ:      &sk,
		Balance:  balance,
		Nonce:    nonce,
		BatchNum: keyDerivationIndex,
	}
}

// L1TxsToCommonL1Txs converts an array of []til.L1Tx to []common.L1Tx
func L1TxsToCommonL1Txs(l1 []L1Tx) []common.L1Tx {
	var r []common.L1Tx
	for i := 0; i < len(l1); i++ {
		r = append(r, l1[i].L1Tx)
	}
	return r
}

// ConfigExtra is the configuration used in FillBlocksExtra to extend the
// blocks returned by til.
type ConfigExtra struct {
	// Address to set as forger for each batch
	BootCoordAddr ethCommon.Address
	// Coordinator user name used to select the corresponding accounts to
	// collect coordinator fees
	CoordUser string
}

// FillBlocksL1UserTxsBatchNum fills the BatchNum of forged L1UserTxs:
// - blocks[].Rollup.L1UserTxs[].BatchNum
func (tc *Context) FillBlocksL1UserTxsBatchNum(blocks []common.BlockData) {
	for i := range blocks {
		block := &blocks[i]
		for j := range block.Rollup.Batches {
			batch := &block.Rollup.Batches[j]
			if batch.L1Batch {
				// Set BatchNum for forged L1UserTxs to til blocks
				bn := batch.Batch.BatchNum
				for k := range blocks {
					block := &blocks[k]
					for l := range block.Rollup.L1UserTxs {
						tx := &block.Rollup.L1UserTxs[l]
						if *tx.ToForgeL1TxsNum == tc.extra.openToForge {
							tx.BatchNum = &bn
						}
					}
				}
				tc.extra.openToForge++
			}
		}
	}
}

// FillBlocksForgedL1UserTxs fills the L1UserTxs of a batch with the L1UserTxs
// that are forged in that batch.  It always sets `EffectiveAmount` = `Amount`
// and `EffectiveDepositAmount` = `DepositAmount`.  This function requires a
// previous call to `FillBlocksExtra`.
// - blocks[].Rollup.L1UserTxs[].BatchNum
// - blocks[].Rollup.L1UserTxs[].EffectiveAmount
// - blocks[].Rollup.L1UserTxs[].EffectiveDepositAmount
// - blocks[].Rollup.L1UserTxs[].EffectiveFromIdx
func (tc *Context) FillBlocksForgedL1UserTxs(blocks []common.BlockData) error {
	for i := range blocks {
		block := &blocks[i]
		for j := range block.Rollup.Batches {
			batch := &block.Rollup.Batches[j]
			if batch.L1Batch {
				batchNum := batch.Batch.BatchNum
				queue := tc.Queues[int(*batch.Batch.ForgeL1TxsNum)]
				batch.L1UserTxs = make([]common.L1Tx, len(queue))
				for k := range queue {
					tx := &batch.L1UserTxs[k]
					*tx = queue[k].L1Tx
					tx.EffectiveAmount = tx.Amount
					tx.EffectiveDepositAmount = tx.DepositAmount
					tx.BatchNum = &batchNum
					_tx, err := common.NewL1Tx(tx)
					if err != nil {
						return common.Wrap(err)
					}
					*tx = *_tx
					if tx.FromIdx == 0 {
						tx.EffectiveFromIdx = tc.extra.idxByTxID[tx.TxID]
					} else {
						tx.EffectiveFromIdx = tx.FromIdx
					}
				}
			}
		}
	}
	return nil
}

// FillBlocksExtra fills extra fields not generated by til in each block, so
// that the blockData is closer to what the HistoryDB stores.  The filled
// fields are:
// - blocks[].Rollup.Batch.EthBlockNum
// - blocks[].Rollup.Batch.ForgerAddr
// - blocks[].Rollup.Batch.ForgeL1TxsNum
// - blocks[].Rollup.Batch.ExitTree
// - blocks[].Rollup.Batch.CreatedAccounts
func (tc *Context) FillBlocksExtra(blocks []common.BlockData, cfg *ConfigExtra) error {
	// Fill extra fields not generated by til in til block
	for i := range blocks {
		block := &blocks[i]
		for j := range block.Rollup.Batches {
			batch := &block.Rollup.Batches[j]
			batch.Batch.EthBlockNum = block.Block.Num
			// til doesn't fill the batch forger addr
			batch.Batch.ForgerAddr = cfg.BootCoordAddr
			if batch.L1Batch {
				toForgeL1TxsNumCpy := tc.extra.toForgeL1TxsNum
				// til doesn't fill the ForgeL1TxsNum
				batch.Batch.ForgeL1TxsNum = &toForgeL1TxsNumCpy
				tc.extra.toForgeL1TxsNum++
			}

			// TODO: default value is nil but the db column type is not nullable
			batch.Batch.GasPrice = big.NewInt(0)
		}
	}

	// Fill CreatedAccounts
	for i := range blocks {
		block := &blocks[i]
		for j := range block.Rollup.Batches {
			batch := &block.Rollup.Batches[j]
			l1Txs := []*common.L1Tx{}
			if batch.L1Batch {
				for k := range tc.Queues[*batch.Batch.ForgeL1TxsNum] {
					l1Txs = append(l1Txs, &tc.Queues[*batch.Batch.ForgeL1TxsNum][k].L1Tx)
				}
			}
			for k := range l1Txs {
				tx := l1Txs[k]
				if tx.Type == common.TxTypeCreateAccountDeposit {
					user, ok := tc.AccountsByIdx[tc.extra.idx]
					if !ok {
						return common.Wrap(fmt.Errorf("created account with idx: %v not found", tc.extra.idx))
					}
					batch.CreatedAccounts = append(batch.CreatedAccounts,
						common.Account{
							Idx:      common.AccountIdx(tc.extra.idx),
							BatchNum: batch.Batch.BatchNum,
							BJJ:      user.BJJ.Public().Compress(),
							EthAddr:  user.Addr,
							Nonce:    0,
							Balance:  big.NewInt(0),
						})
					if !tx.UserOrigin {
						tx.EffectiveFromIdx = common.AccountIdx(tc.extra.idx)
					}
					tc.extra.idxByTxID[tx.TxID] = common.AccountIdx(tc.extra.idx)
					tc.extra.idx++
				}
			}
		}
	}

	// Fill ExitTree (only AccountIdx and Balance)
	for i := range blocks {
		block := &blocks[i]
		for j := range block.Rollup.Batches {
			batch := &block.Rollup.Batches[j]
			if batch.L1Batch {
				for _, _tx := range tc.Queues[*batch.Batch.ForgeL1TxsNum] {
					tx := _tx.L1Tx
					if tx.Type == common.TxTypeForceExit {
						batch.ExitTree =
							append(batch.ExitTree,
								common.ExitInfo{
									BatchNum:   batch.Batch.BatchNum,
									AccountIdx: tx.FromIdx,
									Balance:    tx.Amount,
								})
					}
				}
			}
		}
	}
	return nil
}
