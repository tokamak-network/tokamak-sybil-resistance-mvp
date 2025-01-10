// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sybil

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SybilMetaData contains all meta data concerning the Sybil contract.
var SybilMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"_addTx\",\"inputs\":[{\"name\":\"ethAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"loadAmountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"amountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"toIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_hashFinalNode\",\"inputs\":[{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_hashNode\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAccountDeposit\",\"inputs\":[{\"name\":\"loadAmountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"currentFillingBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"loadAmountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"exit\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"amountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitNullifierMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exitRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explodeMultiple\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"toIdxs\",\"type\":\"uint48[]\",\"internalType\":\"uint48[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forgeBatch\",\"inputs\":[{\"name\":\"newLastIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"newAccountRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newVouchRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newScoreRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newExitRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofA\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"proofB\",\"type\":\"uint256[2][2]\",\"internalType\":\"uint256[2][2]\"},{\"name\":\"proofC\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getL1TransactionQueue\",\"inputs\":[{\"name\":\"queueIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastForgedBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueueLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateRoot\",\"inputs\":[{\"name\":\"batchNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxTx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nLevel\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_poseidon2Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon3Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon4Elements\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastForgedBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastIdx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollupVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"verifierInterface\",\"type\":\"address\",\"internalType\":\"contractVerifierRollupInterface\"},{\"name\":\"maxTx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nLevel\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scoreRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unprocessedBatchesMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unvouch\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"toIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vouch\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"toIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vouchRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawMerkleProof\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"numExitRoot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"siblings\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"idx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ForgeBatch\",\"inputs\":[{\"name\":\"batchNum\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"l1UserTxsLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1UserTxEvent\",\"inputs\":[{\"name\":\"queueIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"position\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"l1UserTx\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawEvent\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint48\",\"indexed\":true,\"internalType\":\"uint48\"},{\"name\":\"numExitRoot\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EthTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFromIdx\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPoseidonAddress\",\"inputs\":[{\"name\":\"elementType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidToIdx\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LoadAmountDoesNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LoadAmountExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SmtProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawAlreadyDone\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506123e88061001d5f395ff3fe6080604052600436106101ba575f3560e01c8063adacd33b116100f2578063c1b190c011610092578063ef8140b511610062578063ef8140b51461058c578063f2fde38b146105ab578063f84f92ee146105ca578063fbb4a00f14610613575f80fd5b8063c1b190c0146104c9578063c25d5789146104e8578063d486645c14610504578063e8bf92ed14610540575f80fd5b8063ba2506df116100cd578063ba2506df14610434578063bbe5a37514610460578063bd8a4a611461047f578063c0b55ae4146104aa575f80fd5b8063adacd33b146103d6578063b1f073d414610401578063b8f7700514610420575f80fd5b806362332ee21161015d5780638195b790116101385780638195b79014610327578063894bc2b8146103465780638da5cb5b14610365578063a5e2ec5b146103ab575f80fd5b806362332ee2146102c9578063715018a6146102e8578063795c6167146102fc575f80fd5b8063212bafd711610198578063212bafd7146102355780632f463f59146102485780633009c59f1461026757806344e0b2ce146102a6575f80fd5b806311917b1d146101be57806311954d3c146101df5780631b78164b146101fe575b5f80fd5b3480156101c9575f80fd5b506101dd6101d83660046119dc565b610626565b005b3480156101ea575f80fd5b506101dd6101f9366004611a53565b61076e565b348015610209575f80fd5b50600254600160d01b900463ffffffff165b60405163ffffffff90911681526020015b60405180910390f35b6101dd610243366004611a98565b610814565b348015610253575f80fd5b506101dd610262366004611ac0565b6108c0565b348015610272575f80fd5b50610298610281366004611b34565b63ffffffff165f9081526004602052604090205490565b60405190815260200161022c565b3480156102b1575f80fd5b5060025461021b90600160d01b900463ffffffff1681565b3480156102d4575f80fd5b506101dd6102e3366004611b54565b610a53565b3480156102f3575f80fd5b506101dd610b70565b348015610307575f80fd5b50610298610316366004611b34565b60046020525f908152604090205481565b348015610332575f80fd5b506101dd610341366004611a98565b610b83565b348015610351575f80fd5b506101dd610360366004611c09565b610c0b565b348015610370575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b03909116815260200161022c565b3480156103b6575f80fd5b506102986103c5366004611b34565b60076020525f908152604090205481565b3480156103e1575f80fd5b506102986103f0366004611b34565b60056020525f908152604090205481565b34801561040c575f80fd5b506101dd61041b366004611cec565b610d3c565b34801561042b575f80fd5b5061021b610f00565b34801561043f575f80fd5b5061045361044e366004611b34565b610f27565b60405161022c9190611d8d565b34801561046b575f80fd5b5061029861047a366004611dbf565b610fcd565b34801561048a575f80fd5b50610298610499366004611b34565b60066020525f908152604090205481565b3480156104b5575f80fd5b506102986104c4366004611dbf565b610ffa565b3480156104d4575f80fd5b506101dd6104e3366004611a53565b611016565b3480156104f3575f80fd5b5060035461021b9063ffffffff1681565b34801561050f575f80fd5b5060025461052990600160a01b900465ffffffffffff1681565b60405165ffffffffffff909116815260200161022c565b34801561054b575f80fd5b50600a54600b54600c54610567926001600160a01b0316919083565b604080516001600160a01b03909416845260208401929092529082015260600161022c565b348015610597575f80fd5b506104536105a6366004611b34565b6110b7565b3480156105b6575f80fd5b506101dd6105c5366004611ddf565b61114e565b3480156105d5575f80fd5b506106036105e4366004611df8565b600960209081525f928352604080842090915290825290205460ff1681565b604051901515815260200161022c565b6101dd610621366004611e12565b611190565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f8115801561066b5750825b90505f8267ffffffffffffffff1660011480156106875750303b155b905081158015610695575080155b156106b35760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106dd57845460ff60401b1916600160401b1785555b6002805465ffffffffffff60a01b191660ff60a01b1790556003805463ffffffff191660011790556107108b8b8b6111ed565b61071b888888611254565b831561076157845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60ff65ffffffffffff831611158061079a575060025465ffffffffffff600160a01b9091048116908316115b156107b85760405163618d793560e01b815260040160405180910390fd5b60ff65ffffffffffff82161115806107e4575060025465ffffffffffff600160a01b9091048116908216115b156108025760405163e9eabe2f60e01b815260040160405180910390fd5b61081033835f6001856108c0565b5050565b5f61081e82611373565b9050600160801b811061084457604051630167494760e71b815260040160405180910390fd5b34811461086457604051635ce5543160e11b815260040160405180910390fd5b60ff65ffffffffffff8416111580610890575060025465ffffffffffff600160a01b9091048116908416115b156108ae5760405163618d793560e01b815260040160405180910390fd5b6108bb3384845f806108c0565b505050565b6040516bffffffffffffffffffffffff19606087901b1660208201526001600160d01b031960d086811b821660348401526001600160d81b031960d887811b8216603a86015286901b16603f84015283901b1660448201525f90604a0160408051601f1981840301815291815260035463ffffffff165f9081526008602052908120805492935090916080919061095690611e2b565b610961929150611e8b565b60035463ffffffff165f90815260086020908152604091829020915192935061098c92859101611e9e565b60408051601f1981840301815291815260035463ffffffff165f908152600860205220906109ba9082611f6c565b5060035460405160ff83169163ffffffff16907fdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119906109fa908690611d8d565b60405180910390a36103e8610a1082600161202c565b10610a4a576003805463ffffffff16905f610a2a8361203f565b91906101000a81548163ffffffff021916908363ffffffff160217905550505b50505050505050565b5f610a5e86336113a7565b90505f610a6a826113e4565b63ffffffff87165f908152600760209081526040808320546009835281842065ffffffffffff89168552909252909120549192509060ff1615610ac05760405163029ec81f60e31b815260040160405180910390fd5b610ad58187878765ffffffffffff1686611453565b610af157604051627e055560e11b815260040160405180910390fd5b63ffffffff87165f90815260096020908152604080832065ffffffffffff881684529091529020805460ff19166001179055610b2c886114df565b60405163ffffffff88169065ffffffffffff8616907f102db758451b2f65238246a452d00c0c4c8f59d8c623aff254111079418e57ec905f90a35050505050505050565b610b786114f1565b610b815f61154c565b565b5f610b8d82611373565b9050600160801b8110610bb35760405163172bd6a160e31b815260040160405180910390fd5b60ff65ffffffffffff8416111580610bdf575060025465ffffffffffff600160a01b9091048116908416115b15610bfd5760405163618d793560e01b815260040160405180910390fd5b6108bb33845f8560016108c0565b60ff65ffffffffffff8316111580610c37575060025465ffffffffffff600160a01b9091048116908316115b15610c555760405163618d793560e01b815260040160405180910390fd5b80515f5b81811015610cfc5760ff65ffffffffffff16838281518110610c7d57610c7d612061565b602002602001015165ffffffffffff16111580610cd65750600260149054906101000a900465ffffffffffff1665ffffffffffff16838281518110610cc457610cc4612061565b602002602001015165ffffffffffff16115b15610cf45760405163e9eabe2f60e01b815260040160405180910390fd5b600101610c59565b505f5b81811015610d3657610d2e33855f6002878681518110610d2157610d21612061565b60200260200101516108c0565b600101610cff565b50505050565b5f610d4a89898989896115bc565b600a546040805160208101825283815290516343753b4d60e01b81529293506001600160a01b03909116916343753b4d91610d8e9188918891889190600401612075565b602060405180830381865afa158015610da9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dcd91906120e1565b610dea576040516309bde33960e01b815260040160405180910390fd5b60028054600160d01b900463ffffffff1690601a610e078361203f565b82546101009290920a63ffffffff8181021990931691831602179091556002805465ffffffffffff60a01b1916600160a01b65ffffffffffff8e160217808255600160d01b9081900483165f9081526004602090815260408083208f9055845484900486168352600582528083208e9055845484900486168352600682528083208d905593549290920490931683526007905281208790559050610ea9611778565b60025460405161ffff83168152919250600160d01b900463ffffffff16907fe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d345886189060200160405180910390a250505050505050505050565b6002546003545f91610f229163ffffffff600160d01b90920482169116612100565b905090565b63ffffffff81165f908152600860205260409020805460609190610f4a90611e2b565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7690611e2b565b8015610fc15780601f10610f9857610100808354040283529160200191610fc1565b820191905f5260205f20905b815481529060010190602001808311610fa457829003601f168201915b50505050509050919050565b5f610fd6611922565b8381526020810183905260016040820152610ff081611841565b9150505b92915050565b5f611003611940565b83815260208101839052610ff081611871565b60ff65ffffffffffff8316111580611042575060025465ffffffffffff600160a01b9091048116908316115b156110605760405163618d793560e01b815260040160405180910390fd5b60ff65ffffffffffff821611158061108c575060025465ffffffffffff600160a01b9091048116908216115b156110aa5760405163e9eabe2f60e01b815260040160405180910390fd5b61081033835f80856108c0565b60086020525f9081526040902080546110cf90611e2b565b80601f01602080910402602001604051908101604052809291908181526020018280546110fb90611e2b565b80156111465780601f1061111d57610100808354040283529160200191611146565b820191905f5260205f20905b81548152906001019060200180831161112957829003601f168201915b505050505081565b6111566114f1565b6001600160a01b03811661118457604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61118d8161154c565b50565b5f61119a82611373565b9050600160801b81106111c057604051630167494760e71b815260040160405180910390fd5b3481146111e057604051635ce5543160e11b815260040160405180910390fd5b610810335f845f806108c0565b6001600160a01b0383166112145760405163043103a360e21b815260040160405180910390fd5b604080516060810182526001600160a01b03909416808552602085018490529301819052600a80546001600160a01b031916909317909255600b55600c55565b6001600160a01b03831661129f57604051631853e3f560e11b8152602060048201526011602482015270706f736569646f6e32456c656d656e747360781b604482015260640161117b565b6001600160a01b0382166112ea57604051631853e3f560e11b8152602060048201526011602482015270706f736569646f6e33456c656d656e747360781b604482015260640161117b565b6001600160a01b03811661133557604051631853e3f560e11b8152602060048201526011602482015270706f736569646f6e34456c656d656e747360781b604482015260640161117b565b5f80546001600160a01b039485166001600160a01b031991821617909155600180549385169382169390931790925560028054919093169116179055565b5f6407ffffffff8216601f602384901c168261139082600a612204565b90505f61139d828561220f565b9695505050505050565b6113af61195e565b6113b761195e565b6001600160c01b03939093168352506001600160a01b031660208201525f60408201819052606082015290565b60025460405163248f667760e01b81525f916001600160a01b03169063248f667790611414908590600401612226565b602060405180830381865afa15801561142f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ff49190612256565b5f8061145f8484610fcd565b90505f8061146e60018861226d565b90505b5f81126114d15787878281811061148a5761148a612061565b6020029190910135925050600186821c811614806114b1576114ac8484610ffa565b6114bb565b6114bb8385610ffa565b93505080806114c99061228c565b915050611471565b505090951495945050505050565b61118d816001600160c01b03166118a1565b336115237f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610b815760405163118cdaa760e01b815233600482015260240161117b565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b600254600160d01b810463ffffffff165f81815260046020908152604080832054600583528184205460069093529083205492949093919291600160a01b90910465ffffffffffff1690859060089082906116189060016122a7565b63ffffffff1663ffffffff1681526020019081526020015f20805461163c90611e2b565b80601f016020809104026020016040519081016040528092919081815260200182805461166890611e2b565b80156116b35780601f1061168a576101008083540402835291602001916116b3565b820191905f5260205f20905b81548152906001019060200180831161169657829003601f168201915b505050505090505f828686868f8f8f8f8f8a6040516020016116de9a999897969594939291906122c4565b60405160208183030381529060405290507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002826040516117209190612336565b602060405180830381855afa15801561173b573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061175e9190612256565b6117689190612351565b9c9b505050505050505050505050565b600254600160d01b900463ffffffff165f908152600860205260408120805482916080916117a590611e2b565b6117b0929150611e8b565b600254600160d01b900463ffffffff165f9081526008602052604081209192506117da919061197c565b60035460025463ffffffff918216916117fc91600160d01b90041660016122a7565b63ffffffff160361183c576003805463ffffffff16905f61181c8361203f565b91906101000a81548163ffffffff021916908363ffffffff160217905550505b919050565b6001546040516304b98e1d60e31b81525f916001600160a01b0316906325cc70e890611414908590600401612364565b5f80546040516314d2f97b60e11b81526001600160a01b03909116906329a5f2f69061141490859060040161238b565b604080515f80825260208201909252339083906040516118c19190612336565b5f6040518083038185875af1925050503d805f81146118fb576040519150601f19603f3d011682016040523d82523d5f602084013e611900565b606091505b505090508061081057604051630db2c7f160e31b815260040160405180910390fd5b60405180606001604052806003906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b50805461198890611e2b565b5f825580601f10611997575050565b601f0160209004905f5260205f209081019061118d91905b808211156119c2575f81556001016119af565b5090565b80356001600160a01b038116811461183c575f80fd5b5f805f805f8060c087890312156119f1575f80fd5b6119fa876119c6565b95506020870135945060408701359350611a16606088016119c6565b9250611a24608088016119c6565b9150611a3260a088016119c6565b90509295509295509295565b803565ffffffffffff8116811461183c575f80fd5b5f8060408385031215611a64575f80fd5b611a6d83611a3e565b9150611a7b60208401611a3e565b90509250929050565b803564ffffffffff8116811461183c575f80fd5b5f8060408385031215611aa9575f80fd5b611ab283611a3e565b9150611a7b60208401611a84565b5f805f805f60a08688031215611ad4575f80fd5b611add866119c6565b9450611aeb60208701611a3e565b9350611af960408701611a84565b9250611b0760608701611a84565b9150611b1560808701611a3e565b90509295509295909350565b803563ffffffff8116811461183c575f80fd5b5f60208284031215611b44575f80fd5b611b4d82611b21565b9392505050565b5f805f805f60808688031215611b68575f80fd5b85356001600160c01b0381168114611b7e575f80fd5b9450611b8c60208701611b21565b9350604086013567ffffffffffffffff80821115611ba8575f80fd5b818801915088601f830112611bbb575f80fd5b813581811115611bc9575f80fd5b8960208260051b8501011115611bdd575f80fd5b602083019550809450505050611b1560608701611a3e565b634e487b7160e01b5f52604160045260245ffd5b5f8060408385031215611c1a575f80fd5b611c2383611a3e565b915060208084013567ffffffffffffffff80821115611c40575f80fd5b818601915086601f830112611c53575f80fd5b813581811115611c6557611c65611bf5565b8060051b604051601f19603f83011681018181108582111715611c8a57611c8a611bf5565b604052918252848201925083810185019189831115611ca7575f80fd5b938501935b82851015611ccc57611cbd85611a3e565b84529385019392850192611cac565b8096505050505050509250929050565b8060408101831015610ff4575f80fd5b5f805f805f805f806101a0898b031215611d04575f80fd5b611d0d89611a3e565b975060208901359650604089013595506060890135945060808901359350611d388a60a08b01611cdc565b925061016089018a811115611d4b575f80fd5b60e08a019250611d5b8b82611cdc565b9150509295985092959890939650565b5f5b83811015611d85578181015183820152602001611d6d565b50505f910152565b602081525f8251806020840152611dab816040850160208701611d6b565b601f01601f19169190910160400192915050565b5f8060408385031215611dd0575f80fd5b50508035926020909101359150565b5f60208284031215611def575f80fd5b611b4d826119c6565b5f8060408385031215611e09575f80fd5b611a6d83611b21565b5f60208284031215611e22575f80fd5b611b4d82611a84565b600181811c90821680611e3f57607f821691505b602082108103611e5d57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f82611e9957611e99611e63565b500490565b5f808454611eab81611e2b565b60018281168015611ec35760018114611ed857611f04565b60ff1984168752821515830287019450611f04565b885f526020805f205f5b85811015611efb5781548a820152908401908201611ee2565b50505082870194505b505050508351611f18818360208801611d6b565b01949350505050565b601f8211156108bb57805f5260205f20601f840160051c81016020851015611f465750805b601f840160051c820191505b81811015611f65575f8155600101611f52565b5050505050565b815167ffffffffffffffff811115611f8657611f86611bf5565b611f9a81611f948454611e2b565b84611f21565b602080601f831160018114611fcd575f8415611fb65750858301515b5f19600386901b1c1916600185901b178555612024565b5f85815260208120601f198616915b82811015611ffb57888601518255948401946001909101908401611fdc565b508582101561201857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b80820180821115610ff457610ff4611e77565b5f63ffffffff80831681810361205757612057611e77565b6001019392505050565b634e487b7160e01b5f52603260045260245ffd5b6101208101604080878437808301865f5b60028110156120a357838284379183019190830190600101612086565b505050808560c0850137506101008201835f5b60018110156120d55781518352602092830192909101906001016120b6565b50505095945050505050565b5f602082840312156120f1575f80fd5b81518015158114611b4d575f80fd5b63ffffffff82811682821603908082111561211d5761211d611e77565b5092915050565b600181815b8085111561215e57815f190482111561214457612144611e77565b8085161561215157918102915b93841c9390800290612129565b509250929050565b5f8261217457506001610ff4565b8161218057505f610ff4565b816001811461219657600281146121a0576121bc565b6001915050610ff4565b60ff8411156121b1576121b1611e77565b50506001821b610ff4565b5060208310610133831016604e8410600b84101617156121df575081810a610ff4565b6121e98383612124565b805f19048211156121fc576121fc611e77565b029392505050565b5f611b4d8383612166565b8082028115828204841417610ff457610ff4611e77565b6080810181835f5b600481101561224d57815183526020928301929091019060010161222e565b50505092915050565b5f60208284031215612266575f80fd5b5051919050565b8181035f83128015838313168383128216171561211d5761211d611e77565b5f600160ff1b82016122a0576122a0611e77565b505f190190565b63ffffffff81811683821601908082111561211d5761211d611e77565b5f65ffffffffffff60d01b808d60d01b1683528b60068401528a6026840152896046840152808960d01b1660668401525086606c83015285608c8301528460ac8301528360cc83015282516123208160ec850160208701611d6b565b9190910160ec019b9a5050505050505050505050565b5f8251612347818460208701611d6b565b9190910192915050565b5f8261235f5761235f611e63565b500690565b6060810181835f5b600381101561224d57815183526020928301929091019060010161236c565b6040810181835f5b600281101561224d57815183526020928301929091019060010161239356fea26469706673582212203d6de0bc237a0ca0e6408dbe8aa5356d2c81cac7a06b76a0a39be206f014336764736f6c63430008170033",
}

