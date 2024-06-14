// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crrrng

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

// BigNumber is an auto generated low-level Go binding around an user-defined struct.
type BigNumber struct {
	Val    []byte
	Bitlen *big.Int
}

// VDFCRRNGCommitValue is an auto generated low-level Go binding around an user-defined struct.
type VDFCRRNGCommitValue struct {
	Commit          BigNumber
	OperatorAddress common.Address
}

// VDFCRRNGOperatorStatusAtRound is an auto generated low-level Go binding around an user-defined struct.
type VDFCRRNGOperatorStatusAtRound struct {
	CommitIndex *big.Int
	Committed   bool
}

// VDFCRRNGValueAtRound is an auto generated low-level Go binding around an user-defined struct.
type VDFCRRNGValueAtRound struct {
	StartTime     *big.Int
	CommitCounts  *big.Int
	Consumer      common.Address
	CommitsString []byte
	Omega         BigNumber
	Stage         uint8
	IsCompleted   bool
}

// CrrrngMetaData contains all meta data concerning the Crrrng contract.
var CrrrngMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"disputePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"avgRecoveOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premiumPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flatFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyLeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BigNumbers__ShouldNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CRRNGCoordinator_InsufficientDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CRRNGCoordinator_NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionInvalidAtThisStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCommittedParticipant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughOperators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughParticipated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStartedRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVerifiedAtTOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OmegaAlreadyCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OmegaNotCompleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecovNotMatchX\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SendFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ShouldNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillInCommitStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TOneNotAtLast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TwoOrMoreCommittedPleaseRecover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"XPrimeNotEqualAtIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"YPrimeNotEqualAtIndex\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"omega\",\"type\":\"bytes\"}],\"name\":\"CalculateOmega\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commitCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"commitVal\",\"type\":\"bytes\"}],\"name\":\"CommitC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recov\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"omega\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"}],\"name\":\"calculateDirectFundingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"disputeLeadershipAtRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateDirectFundingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getCommitValue\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"commit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"internalType\":\"structVDFCRRNG.CommitValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getCostAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getDisputeEndTimeAndIncentiveOfOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getDisputeEndTimeAndLeaderAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSetUpValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getUserStatusAtRound\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"commitIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"committed\",\"type\":\"bool\"}],\"internalType\":\"structVDFCRRNG.OperatorStatusAtRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getValuesAtRound\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitCounts\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commitsString\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"omega\",\"type\":\"tuple\"},{\"internalType\":\"enumVDFCRRNG.Stages\",\"name\":\"stage\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"}],\"internalType\":\"structVDFCRRNG.ValueAtRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber[]\",\"name\":\"v\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"x\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"y\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"operatorWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"reRequestRandomWordAtRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber[]\",\"name\":\"v\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"x\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"y\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"}],\"name\":\"requestRandomWordDirectFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"disputePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"avgRecoveOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premiumPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flatFee\",\"type\":\"uint256\"}],\"name\":\"setSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrrrngABI is the input ABI used to generate the binding from.
// Deprecated: Use CrrrngMetaData.ABI instead.
var CrrrngABI = CrrrngMetaData.ABI

// Crrrng is an auto generated Go binding around an Ethereum contract.
type Crrrng struct {
	CrrrngCaller     // Read-only binding to the contract
	CrrrngTransactor // Write-only binding to the contract
	CrrrngFilterer   // Log filterer for contract events
}

// CrrrngCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrrrngCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrrrngTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrrrngTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrrrngFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrrrngFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrrrngSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrrrngSession struct {
	Contract     *Crrrng           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrrrngCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrrrngCallerSession struct {
	Contract *CrrrngCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CrrrngTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrrrngTransactorSession struct {
	Contract     *CrrrngTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrrrngRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrrrngRaw struct {
	Contract *Crrrng // Generic contract binding to access the raw methods on
}

// CrrrngCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrrrngCallerRaw struct {
	Contract *CrrrngCaller // Generic read-only contract binding to access the raw methods on
}

// CrrrngTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrrrngTransactorRaw struct {
	Contract *CrrrngTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrrrng creates a new instance of Crrrng, bound to a specific deployed contract.
func NewCrrrng(address common.Address, backend bind.ContractBackend) (*Crrrng, error) {
	contract, err := bindCrrrng(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crrrng{CrrrngCaller: CrrrngCaller{contract: contract}, CrrrngTransactor: CrrrngTransactor{contract: contract}, CrrrngFilterer: CrrrngFilterer{contract: contract}}, nil
}

// NewCrrrngCaller creates a new read-only instance of Crrrng, bound to a specific deployed contract.
func NewCrrrngCaller(address common.Address, caller bind.ContractCaller) (*CrrrngCaller, error) {
	contract, err := bindCrrrng(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrrrngCaller{contract: contract}, nil
}

// NewCrrrngTransactor creates a new write-only instance of Crrrng, bound to a specific deployed contract.
func NewCrrrngTransactor(address common.Address, transactor bind.ContractTransactor) (*CrrrngTransactor, error) {
	contract, err := bindCrrrng(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrrrngTransactor{contract: contract}, nil
}

// NewCrrrngFilterer creates a new log filterer instance of Crrrng, bound to a specific deployed contract.
func NewCrrrngFilterer(address common.Address, filterer bind.ContractFilterer) (*CrrrngFilterer, error) {
	contract, err := bindCrrrng(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrrrngFilterer{contract: contract}, nil
}

// bindCrrrng binds a generic wrapper to an already deployed contract.
func bindCrrrng(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrrrngMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crrrng *CrrrngRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crrrng.Contract.CrrrngCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crrrng *CrrrngRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrng.Contract.CrrrngTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crrrng *CrrrngRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crrrng.Contract.CrrrngTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crrrng *CrrrngCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crrrng.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crrrng *CrrrngTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrng.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crrrng *CrrrngTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crrrng.Contract.contract.Transact(opts, method, params...)
}

// CalculateDirectFundingPrice is a free data retrieval call binding the contract method 0x2b5cb59a.
//
// Solidity: function calculateDirectFundingPrice(uint32 _callbackGasLimit) view returns(uint256)
func (_Crrrng *CrrrngCaller) CalculateDirectFundingPrice(opts *bind.CallOpts, _callbackGasLimit uint32) (*big.Int, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "calculateDirectFundingPrice", _callbackGasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDirectFundingPrice is a free data retrieval call binding the contract method 0x2b5cb59a.
//
// Solidity: function calculateDirectFundingPrice(uint32 _callbackGasLimit) view returns(uint256)
func (_Crrrng *CrrrngSession) CalculateDirectFundingPrice(_callbackGasLimit uint32) (*big.Int, error) {
	return _Crrrng.Contract.CalculateDirectFundingPrice(&_Crrrng.CallOpts, _callbackGasLimit)
}

// CalculateDirectFundingPrice is a free data retrieval call binding the contract method 0x2b5cb59a.
//
// Solidity: function calculateDirectFundingPrice(uint32 _callbackGasLimit) view returns(uint256)
func (_Crrrng *CrrrngCallerSession) CalculateDirectFundingPrice(_callbackGasLimit uint32) (*big.Int, error) {
	return _Crrrng.Contract.CalculateDirectFundingPrice(&_Crrrng.CallOpts, _callbackGasLimit)
}

// EstimateDirectFundingPrice is a free data retrieval call binding the contract method 0x6644cb44.
//
// Solidity: function estimateDirectFundingPrice(uint32 _callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Crrrng *CrrrngCaller) EstimateDirectFundingPrice(opts *bind.CallOpts, _callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "estimateDirectFundingPrice", _callbackGasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateDirectFundingPrice is a free data retrieval call binding the contract method 0x6644cb44.
//
// Solidity: function estimateDirectFundingPrice(uint32 _callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Crrrng *CrrrngSession) EstimateDirectFundingPrice(_callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Crrrng.Contract.EstimateDirectFundingPrice(&_Crrrng.CallOpts, _callbackGasLimit, gasPrice)
}

// EstimateDirectFundingPrice is a free data retrieval call binding the contract method 0x6644cb44.
//
// Solidity: function estimateDirectFundingPrice(uint32 _callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Crrrng *CrrrngCallerSession) EstimateDirectFundingPrice(_callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Crrrng.Contract.EstimateDirectFundingPrice(&_Crrrng.CallOpts, _callbackGasLimit, gasPrice)
}

// GetCommitValue is a free data retrieval call binding the contract method 0x20403ed6.
//
// Solidity: function getCommitValue(uint256 _round, uint256 _index) view returns(((bytes,uint256),address))
func (_Crrrng *CrrrngCaller) GetCommitValue(opts *bind.CallOpts, _round *big.Int, _index *big.Int) (VDFCRRNGCommitValue, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getCommitValue", _round, _index)

	if err != nil {
		return *new(VDFCRRNGCommitValue), err
	}

	out0 := *abi.ConvertType(out[0], new(VDFCRRNGCommitValue)).(*VDFCRRNGCommitValue)

	return out0, err

}

// GetCommitValue is a free data retrieval call binding the contract method 0x20403ed6.
//
// Solidity: function getCommitValue(uint256 _round, uint256 _index) view returns(((bytes,uint256),address))
func (_Crrrng *CrrrngSession) GetCommitValue(_round *big.Int, _index *big.Int) (VDFCRRNGCommitValue, error) {
	return _Crrrng.Contract.GetCommitValue(&_Crrrng.CallOpts, _round, _index)
}

// GetCommitValue is a free data retrieval call binding the contract method 0x20403ed6.
//
// Solidity: function getCommitValue(uint256 _round, uint256 _index) view returns(((bytes,uint256),address))
func (_Crrrng *CrrrngCallerSession) GetCommitValue(_round *big.Int, _index *big.Int) (VDFCRRNGCommitValue, error) {
	return _Crrrng.Contract.GetCommitValue(&_Crrrng.CallOpts, _round, _index)
}

// GetCostAtRound is a free data retrieval call binding the contract method 0x705ce484.
//
// Solidity: function getCostAtRound(uint256 round) view returns(uint256)
func (_Crrrng *CrrrngCaller) GetCostAtRound(opts *bind.CallOpts, round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getCostAtRound", round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCostAtRound is a free data retrieval call binding the contract method 0x705ce484.
//
// Solidity: function getCostAtRound(uint256 round) view returns(uint256)
func (_Crrrng *CrrrngSession) GetCostAtRound(round *big.Int) (*big.Int, error) {
	return _Crrrng.Contract.GetCostAtRound(&_Crrrng.CallOpts, round)
}

// GetCostAtRound is a free data retrieval call binding the contract method 0x705ce484.
//
// Solidity: function getCostAtRound(uint256 round) view returns(uint256)
func (_Crrrng *CrrrngCallerSession) GetCostAtRound(round *big.Int) (*big.Int, error) {
	return _Crrrng.Contract.GetCostAtRound(&_Crrrng.CallOpts, round)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address operator) view returns(uint256)
func (_Crrrng *CrrrngCaller) GetDepositAmount(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getDepositAmount", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address operator) view returns(uint256)
func (_Crrrng *CrrrngSession) GetDepositAmount(operator common.Address) (*big.Int, error) {
	return _Crrrng.Contract.GetDepositAmount(&_Crrrng.CallOpts, operator)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address operator) view returns(uint256)
func (_Crrrng *CrrrngCallerSession) GetDepositAmount(operator common.Address) (*big.Int, error) {
	return _Crrrng.Contract.GetDepositAmount(&_Crrrng.CallOpts, operator)
}

// GetDisputeEndTimeAndIncentiveOfOperator is a free data retrieval call binding the contract method 0x28fd65a9.
//
// Solidity: function getDisputeEndTimeAndIncentiveOfOperator(address operator) view returns(uint256, uint256)
func (_Crrrng *CrrrngCaller) GetDisputeEndTimeAndIncentiveOfOperator(opts *bind.CallOpts, operator common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getDisputeEndTimeAndIncentiveOfOperator", operator)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDisputeEndTimeAndIncentiveOfOperator is a free data retrieval call binding the contract method 0x28fd65a9.
//
// Solidity: function getDisputeEndTimeAndIncentiveOfOperator(address operator) view returns(uint256, uint256)
func (_Crrrng *CrrrngSession) GetDisputeEndTimeAndIncentiveOfOperator(operator common.Address) (*big.Int, *big.Int, error) {
	return _Crrrng.Contract.GetDisputeEndTimeAndIncentiveOfOperator(&_Crrrng.CallOpts, operator)
}

// GetDisputeEndTimeAndIncentiveOfOperator is a free data retrieval call binding the contract method 0x28fd65a9.
//
// Solidity: function getDisputeEndTimeAndIncentiveOfOperator(address operator) view returns(uint256, uint256)
func (_Crrrng *CrrrngCallerSession) GetDisputeEndTimeAndIncentiveOfOperator(operator common.Address) (*big.Int, *big.Int, error) {
	return _Crrrng.Contract.GetDisputeEndTimeAndIncentiveOfOperator(&_Crrrng.CallOpts, operator)
}

// GetDisputeEndTimeAndLeaderAtRound is a free data retrieval call binding the contract method 0x26fc6d61.
//
// Solidity: function getDisputeEndTimeAndLeaderAtRound(uint256 round) view returns(uint256, address)
func (_Crrrng *CrrrngCaller) GetDisputeEndTimeAndLeaderAtRound(opts *bind.CallOpts, round *big.Int) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getDisputeEndTimeAndLeaderAtRound", round)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetDisputeEndTimeAndLeaderAtRound is a free data retrieval call binding the contract method 0x26fc6d61.
//
// Solidity: function getDisputeEndTimeAndLeaderAtRound(uint256 round) view returns(uint256, address)
func (_Crrrng *CrrrngSession) GetDisputeEndTimeAndLeaderAtRound(round *big.Int) (*big.Int, common.Address, error) {
	return _Crrrng.Contract.GetDisputeEndTimeAndLeaderAtRound(&_Crrrng.CallOpts, round)
}

// GetDisputeEndTimeAndLeaderAtRound is a free data retrieval call binding the contract method 0x26fc6d61.
//
// Solidity: function getDisputeEndTimeAndLeaderAtRound(uint256 round) view returns(uint256, address)
func (_Crrrng *CrrrngCallerSession) GetDisputeEndTimeAndLeaderAtRound(round *big.Int) (*big.Int, common.Address, error) {
	return _Crrrng.Contract.GetDisputeEndTimeAndLeaderAtRound(&_Crrrng.CallOpts, round)
}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xde4303bc.
//
// Solidity: function getMinimumDepositAmount() view returns(uint256)
func (_Crrrng *CrrrngCaller) GetMinimumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getMinimumDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xde4303bc.
//
// Solidity: function getMinimumDepositAmount() view returns(uint256)
func (_Crrrng *CrrrngSession) GetMinimumDepositAmount() (*big.Int, error) {
	return _Crrrng.Contract.GetMinimumDepositAmount(&_Crrrng.CallOpts)
}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xde4303bc.
//
// Solidity: function getMinimumDepositAmount() view returns(uint256)
func (_Crrrng *CrrrngCallerSession) GetMinimumDepositAmount() (*big.Int, error) {
	return _Crrrng.Contract.GetMinimumDepositAmount(&_Crrrng.CallOpts)
}

// GetNextRound is a free data retrieval call binding the contract method 0x8011d1ed.
//
// Solidity: function getNextRound() view returns(uint256)
func (_Crrrng *CrrrngCaller) GetNextRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getNextRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextRound is a free data retrieval call binding the contract method 0x8011d1ed.
//
// Solidity: function getNextRound() view returns(uint256)
func (_Crrrng *CrrrngSession) GetNextRound() (*big.Int, error) {
	return _Crrrng.Contract.GetNextRound(&_Crrrng.CallOpts)
}

// GetNextRound is a free data retrieval call binding the contract method 0x8011d1ed.
//
// Solidity: function getNextRound() view returns(uint256)
func (_Crrrng *CrrrngCallerSession) GetNextRound() (*big.Int, error) {
	return _Crrrng.Contract.GetNextRound(&_Crrrng.CallOpts)
}

// GetSetUpValues is a free data retrieval call binding the contract method 0x4d541fd3.
//
// Solidity: function getSetUpValues() pure returns(uint256, uint256, uint256, uint256, bytes, bytes, bytes)
func (_Crrrng *CrrrngCaller) GetSetUpValues(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, []byte, []byte, []byte, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getSetUpValues")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new([]byte), *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	out5 := *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	out6 := *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetSetUpValues is a free data retrieval call binding the contract method 0x4d541fd3.
//
// Solidity: function getSetUpValues() pure returns(uint256, uint256, uint256, uint256, bytes, bytes, bytes)
func (_Crrrng *CrrrngSession) GetSetUpValues() (*big.Int, *big.Int, *big.Int, *big.Int, []byte, []byte, []byte, error) {
	return _Crrrng.Contract.GetSetUpValues(&_Crrrng.CallOpts)
}

// GetSetUpValues is a free data retrieval call binding the contract method 0x4d541fd3.
//
// Solidity: function getSetUpValues() pure returns(uint256, uint256, uint256, uint256, bytes, bytes, bytes)
func (_Crrrng *CrrrngCallerSession) GetSetUpValues() (*big.Int, *big.Int, *big.Int, *big.Int, []byte, []byte, []byte, error) {
	return _Crrrng.Contract.GetSetUpValues(&_Crrrng.CallOpts)
}

// GetUserStatusAtRound is a free data retrieval call binding the contract method 0xff235219.
//
// Solidity: function getUserStatusAtRound(address _operator, uint256 _round) view returns((uint256,bool))
func (_Crrrng *CrrrngCaller) GetUserStatusAtRound(opts *bind.CallOpts, _operator common.Address, _round *big.Int) (VDFCRRNGOperatorStatusAtRound, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getUserStatusAtRound", _operator, _round)

	if err != nil {
		return *new(VDFCRRNGOperatorStatusAtRound), err
	}

	out0 := *abi.ConvertType(out[0], new(VDFCRRNGOperatorStatusAtRound)).(*VDFCRRNGOperatorStatusAtRound)

	return out0, err

}

// GetUserStatusAtRound is a free data retrieval call binding the contract method 0xff235219.
//
// Solidity: function getUserStatusAtRound(address _operator, uint256 _round) view returns((uint256,bool))
func (_Crrrng *CrrrngSession) GetUserStatusAtRound(_operator common.Address, _round *big.Int) (VDFCRRNGOperatorStatusAtRound, error) {
	return _Crrrng.Contract.GetUserStatusAtRound(&_Crrrng.CallOpts, _operator, _round)
}

// GetUserStatusAtRound is a free data retrieval call binding the contract method 0xff235219.
//
// Solidity: function getUserStatusAtRound(address _operator, uint256 _round) view returns((uint256,bool))
func (_Crrrng *CrrrngCallerSession) GetUserStatusAtRound(_operator common.Address, _round *big.Int) (VDFCRRNGOperatorStatusAtRound, error) {
	return _Crrrng.Contract.GetUserStatusAtRound(&_Crrrng.CallOpts, _operator, _round)
}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,address,bytes,(bytes,uint256),uint8,bool))
func (_Crrrng *CrrrngCaller) GetValuesAtRound(opts *bind.CallOpts, _round *big.Int) (VDFCRRNGValueAtRound, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getValuesAtRound", _round)

	if err != nil {
		return *new(VDFCRRNGValueAtRound), err
	}

	out0 := *abi.ConvertType(out[0], new(VDFCRRNGValueAtRound)).(*VDFCRRNGValueAtRound)

	return out0, err

}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,address,bytes,(bytes,uint256),uint8,bool))
func (_Crrrng *CrrrngSession) GetValuesAtRound(_round *big.Int) (VDFCRRNGValueAtRound, error) {
	return _Crrrng.Contract.GetValuesAtRound(&_Crrrng.CallOpts, _round)
}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,address,bytes,(bytes,uint256),uint8,bool))
func (_Crrrng *CrrrngCallerSession) GetValuesAtRound(_round *big.Int) (VDFCRRNGValueAtRound, error) {
	return _Crrrng.Contract.GetValuesAtRound(&_Crrrng.CallOpts, _round)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crrrng *CrrrngCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crrrng *CrrrngSession) Owner() (common.Address, error) {
	return _Crrrng.Contract.Owner(&_Crrrng.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crrrng *CrrrngCallerSession) Owner() (common.Address, error) {
	return _Crrrng.Contract.Owner(&_Crrrng.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrng *CrrrngTransactor) Commit(opts *bind.TransactOpts, round *big.Int, c BigNumber) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "commit", round, c)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrng *CrrrngSession) Commit(round *big.Int, c BigNumber) (*types.Transaction, error) {
	return _Crrrng.Contract.Commit(&_Crrrng.TransactOpts, round, c)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrng *CrrrngTransactorSession) Commit(round *big.Int, c BigNumber) (*types.Transaction, error) {
	return _Crrrng.Contract.Commit(&_Crrrng.TransactOpts, round, c)
}

// DisputeLeadershipAtRound is a paid mutator transaction binding the contract method 0xda69041b.
//
// Solidity: function disputeLeadershipAtRound(uint256 round) returns()
func (_Crrrng *CrrrngTransactor) DisputeLeadershipAtRound(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "disputeLeadershipAtRound", round)
}

// DisputeLeadershipAtRound is a paid mutator transaction binding the contract method 0xda69041b.
//
// Solidity: function disputeLeadershipAtRound(uint256 round) returns()
func (_Crrrng *CrrrngSession) DisputeLeadershipAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.DisputeLeadershipAtRound(&_Crrrng.TransactOpts, round)
}

