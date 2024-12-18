package til

import (
	"math/big"
	"testing"
	"tokamak-sybil-resistance/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateKeys(t *testing.T) {
	tc := NewContext(0, common.RollupConstMaxL1UserTx)
	usernames := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	assert.Equal(t, 0, len(tc.Accounts))
	tc.generateKeys(usernames)
	assert.Equal(t, len(usernames), len(tc.Accounts))
}

func TestGenerateBlocksNoBatches(t *testing.T) {
	set := `
		Type: Blockchain

		CreateAccountDeposit A: 11
		CreateAccountDeposit B: 22

		> block

		Deposit A: 6
		CreateVouch A-B
		CreateVouch B-A
		DeleteVouch A-B

		> block
	`
	tc := NewContext(0, common.RollupConstMaxL1UserTx)
	blocks, err := tc.GenerateBlocks(set)
	require.NoError(t, err)

	assert.Equal(t, 2, len(blocks))

	// assert.Equal(t, 0, len(blocks[0].Rollup.Batches))
	assert.Equal(t, 2, len(blocks[0].Rollup.L1UserTxs))

	assert.Equal(t, 0, len(blocks[1].Rollup.Batches))
	assert.Equal(t, 4, len(blocks[1].Rollup.L1UserTxs)) // Vouch txs are not L1UserTxs
}

func TestGenerateBlocksWithBatches(t *testing.T) {
	set := `
		Type: Blockchain
	
		CreateAccountDeposit A: 10
		CreateAccountDeposit B: 5
		Deposit A: 6
		CreateAccountDeposit C: 5
		CreateAccountDeposit D: 5

		> batchL1 // batchNum = 1
		> batchL1 // batchNum = 2

		Deposit A: 3
		CreateVouch A-B
		CreateVouch B-A
		CreateVouch A-C
		DeleteVouch A-B

		// set new batch
		> batchL1 // batchNum = 3

		> block

		// Exits
		CreateVouch C-D
		// Exit A: 5

		> batchL1 // batchNum = 4
		> block
	`
	tc := NewContext(0, common.RollupConstMaxL1UserTx)
	blocks, err := tc.GenerateBlocks(set)
	require.NoError(t, err)
	assert.Equal(t, 2, len(blocks))
	assert.Equal(t, 3, len(blocks[0].Rollup.Batches))
	assert.Equal(t, 10, len(blocks[0].Rollup.L1UserTxs))
	// assert.Equal(t, 0, len(blocks[0].Rollup.Batches[2].L2Txs))

	assert.Equal(t, 1, len(blocks[1].Rollup.Batches))
	assert.Equal(t, 1, len(blocks[1].Rollup.L1UserTxs))
	// assert.Equal(t, 1, len(blocks[1].Rollup.Batches[0].L2Txs))

	// Check expected values generated by each line
	// #0: Deposit A: 10
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[0], common.TxTypeCreateAccountDeposit,
		"A", "", big.NewInt(10), nil)
	// #1: Deposit B: 5
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[1], common.TxTypeCreateAccountDeposit,
		"B", "", big.NewInt(5), nil)
	// #2: Deposit A: 16
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[2], common.TxTypeDeposit,
		"A", "", big.NewInt(6), nil)
	// #3: Deposit C: 5
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[3], common.TxTypeCreateAccountDeposit,
		"C", "", big.NewInt(5), nil)
	// #4: Deposit D: 5
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[4], common.TxTypeCreateAccountDeposit,
		"D", "", big.NewInt(5), nil)
	// #5: Deposit A: 16
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[5], common.TxTypeDeposit,
		"A", "", big.NewInt(3), nil)
	// #6: CreateVouch A-B
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[6], common.TxTypeCreateVouch, "A",
		"B", nil, nil)
	// #7: CreateVouch B-A
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[7], common.TxTypeCreateVouch, "B",
		"A", nil, nil)
	// #8: CreateVouch A-C
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[8], common.TxTypeCreateVouch, "A",
		"C", nil, nil)
	// #9: DeleteVouch A-B
	tc.checkL1TxParams(t, blocks[0].Rollup.L1UserTxs[9], common.TxTypeDeleteVouch, "A",
		"B", nil, nil)
	// #10: CreateVouch C-D
	tc.checkL1TxParams(t, blocks[1].Rollup.L1UserTxs[0], common.TxTypeCreateVouch, "C",
		"D", nil, nil)
}

func (tc *Context) checkL1TxParams(t *testing.T, tx common.L1Tx, typ common.TxType,
	from, to string, depositAmount, amount *big.Int) {
	assert.Equal(t, typ, tx.Type)
	if tx.FromIdx != common.AccountIdx(0) {
		assert.Equal(t, tc.Accounts[from].Idx, tx.FromIdx)
	}
	assert.Equal(t, tc.Accounts[from].Addr.Hex(), tx.FromEthAddr.Hex())
	assert.Equal(t, tc.Accounts[from].BJJ.Public().Compress(), tx.FromBJJ)
	if tx.ToIdx != common.AccountIdx(0) {
		assert.Equal(t, tc.Accounts[to].Idx, tx.ToIdx)
	}
	if depositAmount != nil {
		assert.Equal(t, depositAmount, tx.DepositAmount)
	}
	if amount != nil {
		assert.Equal(t, amount, tx.Amount)
	}
}