// SybilABI is the input ABI used to generate the binding from.
// Deprecated: Use SybilMetaData.ABI instead.
var SybilABI = SybilMetaData.ABI

// SybilBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SybilMetaData.Bin instead.
var SybilBin = SybilMetaData.Bin

// DeploySybil deploys a new Ethereum contract, binding an instance of Sybil to it.
func DeploySybil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sybil, error) {
	parsed, err := SybilMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SybilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sybil{SybilCaller: SybilCaller{contract: contract}, SybilTransactor: SybilTransactor{contract: contract}, SybilFilterer: SybilFilterer{contract: contract}}, nil
}

// Sybil is an auto generated Go binding around an Ethereum contract.
type Sybil struct {
	SybilCaller     // Read-only binding to the contract
	SybilTransactor // Write-only binding to the contract
	SybilFilterer   // Log filterer for contract events
}

// SybilCaller is an auto generated read-only Go binding around an Ethereum contract.
type SybilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SybilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SybilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SybilSession struct {
	Contract     *Sybil            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SybilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SybilCallerSession struct {
	Contract *SybilCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SybilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SybilTransactorSession struct {
	Contract     *SybilTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SybilRaw is an auto generated low-level Go binding around an Ethereum contract.
type SybilRaw struct {
	Contract *Sybil // Generic contract binding to access the raw methods on
}

// SybilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SybilCallerRaw struct {
	Contract *SybilCaller // Generic read-only contract binding to access the raw methods on
}

// SybilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SybilTransactorRaw struct {
	Contract *SybilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSybil creates a new instance of Sybil, bound to a specific deployed contract.
func NewSybil(address common.Address, backend bind.ContractBackend) (*Sybil, error) {
	contract, err := bindSybil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sybil{SybilCaller: SybilCaller{contract: contract}, SybilTransactor: SybilTransactor{contract: contract}, SybilFilterer: SybilFilterer{contract: contract}}, nil
}

// NewSybilCaller creates a new read-only instance of Sybil, bound to a specific deployed contract.
func NewSybilCaller(address common.Address, caller bind.ContractCaller) (*SybilCaller, error) {
	contract, err := bindSybil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SybilCaller{contract: contract}, nil
}

// NewSybilTransactor creates a new write-only instance of Sybil, bound to a specific deployed contract.
func NewSybilTransactor(address common.Address, transactor bind.ContractTransactor) (*SybilTransactor, error) {
	contract, err := bindSybil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SybilTransactor{contract: contract}, nil
}

// NewSybilFilterer creates a new log filterer instance of Sybil, bound to a specific deployed contract.
func NewSybilFilterer(address common.Address, filterer bind.ContractFilterer) (*SybilFilterer, error) {
	contract, err := bindSybil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SybilFilterer{contract: contract}, nil
}

// bindSybil binds a generic wrapper to an already deployed contract.
func bindSybil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SybilMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sybil *SybilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sybil.Contract.SybilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sybil *SybilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.Contract.SybilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sybil *SybilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sybil.Contract.SybilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sybil *SybilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sybil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sybil *SybilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sybil *SybilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sybil.Contract.contract.Transact(opts, method, params...)
}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Sybil *SybilCaller) HashFinalNode(opts *bind.CallOpts, key *big.Int, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "_hashFinalNode", key, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Sybil *SybilSession) HashFinalNode(key *big.Int, value *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashFinalNode(&_Sybil.CallOpts, key, value)
}