// DisputeLeadershipAtRound is a paid mutator transaction binding the contract method 0xda69041b.
//
// Solidity: function disputeLeadershipAtRound(uint256 round) returns()
func (_Crrrng *CrrrngTransactorSession) DisputeLeadershipAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.DisputeLeadershipAtRound(&_Crrrng.TransactOpts, round)
}

// Initialize is a paid mutator transaction binding the contract method 0x8efd09d1.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrng *CrrrngTransactor) Initialize(opts *bind.TransactOpts, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "initialize", v, x, y)
}

// Initialize is a paid mutator transaction binding the contract method 0x8efd09d1.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrng *CrrrngSession) Initialize(v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrng.Contract.Initialize(&_Crrrng.TransactOpts, v, x, y)
}

// Initialize is a paid mutator transaction binding the contract method 0x8efd09d1.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrng *CrrrngTransactorSession) Initialize(v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrng.Contract.Initialize(&_Crrrng.TransactOpts, v, x, y)
}

// OperatorDeposit is a paid mutator transaction binding the contract method 0xf23fa3d3.
//
// Solidity: function operatorDeposit() payable returns()
func (_Crrrng *CrrrngTransactor) OperatorDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "operatorDeposit")
}

// OperatorDeposit is a paid mutator transaction binding the contract method 0xf23fa3d3.
//
// Solidity: function operatorDeposit() payable returns()
func (_Crrrng *CrrrngSession) OperatorDeposit() (*types.Transaction, error) {
	return _Crrrng.Contract.OperatorDeposit(&_Crrrng.TransactOpts)
}