// func (tc *Context) checkL2TxParams(t *testing.T, tx common.L2Tx, typ common.TxType,
// 	from, to string, amount *big.Int, batchNum common.BatchNum) {
// 	assert.Equal(t, typ, tx.Type)
// 	assert.Equal(t, tc.Accounts[from].Idx, tx.FromIdx)
// 	if tx.Type != common.TxTypeExit {
// 		assert.Equal(t, tc.Accounts[to].Idx, tx.ToIdx)
// 	}
// 	if amount != nil {
// 		assert.Equal(t, amount, tx.Amount)
// 	}
// 	assert.Equal(t, batchNum, tx.BatchNum)
// }

// func TestGeneratePoolL2Txs(t *testing.T) {
// 	set := `
// 		Type: Blockchain
// 		CreateAccountDeposit A: 10
// 		CreateAccountDeposit C: 5
// 		CreateAccountDeposit B: 5
// 		CreateAccountDeposit D: 0
// 		> batchL1
// 		> batchL1
// 	`
// 	tc := NewContext(0, common.RollupConstMaxL1UserTx)
// 	_, err := tc.GenerateBlocks(set)
// 	require.NoError(t, err)
// 	set = `
// 		Type: PoolL2
// 		PoolCreateVouch A-B
// 		PoolCreateVouch A-C
// 		PoolDeleteVouch A-C
// 		PoolExit A: 3
// 	`
// 	poolL2Txs, err := tc.GeneratePoolL2Txs(set)
// 	require.NoError(t, err)
// 	assert.Equal(t, 4, len(poolL2Txs))
// 	assert.Equal(t, common.TxTypeCreateVouch, poolL2Txs[0].Type)
// 	assert.Equal(t, common.TxTypeDeleteVouch, poolL2Txs[2].Type)
// 	assert.Equal(t, common.TxTypeExit, poolL2Txs[3].Type)
// 	assert.Equal(t, common.Nonce(0), poolL2Txs[0].Nonce)
// 	assert.Equal(t, common.Nonce(1), poolL2Txs[1].Nonce)
// 	assert.Equal(t, common.Nonce(2), poolL2Txs[2].Nonce)
// 	assert.Equal(t, common.Nonce(3), poolL2Txs[3].Nonce)

// 	// load another set in the same Context
// 	set = `
// 		Type: PoolL2
// 		PoolExit B: 3
// 	`
// 	poolL2Txs, err = tc.GeneratePoolL2Txs(set)
// 	require.NoError(t, err)
// 	assert.Equal(t, common.Nonce(0), poolL2Txs[0].Nonce)
// }

func TestGenerateErrors(t *testing.T) {
	// check transactions when account is not created yet
	set := `
		Type: Blockchain
		CreateAccountDeposit A: 10
		> batchL1
		CreateAccountDeposit B
		> batchL1
	`
	tc := NewContext(0, common.RollupConstMaxL1UserTx)
	_, err := tc.GenerateBlocks(set)
	require.Equal(t, "line 4: CreateAccountDepositB> batchL1\n, err: "+
		"expected ':', found '>'", err.Error())

	set = `
		Type: Blockchain
		CreateAccountDeposit A: 10
		> batchL1
		> batchL1
	`
	tc = NewContext(0, common.RollupConstMaxL1UserTx)
	_, err = tc.GenerateBlocks(set)
	require.NoError(t, err)

	// check nonces
	set = `
		Type: Blockchain
		CreateAccountDeposit A: 10
		> batchL1
		// Exit A: 3
		> batchL1
	`
	tc = NewContext(0, common.RollupConstMaxL1UserTx)
	_, err = tc.GenerateBlocks(set)
	require.NoError(t, err)
	// assert.Equal(t, common.Nonce(1), tc.Accounts["A"].Nonce)
	assert.Equal(t, common.AccountIdx(256), tc.Accounts["A"].Idx)

	// check vouching syntax errors on L1
	set = `
		Type: Blockchain
		CreateVouch A
		> batchL1
	`
	tc = NewContext(0, common.RollupConstMaxL1UserTx)
	_, err = tc.GenerateBlocks(set)
	require.Equal(t, "line 2: CreateVouchA> batchL1\n, err: expected '-', found '>'", err.Error())

	set = `
		Type: Blockchain
		CreateVouch A-
	`
	tc = NewContext(0, common.RollupConstMaxL1UserTx)
	_, err = tc.GenerateBlocks(set)
	require.Equal(t, "line 2: CreateVouchA-, err: expected 'to' account name, found ''", err.Error())

	// // check vouching syntax errors on L2
	// set = `
	// 	Type: PoolL2
	// 	DeleteVouch A
	// 	> batch
	// `
	// tc = NewContext(0, common.RollupConstMaxL1UserTx)
	// _, err = tc.GenerateBlocks(set)
	// require.Equal(t, "line 2: DeleteVouch, err: unexpected PoolL2 tx type: DeleteVouch", err.Error())

	// set = `
	// 	Type: PoolL2
	// 	PoolCreateVouch A
	// 	> batch
	// `
	// tc = NewContext(0, common.RollupConstMaxL1UserTx)
	// _, err = tc.GenerateBlocks(set)
	// require.Equal(t, "line 2: PoolCreateVouchA> batch\n, err: expected '-', found '>'", err.Error())

	// set = `
	// 	Type: PoolL2
	// 	PoolDeleteVouch A-
	// `
	// tc = NewContext(0, common.RollupConstMaxL1UserTx)
	// _, err = tc.GenerateBlocks(set)
	// require.Equal(t, "line 2: PoolDeleteVouchA-, err: expected 'to' account name, found ''", err.Error())
}