// HashFinalNode is a free data retrieval call binding the contract method 0xbbe5a375.
//
// Solidity: function _hashFinalNode(uint256 key, uint256 value) view returns(uint256)
func (_Sybil *SybilCallerSession) HashFinalNode(key *big.Int, value *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashFinalNode(&_Sybil.CallOpts, key, value)
}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Sybil *SybilCaller) HashNode(opts *bind.CallOpts, left *big.Int, right *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "_hashNode", left, right)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Sybil *SybilSession) HashNode(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashNode(&_Sybil.CallOpts, left, right)
}

// HashNode is a free data retrieval call binding the contract method 0xc0b55ae4.
//
// Solidity: function _hashNode(uint256 left, uint256 right) view returns(uint256)
func (_Sybil *SybilCallerSession) HashNode(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Sybil.Contract.HashNode(&_Sybil.CallOpts, left, right)
}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) AccountRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "accountRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) AccountRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.AccountRootMap(&_Sybil.CallOpts, arg0)
}

// AccountRootMap is a free data retrieval call binding the contract method 0x795c6167.
//
// Solidity: function accountRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) AccountRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.AccountRootMap(&_Sybil.CallOpts, arg0)
}

// CurrentFillingBatch is a free data retrieval call binding the contract method 0xc25d5789.
//
// Solidity: function currentFillingBatch() view returns(uint32)
func (_Sybil *SybilCaller) CurrentFillingBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "currentFillingBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentFillingBatch is a free data retrieval call binding the contract method 0xc25d5789.
//
// Solidity: function currentFillingBatch() view returns(uint32)
func (_Sybil *SybilSession) CurrentFillingBatch() (uint32, error) {
	return _Sybil.Contract.CurrentFillingBatch(&_Sybil.CallOpts)
}