// OperatorDeposit is a paid mutator transaction binding the contract method 0xf23fa3d3.
//
// Solidity: function operatorDeposit() payable returns()
func (_Crrrng *CrrrngTransactorSession) OperatorDeposit() (*types.Transaction, error) {
	return _Crrrng.Contract.OperatorDeposit(&_Crrrng.TransactOpts)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0x9238a02f.
//
// Solidity: function operatorWithdraw(uint256 amount) returns()
func (_Crrrng *CrrrngTransactor) OperatorWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "operatorWithdraw", amount)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0x9238a02f.
//
// Solidity: function operatorWithdraw(uint256 amount) returns()
func (_Crrrng *CrrrngSession) OperatorWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.OperatorWithdraw(&_Crrrng.TransactOpts, amount)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0x9238a02f.
//
// Solidity: function operatorWithdraw(uint256 amount) returns()
func (_Crrrng *CrrrngTransactorSession) OperatorWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.OperatorWithdraw(&_Crrrng.TransactOpts, amount)
}

// ReRequestRandomWordAtRound is a paid mutator transaction binding the contract method 0x3b246b1d.
//
// Solidity: function reRequestRandomWordAtRound(uint256 round) returns()
func (_Crrrng *CrrrngTransactor) ReRequestRandomWordAtRound(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "reRequestRandomWordAtRound", round)
}

