// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
var SybilMetaData = &bind.metadata{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"verifiers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxTxs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"nLevels\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_forgeL1L2BatchTimeout\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_poseidon2Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon3Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon4Elements\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ABSOLUTE_MAX_L1L2BATCHTIMEOUT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_hashFinalNode\",\"inputs\":[{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_hashNode\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addL1Transaction\",\"inputs\":[{\"name\":\"babyPubKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"loadAmountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"amountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"toIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"exitNullifierMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exitRootsMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forgeBatch\",\"inputs\":[{\"name\":\"newLastIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"newStRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newVouchRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newScoreRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newExitRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifierIdx\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"l1Batch\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"proofA\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"proofB\",\"type\":\"uint256[2][2]\",\"internalType\":\"uint256[2][2]\"},{\"name\":\"proofC\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"input\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forgeL1L2BatchTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getL1TransactionQueue\",\"inputs\":[{\"name\":\"queueIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastForgedBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueueLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateRoot\",\"inputs\":[{\"name\":\"batchNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"verifiers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"maxTxs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"nLevels\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_forgeL1L2BatchTimeout\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_poseidon2Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon3Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon4Elements\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1L2TxsDataHashMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastForgedBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastIdx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastL1L2Batch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mapL1TxQueue\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextL1FillingQueue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextL1ToForgeQueue\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollupVerifiers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"verifierInterface\",\"type\":\"address\",\"internalType\":\"contractVerifierRollupInterface\"},{\"name\":\"maxTxs\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nLevels\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scoreRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setForgeL1L2BatchTimeout\",\"inputs\":[{\"name\":\"newTimeout\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stateRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vouchRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawMerkleProof\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"babyPubKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numExitRoot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"siblings\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"idx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ForgeBatch\",\"inputs\":[{\"name\":\"batchNum\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"l1UserTxsLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialize\",\"inputs\":[{\"name\":\"forgeL1L2BatchTimeout\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1UserTxEvent\",\"inputs\":[{\"name\":\"queueIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"position\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"l1UserTx\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateForgeL1L2BatchTimeout\",\"inputs\":[{\"name\":\"newForgeL1L2BatchTimeout\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawEvent\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint48\",\"indexed\":true,\"internalType\":\"uint48\"},{\"name\":\"numExitRoot\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BatchTimeoutExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EthTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InternalTxNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCreateAccountTransaction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDepositTransaction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidForceExitTransaction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidForceExplodeTransaction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPoseidonAddress\",\"inputs\":[{\"name\":\"elementType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTransactionParameters\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LoadAmountDoesNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LoadAmountExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SmtProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawAlreadyDone\",\"inputs\":[]}]",
}

// SybilABI is the input ABI used to generate the binding from.
// Deprecated: Use SybilMetaData.ABI instead.
var SybilABI = SybilMetaData.ABI

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

// ABSOLUTEMAXL1L2BATCHTIMEOUT is a free data retrieval call binding the contract method 0x95a09f2a.
//
// Solidity: function ABSOLUTE_MAX_L1L2BATCHTIMEOUT() view returns(uint8)
func (_Sybil *SybilCaller) ABSOLUTEMAXL1L2BATCHTIMEOUT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "ABSOLUTE_MAX_L1L2BATCHTIMEOUT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ABSOLUTEMAXL1L2BATCHTIMEOUT is a free data retrieval call binding the contract method 0x95a09f2a.
//
// Solidity: function ABSOLUTE_MAX_L1L2BATCHTIMEOUT() view returns(uint8)
func (_Sybil *SybilSession) ABSOLUTEMAXL1L2BATCHTIMEOUT() (uint8, error) {
	return _Sybil.Contract.ABSOLUTEMAXL1L2BATCHTIMEOUT(&_Sybil.CallOpts)
}

// ABSOLUTEMAXL1L2BATCHTIMEOUT is a free data retrieval call binding the contract method 0x95a09f2a.
//
// Solidity: function ABSOLUTE_MAX_L1L2BATCHTIMEOUT() view returns(uint8)
func (_Sybil *SybilCallerSession) ABSOLUTEMAXL1L2BATCHTIMEOUT() (uint8, error) {
	return _Sybil.Contract.ABSOLUTEMAXL1L2BATCHTIMEOUT(&_Sybil.CallOpts)
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

// ExitRootsMap is a free data retrieval call binding the contract method 0x3ee641ea.
//
// Solidity: function exitRootsMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) ExitRootsMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "exitRootsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitRootsMap is a free data retrieval call binding the contract method 0x3ee641ea.
//
// Solidity: function exitRootsMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) ExitRootsMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ExitRootsMap(&_Sybil.CallOpts, arg0)
}

