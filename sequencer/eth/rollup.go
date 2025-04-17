package eth

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"tokamak-sybil-resistance/common"
	Sybil "tokamak-sybil-resistance/eth/contracts"
	"tokamak-sybil-resistance/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// QueueStruct is the queue of L1Txs for a batch
type QueueStruct struct {
	L1TxQueue    []common.L1Tx
	TotalL1TxFee *big.Int
}

// NewQueueStruct creates a new clear QueueStruct.
func NewQueueStruct() *QueueStruct {
	return &QueueStruct{
		L1TxQueue:    make([]common.L1Tx, 0),
		TotalL1TxFee: big.NewInt(0),
	}
}

// RollupState represents the state of the Rollup in the Smart Contract
type RollupState struct {
	AccountRoot *big.Int
	VouchRoot   *big.Int
	ScoreRoot   *big.Int
	ExitRoots   []*big.Int
	// ExitNullifierMap       map[[256 / 8]byte]bool
	ExitNullifierMap       map[int64]map[int64]bool // batchNum -> idx -> bool
	MapL1TxQueue           map[int64]*QueueStruct
	LastL1L2Batch          int64
	CurrentToForgeL1TxsNum int64
	LastToForgeL1TxsNum    int64
	CurrentIdx             int64
}

// RollupEventInitialize is the InitializeHermezEvent event of the
// Smart Contract
type RollupEventInitialize struct {
	ForgeL1L2BatchTimeout uint8
	FeeAddToken           *big.Int
	WithdrawalDelay       uint64
}

// RollupEventL1UserTx is an event of the Rollup Smart Contract
type RollupEventL1UserTx struct {
	ToForgeL1TxsNum int64 // QueueIndex       *big.Int
	Position        int   // TransactionIndex *big.Int
	L1UserTx        common.L1Tx
}

// RollupEventL1UserTxAux is an event of the Rollup Smart Contract
type rollupEventL1UserTxAux struct {
	ToForgeL1TxsNum uint64 // QueueIndex       *big.Int
	Position        uint8  // TransactionIndex *big.Int
	L1Tx            []byte
}

// RollupEventForgeBatch is an event of the Rollup Smart Contract
type RollupEventForgeBatch struct {
	BatchNum int64
	// Sender    ethCommon.Address
	EthTxHash    ethCommon.Hash
	L1UserTxsLen uint16
	GasUsed      uint64
	GasPrice     *big.Int
}

// RollupEventUpdateForgeL1L2BatchTimeout is an event of the Rollup Smart Contract
type RollupEventUpdateForgeL1L2BatchTimeout struct {
	NewForgeL1L2BatchTimeout int64
}

// RollupEventUpdateFeeAddToken is an event of the Rollup Smart Contract
type RollupEventUpdateFeeAddToken struct {
	NewFeeAddToken *big.Int
}

// RollupEventWithdraw is an event of the Rollup Smart Contract
type RollupEventWithdraw struct {
	Idx             uint64
	NumExitRoot     uint64
	InstantWithdraw bool
	TxHash          ethCommon.Hash // Hash of the transaction that generated this event
}

// RollupEventUpdateBucketWithdraw is an event of the Rollup Smart Contract
type RollupEventUpdateBucketWithdraw struct {
	NumBucket   int
	BlockStamp  int64 // blockNum
	Withdrawals *big.Int
}

type rollupEventUpdateBucketWithdrawAux struct {
	NumBucket   uint8
	BlockStamp  *big.Int
	Withdrawals *big.Int
}

// RollupEventUpdateWithdrawalDelay is an event of the Rollup Smart Contract
type RollupEventUpdateWithdrawalDelay struct {
	NewWithdrawalDelay uint64
}

// RollupUpdateBucketsParameters are the bucket parameters used in an update
type RollupUpdateBucketsParameters struct {
	CeilUSD         *big.Int
	BlockStamp      *big.Int
	Withdrawals     *big.Int
	RateBlocks      *big.Int
	RateWithdrawals *big.Int
	MaxWithdrawals  *big.Int
}

// type rollupEventUpdateBucketsParametersAux struct {
// 	ArrayBuckets []*big.Int
// }

// RollupEventUpdateBucketsParameters is an event of the Rollup Smart Contract
type RollupEventUpdateBucketsParameters struct {
	ArrayBuckets []RollupUpdateBucketsParameters
	SafeMode     bool
}

// RollupEventUpdateTokenExchange is an event of the Rollup Smart Contract
type RollupEventUpdateTokenExchange struct {
	AddressArray []ethCommon.Address
	ValueArray   []uint64
}

// RollupEventSafeMode is an event of the Rollup Smart Contract
type RollupEventSafeMode struct{}

