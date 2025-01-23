package api

import (
	"tokamak-sybil-resistance/common"

	ethCommon "github.com/ethereum/go-ethereum/common"
)

func getConfigTest(chainID uint64) NetworkConfig {
	var config NetworkConfig

	var rollupPublicConstants common.RollupConstants
	rollupPublicConstants.AbsoluteMaxL1BatchTimeout = 240
	var verifier common.RollupVerifierStruct

	verifier.MaxTx = 512
	verifier.NLevels = 32
	rollupPublicConstants.Verifiers = append(rollupPublicConstants.Verifiers, verifier)

	config.RollupConstants = rollupPublicConstants

	config.ChainID = chainID
	config.TonAddress = ethCommon.HexToAddress("0xc344E203a046Da13b0B4467EB7B3629D0C99F6E6")

	return config
}