// CurrentFillingBatch is a free data retrieval call binding the contract method 0xc25d5789.
//
// Solidity: function currentFillingBatch() view returns(uint32)
func (_Sybil *SybilCallerSession) CurrentFillingBatch() (uint32, error) {
	return _Sybil.Contract.CurrentFillingBatch(&_Sybil.CallOpts)
}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Sybil *SybilCaller) ExitNullifierMap(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "exitNullifierMap", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Sybil *SybilSession) ExitNullifierMap(arg0 uint32, arg1 *big.Int) (bool, error) {
	return _Sybil.Contract.ExitNullifierMap(&_Sybil.CallOpts, arg0, arg1)
}

// ExitNullifierMap is a free data retrieval call binding the contract method 0xf84f92ee.
//
// Solidity: function exitNullifierMap(uint32 , uint48 ) view returns(bool)
func (_Sybil *SybilCallerSession) ExitNullifierMap(arg0 uint32, arg1 *big.Int) (bool, error) {
	return _Sybil.Contract.ExitNullifierMap(&_Sybil.CallOpts, arg0, arg1)
}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) ExitRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "exitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) ExitRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ExitRootMap(&_Sybil.CallOpts, arg0)
}

// ExitRootMap is a free data retrieval call binding the contract method 0xa5e2ec5b.
//
// Solidity: function exitRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) ExitRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ExitRootMap(&_Sybil.CallOpts, arg0)
}