// RollupEvents is the list of events in a block of the Rollup Smart Contract
type RollupEvents struct {
	L1UserTx   []RollupEventL1UserTx
	ForgeBatch []RollupEventForgeBatch
}

// NewRollupEvents creates an empty RollupEvents with the slices initialized.
func NewRollupEvents() RollupEvents {
	return RollupEvents{
		L1UserTx:   make([]RollupEventL1UserTx, 0),
		ForgeBatch: make([]RollupEventForgeBatch, 0),
	}
}

// RollupForgeBatchArgs are the arguments to the ForgeBatch function in the Rollup Smart Contract
type RollupForgeBatchArgs struct {
	NewAccountRoot *big.Int
	NewScoreRoot   *big.Int
	NewVouchRoot   *big.Int
	NewExitRoot    *big.Int
	L1UserTxs      []common.L1Tx
	ProofA         [2]*big.Int
	ProofB         [2][2]*big.Int
	ProofC         [2]*big.Int
}

// RollupForgeBatchArgsAux are the arguments to the ForgeBatch function in the Rollup Smart Contract
type rollupForgeBatchArgsAux struct {
	NewAccountRoot *big.Int
	NewVouchRoot   *big.Int
	NewScoreRoot   *big.Int
	L1L2TxsData    []byte
	// Circuit selector
	ProofA [2]*big.Int
	ProofB [2][2]*big.Int
	ProofC [2]*big.Int
}

// TODO: Update interfaces and the functions
// RollupInterface is the inteface to to Rollup Smart Contract
type RollupInterface interface {
	//
	// Smart Contract Methods
	//

	// Public Functions

	RollupForgeBatch(*RollupForgeBatchArgs, *bind.TransactOpts) (*types.Transaction, error)

	// Viewers
	RollupLastForgedBatch() (int64, error)

	//
	// Smart Contract Status
	//

	RollupConstants() (*common.RollupConstants, error)
	RollupEventsByBlock(blockNum int64, blockHash *ethCommon.Hash) (*RollupEvents, error)
	RollupForgeBatchArgs(ethCommon.Hash, uint16) (*RollupForgeBatchArgs, *ethCommon.Address, error)
}

//
// Implementation
//

// RollupClient is the implementation of the interface to the Rollup Smart Contract in ethereum.
type RollupClient struct {
	client      *EthereumClient
	chainID     *big.Int
	address     ethCommon.Address
	sybil       *Sybil.Sybil
	contractAbi abi.ABI
	opts        *bind.CallOpts
	consts      *common.RollupConstants
}

// RollupVariables returns the RollupVariables from the initialize event
func (ei *RollupEventInitialize) RollupVariables() *common.RollupVariables {
	return &common.RollupVariables{
		EthBlockNum:           0,
		ForgeL1L2BatchTimeout: int64(ei.ForgeL1L2BatchTimeout),
	}
}

// NewRollupClient creates a new RollupClient
func NewRollupClient(client *EthereumClient, address ethCommon.Address) (*RollupClient, error) {
	contractAbi, err := abi.JSON(strings.NewReader(string(Sybil.SybilABI)))
	if err != nil {
		return nil, common.Wrap(err)
	}
	sybil, err := Sybil.NewSybil(address, client.Client())
	if err != nil {
		return nil, common.Wrap(err)
	}
	chainID, err := client.EthChainID()
	if err != nil {
		return nil, common.Wrap(err)
	}
	c := &RollupClient{
		client:      client,
		chainID:     chainID,
		address:     address,
		sybil:       sybil,
		contractAbi: contractAbi,
		opts:        newCallOpts(),
	}
	consts, err := c.RollupConstants()
	if err != nil {
		return nil, common.Wrap(fmt.Errorf("RollupConstants at %v: %w", address, err))
	}
	c.consts = consts
	return c, nil
}

// RollupConstants returns the Constants of the Rollup Smart Contract
func (c *RollupClient) RollupConstants() (rollupConstants *common.RollupConstants, err error) {
	rollupConstants = new(common.RollupConstants)
	if err := c.client.Call(func(ec *ethclient.Client) error {
		maxTx, err := strconv.ParseInt(os.Getenv("MAXTX"), 10, 64)
		if err != nil {
			return common.Wrap(fmt.Errorf("failed to parse MAXTX: %w", err))
		}

		nLevels, err := strconv.ParseInt(os.Getenv("NLEVEL"), 10, 64)
		if err != nil {
			return common.Wrap(fmt.Errorf("failed to parse NLEVEL: %w", err))
		}

		var newRollupVerifier common.RollupVerifierStruct
		newRollupVerifier.MaxTx = maxTx
		newRollupVerifier.NLevels = nLevels
		rollupConstants.Verifiers = append(rollupConstants.Verifiers,
			newRollupVerifier)
		return nil
	}); err != nil {
		return nil, common.Wrap(err)
	}
	return rollupConstants, nil
}