// ReRequestRandomWordAtRound is a paid mutator transaction binding the contract method 0x3b246b1d.
//
// Solidity: function reRequestRandomWordAtRound(uint256 round) returns()
func (_Crrrng *CrrrngSession) ReRequestRandomWordAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.ReRequestRandomWordAtRound(&_Crrrng.TransactOpts, round)
}

// ReRequestRandomWordAtRound is a paid mutator transaction binding the contract method 0x3b246b1d.
//
// Solidity: function reRequestRandomWordAtRound(uint256 round) returns()
func (_Crrrng *CrrrngTransactorSession) ReRequestRandomWordAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.ReRequestRandomWordAtRound(&_Crrrng.TransactOpts, round)
}

// Recover is a paid mutator transaction binding the contract method 0x37619135.
//
// Solidity: function recover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrng *CrrrngTransactor) Recover(opts *bind.TransactOpts, round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "recover", round, v, x, y)
}

// Recover is a paid mutator transaction binding the contract method 0x37619135.
//
// Solidity: function recover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrng *CrrrngSession) Recover(round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrng.Contract.Recover(&_Crrrng.TransactOpts, round, v, x, y)
}

// Recover is a paid mutator transaction binding the contract method 0x37619135.
//
// Solidity: function recover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrng *CrrrngTransactorSession) Recover(round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrng.Contract.Recover(&_Crrrng.TransactOpts, round, v, x, y)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crrrng *CrrrngTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crrrng *CrrrngSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crrrng.Contract.RenounceOwnership(&_Crrrng.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crrrng *CrrrngTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crrrng.Contract.RenounceOwnership(&_Crrrng.TransactOpts)
}

