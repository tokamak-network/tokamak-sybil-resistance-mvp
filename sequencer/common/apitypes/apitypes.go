package apitypes

import (
	"math/big"
	"tokamak-sybil-resistance/common"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

type BigIntStr string

// CollectedFeesAPI is send common.batch.CollectedFee through the API
type CollectedFeesAPI map[common.TokenID]BigIntStr

// NewBigIntStr creates a *BigIntStr from a *big.Int.
// If the provided bigInt is nil the returned *BigIntStr will also be nil
func NewBigIntStr(bigInt *big.Int) *BigIntStr {
	if bigInt == nil {
		return nil
	}
	bigIntStr := BigIntStr(bigInt.String())
	return &bigIntStr
}

// TonEthAddr is used to scan/value Ethereum Address directly into strings that follow the Ethereum address ton format (^ton:0x[a-fA-F0-9]{40}$) from/to sql DBs.
// It assumes that Ethereum Address are inserted/fetched to/from the DB using the default Scan/Value interface
type TonEthAddr string

// NewTonEthAddr creates a TonEthAddr from an Ethereum addr
func NewTonEthAddr(addr ethCommon.Address) TonEthAddr {
	return TonEthAddr("ton:" + addr.String())
}

// TonIdx is used to value common.Idx directly into strings that follow the Idx key ton format (ton:tokenSymbol:idx) to sql DBs.
// Note that this can only be used to insert to DB since there is no way to automatically read from the DB since it needs the tokenSymbol
type TonIdx string
