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
	ABI: "[{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_addTx\",\"inputs\":[{\"name\":\"ethAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"loadAmountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"amountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"toIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_hashFinalNode\",\"inputs\":[{\"name\":\"key\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_hashNode\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAccountDeposit\",\"inputs\":[{\"name\":\"loadAmountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"currentFillingBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"loadAmountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"exit\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"amountF\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitNullifierMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exitRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explodeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explodeMultiple\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"toIdxs\",\"type\":\"uint48[]\",\"internalType\":\"uint48[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forgeBatch\",\"inputs\":[{\"name\":\"newLastIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"newAccountRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newVouchRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newScoreRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newExitRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofA\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"proofB\",\"type\":\"uint256[2][2]\",\"internalType\":\"uint256[2][2]\"},{\"name\":\"proofC\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getL1TransactionQueue\",\"inputs\":[{\"name\":\"queueIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLastForgedBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueueLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateRoot\",\"inputs\":[{\"name\":\"batchNum\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxTx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nLevel\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_poseidon2Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poseidon3Elements\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_adminRole\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastForgedBatch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastIdx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"scoreRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unprocessedBatchesMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unvouch\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"toIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateExplodeAmount\",\"inputs\":[{\"name\":\"_explodeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMinBalance\",\"inputs\":[{\"name\":\"_minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"verifierInterface\",\"type\":\"address\",\"internalType\":\"contractIVerifier\"},{\"name\":\"maxTx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nLevel\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vouch\",\"inputs\":[{\"name\":\"fromIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"toIdx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vouchRootMap\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawMerkleProof\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"numExitRoot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"siblings\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"idx\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExplodeAmountUpdated\",\"inputs\":[{\"name\":\"explodeAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForgeBatch\",\"inputs\":[{\"name\":\"batchNum\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"l1UserTxsLen\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1UserTxEvent\",\"inputs\":[{\"name\":\"queueIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"position\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"l1UserTx\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBalanceUpdated\",\"inputs\":[{\"name\":\"minBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawEvent\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint48\",\"indexed\":true,\"internalType\":\"uint48\"},{\"name\":\"numExitRoot\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AmountExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EthTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFromIdx\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPoseidonAddress\",\"inputs\":[{\"name\":\"elementType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidToIdx\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LoadAmountDoesNotMatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LoadAmountExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmtProofInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawAlreadyDone\",\"inputs\":[]}]",
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

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilSession) ADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.ADMINROLE(&_Sybil.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCallerSession) ADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.ADMINROLE(&_Sybil.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.DEFAULTADMINROLE(&_Sybil.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sybil *SybilCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Sybil.Contract.DEFAULTADMINROLE(&_Sybil.CallOpts)
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

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Sybil *SybilCaller) ExplodeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "explodeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Sybil *SybilSession) ExplodeAmount() (*big.Int, error) {
	return _Sybil.Contract.ExplodeAmount(&_Sybil.CallOpts)
}

