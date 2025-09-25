package til

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBlockchainTxs(t *testing.T) {
	s := `
		Type: Blockchain

		Deposit A: 10
		Deposit A: 20
		Deposit B: 5
		CreateVouch A-B

		// set new batch
		> batchL1

		CreateAccountDeposit C: 5
		CreateVouch B-C
		Deposit User0: 20
		Deposit User1: 20

		> batchL1
		> block

		DeleteVouch A-B
	`

	parser := newParser(strings.NewReader(s))
	instructions, err := parser.parse()
	require.NoError(t, err)
	assert.Equal(t, 12, len(instructions.instructions))
	assert.Equal(t, 5, len(instructions.users))

	if Debug {
		for _, instruction := range instructions.instructions {
			fmt.Println(instruction.raw())
		}
	}

	assert.Equal(t, TypeNewBatchL1, instructions.instructions[4].Typ)
	assert.Equal(t, "DepositUser0:20", instructions.instructions[7].raw())
	assert.Equal(t, "CreateVouchA-B", instructions.instructions[3].raw())
	assert.Equal(t, "DeleteVouchA-B", instructions.instructions[11].raw())
}

func TestParseErrors(t *testing.T) {
	s := `
		Type: Blockchain
		Deposit A:: 10
	`
	parser := newParser(strings.NewReader(s))
	_, err := parser.parse()
	assert.Equal(t, "line 2: DepositA:: 10\n, err: can not parse number for Amount: :", err.Error())

	s = `
		Type: Blockchain
		Deposit A: 10 20
	`
	parser = newParser(strings.NewReader(s))
	_, err = parser.parse()
	assert.Equal(t, "line 3: 20, err: unexpected Blockchain tx type: 20", err.Error())

	s = `
		Type: Blockchain
		> btch
	`
	parser = newParser(strings.NewReader(s))
	_, err = parser.parse()
	assert.Equal(t,
		"line 2: >, err: unexpected '> btch', expected '> batch' or '> block'",
		err.Error())

	// check definition of set Type
	s = `PoolExit A: 10`
	parser = newParser(strings.NewReader(s))
	_, err = parser.parse()
	assert.Equal(t, "line 1: PoolExit, err: set type not defined", err.Error())

	s = `Type: PoolL1`
	parser = newParser(strings.NewReader(s))
	_, err = parser.parse()
	assert.Equal(t,
		"line 1: Type:, err: invalid set type: 'PoolL1'. Valid set types: 'Blockchain', 'PoolL2'",
		err.Error())

	s = `Type: PoolL1
		Type: Blockchain`
	parser = newParser(strings.NewReader(s))
	_, err = parser.parse()
	assert.Equal(t,
		"line 1: Type:, err: invalid set type: 'PoolL1'. Valid set types: 'Blockchain', 'PoolL2'",
		err.Error())
}
