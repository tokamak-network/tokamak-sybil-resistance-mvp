package synchronizer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"testing"
	"tokamak-sybil-resistance/common"
	"tokamak-sybil-resistance/database/historydb"
	"tokamak-sybil-resistance/database/statedb"
	"tokamak-sybil-resistance/test"
	"tokamak-sybil-resistance/test/til"

	dbUtils "tokamak-sybil-resistance/database"

	ethCommon "github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type timer struct {
	time int64
}

func (t *timer) Time() int64 {
	currentTime := t.time
	t.time++
	return currentTime
}

func accountsCmp(accounts []common.Account) func(i, j int) bool {
	return func(i, j int) bool { return accounts[i].Idx < accounts[j].Idx }
}

// Check Sync output and HistoryDB state against expected values generated by
// til
func checkSyncBlock(t *testing.T, s *Synchronizer, blockNum int, block,
	syncBlock *common.BlockData) {
	// Check Blocks
	dbBlocks, err := s.historyDB.GetAllBlocks()
	require.NoError(t, err)
	dbBlocks = dbBlocks[1:] // ignore block 0, added by default in the DB
	assert.Equal(t, blockNum, len(dbBlocks))
	assert.Equal(t, int64(blockNum), dbBlocks[blockNum-1].Num)
	assert.NotEqual(t, dbBlocks[blockNum-1].Hash, dbBlocks[blockNum-2].Hash)
	assert.Greater(t, dbBlocks[blockNum-1].Timestamp.Unix(), dbBlocks[blockNum-2].Timestamp.Unix())

	// Check submitted L1UserTxs
	assert.Equal(t, len(block.Rollup.L1UserTxs), len(syncBlock.Rollup.L1UserTxs))
	dbL1UserTxs, err := s.historyDB.GetAllL1UserTxs()
	require.NoError(t, err)
	// Ignore BatchNum in syncBlock.L1UserTxs because this value is set by
	// the HistoryDB. Also ignore EffectiveAmount & EffectiveDepositAmount
	// because this value is set by StateDB.ProcessTxs.
	for i := range syncBlock.Rollup.L1UserTxs {
		syncBlock.Rollup.L1UserTxs[i].BatchNum = block.Rollup.L1UserTxs[i].BatchNum
		assert.Nil(t, syncBlock.Rollup.L1UserTxs[i].EffectiveDepositAmount)
		assert.Nil(t, syncBlock.Rollup.L1UserTxs[i].EffectiveAmount)
	}
	assert.Equal(t, block.Rollup.L1UserTxs, syncBlock.Rollup.L1UserTxs)
	for _, tx := range block.Rollup.L1UserTxs {
		var dbTx *common.L1Tx
		// Find tx in DB output
		for _, _dbTx := range dbL1UserTxs {
			if *tx.ToForgeL1TxsNum == *_dbTx.ToForgeL1TxsNum &&
				tx.Position == _dbTx.Position {
				dbTx = new(common.L1Tx)
				*dbTx = _dbTx
				// NOTE: Overwrite EffectiveFromIdx in L1UserTx
				// from db because we don't expect
				// EffectiveFromIdx to be set yet, as this tx
				// is not in yet forged
				dbTx.EffectiveFromIdx = 0
				break
			}
		}
		// If the tx has been forged in this block, this will be
		// reflected in the DB, and so the Effective values will be
		// already set
		if dbTx.BatchNum != nil {
			tx.EffectiveAmount = tx.Amount
			tx.EffectiveDepositAmount = tx.DepositAmount
		}
		assert.Equal(t, &tx, dbTx) //nolint:gosec
	}

	// Check Batches
	assert.Equal(t, len(block.Rollup.Batches), len(syncBlock.Rollup.Batches))
	dbBatches, err := s.historyDB.GetAllBatches()
	require.NoError(t, err)

	require.NoError(t, err)
	dbExits, err := s.historyDB.GetAllExits()
	require.NoError(t, err)
	// dbL1CoordinatorTxs := []common.L1Tx{}
	for i, batch := range block.Rollup.Batches {
		var dbBatch *common.Batch
		// Find batch in DB output
		for _, _dbBatch := range dbBatches {
			if batch.Batch.BatchNum == _dbBatch.BatchNum {
				dbBatch = new(common.Batch)
				*dbBatch = _dbBatch
				dbBatch.GasPrice = batch.Batch.GasPrice
				break
			}
		}
		syncBatch := syncBlock.Rollup.Batches[i]

		// We don't care about TotalFeesUSD.  Use the syncBatch that
		// has a TotalFeesUSD inserted by the HistoryDB
		batch.Batch.TotalFeesUSD = syncBatch.Batch.TotalFeesUSD
		assert.Equal(t, batch.CreatedAccounts, syncBatch.CreatedAccounts)
		batch.Batch.NumAccounts = len(batch.CreatedAccounts)

		// Test field by field to facilitate debugging of errors
		assert.Equal(t, len(batch.L1UserTxs), len(syncBatch.L1UserTxs))
		// NOTE: EffectiveFromIdx is set to til L1UserTxs in
		// `FillBlocksForgedL1UserTxs` function
		for j := range syncBatch.L1UserTxs {
			assert.NotEqual(t, 0, syncBatch.L1UserTxs[j].EffectiveFromIdx)
		}
		assert.Equal(t, batch.L1UserTxs, syncBatch.L1UserTxs)

		// In exit tree, we only check AccountIdx and Balance, because
		// it's what we have precomputed before.
		require.Equal(t, len(batch.ExitTree), len(syncBatch.ExitTree))
		for j := range batch.ExitTree {
			exit := &batch.ExitTree[j]
			assert.Equal(t, exit.AccountIdx, syncBatch.ExitTree[j].AccountIdx)
			assert.Equal(t, exit.Balance, syncBatch.ExitTree[j].Balance)
			*exit = syncBatch.ExitTree[j]
		}
		assert.Equal(t, batch.Batch, syncBatch.Batch)
		// Ignore updated accounts
		syncBatch.UpdatedAccounts = nil
		assert.Equal(t, batch, syncBatch)
		assert.Equal(t, &batch.Batch, dbBatch) //nolint:gosec

		// Check forged L1UserTxs from DB, and check effective values
		// in sync output
		for j, tx := range batch.L1UserTxs {
			var dbTx *common.L1Tx
			// Find tx in DB output
			for _, _dbTx := range dbL1UserTxs {
				if *tx.BatchNum == *_dbTx.BatchNum &&
					tx.Position == _dbTx.Position {
					dbTx = new(common.L1Tx)
					*dbTx = _dbTx
					break
				}
			}
			// TODO: sync rollup smart contract
			// assert.Equal(t, &tx, dbTx) //nolint:gosec

			syncTx := &syncBlock.Rollup.Batches[i].L1UserTxs[j]
			assert.Equal(t, syncTx.DepositAmount, syncTx.EffectiveDepositAmount)
			assert.Equal(t, syncTx.Amount, syncTx.EffectiveAmount)
		}

		// Check Exits from DB
		for _, exit := range batch.ExitTree {
			var dbExit *common.ExitInfo
			// Find exit in DB output
			for _, _dbExit := range dbExits {
				if exit.BatchNum == _dbExit.BatchNum &&
					exit.AccountIdx == _dbExit.AccountIdx {
					dbExit = new(common.ExitInfo)
					*dbExit = _dbExit
					break
				}
			}
			// Compare MerkleProof in JSON because unmarshaled 0
			// big.Int leaves the internal big.Int array at nil,
			// and gives trouble when comparing big.Int with
			// internal big.Int array != nil but empty.
			mtp, err := json.Marshal(exit.MerkleProof)
			require.NoError(t, err)
			dbMtp, err := json.Marshal(dbExit.MerkleProof)
			require.NoError(t, err)
			assert.Equal(t, mtp, dbMtp)
			dbExit.MerkleProof = exit.MerkleProof
			assert.Equal(t, &exit, dbExit) //nolint:gosec
		}
	}

	// Compare accounts from HistoryDB with StateDB (they should match)
	dbAccounts, err := s.historyDB.GetAllAccounts()
	require.NoError(t, err)
	sdbAccounts, err := s.stateDB.TestGetAccounts()
	require.NoError(t, err)
	assertEqualAccountsHistoryDBStateDB(t, dbAccounts, sdbAccounts)
}