func TestGenerateBlocksFromInstructions(t *testing.T) {
	// Generate block from instructions
	setInst := []Instruction{}

	i := 0
	da := big.NewInt(10)
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "CreateAccountDeposit A: 10",
		Typ:           common.TxTypeCreateAccountDeposit,
		From:          "A",
		DepositAmount: da,
	})

	i++
	da = big.NewInt(10)
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "CreateAccountDeposit B: 10",
		Typ:           common.TxTypeCreateAccountDeposit,
		From:          "B",
		DepositAmount: da,
	})

	i++
	da = big.NewInt(6)
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "Deposit A: 6",
		Typ:           common.TxTypeDeposit,
		From:          "A",
		DepositAmount: da,
	})

	i++
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "CreateVouch A-B",
		Typ:  common.TxTypeCreateVouch,
		From: "A",
		To:   "B",
	})

	i++
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "CreateVouch B-A",
		Typ:  common.TxTypeCreateVouch,
		From: "B",
		To:   "A",
	})

	i++
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "> batchL1",
		Typ: TypeNewBatchL1,
	})

	// i++
	// a := big.NewInt(3)
	// setInst = append(setInst, Instruction{
	// 	LineNum: i,
	// 	// Literal: "Exit A: 3",
	// 	Typ:    common.TxTypeExit,
	// 	From:   "A",
	// 	Amount: a,
	// 	Fee:    1,
	// })

	i++
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "DeleteVouch A-B",
		Typ:  common.TxTypeDeleteVouch,
		From: "A",
		To:   "B",
	})

	i++
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "> batch",
		Typ: TypeNewBatchL1,
	})

	i++
	setInst = append(setInst, Instruction{
		LineNum: i,
		// Literal: "> block",
		Typ: TypeNewBlock,
	})

	tc := NewContext(0, common.RollupConstMaxL1UserTx)
	blockFromInstructions, err := tc.GenerateBlocksFromInstructions(setInst)
	require.NoError(t, err)

	// Generate block from string
	setString := `
		Type: Blockchain
		CreateAccountDeposit A: 10
		CreateAccountDeposit B: 10
		Deposit A: 6
		CreateVouch A-B
		CreateVouch B-A
		> batchL1
		// Exit A: 3
		DeleteVouch A-B
		> batchL1
		> block
	`
	tc = NewContext(0, common.RollupConstMaxL1UserTx)
	blockFromString, err := tc.GenerateBlocks(setString)
	require.NoError(t, err)

	// Generated data should be equivalent, except for Eth Addrs and BJJs
	for i, strBatch := range blockFromString[0].Rollup.Batches {
		// instBatch := blockFromInstructions[0].Rollup.Batches[i]
		// for j := 0; j < len(strBatch.L1CoordinatorTxs); j++ {
		// 	blockFromInstructions[0].Rollup.Batches[i].L1CoordinatorTxs[j].FromEthAddr =
		// 		blockFromString[0].Rollup.Batches[i].L1CoordinatorTxs[j].FromEthAddr
		// 	blockFromInstructions[0].Rollup.Batches[i].L1CoordinatorTxs[j].FromBJJ =
		// 		blockFromString[0].Rollup.Batches[i].L1CoordinatorTxs[j].FromBJJ
		// }
		for j := 0; j < len(strBatch.L1UserTxs); j++ {
			blockFromInstructions[0].Rollup.Batches[i].L1UserTxs[j].FromEthAddr =
				blockFromString[0].Rollup.Batches[i].L1UserTxs[j].FromEthAddr
			blockFromInstructions[0].Rollup.Batches[i].L1UserTxs[j].FromBJJ =
				blockFromString[0].Rollup.Batches[i].L1UserTxs[j].FromBJJ
		}
	}
	for i := 0; i < len(blockFromString[0].Rollup.L1UserTxs); i++ {
		blockFromInstructions[0].Rollup.L1UserTxs[i].FromEthAddr =
			blockFromString[0].Rollup.L1UserTxs[i].FromEthAddr
		blockFromInstructions[0].Rollup.L1UserTxs[i].FromBJJ =
			blockFromString[0].Rollup.L1UserTxs[i].FromBJJ
	}
	assert.Equal(t, blockFromString, blockFromInstructions)
}
