package api

import (
	"math/big"
	"tokamak-sybil-resistance/common"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

type rollupConstants struct {
	PublicConstants       common.RollupConstants `json:"publicConstants"`
	ReservedIdx           int                    `json:"reservedIdx"`
	LimitDepositAmount    *big.Int               `json:"limitDepositAmount"`
	L1UserTotalBytes      int                    `json:"l1UserTotalBytes"`
	MaxL1UserTx           int                    `json:"maxL1UserTx"`
	MaxL1Tx               int                    `json:"maxL1Tx"`
	InputSHAConstantBytes int                    `json:"inputSHAConstantBytes"`
	ExchangeMultiplier    int                    `json:"exchangeMultiplier"`
}

type configAPI struct {
	ChainID         uint64          `json:"chainId"`
	RollupConstants rollupConstants `json:"hermez"`
}

func newRollupConstants(publicConstants common.RollupConstants) *rollupConstants {
	return &rollupConstants{
		PublicConstants:       publicConstants,
		ReservedIdx:           common.RollupConstReservedIDx,
		LimitDepositAmount:    common.RollupConstLimitDepositAmount,
		L1UserTotalBytes:      common.RollupConstL1UserTotalBytes,
		MaxL1UserTx:           common.RollupConstMaxL1UserTx,
		MaxL1Tx:               common.RollupConstMaxL1Tx,
		InputSHAConstantBytes: common.RollupConstInputSHAConstantBytes,
		ExchangeMultiplier:    common.RollupConstExchangeMultiplier,
	}
}

// NetworkConfig of the API
type NetworkConfig struct {
	RollupConstants common.RollupConstants
	ChainID         uint64
	TonAddress      ethCommon.Address
}