// RequestRandomWordDirectFunding is a paid mutator transaction binding the contract method 0x7072b977.
//
// Solidity: function requestRandomWordDirectFunding(uint32 callbackGasLimit) payable returns(uint256)
func (_Crrrng *CrrrngTransactor) RequestRandomWordDirectFunding(opts *bind.TransactOpts, callbackGasLimit uint32) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "requestRandomWordDirectFunding", callbackGasLimit)
}

// RequestRandomWordDirectFunding is a paid mutator transaction binding the contract method 0x7072b977.
//
// Solidity: function requestRandomWordDirectFunding(uint32 callbackGasLimit) payable returns(uint256)
func (_Crrrng *CrrrngSession) RequestRandomWordDirectFunding(callbackGasLimit uint32) (*types.Transaction, error) {
	return _Crrrng.Contract.RequestRandomWordDirectFunding(&_Crrrng.TransactOpts, callbackGasLimit)
}

// RequestRandomWordDirectFunding is a paid mutator transaction binding the contract method 0x7072b977.
//
// Solidity: function requestRandomWordDirectFunding(uint32 callbackGasLimit) payable returns(uint256)
func (_Crrrng *CrrrngTransactorSession) RequestRandomWordDirectFunding(callbackGasLimit uint32) (*types.Transaction, error) {
	return _Crrrng.Contract.RequestRandomWordDirectFunding(&_Crrrng.TransactOpts, callbackGasLimit)
}

// SetSettings is a paid mutator transaction binding the contract method 0x84fa32ac.
//
// Solidity: function setSettings(uint256 disputePeriod, uint256 minimumDepositAmount, uint256 avgRecoveOverhead, uint256 premiumPercentage, uint256 flatFee) returns()
func (_Crrrng *CrrrngTransactor) SetSettings(opts *bind.TransactOpts, disputePeriod *big.Int, minimumDepositAmount *big.Int, avgRecoveOverhead *big.Int, premiumPercentage *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "setSettings", disputePeriod, minimumDepositAmount, avgRecoveOverhead, premiumPercentage, flatFee)
}

// SetSettings is a paid mutator transaction binding the contract method 0x84fa32ac.
//
// Solidity: function setSettings(uint256 disputePeriod, uint256 minimumDepositAmount, uint256 avgRecoveOverhead, uint256 premiumPercentage, uint256 flatFee) returns()
func (_Crrrng *CrrrngSession) SetSettings(disputePeriod *big.Int, minimumDepositAmount *big.Int, avgRecoveOverhead *big.Int, premiumPercentage *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.SetSettings(&_Crrrng.TransactOpts, disputePeriod, minimumDepositAmount, avgRecoveOverhead, premiumPercentage, flatFee)
}