// RollupLastForgedBatch is the interface to call the smart contract function
func (c *RollupClient) RollupLastForgedBatch() (lastForgedBatch int64, err error) {
	if err := c.client.Call(func(ec *ethclient.Client) error {
		_lastForgedBatch, err := c.sybil.LastForgedBatch(c.opts)
		lastForgedBatch = int64(_lastForgedBatch)
		return common.Wrap(err)
	}); err != nil {
		return 0, common.Wrap(err)
	}
	return lastForgedBatch, nil
}

var (
	logSYBL1Tx = crypto.Keccak256Hash([]byte(
		"TxEvent(uint32,uint8,bytes)"))
	logSYBForgeBatch = crypto.Keccak256Hash([]byte(
		"ForgeBatch(uint32,uint16)"))
)

// RollupEventsByBlock returns the events in a block that happened in the
// Rollup Smart Contract.
// To query by blockNum, set blockNum >= 0 and blockHash == nil.
// To query by blockHash set blockHash != nil, and blockNum will be ignored.
// If there are no events in that block the result is nil.
func (c *RollupClient) RollupEventsByBlock(blockNum int64,
	blockHash *ethCommon.Hash) (*RollupEvents, error) {
	var rollupEvents RollupEvents

	var blockNumBigInt *big.Int
	if blockHash == nil {
		blockNumBigInt = big.NewInt(blockNum)
	}
	query := ethereum.FilterQuery{
		BlockHash: blockHash,
		FromBlock: blockNumBigInt,
		ToBlock:   blockNumBigInt,
		Addresses: []ethCommon.Address{
			c.address,
		},
		Topics: [][]ethCommon.Hash{},
	}
	logs, err := c.client.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, common.Wrap(err)
	}
	if len(logs) == 0 {
		return nil, nil
	}

	for _, vLog := range logs {
		if blockHash != nil && vLog.BlockHash != *blockHash {
			log.Errorw("Block hash mismatch", "expected", blockHash.String(), "got", vLog.BlockHash.String())
			return nil, common.Wrap(ErrBlockHashMismatchEvent)
		}
		switch vLog.Topics[0] {
		case logSYBL1Tx:
			var L1UserTxAux rollupEventL1UserTxAux
			var L1UserTx RollupEventL1UserTx
			err := c.contractAbi.UnpackIntoInterface(&L1UserTxAux, "TxEvent", vLog.Data)
			if err != nil {
				return nil, common.Wrap(err)
			}
			L1Tx, err := common.L1UserTxFromBytes(L1UserTxAux.L1Tx)
			if err != nil {
				return nil, common.Wrap(err)
			}
			toForgeL1TxsNum := new(big.Int).SetBytes(vLog.Topics[1][:]).Int64()
			L1Tx.ToForgeL1TxsNum = &toForgeL1TxsNum
			L1Tx.Position = int(new(big.Int).SetBytes(vLog.Topics[2][:]).Int64())
			L1Tx.EthTxHash = vLog.TxHash
			//Get l1Fee in eth wei spent in the l1 tx
			// tx, _, err := c.client.client.TransactionByHash(context.Background(), vLog.TxHash)
			// if err != nil {
			// 	return nil, common.Wrap(fmt.Errorf("failed to get TransactionByHash, hash: %s, err: %w", vLog.TxHash.String(), err))
			// }
			// l1Fee := new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(tx.Gas()))
			// L1Tx.L1Fee = l1Fee
			L1UserTx.L1UserTx = *L1Tx
			rollupEvents.L1UserTx = append(rollupEvents.L1UserTx, L1UserTx)
		case logSYBForgeBatch:
			var forgeBatch RollupEventForgeBatch
			err := c.contractAbi.UnpackIntoInterface(&forgeBatch, "ForgeBatch", vLog.Data)
			if err != nil {
				return nil, common.Wrap(err)
			}
			forgeBatch.BatchNum = new(big.Int).SetBytes(vLog.Topics[1][:]).Int64()
			forgeBatch.EthTxHash = vLog.TxHash
			//Check tx info using EthTxHash to get gasprice and gas used
			tx, _, err := c.client.client.TransactionByHash(context.Background(), vLog.TxHash)
			if err != nil {
				return nil, common.Wrap(fmt.Errorf("failed to get TransactionByHash, hash: %s, err: %w", vLog.TxHash.String(), err))
			}
			forgeBatch.GasPrice = tx.GasPrice()
			// Get gas used from TxReceipt
			txReceipt, err := c.client.client.TransactionReceipt(context.Background(), vLog.TxHash)
			if err != nil {
				return nil, common.Wrap(fmt.Errorf("failed to get TransactionByHash, hash: %s, err: %w", vLog.TxHash.String(), err))
			}
			forgeBatch.GasUsed = txReceipt.GasUsed
			rollupEvents.ForgeBatch = append(rollupEvents.ForgeBatch, forgeBatch)
		}
	}
	return &rollupEvents, nil
}