func assertEqualAccountsHistoryDBStateDB(t *testing.T, hdbAccs, sdbAccs []common.Account) {
	assert.Equal(t, len(hdbAccs), len(sdbAccs))
	sort.SliceStable(hdbAccs, accountsCmp(hdbAccs))
	sort.SliceStable(sdbAccs, accountsCmp(sdbAccs))
	for i := range hdbAccs {
		hdbAcc := hdbAccs[i]
		sdbAcc := sdbAccs[i]
		assert.Equal(t, hdbAcc.Idx, sdbAcc.Idx)
		assert.Equal(t, hdbAcc.EthAddr, sdbAcc.EthAddr)
		assert.Equal(t, hdbAcc.BJJ, sdbAcc.BJJ)
	}
}

var chainID uint64 = 0
var deleteme = []string{}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	for _, dir := range deleteme {
		if err := os.RemoveAll(dir); err != nil {
			panic(err)
		}
	}
	os.Exit(exitVal)
}

func newTestModules(t *testing.T) (*statedb.StateDB, *historydb.HistoryDB) {
	// Int State DB
	dir, err := os.MkdirTemp("", "tmpdb")
	require.NoError(t, err)
	deleteme = append(deleteme, dir)

	stateDB, err := statedb.NewStateDB(statedb.Config{Path: dir, Keep: 128,
		Type: statedb.TypeSynchronizer, NLevels: 32})
	require.NoError(t, err)

	// Init History DB
	db, err := dbUtils.InitSQLDB()
	require.NoError(t, err)
	historyDB := historydb.NewHistoryDB(db, db /*, nil*/)

	t.Cleanup(func() {
		test.MigrationsDownTest(historyDB.DB())
		stateDB.Close()
	})

	return stateDB, historyDB
}