// SetSettings is a paid mutator transaction binding the contract method 0x84fa32ac.
//
// Solidity: function setSettings(uint256 disputePeriod, uint256 minimumDepositAmount, uint256 avgRecoveOverhead, uint256 premiumPercentage, uint256 flatFee) returns()
func (_Crrrng *CrrrngTransactorSession) SetSettings(disputePeriod *big.Int, minimumDepositAmount *big.Int, avgRecoveOverhead *big.Int, premiumPercentage *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.SetSettings(&_Crrrng.TransactOpts, disputePeriod, minimumDepositAmount, avgRecoveOverhead, premiumPercentage, flatFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crrrng *CrrrngTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crrrng *CrrrngSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crrrng.Contract.TransferOwnership(&_Crrrng.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crrrng *CrrrngTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crrrng.Contract.TransferOwnership(&_Crrrng.TransactOpts, newOwner)
}

// CrrrngCalculateOmegaIterator is returned from FilterCalculateOmega and is used to iterate over the raw logs and unpacked data for CalculateOmega events raised by the Crrrng contract.
type CrrrngCalculateOmegaIterator struct {
	Event *CrrrngCalculateOmega // Event containing the contract specifics and raw log

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
func (it *CrrrngCalculateOmegaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngCalculateOmega)
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
		it.Event = new(CrrrngCalculateOmega)
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
func (it *CrrrngCalculateOmegaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngCalculateOmegaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngCalculateOmega represents a CalculateOmega event raised by the Crrrng contract.
type CrrrngCalculateOmega struct {
	Round *big.Int
	Omega []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCalculateOmega is a free log retrieval operation binding the contract event 0xc91805177c38039c1218eaddf1b899b1f1b45e155afc9e054d4a5ce4115324ea.
//
// Solidity: event CalculateOmega(uint256 round, bytes omega)
func (_Crrrng *CrrrngFilterer) FilterCalculateOmega(opts *bind.FilterOpts) (*CrrrngCalculateOmegaIterator, error) {

	logs, sub, err := _Crrrng.contract.FilterLogs(opts, "CalculateOmega")
	if err != nil {
		return nil, err
	}
	return &CrrrngCalculateOmegaIterator{contract: _Crrrng.contract, event: "CalculateOmega", logs: logs, sub: sub}, nil
}

// WatchCalculateOmega is a free log subscription operation binding the contract event 0xc91805177c38039c1218eaddf1b899b1f1b45e155afc9e054d4a5ce4115324ea.
//
// Solidity: event CalculateOmega(uint256 round, bytes omega)
func (_Crrrng *CrrrngFilterer) WatchCalculateOmega(opts *bind.WatchOpts, sink chan<- *CrrrngCalculateOmega) (event.Subscription, error) {

	logs, sub, err := _Crrrng.contract.WatchLogs(opts, "CalculateOmega")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngCalculateOmega)
				if err := _Crrrng.contract.UnpackLog(event, "CalculateOmega", log); err != nil {
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

// ParseCalculateOmega is a log parse operation binding the contract event 0xc91805177c38039c1218eaddf1b899b1f1b45e155afc9e054d4a5ce4115324ea.
//
// Solidity: event CalculateOmega(uint256 round, bytes omega)
func (_Crrrng *CrrrngFilterer) ParseCalculateOmega(log types.Log) (*CrrrngCalculateOmega, error) {
	event := new(CrrrngCalculateOmega)
	if err := _Crrrng.contract.UnpackLog(event, "CalculateOmega", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngCommitCIterator is returned from FilterCommitC and is used to iterate over the raw logs and unpacked data for CommitC events raised by the Crrrng contract.
type CrrrngCommitCIterator struct {
	Event *CrrrngCommitC // Event containing the contract specifics and raw log

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
func (it *CrrrngCommitCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngCommitC)
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
		it.Event = new(CrrrngCommitC)
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
func (it *CrrrngCommitCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngCommitCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngCommitC represents a CommitC event raised by the Crrrng contract.
type CrrrngCommitC struct {
	CommitCount *big.Int
	CommitVal   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommitC is a free log retrieval operation binding the contract event 0xe2bd9c5fc4c79023527f33e9ed312ce0178f7adfc724f21a8bf76f43d93042c4.
//
// Solidity: event CommitC(uint256 commitCount, bytes commitVal)
func (_Crrrng *CrrrngFilterer) FilterCommitC(opts *bind.FilterOpts) (*CrrrngCommitCIterator, error) {

	logs, sub, err := _Crrrng.contract.FilterLogs(opts, "CommitC")
	if err != nil {
		return nil, err
	}
	return &CrrrngCommitCIterator{contract: _Crrrng.contract, event: "CommitC", logs: logs, sub: sub}, nil
}

// WatchCommitC is a free log subscription operation binding the contract event 0xe2bd9c5fc4c79023527f33e9ed312ce0178f7adfc724f21a8bf76f43d93042c4.
//
// Solidity: event CommitC(uint256 commitCount, bytes commitVal)
func (_Crrrng *CrrrngFilterer) WatchCommitC(opts *bind.WatchOpts, sink chan<- *CrrrngCommitC) (event.Subscription, error) {

	logs, sub, err := _Crrrng.contract.WatchLogs(opts, "CommitC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngCommitC)
				if err := _Crrrng.contract.UnpackLog(event, "CommitC", log); err != nil {
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

// ParseCommitC is a log parse operation binding the contract event 0xe2bd9c5fc4c79023527f33e9ed312ce0178f7adfc724f21a8bf76f43d93042c4.
//
// Solidity: event CommitC(uint256 commitCount, bytes commitVal)
func (_Crrrng *CrrrngFilterer) ParseCommitC(log types.Log) (*CrrrngCommitC, error) {
	event := new(CrrrngCommitC)
	if err := _Crrrng.contract.UnpackLog(event, "CommitC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crrrng contract.
type CrrrngOwnershipTransferredIterator struct {
	Event *CrrrngOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrrrngOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngOwnershipTransferred)
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
		it.Event = new(CrrrngOwnershipTransferred)
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
func (it *CrrrngOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngOwnershipTransferred represents a OwnershipTransferred event raised by the Crrrng contract.
type CrrrngOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crrrng *CrrrngFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrrrngOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crrrng.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrrrngOwnershipTransferredIterator{contract: _Crrrng.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crrrng *CrrrngFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrrrngOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crrrng.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngOwnershipTransferred)
				if err := _Crrrng.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Crrrng *CrrrngFilterer) ParseOwnershipTransferred(log types.Log) (*CrrrngOwnershipTransferred, error) {
	event := new(CrrrngOwnershipTransferred)
	if err := _Crrrng.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngRandomWordsRequestedIterator is returned from FilterRandomWordsRequested and is used to iterate over the raw logs and unpacked data for RandomWordsRequested events raised by the Crrrng contract.
type CrrrngRandomWordsRequestedIterator struct {
	Event *CrrrngRandomWordsRequested // Event containing the contract specifics and raw log

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
func (it *CrrrngRandomWordsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngRandomWordsRequested)
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
		it.Event = new(CrrrngRandomWordsRequested)
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
func (it *CrrrngRandomWordsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngRandomWordsRequested represents a RandomWordsRequested event raised by the Crrrng contract.
type CrrrngRandomWordsRequested struct {
	Round  *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsRequested is a free log retrieval operation binding the contract event 0x5afc92f47019f6f981dbe48fba0d8bc10298b161505d2ead5ad4b975bfde77f4.
//
// Solidity: event RandomWordsRequested(uint256 round, address sender)
func (_Crrrng *CrrrngFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts) (*CrrrngRandomWordsRequestedIterator, error) {

	logs, sub, err := _Crrrng.contract.FilterLogs(opts, "RandomWordsRequested")
	if err != nil {
		return nil, err
	}
	return &CrrrngRandomWordsRequestedIterator{contract: _Crrrng.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

// WatchRandomWordsRequested is a free log subscription operation binding the contract event 0x5afc92f47019f6f981dbe48fba0d8bc10298b161505d2ead5ad4b975bfde77f4.
//
// Solidity: event RandomWordsRequested(uint256 round, address sender)
func (_Crrrng *CrrrngFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *CrrrngRandomWordsRequested) (event.Subscription, error) {

	logs, sub, err := _Crrrng.contract.WatchLogs(opts, "RandomWordsRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngRandomWordsRequested)
				if err := _Crrrng.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
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

// ParseRandomWordsRequested is a log parse operation binding the contract event 0x5afc92f47019f6f981dbe48fba0d8bc10298b161505d2ead5ad4b975bfde77f4.
//
// Solidity: event RandomWordsRequested(uint256 round, address sender)
func (_Crrrng *CrrrngFilterer) ParseRandomWordsRequested(log types.Log) (*CrrrngRandomWordsRequested, error) {
	event := new(CrrrngRandomWordsRequested)
	if err := _Crrrng.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the Crrrng contract.
type CrrrngRecoveredIterator struct {
	Event *CrrrngRecovered // Event containing the contract specifics and raw log

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
func (it *CrrrngRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngRecovered)
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
		it.Event = new(CrrrngRecovered)
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
func (it *CrrrngRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngRecovered represents a Recovered event raised by the Crrrng contract.
type CrrrngRecovered struct {
	Round   *big.Int
	Recov   []byte
	Omega   []byte
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x132f03da24c1a79bae250687b1f230295572741fd5adb0c65b80d78b2d8df818.
//
// Solidity: event Recovered(uint256 round, bytes recov, bytes omega, bool success)
func (_Crrrng *CrrrngFilterer) FilterRecovered(opts *bind.FilterOpts) (*CrrrngRecoveredIterator, error) {

	logs, sub, err := _Crrrng.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &CrrrngRecoveredIterator{contract: _Crrrng.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x132f03da24c1a79bae250687b1f230295572741fd5adb0c65b80d78b2d8df818.
//
// Solidity: event Recovered(uint256 round, bytes recov, bytes omega, bool success)
func (_Crrrng *CrrrngFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *CrrrngRecovered) (event.Subscription, error) {

	logs, sub, err := _Crrrng.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngRecovered)
				if err := _Crrrng.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x132f03da24c1a79bae250687b1f230295572741fd5adb0c65b80d78b2d8df818.
//
// Solidity: event Recovered(uint256 round, bytes recov, bytes omega, bool success)
func (_Crrrng *CrrrngFilterer) ParseRecovered(log types.Log) (*CrrrngRecovered, error) {
	event := new(CrrrngRecovered)
	if err := _Crrrng.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