// GetL1TransactionQueue is a free data retrieval call binding the contract method 0xba2506df.
//
// Solidity: function getL1TransactionQueue(uint32 queueIndex) view returns(bytes)
func (_Sybil *SybilCaller) GetL1TransactionQueue(opts *bind.CallOpts, queueIndex uint32) ([]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "getL1TransactionQueue", queueIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetL1TransactionQueue is a free data retrieval call binding the contract method 0xba2506df.
//
// Solidity: function getL1TransactionQueue(uint32 queueIndex) view returns(bytes)
func (_Sybil *SybilSession) GetL1TransactionQueue(queueIndex uint32) ([]byte, error) {
	return _Sybil.Contract.GetL1TransactionQueue(&_Sybil.CallOpts, queueIndex)
}

// GetL1TransactionQueue is a free data retrieval call binding the contract method 0xba2506df.
//
// Solidity: function getL1TransactionQueue(uint32 queueIndex) view returns(bytes)
func (_Sybil *SybilCallerSession) GetL1TransactionQueue(queueIndex uint32) ([]byte, error) {
	return _Sybil.Contract.GetL1TransactionQueue(&_Sybil.CallOpts, queueIndex)
}

// GetLastForgedBatch is a free data retrieval call binding the contract method 0x1b78164b.
//
// Solidity: function getLastForgedBatch() view returns(uint32)
func (_Sybil *SybilCaller) GetLastForgedBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "getLastForgedBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLastForgedBatch is a free data retrieval call binding the contract method 0x1b78164b.
//
// Solidity: function getLastForgedBatch() view returns(uint32)
func (_Sybil *SybilSession) GetLastForgedBatch() (uint32, error) {
	return _Sybil.Contract.GetLastForgedBatch(&_Sybil.CallOpts)
}

// GetLastForgedBatch is a free data retrieval call binding the contract method 0x1b78164b.
//
// Solidity: function getLastForgedBatch() view returns(uint32)
func (_Sybil *SybilCallerSession) GetLastForgedBatch() (uint32, error) {
	return _Sybil.Contract.GetLastForgedBatch(&_Sybil.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint32)
func (_Sybil *SybilCaller) GetQueueLength(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "getQueueLength")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint32)
func (_Sybil *SybilSession) GetQueueLength() (uint32, error) {
	return _Sybil.Contract.GetQueueLength(&_Sybil.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint32)
func (_Sybil *SybilCallerSession) GetQueueLength() (uint32, error) {
	return _Sybil.Contract.GetQueueLength(&_Sybil.CallOpts)
}

// GetStateRoot is a free data retrieval call binding the contract method 0x3009c59f.
//
// Solidity: function getStateRoot(uint32 batchNum) view returns(uint256)
func (_Sybil *SybilCaller) GetStateRoot(opts *bind.CallOpts, batchNum uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "getStateRoot", batchNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateRoot is a free data retrieval call binding the contract method 0x3009c59f.
//
// Solidity: function getStateRoot(uint32 batchNum) view returns(uint256)
func (_Sybil *SybilSession) GetStateRoot(batchNum uint32) (*big.Int, error) {
	return _Sybil.Contract.GetStateRoot(&_Sybil.CallOpts, batchNum)
}

// GetStateRoot is a free data retrieval call binding the contract method 0x3009c59f.
//
// Solidity: function getStateRoot(uint32 batchNum) view returns(uint256)
func (_Sybil *SybilCallerSession) GetStateRoot(batchNum uint32) (*big.Int, error) {
	return _Sybil.Contract.GetStateRoot(&_Sybil.CallOpts, batchNum)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Sybil *SybilCaller) LastForgedBatch(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastForgedBatch")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Sybil *SybilSession) LastForgedBatch() (uint32, error) {
	return _Sybil.Contract.LastForgedBatch(&_Sybil.CallOpts)
}

// LastForgedBatch is a free data retrieval call binding the contract method 0x44e0b2ce.
//
// Solidity: function lastForgedBatch() view returns(uint32)
func (_Sybil *SybilCallerSession) LastForgedBatch() (uint32, error) {
	return _Sybil.Contract.LastForgedBatch(&_Sybil.CallOpts)
}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Sybil *SybilCaller) LastIdx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastIdx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Sybil *SybilSession) LastIdx() (*big.Int, error) {
	return _Sybil.Contract.LastIdx(&_Sybil.CallOpts)
}

// LastIdx is a free data retrieval call binding the contract method 0xd486645c.
//
// Solidity: function lastIdx() view returns(uint48)
func (_Sybil *SybilCallerSession) LastIdx() (*big.Int, error) {
	return _Sybil.Contract.LastIdx(&_Sybil.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sybil *SybilCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sybil *SybilSession) Owner() (common.Address, error) {
	return _Sybil.Contract.Owner(&_Sybil.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sybil *SybilCallerSession) Owner() (common.Address, error) {
	return _Sybil.Contract.Owner(&_Sybil.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilCaller) RollupVerifier(opts *bind.CallOpts) (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "rollupVerifier")

	outstruct := new(struct {
		VerifierInterface common.Address
		MaxTx             *big.Int
		NLevel            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VerifierInterface = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MaxTx = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NLevel = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilSession) RollupVerifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Sybil.Contract.RollupVerifier(&_Sybil.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilCallerSession) RollupVerifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Sybil.Contract.RollupVerifier(&_Sybil.CallOpts)
}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) ScoreRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "scoreRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) ScoreRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ScoreRootMap(&_Sybil.CallOpts, arg0)
}

// ScoreRootMap is a free data retrieval call binding the contract method 0xbd8a4a61.
//
// Solidity: function scoreRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) ScoreRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ScoreRootMap(&_Sybil.CallOpts, arg0)
}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xef8140b5.
//
// Solidity: function unprocessedBatchesMap(uint32 ) view returns(bytes)
func (_Sybil *SybilCaller) UnprocessedBatchesMap(opts *bind.CallOpts, arg0 uint32) ([]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "unprocessedBatchesMap", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xef8140b5.
//
// Solidity: function unprocessedBatchesMap(uint32 ) view returns(bytes)
func (_Sybil *SybilSession) UnprocessedBatchesMap(arg0 uint32) ([]byte, error) {
	return _Sybil.Contract.UnprocessedBatchesMap(&_Sybil.CallOpts, arg0)
}

// UnprocessedBatchesMap is a free data retrieval call binding the contract method 0xef8140b5.
//
// Solidity: function unprocessedBatchesMap(uint32 ) view returns(bytes)
func (_Sybil *SybilCallerSession) UnprocessedBatchesMap(arg0 uint32) ([]byte, error) {
	return _Sybil.Contract.UnprocessedBatchesMap(&_Sybil.CallOpts, arg0)
}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) VouchRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "vouchRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) VouchRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.VouchRootMap(&_Sybil.CallOpts, arg0)
}