func newBigInt(s string) *big.Int {
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic(fmt.Errorf("Can't set big.Int from %s", s))
	}
	return v
}

func TestSyncGeneral(t *testing.T) {
	stateDB, historyDB := newTestModules(t)

	// Init eth client
	var timer timer
	clientSetup := test.NewClientSetupExample()
	clientSetup.ChainID = big.NewInt(int64(chainID))
	bootCoordAddr := ethCommon.HexToAddress("0xE39fEc6224708f0772D2A74fd3f9055A90E0A9f2")
	client := test.NewClient(true, &timer, &ethCommon.Address{}, clientSetup)

	// Create Synchronizer
	s, err := NewSynchronizer(
		client,
		historyDB,
		stateDB,
		Config{
			StatsUpdateBlockNumDiffThreshold: 100,
			StatsUpdateFrequencyDivider:      100,
			StartBlockNum:                    1,
		},
	)
	require.NoError(t, err)

	ctx := context.Background()

	//
	// First Sync from an initial state
	//
	stats := s.Stats()
	assert.Equal(t, false, stats.Synced())

	// Test Sync for rollup genesis block
	syncBlock, discards, err := s.Sync(ctx, nil)
	require.NoError(t, err)
	require.Nil(t, discards)
	require.NotNil(t, syncBlock)
	require.Nil(t, syncBlock.Rollup.Vars)
	assert.Equal(t, int64(1), syncBlock.Block.Num)
	stats = s.Stats()
	assert.Equal(t, int64(1), stats.Eth.FirstBlockNum)
	assert.Equal(t, int64(1), stats.Eth.LastBlock.Num)
	assert.Equal(t, int64(1), stats.Sync.LastBlock.Num)
	// vars := s.SCVars()
	// assert.Equal(t, *clientSetup.RollupVariables, vars.Rollup)

	dbBlocks, err := s.historyDB.GetAllBlocks()
	require.NoError(t, err)
	assert.Equal(t, 2, len(dbBlocks))
	assert.Equal(t, int64(1), dbBlocks[1].Num)

	// Sync again and expect no new blocks
	syncBlock, discards, err = s.Sync(ctx, nil)
	require.NoError(t, err)
	require.Nil(t, discards)
	require.Nil(t, syncBlock)

	//
	// Generate blockchain and smart contract data, and fill the test smart contracts
	//

	// Generate blockchain data with til
	set1 := `
		Type: Blockchain

		CreateAccountDeposit A: 2000 // Idx=256+0=256
		CreateAccountDeposit B: 500  // Idx=256+1=257
		CreateAccountDeposit C: 2000 // Idx=256+2=258
		CreateAccountDeposit D: 500  // Idx=256+3=259

		> batchL1 // forge L1UserTxs{nil}, freeze defined L1UserTxs{4}
		> batchL1 // forge defined L1UserTxs{4}, freeze L1UserTxs{nil}
		> block // blockNum=2

		ForceExit A: 100
		ForceExit B: 80

		CreateVouch C-A
		CreateVouch C-B
		CreateVouch C-D
		// Exit C: 50
		// Exit D: 30

		> batchL1 // forge L1UserTxs{nil}, freeze defined L1UserTxs{2}
		> batchL1 // forge L1UserTxs{2}, freeze defined L1UserTxs{nil}
		> block // blockNum=3
	`
	tc := til.NewContext(chainID, common.RollupConstMaxL1UserTx)
	tilCfgExtra := til.ConfigExtra{
		BootCoordAddr: bootCoordAddr,
		CoordUser:     "A",
	}
	blocks, err := tc.GenerateBlocks(set1)
	require.NoError(t, err)
	// Sanity check
	require.Equal(t, 2, len(blocks))
	// blocks 0 (blockNum=2)
	i := 0
	require.Equal(t, 2, int(blocks[i].Block.Num))
	require.Equal(t, 4, len(blocks[i].Rollup.L1UserTxs))
	require.Equal(t, 2, len(blocks[i].Rollup.Batches))
	// require.Equal(t, 2, len(blocks[i].Rollup.Batches[0].L1CoordinatorTxs))

	// Set StateRoots for batches manually (til doesn't set it)
	blocks[i].Rollup.Batches[0].Batch.AccountRoot =
		newBigInt("11432094872416618651837327395264042968926668786266585816625577088890451620254")
	blocks[i].Rollup.Batches[0].Batch.VouchRoot =
		newBigInt("11432094872416618651837327395264042968926668786266585816625577088890451620254")
	blocks[i].Rollup.Batches[0].Batch.ScoreRoot =
		newBigInt("11432094872416618651837327395264042968926668786266585816625577088890451620254")

	blocks[i].Rollup.Batches[1].Batch.AccountRoot =
		newBigInt("16914212635847451457076355431350059348585556180740555407203882688922702410093")
	blocks[i].Rollup.Batches[1].Batch.VouchRoot =
		newBigInt("16914212635847451457076355431350059348585556180740555407203882688922702410093")
	blocks[i].Rollup.Batches[1].Batch.ScoreRoot =
		newBigInt("16914212635847451457076355431350059348585556180740555407203882688922702410093")

		// blocks 1 (blockNum=3)
	i = 1
	require.Equal(t, 3, int(blocks[i].Block.Num))
	require.Equal(t, 5, len(blocks[i].Rollup.L1UserTxs))
	require.Equal(t, 2, len(blocks[i].Rollup.Batches))
	// Set StateRoots for batches manually (til doesn't set it)
	blocks[i].Rollup.Batches[0].Batch.AccountRoot =
		newBigInt("13535760140937349829640752733057594576151546047374619177689224612061148090678")
	blocks[i].Rollup.Batches[0].Batch.VouchRoot =
		newBigInt("13535760140937349829640752733057594576151546047374619177689224612061148090678")
	blocks[i].Rollup.Batches[0].Batch.ScoreRoot =
		newBigInt("13535760140937349829640752733057594576151546047374619177689224612061148090678")

	blocks[i].Rollup.Batches[1].Batch.AccountRoot =
		newBigInt("19413739476363469870744893742469056615496274423228302914851564791727474664804")
	blocks[i].Rollup.Batches[1].Batch.VouchRoot =
		newBigInt("19413739476363469870744893742469056615496274423228302914851564791727474664804")
	blocks[i].Rollup.Batches[1].Batch.ScoreRoot =
		newBigInt("19413739476363469870744893742469056615496274423228302914851564791727474664804")

	err = tc.FillBlocksExtra(blocks, &tilCfgExtra)
	require.NoError(t, err)
	tc.FillBlocksL1UserTxsBatchNum(blocks)
	err = tc.FillBlocksForgedL1UserTxs(blocks)
	require.NoError(t, err)

	// Add block data to the smart contracts
	err = client.CtlAddBlocks(blocks)
	require.NoError(t, err)

	//
	// Sync to synchronize the current state from the test smart contracts,
	// and check the outcome
	//

	// Block 2

	syncBlock, discards, err = s.Sync(ctx, nil)
	require.NoError(t, err)
	require.Nil(t, discards)
	require.NotNil(t, syncBlock)
	assert.Nil(t, syncBlock.Rollup.Vars)
	assert.Equal(t, int64(2), syncBlock.Block.Num)
	stats = s.Stats()
	assert.Equal(t, int64(1), stats.Eth.FirstBlockNum)
	assert.Equal(t, int64(3), stats.Eth.LastBlock.Num)
	assert.Equal(t, int64(2), stats.Sync.LastBlock.Num)
	// Set ethereum transaction hash (til doesn't set it)
	blocks[0].Rollup.Batches[0].Batch.EthTxHash = syncBlock.Rollup.Batches[0].Batch.EthTxHash
	blocks[0].Rollup.Batches[1].Batch.EthTxHash = syncBlock.Rollup.Batches[1].Batch.EthTxHash
	blocks[0].Rollup.Batches[0].Batch.GasPrice = syncBlock.Rollup.Batches[0].Batch.GasPrice
	blocks[0].Rollup.Batches[1].Batch.GasPrice = syncBlock.Rollup.Batches[1].Batch.GasPrice

	checkSyncBlock(t, s, 2, &blocks[0], syncBlock)

	// Block 3

	syncBlock, discards, err = s.Sync(ctx, nil)
	require.NoError(t, err)
	require.Nil(t, discards)
	require.NotNil(t, syncBlock)
	assert.Nil(t, syncBlock.Rollup.Vars)
	assert.Equal(t, int64(3), syncBlock.Block.Num)
	stats = s.Stats()
	assert.Equal(t, int64(1), stats.Eth.FirstBlockNum)
	assert.Equal(t, int64(3), stats.Eth.LastBlock.Num)
	assert.Equal(t, int64(3), stats.Sync.LastBlock.Num)
	// Set ethereum transaction hash (til doesn't set it)
	blocks[1].Rollup.Batches[0].Batch.EthTxHash = syncBlock.Rollup.Batches[0].Batch.EthTxHash
	blocks[1].Rollup.Batches[1].Batch.EthTxHash = syncBlock.Rollup.Batches[1].Batch.EthTxHash
	blocks[1].Rollup.Batches[0].Batch.GasPrice = syncBlock.Rollup.Batches[0].Batch.GasPrice
	blocks[1].Rollup.Batches[1].Batch.GasPrice = syncBlock.Rollup.Batches[1].Batch.GasPrice

	checkSyncBlock(t, s, 3, &blocks[1], syncBlock)

	// Block 4
	// Generate 2 withdraws manually
	_, err = client.RollupWithdrawMerkleProof(tc.Accounts["A"].BJJ.Public().Compress(), 4, 256,
		big.NewInt(100), []*big.Int{}, true)
	require.NoError(t, err)
	_, err = client.RollupWithdrawMerkleProof(tc.Accounts["C"].BJJ.Public().Compress(), 3, 258,
		big.NewInt(50), []*big.Int{}, false)
	require.NoError(t, err)
	client.CtlMineBlock()

	syncBlock, discards, err = s.Sync(ctx, nil)
	require.NoError(t, err)
	require.Nil(t, discards)
	require.NotNil(t, syncBlock)
	assert.Nil(t, syncBlock.Rollup.Vars)
	assert.Equal(t, int64(4), syncBlock.Block.Num)
	stats = s.Stats()
	assert.Equal(t, int64(1), stats.Eth.FirstBlockNum)
	assert.Equal(t, int64(4), stats.Eth.LastBlock.Num)
	assert.Equal(t, int64(4), stats.Sync.LastBlock.Num)

	// dbExits, err := s.historyDB.GetAllExits()
	// require.NoError(t, err)
	// foundA1, foundC1 := false, false

	// for _, exit := range dbExits {
	// 	if exit.AccountIdx == 256 && exit.BatchNum == 4 {
	// 		foundA1 = true
	// 	}
	// 	if exit.AccountIdx == 258 && exit.BatchNum == 3 {
	// 		foundC1 = true
	// 	}
	// }

	// assert.True(t, foundA1)
	// assert.True(t, foundC1)

	// Block 5
	// Update variables manually
	rollupVars, err := s.historyDB.GetSCVars()
	require.NoError(t, err)
	rollupVars.ForgeL1L2BatchTimeout = 42
	_, err = client.RollupUpdateForgeL1BatchTimeout(rollupVars.ForgeL1L2BatchTimeout)
	require.NoError(t, err)

	client.CtlMineBlock()

	syncBlock, discards, err = s.Sync(ctx, nil)
	require.NoError(t, err)
	require.Nil(t, discards)
	require.NotNil(t, syncBlock)
	assert.NotNil(t, syncBlock.Rollup.Vars)
	assert.Equal(t, int64(5), syncBlock.Block.Num)
	stats = s.Stats()
	assert.Equal(t, int64(1), stats.Eth.FirstBlockNum)
	assert.Equal(t, int64(5), stats.Eth.LastBlock.Num)
	assert.Equal(t, int64(5), stats.Sync.LastBlock.Num)

	dbRollupVars, err := s.historyDB.GetSCVars()
	require.NoError(t, err)
	// Set EthBlockNum for Vars to the blockNum in which they were updated (should be 5)
	rollupVars.EthBlockNum = syncBlock.Block.Num
	assert.Equal(t, rollupVars, dbRollupVars)
}