// RollupForgeBatch is the interface to call the smart contract function
func (c *RollupClient) RollupForgeBatch(args *RollupForgeBatchArgs, auth *bind.TransactOpts) (tx *types.Transaction, err error) {
	if auth == nil {
		auth, err = c.client.NewAuth()
		if err != nil {
			return nil, err
		}
		auth.GasLimit = 1000000
	}

	// nLevels := c.consts.Verifiers[args.VerifierIdx].NLevels //check verifiers

	// newLastIdx := big.NewInt(int64(args.NewLastIdx))

	// var l1TxData []byte
	// for i := 0; i < len(args.L1UserTxs); i++ {
	// 	l1User := args.L1UserTxs[i]
	// 	bytesl1User, err := l1User.BytesDataAvailability(uint32(nLevels))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	l1TxData = append(l1TxData, bytesl1User[:]...)
	// }

	// TODO: Need to send ZK Proof here on last param
	tx, err = c.sybil.ForgeBatch(
		auth,
		args.NewAccountRoot,
		args.NewVouchRoot,
		args.NewScoreRoot,
		args.NewExitRoot,
		args.ProofA,
		args.ProofB,
		args.ProofC,
	)
	if err != nil {
		return nil, fmt.Errorf("Sybil.ForgeBatch: %w", err)
	}
	return tx, nil
}

// RollupForgeBatchArgs returns the arguments used in a ForgeBatch call in the
// Rollup Smart Contract in the given transaction, and the sender address.
func (c *RollupClient) RollupForgeBatchArgs(ethTxHash ethCommon.Hash,
	l1UserTxsLen uint16) (*RollupForgeBatchArgs, *ethCommon.Address, error) {
	tx, _, err := c.client.client.TransactionByHash(context.Background(), ethTxHash)
	if err != nil {
		return nil, nil, common.Wrap(fmt.Errorf("TransactionByHash: %w", err))
	}
	txData := tx.Data()

	method, err := c.contractAbi.MethodById(txData[:4])
	if err != nil {
		return nil, nil, common.Wrap(err)
	}
	receipt, err := c.client.client.TransactionReceipt(context.Background(), ethTxHash)
	if err != nil {
		return nil, nil, common.Wrap(err)
	}
	sender, err := c.client.client.TransactionSender(context.Background(), tx,
		receipt.Logs[0].BlockHash, receipt.Logs[0].Index)
	if err != nil {
		return nil, nil, common.Wrap(err)
	}
	var aux rollupForgeBatchArgsAux
	if values, err := method.Inputs.Unpack(txData[4:]); err != nil {
		return nil, nil, common.Wrap(err)
	} else if err := method.Inputs.Copy(&aux, values); err != nil {
		return nil, nil, common.Wrap(err)
	}
	rollupForgeBatchArgs := RollupForgeBatchArgs{
		NewAccountRoot: aux.NewAccountRoot,
		NewVouchRoot:   aux.NewVouchRoot,
		NewScoreRoot:   aux.NewScoreRoot,
		ProofA:         aux.ProofA,
		ProofB:         aux.ProofB,
		ProofC:         aux.ProofC,
	}
	// TODO: restruct the logic to fetch NLevels -> constants
	nLevels := c.consts.Verifiers[rollupForgeBatchArgs.VerifierIdx].NLevels
	lenL1TxsBytes := int((nLevels/8)*2 + common.Float40BytesLength + 1) //nolint:gomnd
	numBytesL1TxUser := int(l1UserTxsLen) * lenL1TxsBytes
	l1UserTxsData := []byte{}
	if l1UserTxsLen > 0 {
		l1UserTxsData = aux.L1L2TxsData[:numBytesL1TxUser]
	}
	for i := 0; i < int(l1UserTxsLen); i++ {
		l1Tx, err :=
			common.L1TxFromDataAvailability(l1UserTxsData[i*lenL1TxsBytes:(i+1)*lenL1TxsBytes],
				uint32(nLevels))
		if err != nil {
			return nil, nil, common.Wrap(err)
		}
		rollupForgeBatchArgs.L1UserTxs = append(rollupForgeBatchArgs.L1UserTxs, *l1Tx)
	}
	return &rollupForgeBatchArgs, &sender, nil
}