// VouchRootMap is a free data retrieval call binding the contract method 0xadacd33b.
//
// Solidity: function vouchRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) VouchRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.VouchRootMap(&_Sybil.CallOpts, arg0)
}

// AddTx is a paid mutator transaction binding the contract method 0x2f463f59.
//
// Solidity: function _addTx(address ethAddress, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) returns()
func (_Sybil *SybilTransactor) AddTx(opts *bind.TransactOpts, ethAddress common.Address, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "_addTx", ethAddress, fromIdx, loadAmountF, amountF, toIdx)
}

// AddTx is a paid mutator transaction binding the contract method 0x2f463f59.
//
// Solidity: function _addTx(address ethAddress, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) returns()
func (_Sybil *SybilSession) AddTx(ethAddress common.Address, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.AddTx(&_Sybil.TransactOpts, ethAddress, fromIdx, loadAmountF, amountF, toIdx)
}

// AddTx is a paid mutator transaction binding the contract method 0x2f463f59.
//
// Solidity: function _addTx(address ethAddress, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) returns()
func (_Sybil *SybilTransactorSession) AddTx(ethAddress common.Address, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.AddTx(&_Sybil.TransactOpts, ethAddress, fromIdx, loadAmountF, amountF, toIdx)
}

// CreateAccountDeposit is a paid mutator transaction binding the contract method 0xfbb4a00f.
//
// Solidity: function createAccountDeposit(uint40 loadAmountF) payable returns()
func (_Sybil *SybilTransactor) CreateAccountDeposit(opts *bind.TransactOpts, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "createAccountDeposit", loadAmountF)
}

// CreateAccountDeposit is a paid mutator transaction binding the contract method 0xfbb4a00f.
//
// Solidity: function createAccountDeposit(uint40 loadAmountF) payable returns()
func (_Sybil *SybilSession) CreateAccountDeposit(loadAmountF *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.CreateAccountDeposit(&_Sybil.TransactOpts, loadAmountF)
}

// CreateAccountDeposit is a paid mutator transaction binding the contract method 0xfbb4a00f.
//
// Solidity: function createAccountDeposit(uint40 loadAmountF) payable returns()
func (_Sybil *SybilTransactorSession) CreateAccountDeposit(loadAmountF *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.CreateAccountDeposit(&_Sybil.TransactOpts, loadAmountF)
}

// Deposit is a paid mutator transaction binding the contract method 0x212bafd7.
//
// Solidity: function deposit(uint48 fromIdx, uint40 loadAmountF) payable returns()
func (_Sybil *SybilTransactor) Deposit(opts *bind.TransactOpts, fromIdx *big.Int, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "deposit", fromIdx, loadAmountF)
}

// Deposit is a paid mutator transaction binding the contract method 0x212bafd7.
//
// Solidity: function deposit(uint48 fromIdx, uint40 loadAmountF) payable returns()
func (_Sybil *SybilSession) Deposit(fromIdx *big.Int, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Deposit(&_Sybil.TransactOpts, fromIdx, loadAmountF)
}

// Deposit is a paid mutator transaction binding the contract method 0x212bafd7.
//
// Solidity: function deposit(uint48 fromIdx, uint40 loadAmountF) payable returns()
func (_Sybil *SybilTransactorSession) Deposit(fromIdx *big.Int, loadAmountF *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Deposit(&_Sybil.TransactOpts, fromIdx, loadAmountF)
}