// ExplodeAmount is a free data retrieval call binding the contract method 0x1dbceceb.
//
// Solidity: function explodeAmount() view returns(uint256)
func (_Sybil *SybilCallerSession) ExplodeAmount() (*big.Int, error) {
	return _Sybil.Contract.ExplodeAmount(&_Sybil.CallOpts)
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

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sybil *SybilCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sybil *SybilSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Sybil.Contract.GetRoleAdmin(&_Sybil.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sybil *SybilCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Sybil.Contract.GetRoleAdmin(&_Sybil.CallOpts, role)
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

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sybil *SybilCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sybil *SybilSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Sybil.Contract.HasRole(&_Sybil.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sybil *SybilCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Sybil.Contract.HasRole(&_Sybil.CallOpts, role, account)
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

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Sybil *SybilCaller) MinBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "minBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Sybil *SybilSession) MinBalance() (*big.Int, error) {
	return _Sybil.Contract.MinBalance(&_Sybil.CallOpts)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() view returns(uint256)
func (_Sybil *SybilCallerSession) MinBalance() (*big.Int, error) {
	return _Sybil.Contract.MinBalance(&_Sybil.CallOpts)
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sybil *SybilCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sybil *SybilSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sybil.Contract.SupportsInterface(&_Sybil.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sybil *SybilCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sybil.Contract.SupportsInterface(&_Sybil.CallOpts, interfaceId)
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

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilCaller) Verifier(opts *bind.CallOpts) (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "verifier")

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

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilSession) Verifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Sybil.Contract.Verifier(&_Sybil.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address verifierInterface, uint256 maxTx, uint256 nLevel)
func (_Sybil *SybilCallerSession) Verifier() (struct {
	VerifierInterface common.Address
	MaxTx             *big.Int
	NLevel            *big.Int
}, error) {
	return _Sybil.Contract.Verifier(&_Sybil.CallOpts)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sybil *SybilSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.GrantRole(&_Sybil.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.GrantRole(&_Sybil.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Sybil *SybilTransactor) Initialize(opts *bind.TransactOpts, _verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "initialize", _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Sybil *SybilSession) Initialize(_verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// Initialize is a paid mutator transaction binding the contract method 0x11917b1d.
//
// Solidity: function initialize(address _verifier, uint256 maxTx, uint256 nLevel, address _poseidon2Elements, address _poseidon3Elements, address _adminRole) returns()
func (_Sybil *SybilTransactorSession) Initialize(_verifier common.Address, maxTx *big.Int, nLevel *big.Int, _poseidon2Elements common.Address, _poseidon3Elements common.Address, _adminRole common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Initialize(&_Sybil.TransactOpts, _verifier, maxTx, nLevel, _poseidon2Elements, _poseidon3Elements, _adminRole)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Sybil *SybilTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Sybil *SybilSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RenounceRole(&_Sybil.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Sybil *SybilTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RenounceRole(&_Sybil.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sybil *SybilSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RevokeRole(&_Sybil.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sybil *SybilTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.RevokeRole(&_Sybil.TransactOpts, role, account)
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

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Sybil *SybilTransactor) UpdateExplodeAmount(opts *bind.TransactOpts, _explodeAmount *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "updateExplodeAmount", _explodeAmount)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Sybil *SybilSession) UpdateExplodeAmount(_explodeAmount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateExplodeAmount(&_Sybil.TransactOpts, _explodeAmount)
}

// UpdateExplodeAmount is a paid mutator transaction binding the contract method 0xaa4f9116.
//
// Solidity: function updateExplodeAmount(uint256 _explodeAmount) returns()
func (_Sybil *SybilTransactorSession) UpdateExplodeAmount(_explodeAmount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateExplodeAmount(&_Sybil.TransactOpts, _explodeAmount)
}

// UpdateMinBalance is a paid mutator transaction binding the contract method 0xd83567ab.
//
// Solidity: function updateMinBalance(uint256 _minBalance) returns()
func (_Sybil *SybilTransactor) UpdateMinBalance(opts *bind.TransactOpts, _minBalance *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "updateMinBalance", _minBalance)
}

// UpdateMinBalance is a paid mutator transaction binding the contract method 0xd83567ab.
//
// Solidity: function updateMinBalance(uint256 _minBalance) returns()
func (_Sybil *SybilSession) UpdateMinBalance(_minBalance *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateMinBalance(&_Sybil.TransactOpts, _minBalance)
}

// UpdateMinBalance is a paid mutator transaction binding the contract method 0xd83567ab.
//
// Solidity: function updateMinBalance(uint256 _minBalance) returns()
func (_Sybil *SybilTransactorSession) UpdateMinBalance(_minBalance *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.UpdateMinBalance(&_Sybil.TransactOpts, _minBalance)
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

// SybilExplodeAmountUpdatedIterator is returned from FilterExplodeAmountUpdated and is used to iterate over the raw logs and unpacked data for ExplodeAmountUpdated events raised by the Sybil contract.
type SybilExplodeAmountUpdatedIterator struct {
	Event *SybilExplodeAmountUpdated // Event containing the contract specifics and raw log

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
func (it *SybilExplodeAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilExplodeAmountUpdated)
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
		it.Event = new(SybilExplodeAmountUpdated)
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
func (it *SybilExplodeAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilExplodeAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilExplodeAmountUpdated represents a ExplodeAmountUpdated event raised by the Sybil contract.
type SybilExplodeAmountUpdated struct {
	ExplodeAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExplodeAmountUpdated is a free log retrieval operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Sybil *SybilFilterer) FilterExplodeAmountUpdated(opts *bind.FilterOpts) (*SybilExplodeAmountUpdatedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ExplodeAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &SybilExplodeAmountUpdatedIterator{contract: _Sybil.contract, event: "ExplodeAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchExplodeAmountUpdated is a free log subscription operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Sybil *SybilFilterer) WatchExplodeAmountUpdated(opts *bind.WatchOpts, sink chan<- *SybilExplodeAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ExplodeAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilExplodeAmountUpdated)
				if err := _Sybil.contract.UnpackLog(event, "ExplodeAmountUpdated", log); err != nil {
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

// ParseExplodeAmountUpdated is a log parse operation binding the contract event 0xe4d07ddba7bee2524330d02dcb17edc18e025341f913484d72c85301628b4a79.
//
// Solidity: event ExplodeAmountUpdated(uint256 explodeAmount)
func (_Sybil *SybilFilterer) ParseExplodeAmountUpdated(log types.Log) (*SybilExplodeAmountUpdated, error) {
	event := new(SybilExplodeAmountUpdated)
	if err := _Sybil.contract.UnpackLog(event, "ExplodeAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// SybilMinBalanceUpdatedIterator is returned from FilterMinBalanceUpdated and is used to iterate over the raw logs and unpacked data for MinBalanceUpdated events raised by the Sybil contract.
type SybilMinBalanceUpdatedIterator struct {
	Event *SybilMinBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *SybilMinBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilMinBalanceUpdated)
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
		it.Event = new(SybilMinBalanceUpdated)
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
func (it *SybilMinBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilMinBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilMinBalanceUpdated represents a MinBalanceUpdated event raised by the Sybil contract.
type SybilMinBalanceUpdated struct {
	MinBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinBalanceUpdated is a free log retrieval operation binding the contract event 0x4e1cd0a17dbc393262d4d9b66380671f5273c5f0a34fed0ed36c50ba6b1f0e16.
//
// Solidity: event MinBalanceUpdated(uint256 minBalance)
func (_Sybil *SybilFilterer) FilterMinBalanceUpdated(opts *bind.FilterOpts) (*SybilMinBalanceUpdatedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "MinBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &SybilMinBalanceUpdatedIterator{contract: _Sybil.contract, event: "MinBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchMinBalanceUpdated is a free log subscription operation binding the contract event 0x4e1cd0a17dbc393262d4d9b66380671f5273c5f0a34fed0ed36c50ba6b1f0e16.
//
// Solidity: event MinBalanceUpdated(uint256 minBalance)
func (_Sybil *SybilFilterer) WatchMinBalanceUpdated(opts *bind.WatchOpts, sink chan<- *SybilMinBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "MinBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilMinBalanceUpdated)
				if err := _Sybil.contract.UnpackLog(event, "MinBalanceUpdated", log); err != nil {
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

// ParseMinBalanceUpdated is a log parse operation binding the contract event 0x4e1cd0a17dbc393262d4d9b66380671f5273c5f0a34fed0ed36c50ba6b1f0e16.
//
// Solidity: event MinBalanceUpdated(uint256 minBalance)
func (_Sybil *SybilFilterer) ParseMinBalanceUpdated(log types.Log) (*SybilMinBalanceUpdated, error) {
	event := new(SybilMinBalanceUpdated)
	if err := _Sybil.contract.UnpackLog(event, "MinBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Sybil contract.
type SybilRoleAdminChangedIterator struct {
	Event *SybilRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SybilRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilRoleAdminChanged)
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
		it.Event = new(SybilRoleAdminChanged)
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
func (it *SybilRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilRoleAdminChanged represents a RoleAdminChanged event raised by the Sybil contract.
type SybilRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sybil *SybilFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SybilRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SybilRoleAdminChangedIterator{contract: _Sybil.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sybil *SybilFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SybilRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilRoleAdminChanged)
				if err := _Sybil.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sybil *SybilFilterer) ParseRoleAdminChanged(log types.Log) (*SybilRoleAdminChanged, error) {
	event := new(SybilRoleAdminChanged)
	if err := _Sybil.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Sybil contract.
type SybilRoleGrantedIterator struct {
	Event *SybilRoleGranted // Event containing the contract specifics and raw log

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
func (it *SybilRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilRoleGranted)
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
		it.Event = new(SybilRoleGranted)
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
func (it *SybilRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilRoleGranted represents a RoleGranted event raised by the Sybil contract.
type SybilRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SybilRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SybilRoleGrantedIterator{contract: _Sybil.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SybilRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilRoleGranted)
				if err := _Sybil.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) ParseRoleGranted(log types.Log) (*SybilRoleGranted, error) {
	event := new(SybilRoleGranted)
	if err := _Sybil.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Sybil contract.
type SybilRoleRevokedIterator struct {
	Event *SybilRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SybilRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilRoleRevoked)
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
		it.Event = new(SybilRoleRevoked)
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
func (it *SybilRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilRoleRevoked represents a RoleRevoked event raised by the Sybil contract.
type SybilRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SybilRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SybilRoleRevokedIterator{contract: _Sybil.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SybilRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilRoleRevoked)
				if err := _Sybil.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sybil *SybilFilterer) ParseRoleRevoked(log types.Log) (*SybilRoleRevoked, error) {
	event := new(SybilRoleRevoked)
	if err := _Sybil.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