// ExitRootsMap is a free data retrieval call binding the contract method 0x3ee641ea.
//
// Solidity: function exitRootsMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) ExitRootsMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.ExitRootsMap(&_Sybil.CallOpts, arg0)
}

// ForgeL1L2BatchTimeout is a free data retrieval call binding the contract method 0xa3275838.
//
// Solidity: function forgeL1L2BatchTimeout() view returns(uint8)
func (_Sybil *SybilCaller) ForgeL1L2BatchTimeout(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "forgeL1L2BatchTimeout")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ForgeL1L2BatchTimeout is a free data retrieval call binding the contract method 0xa3275838.
//
// Solidity: function forgeL1L2BatchTimeout() view returns(uint8)
func (_Sybil *SybilSession) ForgeL1L2BatchTimeout() (uint8, error) {
	return _Sybil.Contract.ForgeL1L2BatchTimeout(&_Sybil.CallOpts)
}

// ForgeL1L2BatchTimeout is a free data retrieval call binding the contract method 0xa3275838.
//
// Solidity: function forgeL1L2BatchTimeout() view returns(uint8)
func (_Sybil *SybilCallerSession) ForgeL1L2BatchTimeout() (uint8, error) {
	return _Sybil.Contract.ForgeL1L2BatchTimeout(&_Sybil.CallOpts)
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

// L1L2TxsDataHashMap is a free data retrieval call binding the contract method 0xce5ec65a.
//
// Solidity: function l1L2TxsDataHashMap(uint32 ) view returns(bytes32)
func (_Sybil *SybilCaller) L1L2TxsDataHashMap(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "l1L2TxsDataHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1L2TxsDataHashMap is a free data retrieval call binding the contract method 0xce5ec65a.
//
// Solidity: function l1L2TxsDataHashMap(uint32 ) view returns(bytes32)
func (_Sybil *SybilSession) L1L2TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Sybil.Contract.L1L2TxsDataHashMap(&_Sybil.CallOpts, arg0)
}

// L1L2TxsDataHashMap is a free data retrieval call binding the contract method 0xce5ec65a.
//
// Solidity: function l1L2TxsDataHashMap(uint32 ) view returns(bytes32)
func (_Sybil *SybilCallerSession) L1L2TxsDataHashMap(arg0 uint32) ([32]byte, error) {
	return _Sybil.Contract.L1L2TxsDataHashMap(&_Sybil.CallOpts, arg0)
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

// LastL1L2Batch is a free data retrieval call binding the contract method 0x84ef9ed4.
//
// Solidity: function lastL1L2Batch() view returns(uint64)
func (_Sybil *SybilCaller) LastL1L2Batch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastL1L2Batch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastL1L2Batch is a free data retrieval call binding the contract method 0x84ef9ed4.
//
// Solidity: function lastL1L2Batch() view returns(uint64)
func (_Sybil *SybilSession) LastL1L2Batch() (uint64, error) {
	return _Sybil.Contract.LastL1L2Batch(&_Sybil.CallOpts)
}

// LastL1L2Batch is a free data retrieval call binding the contract method 0x84ef9ed4.
//
// Solidity: function lastL1L2Batch() view returns(uint64)
func (_Sybil *SybilCallerSession) LastL1L2Batch() (uint64, error) {
	return _Sybil.Contract.LastL1L2Batch(&_Sybil.CallOpts)
}

// MapL1TxQueue is a free data retrieval call binding the contract method 0xdc3e718e.
//
// Solidity: function mapL1TxQueue(uint32 ) view returns(bytes)
func (_Sybil *SybilCaller) MapL1TxQueue(opts *bind.CallOpts, arg0 uint32) ([]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "mapL1TxQueue", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MapL1TxQueue is a free data retrieval call binding the contract method 0xdc3e718e.
//
// Solidity: function mapL1TxQueue(uint32 ) view returns(bytes)
func (_Sybil *SybilSession) MapL1TxQueue(arg0 uint32) ([]byte, error) {
	return _Sybil.Contract.MapL1TxQueue(&_Sybil.CallOpts, arg0)
}

// MapL1TxQueue is a free data retrieval call binding the contract method 0xdc3e718e.
//
// Solidity: function mapL1TxQueue(uint32 ) view returns(bytes)
func (_Sybil *SybilCallerSession) MapL1TxQueue(arg0 uint32) ([]byte, error) {
	return _Sybil.Contract.MapL1TxQueue(&_Sybil.CallOpts, arg0)
}

// NextL1FillingQueue is a free data retrieval call binding the contract method 0x0ee8e52b.
//
// Solidity: function nextL1FillingQueue() view returns(uint32)
func (_Sybil *SybilCaller) NextL1FillingQueue(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "nextL1FillingQueue")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextL1FillingQueue is a free data retrieval call binding the contract method 0x0ee8e52b.
//
// Solidity: function nextL1FillingQueue() view returns(uint32)
func (_Sybil *SybilSession) NextL1FillingQueue() (uint32, error) {
	return _Sybil.Contract.NextL1FillingQueue(&_Sybil.CallOpts)
}

// NextL1FillingQueue is a free data retrieval call binding the contract method 0x0ee8e52b.
//
// Solidity: function nextL1FillingQueue() view returns(uint32)
func (_Sybil *SybilCallerSession) NextL1FillingQueue() (uint32, error) {
	return _Sybil.Contract.NextL1FillingQueue(&_Sybil.CallOpts)
}

// NextL1ToForgeQueue is a free data retrieval call binding the contract method 0xd0f32e67.
//
// Solidity: function nextL1ToForgeQueue() view returns(uint32)
func (_Sybil *SybilCaller) NextL1ToForgeQueue(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "nextL1ToForgeQueue")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextL1ToForgeQueue is a free data retrieval call binding the contract method 0xd0f32e67.
//
// Solidity: function nextL1ToForgeQueue() view returns(uint32)
func (_Sybil *SybilSession) NextL1ToForgeQueue() (uint32, error) {
	return _Sybil.Contract.NextL1ToForgeQueue(&_Sybil.CallOpts)
}

// NextL1ToForgeQueue is a free data retrieval call binding the contract method 0xd0f32e67.
//
// Solidity: function nextL1ToForgeQueue() view returns(uint32)
func (_Sybil *SybilCallerSession) NextL1ToForgeQueue() (uint32, error) {
	return _Sybil.Contract.NextL1ToForgeQueue(&_Sybil.CallOpts)
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

// RollupVerifiers is a free data retrieval call binding the contract method 0x38330200.
//
// Solidity: function rollupVerifiers(uint256 ) view returns(address verifierInterface, uint256 maxTxs, uint256 nLevels)
func (_Sybil *SybilCaller) RollupVerifiers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	VerifierInterface common.Address
	MaxTxs            *big.Int
	NLevels           *big.Int
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "rollupVerifiers", arg0)

	outstruct := new(struct {
		VerifierInterface common.Address
		MaxTxs            *big.Int
		NLevels           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VerifierInterface = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MaxTxs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NLevels = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RollupVerifiers is a free data retrieval call binding the contract method 0x38330200.
//
// Solidity: function rollupVerifiers(uint256 ) view returns(address verifierInterface, uint256 maxTxs, uint256 nLevels)
func (_Sybil *SybilSession) RollupVerifiers(arg0 *big.Int) (struct {
	VerifierInterface common.Address
	MaxTxs            *big.Int
	NLevels           *big.Int
}, error) {
	return _Sybil.Contract.RollupVerifiers(&_Sybil.CallOpts, arg0)
}

// RollupVerifiers is a free data retrieval call binding the contract method 0x38330200.
//
// Solidity: function rollupVerifiers(uint256 ) view returns(address verifierInterface, uint256 maxTxs, uint256 nLevels)
func (_Sybil *SybilCallerSession) RollupVerifiers(arg0 *big.Int) (struct {
	VerifierInterface common.Address
	MaxTxs            *big.Int
	NLevels           *big.Int
}, error) {
	return _Sybil.Contract.RollupVerifiers(&_Sybil.CallOpts, arg0)
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

// StateRootMap is a free data retrieval call binding the contract method 0x9e00d7ea.
//
// Solidity: function stateRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCaller) StateRootMap(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "stateRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateRootMap is a free data retrieval call binding the contract method 0x9e00d7ea.
//
// Solidity: function stateRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilSession) StateRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.StateRootMap(&_Sybil.CallOpts, arg0)
}

// StateRootMap is a free data retrieval call binding the contract method 0x9e00d7ea.
//
// Solidity: function stateRootMap(uint32 ) view returns(uint256)
func (_Sybil *SybilCallerSession) StateRootMap(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.StateRootMap(&_Sybil.CallOpts, arg0)
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

// AddL1Transaction is a paid mutator transaction binding the contract method 0x29b1ac6b.
//
// Solidity: function addL1Transaction(string babyPubKey, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) payable returns()
func (_Sybil *SybilTransactor) AddL1Transaction(opts *bind.TransactOpts, babyPubKey string, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "addL1Transaction", babyPubKey, fromIdx, loadAmountF, amountF, toIdx)
}

// AddL1Transaction is a paid mutator transaction binding the contract method 0x29b1ac6b.
//
// Solidity: function addL1Transaction(string babyPubKey, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) payable returns()
func (_Sybil *SybilSession) AddL1Transaction(babyPubKey string, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.AddL1Transaction(&_Sybil.TransactOpts, babyPubKey, fromIdx, loadAmountF, amountF, toIdx)
}

// AddL1Transaction is a paid mutator transaction binding the contract method 0x29b1ac6b.
//
// Solidity: function addL1Transaction(string babyPubKey, uint48 fromIdx, uint40 loadAmountF, uint40 amountF, uint48 toIdx) payable returns()
func (_Sybil *SybilTransactorSession) AddL1Transaction(babyPubKey string, fromIdx *big.Int, loadAmountF *big.Int, amountF *big.Int, toIdx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.AddL1Transaction(&_Sybil.TransactOpts, babyPubKey, fromIdx, loadAmountF, amountF, toIdx)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x8112639b.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newStRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, uint8 verifierIdx, bool l1Batch, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint256 input) returns()
func (_Sybil *SybilTransactor) ForgeBatch(opts *bind.TransactOpts, newLastIdx *big.Int, newStRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, verifierIdx uint8, l1Batch bool, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, input *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "forgeBatch", newLastIdx, newStRoot, newVouchRoot, newScoreRoot, newExitRoot, verifierIdx, l1Batch, proofA, proofB, proofC, input)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x8112639b.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newStRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, uint8 verifierIdx, bool l1Batch, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint256 input) returns()
func (_Sybil *SybilSession) ForgeBatch(newLastIdx *big.Int, newStRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, verifierIdx uint8, l1Batch bool, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, input *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ForgeBatch(&_Sybil.TransactOpts, newLastIdx, newStRoot, newVouchRoot, newScoreRoot, newExitRoot, verifierIdx, l1Batch, proofA, proofB, proofC, input)
}

// ForgeBatch is a paid mutator transaction binding the contract method 0x8112639b.
//
// Solidity: function forgeBatch(uint48 newLastIdx, uint256 newStRoot, uint256 newVouchRoot, uint256 newScoreRoot, uint256 newExitRoot, uint8 verifierIdx, bool l1Batch, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint256 input) returns()
func (_Sybil *SybilTransactorSession) ForgeBatch(newLastIdx *big.Int, newStRoot *big.Int, newVouchRoot *big.Int, newScoreRoot *big.Int, newExitRoot *big.Int, verifierIdx uint8, l1Batch bool, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, input *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.ForgeBatch(&_Sybil.TransactOpts, newLastIdx, newStRoot, newVouchRoot, newScoreRoot, newExitRoot, verifierIdx, l1Batch, proofA, proofB, proofC, input)
}

// Initialize is a paid mutator transaction binding the contract method 0x0f1c7003.
//
// Solidity: function initialize(address[] verifiers, uint256[] maxTxs, uint256[] nLevels, uint8 _forgeL1L2BatchTimeout, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Sybil *SybilTransactor) Initialize(opts *bind.TransactOpts, verifiers []common.Address, maxTxs []*big.Int, nLevels []*big.Int, _forgeL1L2BatchTimeout uint8, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "initialize", verifiers, maxTxs, nLevels, _forgeL1L2BatchTimeout, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0x0f1c7003.
//
// Solidity: function initialize(address[] verifiers, uint256[] maxTxs, uint256[] nLevels, uint8 _forgeL1L2BatchTimeout, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Sybil *SybilSession) Initialize(verifiers []common.Address, maxTxs []*big.Int, nLevels []*big.Int, _forgeL1L2BatchTimeout uint8, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, verifiers, maxTxs, nLevels, _forgeL1L2BatchTimeout, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
}

// Initialize is a paid mutator transaction binding the contract method 0x0f1c7003.
//
// Solidity: function initialize(address[] verifiers, uint256[] maxTxs, uint256[] nLevels, uint8 _forgeL1L2BatchTimeout, address _poseidon2Elements, address _poseidon3Elements, address _poseidon4Elements) returns()
func (_Sybil *SybilTransactorSession) Initialize(verifiers []common.Address, maxTxs []*big.Int, nLevels []*big.Int, _forgeL1L2BatchTimeout uint8, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _poseidon4Elements common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, verifiers, maxTxs, nLevels, _forgeL1L2BatchTimeout, _poseidon2Elements, _poseidon3Elements, _poseidon4Elements)
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

// SetForgeL1L2BatchTimeout is a paid mutator transaction binding the contract method 0x40105466.
//
// Solidity: function setForgeL1L2BatchTimeout(uint8 newTimeout) returns()
func (_Sybil *SybilTransactor) SetForgeL1L2BatchTimeout(opts *bind.TransactOpts, newTimeout uint8) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "setForgeL1L2BatchTimeout", newTimeout)
}

// SetForgeL1L2BatchTimeout is a paid mutator transaction binding the contract method 0x40105466.
//
// Solidity: function setForgeL1L2BatchTimeout(uint8 newTimeout) returns()
func (_Sybil *SybilSession) SetForgeL1L2BatchTimeout(newTimeout uint8) (*types.Transaction, error) {
	return _Sybil.Contract.SetForgeL1L2BatchTimeout(&_Sybil.TransactOpts, newTimeout)
}

// SetForgeL1L2BatchTimeout is a paid mutator transaction binding the contract method 0x40105466.
//
// Solidity: function setForgeL1L2BatchTimeout(uint8 newTimeout) returns()
func (_Sybil *SybilTransactorSession) SetForgeL1L2BatchTimeout(newTimeout uint8) (*types.Transaction, error) {
	return _Sybil.Contract.SetForgeL1L2BatchTimeout(&_Sybil.TransactOpts, newTimeout)
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

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0xc285d345.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint256 babyPubKey, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Sybil *SybilTransactor) WithdrawMerkleProof(opts *bind.TransactOpts, amount *big.Int, babyPubKey *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "withdrawMerkleProof", amount, babyPubKey, numExitRoot, siblings, idx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0xc285d345.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint256 babyPubKey, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Sybil *SybilSession) WithdrawMerkleProof(amount *big.Int, babyPubKey *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.WithdrawMerkleProof(&_Sybil.TransactOpts, amount, babyPubKey, numExitRoot, siblings, idx)
}

// WithdrawMerkleProof is a paid mutator transaction binding the contract method 0xc285d345.
//
// Solidity: function withdrawMerkleProof(uint192 amount, uint256 babyPubKey, uint32 numExitRoot, uint256[] siblings, uint48 idx) returns()
func (_Sybil *SybilTransactorSession) WithdrawMerkleProof(amount *big.Int, babyPubKey *big.Int, numExitRoot uint32, siblings []*big.Int, idx *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.WithdrawMerkleProof(&_Sybil.TransactOpts, amount, babyPubKey, numExitRoot, siblings, idx)
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

// SybilInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Sybil contract.
type SybilInitializeIterator struct {
	Event *SybilInitialize // Event containing the contract specifics and raw log

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
func (it *SybilInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilInitialize)
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
		it.Event = new(SybilInitialize)
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
func (it *SybilInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilInitialize represents a Initialize event raised by the Sybil contract.
type SybilInitialize struct {
	ForgeL1L2BatchTimeout uint8
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0xd2b214d5e2d2f958eb3b30690fa010715ebfdb9438837a496031fd1d0462e593.
//
// Solidity: event Initialize(uint8 forgeL1L2BatchTimeout)
func (_Sybil *SybilFilterer) FilterInitialize(opts *bind.FilterOpts) (*SybilInitializeIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &SybilInitializeIterator{contract: _Sybil.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0xd2b214d5e2d2f958eb3b30690fa010715ebfdb9438837a496031fd1d0462e593.
//
// Solidity: event Initialize(uint8 forgeL1L2BatchTimeout)
func (_Sybil *SybilFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *SybilInitialize) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilInitialize)
				if err := _Sybil.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0xd2b214d5e2d2f958eb3b30690fa010715ebfdb9438837a496031fd1d0462e593.
//
// Solidity: event Initialize(uint8 forgeL1L2BatchTimeout)
func (_Sybil *SybilFilterer) ParseInitialize(log types.Log) (*SybilInitialize, error) {
	event := new(SybilInitialize)
	if err := _Sybil.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// SybilUpdateForgeL1L2BatchTimeoutIterator is returned from FilterUpdateForgeL1L2BatchTimeout and is used to iterate over the raw logs and unpacked data for UpdateForgeL1L2BatchTimeout events raised by the Sybil contract.
type SybilUpdateForgeL1L2BatchTimeoutIterator struct {
	Event *SybilUpdateForgeL1L2BatchTimeout // Event containing the contract specifics and raw log

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
func (it *SybilUpdateForgeL1L2BatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilUpdateForgeL1L2BatchTimeout)
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
		it.Event = new(SybilUpdateForgeL1L2BatchTimeout)
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
func (it *SybilUpdateForgeL1L2BatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilUpdateForgeL1L2BatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilUpdateForgeL1L2BatchTimeout represents a UpdateForgeL1L2BatchTimeout event raised by the Sybil contract.
type SybilUpdateForgeL1L2BatchTimeout struct {
	NewForgeL1L2BatchTimeout uint8
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdateForgeL1L2BatchTimeout is a free log retrieval operation binding the contract event 0xff6221781ac525b04585dbb55cd2ebd2a92c828ca3e42b23813a1137ac974431.
//
// Solidity: event UpdateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout)
func (_Sybil *SybilFilterer) FilterUpdateForgeL1L2BatchTimeout(opts *bind.FilterOpts) (*SybilUpdateForgeL1L2BatchTimeoutIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "UpdateForgeL1L2BatchTimeout")
	if err != nil {
		return nil, err
	}
	return &SybilUpdateForgeL1L2BatchTimeoutIterator{contract: _Sybil.contract, event: "UpdateForgeL1L2BatchTimeout", logs: logs, sub: sub}, nil
}

// WatchUpdateForgeL1L2BatchTimeout is a free log subscription operation binding the contract event 0xff6221781ac525b04585dbb55cd2ebd2a92c828ca3e42b23813a1137ac974431.
//
// Solidity: event UpdateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout)
func (_Sybil *SybilFilterer) WatchUpdateForgeL1L2BatchTimeout(opts *bind.WatchOpts, sink chan<- *SybilUpdateForgeL1L2BatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "UpdateForgeL1L2BatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilUpdateForgeL1L2BatchTimeout)
				if err := _Sybil.contract.UnpackLog(event, "UpdateForgeL1L2BatchTimeout", log); err != nil {
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

// ParseUpdateForgeL1L2BatchTimeout is a log parse operation binding the contract event 0xff6221781ac525b04585dbb55cd2ebd2a92c828ca3e42b23813a1137ac974431.
//
// Solidity: event UpdateForgeL1L2BatchTimeout(uint8 newForgeL1L2BatchTimeout)
func (_Sybil *SybilFilterer) ParseUpdateForgeL1L2BatchTimeout(log types.Log) (*SybilUpdateForgeL1L2BatchTimeout, error) {
	event := new(SybilUpdateForgeL1L2BatchTimeout)
	if err := _Sybil.contract.UnpackLog(event, "UpdateForgeL1L2BatchTimeout", log); err != nil {
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