// Exit is a paid mutator transaction binding the contract method 0x8195b790.
//
// Solidity: function exit(uint48 fromIdx, uint40 amountF) returns()
func (_Sybil *SybilTransactor) Exit(opts *bind.TransactOpts, fromIdx *big.Int, amountF *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "exit", fromIdx, amountF)
}

// Exit is a paid mutator transaction binding the contract method 0x8195b790.
//
// Solidity: function exit(uint48 fromIdx, uint40 amountF) returns()
func (_Sybil *SybilSession) Exit(fromIdx *big.Int, amountF *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Exit(&_Sybil.TransactOpts, fromIdx, amountF)
}

// Exit is a paid mutator transaction binding the contract method 0x8195b790.
//
// Solidity: function exit(uint48 fromIdx, uint40 amountF) returns()
func (_Sybil *SybilTransactorSession) Exit(fromIdx *big.Int, amountF *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Exit(&_Sybil.TransactOpts, fromIdx, amountF)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0x894bc2b8.
//
// Solidity: function explodeMultiple(uint48 fromIdx, uint48[] toIdxs) returns()
func (_Sybil *SybilTransactor) ExplodeMultiple(opts *bind.TransactOpts, fromIdx *big.Int, toIdxs []*big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "explodeMultiple", fromIdx, toIdxs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0x894bc2b8.
//
// Solidity: function explodeMultiple(uint48 fromIdx, uint48[] toIdxs) returns()
func (_Sybil *SybilSession) ExplodeMultiple(fromIdx *big.Int, toIdxs []*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ExplodeMultiple(&_Sybil.TransactOpts, fromIdx, toIdxs)
}

// ExplodeMultiple is a paid mutator transaction binding the contract method 0x894bc2b8.
//
// Solidity: function explodeMultiple(uint48 fromIdx, uint48[] toIdxs) returns()
func (_Sybil *SybilTransactorSession) ExplodeMultiple(fromIdx *big.Int, toIdxs []*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ExplodeMultiple(&_Sybil.TransactOpts, fromIdx, toIdxs)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xb1f073d4.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Sybil *SybilTransactor) ForgeBatch(opts *bind.TransactOpts, newLastIdx *big.Int, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "forgeBatch", newLastIdx, newAccountRoot, newVouchRoot, newScoreRoot, newExitRoot, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xb1f073d4.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Sybil *SybilSession) ForgeBatch(newLastIdx *big.Int, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ForgeBatch(&_Sybil.TransactOpts, newLastIdx, newAccountRoot, newVouchRoot, newScoreRoot, newExitRoot, proofA, proofB, proofC)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0xb1f073d4.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newAccountRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Sybil *SybilTransactorSession) ForgeBatch(newLastIdx *big.Int, newAccountRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ForgeBatch(&_Sybil.TransactOpts, newLastIdx, newAccountRoot, newVouchRoot, newScoreRoot, newExitRoot, proofA, proofB, proofC)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Sybil *SybilTransactor) Initialize(opts *bind.TransactOpts, verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "initialize", verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Sybil *SybilSession) Initialize(verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Sybil *SybilTransactorSession) Initialize(verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sybil *SybilTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sybil *SybilSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sybil.Contract.RenounceOwnership(&_Sybil.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sybil *SybilTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sybil.Contract.RenounceOwnership(&_Sybil.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sybil *SybilTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sybil *SybilSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.TransferOwnership(&_Sybil.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sybil *SybilTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.TransferOwnership(&_Sybil.TransactOpts, newOwner)
}

// Unvouch is a paid mutator transaction binding the contract method 0xc1b190c0.
//
// Solidity: function unvouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Sybil *SybilTransactor) Unvouch(opts *bind.TransactOpts, fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "unvouch", fromIdx, toIdx)
}

// Unvouch is a paid mutator transaction binding the contract method 0xc1b190c0.
//
// Solidity: function unvouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Sybil *SybilSession) Unvouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Unvouch(&_Sybil.TransactOpts, fromIdx, toIdx)
}

// Unvouch is a paid mutator transaction binding the contract method 0xc1b190c0.
//
// Solidity: function unvouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Sybil *SybilTransactorSession) Unvouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Unvouch(&_Sybil.TransactOpts, fromIdx, toIdx)
}

// Vouch is a paid mutator transaction binding the contract method 0x11954d3c.
//
// Solidity: function vouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Sybil *SybilTransactor) Vouch(opts *bind.TransactOpts, fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "vouch", fromIdx, toIdx)
}

// Vouch is a paid mutator transaction binding the contract method 0x11954d3c.
//
// Solidity: function vouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Sybil *SybilSession) Vouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Vouch(&_Sybil.TransactOpts, fromIdx, toIdx)
}

// Vouch is a paid mutator transaction binding the contract method 0x11954d3c.
//
// Solidity: function vouch(uint48 fromIdx, uint48 toIdx) returns()
func (_Sybil *SybilTransactorSession) Vouch(fromIdx *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Vouch(&_Sybil.TransactOpts, fromIdx, toIdx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0x62332ee2.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Sybil *SybilTransactor) WithdrawMerkleProof(opts *bind.TransactOpts, amount *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "withdrawMerkleProof", amount, numExitRoot, siblings, idx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0x62332ee2.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Sybil *SybilSession) WithdrawMerkleProof(amount *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.WithdrawMerkleProof(&_Sybil.TransactOpts, amount, numExitRoot, siblings, idx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0x62332ee2.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Sybil *SybilTransactorSession) WithdrawMerkleProof(amount *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.WithdrawMerkleProof(&_Sybil.TransactOpts, amount, numExitRoot, siblings, idx)
}

// SybilForgeBatchIterator is returned from FilterForgeBatch and is used to iterate over the raw logs and unpacked data for ForgeBatch events raised by the Sybil contract.
type SybilForgeBatchIterator struct {
	Event *SybilForgeBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilForgeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilForgeBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilForgeBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilForgeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilForgeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilForgeBatch represents a ForgeBatch event raised by the Sybil contract.
type SybilForgeBatch struct {
	BatchNum     uint32
	L1UserTxsLen uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterForgeBatch is a free log retrieval operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Sybil *SybilFilterer) FilterForgeBatch(opts *bind.FilterOpts, batchNum []uint32) (*SybilForgeBatchIterator, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ForgeBatch", batchNumRule)
	if err != nil {
		return nil, err
	}
	return &SybilForgeBatchIterator{contract: _Sybil.contract, event: "ForgeBatch", logs: logs, sub: sub}, nil
}

// WatchForgeBatch is a free log subscription operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Sybil *SybilFilterer) WatchForgeBatch(opts *bind.WatchOpts, sink chan<- *SybilForgeBatch, batchNum []uint32) (event.Subscription, error) {

	var batchNumRule []interface{}
	for _, batchNumItem := range batchNum {
		batchNumRule = append(batchNumRule, batchNumItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ForgeBatch", batchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilForgeBatch)
				if err := _Sybil.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForgeBatch is a log parse operation binding the contract event 0xe00040c8a3b0bf905636c26924e90520eafc5003324138236fddee2d34588618.
//
// Solidity: event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen)
func (_Sybil *SybilFilterer) ParseForgeBatch(log types.Log) (*SybilForgeBatch, error) {
	event := new(SybilForgeBatch)
	if err := _Sybil.contract.UnpackLog(event, "ForgeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Sybil contract.
type SybilInitializedIterator struct {
	Event *SybilInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilInitialized represents a Initialized event raised by the Sybil contract.
type SybilInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Sybil *SybilFilterer) FilterInitialized(opts *bind.FilterOpts) (*SybilInitializedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SybilInitializedIterator{contract: _Sybil.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Sybil *SybilFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SybilInitialized) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilInitialized)
				if err := _Sybil.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Sybil *SybilFilterer) ParseInitialized(log types.Log) (*SybilInitialized, error) {
	event := new(SybilInitialized)
	if err := _Sybil.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilL1UserTxEventIterator is returned from FilterL1UserTxEvent and is used to iterate over the raw logs and unpacked data for L1UserTxEvent events raised by the Sybil contract.
type SybilL1UserTxEventIterator struct {
	Event *SybilL1UserTxEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilL1UserTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilL1UserTxEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilL1UserTxEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilL1UserTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilL1UserTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilL1UserTxEvent represents a L1UserTxEvent event raised by the Sybil contract.
type SybilL1UserTxEvent struct {
	QueueIndex uint32
	Position   uint8
	L1UserTx   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL1UserTxEvent is a free log retrieval operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Sybil *SybilFilterer) FilterL1UserTxEvent(opts *bind.FilterOpts, queueIndex []uint32, position []uint8) (*SybilL1UserTxEventIterator, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "L1UserTxEvent", queueIndexRule, positionRule)
	if err != nil {
		return nil, err
	}
	return &SybilL1UserTxEventIterator{contract: _Sybil.contract, event: "L1UserTxEvent", logs: logs, sub: sub}, nil
}

// WatchL1UserTxEvent is a free log subscription operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Sybil *SybilFilterer) WatchL1UserTxEvent(opts *bind.WatchOpts, sink chan<- *SybilL1UserTxEvent, queueIndex []uint32, position []uint8) (event.Subscription, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "L1UserTxEvent", queueIndexRule, positionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilL1UserTxEvent)
				if err := _Sybil.contract.UnpackLog(event, "L1UserTxEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseL1UserTxEvent is a log parse operation binding the contract event 0xdd5c7c5ea02d3c5d1621513faa6de53d474ee6f111eda6352a63e3dfe8c40119.
//
// Solidity: event L1UserTxEvent(uint32 indexed queueIndex, uint8 indexed position, bytes l1UserTx)
func (_Sybil *SybilFilterer) ParseL1UserTxEvent(log types.Log) (*SybilL1UserTxEvent, error) {
	event := new(SybilL1UserTxEvent)
	if err := _Sybil.contract.UnpackLog(event, "L1UserTxEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sybil contract.
type SybilOwnershipTransferredIterator struct {
	Event *SybilOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilOwnershipTransferred represents a OwnershipTransferred event raised by the Sybil contract.
type SybilOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sybil *SybilFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SybilOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SybilOwnershipTransferredIterator{contract: _Sybil.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sybil *SybilFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SybilOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilOwnershipTransferred)
				if err := _Sybil.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sybil *SybilFilterer) ParseOwnershipTransferred(log types.Log) (*SybilOwnershipTransferred, error) {
	event := new(SybilOwnershipTransferred)
	if err := _Sybil.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Sybil contract.
type SybilWithdrawEventIterator struct {
	Event *SybilWithdrawEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SybilWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilWithdrawEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SybilWithdrawEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SybilWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilWithdrawEvent represents a WithdrawEvent event raised by the Sybil contract.
type SybilWithdrawEvent struct {
	Idx         *big.Int
	NumExitRoot uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x102db758451b2f65238246a452d00c0c4c8f59d8c623aff254111079418e57ec.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot)
func (_Sybil *SybilFilterer) FilterWithdrawEvent(opts *bind.FilterOpts, idx []*big.Int, numExitRoot []uint32) (*SybilWithdrawEventIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}
	var numExitRootRule []interface{}
	for _, numExitRootItem := range numExitRoot {
		numExitRootRule = append(numExitRootRule, numExitRootItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "WithdrawEvent", idxRule, numExitRootRule)
	if err != nil {
		return nil, err
	}
	return &SybilWithdrawEventIterator{contract: _Sybil.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x102db758451b2f65238246a452d00c0c4c8f59d8c623aff254111079418e57ec.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot)
func (_Sybil *SybilFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *SybilWithdrawEvent, idx []*big.Int, numExitRoot []uint32) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}
	var numExitRootRule []interface{}
	for _, numExitRootItem := range numExitRoot {
		numExitRootRule = append(numExitRootRule, numExitRootItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "WithdrawEvent", idxRule, numExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilWithdrawEvent)
				if err := _Sybil.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawEvent is a log parse operation binding the contract event 0x102db758451b2f65238246a452d00c0c4c8f59d8c623aff254111079418e57ec.
//
// Solidity: event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot)
func (_Sybil *SybilFilterer) ParseWithdrawEvent(log types.Log) (*SybilWithdrawEvent, error) {
	event := new(SybilWithdrawEvent)
	if err := _Sybil.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
