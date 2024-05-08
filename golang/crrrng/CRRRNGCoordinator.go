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

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	C                  Struct0
	A                  Struct0
	ParticipantAddress common.Address
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Val    []byte
	Bitlen *big.Int
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	Index     *big.Int
	Committed bool
	Revealed  bool
}

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	StartTime          *big.Int
	NumOfPariticipants *big.Int
	Count              *big.Int
	Consumer           common.Address
	BStar              []byte
	CommitsString      []byte
	Omega              Struct0
	Stage              uint8
	IsCompleted        bool
	IsAllRevealed      bool
}

// CrrrngMetaData contains all meta data concerning the Crrrng contract.
var CrrrngMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"AlreadyCommitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BigNumbers__ShouldNotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FunctionInvalidAtThisStage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofsLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModExpRevealNotMatchCommit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoneParticipated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCommittedParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotStartedRound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotVerifiedAtTOne\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OmegaAlreadyCompleted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecovNotMatchX\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuard\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShouldNotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillInCommitStage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TOneNotAtLast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TwoOrMoreCommittedPleaseRecover\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"XPrimeNotEqualAtIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\"}]},{\"type\":\"error\",\"name\":\"YPrimeNotEqualAtIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\"}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"CalculateOmega\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"omega\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"CommitC\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"commitCount\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"commitVal\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"RandomWordsRequested\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\",\"indexed\":false},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Recovered\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"recov\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"omega\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"success\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"RevealA\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"revealLeftCount\",\"indexed\":false},{\"type\":\"bytes\",\"name\":\"aVal\",\"indexed\":false}]},{\"type\":\"function\",\"name\":\"calculateOmega\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"commit\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\"},{\"type\":\"tuple\",\"name\":\"c\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"getCommitRevealValues\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"_round\"},{\"type\":\"uint256\",\"name\":\"_index\"}],\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"components\":[{\"type\":\"tuple\",\"name\":\"c\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"tuple\",\"name\":\"a\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"address\",\"name\":\"participantAddress\"}]}]},{\"type\":\"function\",\"name\":\"getNextRound\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}]},{\"type\":\"function\",\"name\":\"getSetUpValues\",\"constant\":true,\"stateMutability\":\"pure\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"},{\"type\":\"uint256\",\"name\":\"\"},{\"type\":\"uint256\",\"name\":\"\"},{\"type\":\"uint256\",\"name\":\"\"},{\"type\":\"bytes\",\"name\":\"\"},{\"type\":\"bytes\",\"name\":\"\"},{\"type\":\"bytes\",\"name\":\"\"}]},{\"type\":\"function\",\"name\":\"getUserInfosAtRound\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"uint256\",\"name\":\"_round\"}],\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"components\":[{\"type\":\"uint256\",\"name\":\"index\"},{\"type\":\"bool\",\"name\":\"committed\"},{\"type\":\"bool\",\"name\":\"revealed\"}]}]},{\"type\":\"function\",\"name\":\"getValuesAtRound\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"_round\"}],\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"components\":[{\"type\":\"uint256\",\"name\":\"startTime\"},{\"type\":\"uint256\",\"name\":\"numOfPariticipants\"},{\"type\":\"uint256\",\"name\":\"count\"},{\"type\":\"address\",\"name\":\"consumer\"},{\"type\":\"bytes\",\"name\":\"bStar\"},{\"type\":\"bytes\",\"name\":\"commitsString\"},{\"type\":\"tuple\",\"name\":\"omega\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"uint8\",\"name\":\"stage\"},{\"type\":\"bool\",\"name\":\"isCompleted\"},{\"type\":\"bool\",\"name\":\"isAllRevealed\"}]}]},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"tuple[]\",\"name\":\"v\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"tuple\",\"name\":\"x\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"tuple\",\"name\":\"y\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"bytes\",\"name\":\"bigNumTwoPowerOfDelta\"},{\"type\":\"uint256\",\"name\":\"delta\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"reRequestRandomWordAtRound\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"recover\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\"},{\"type\":\"tuple[]\",\"name\":\"v\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"tuple\",\"name\":\"x\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"tuple\",\"name\":\"y\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]},{\"type\":\"bytes\",\"name\":\"bigNumTwoPowerOfDelta\"},{\"type\":\"uint256\",\"name\":\"delta\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"requestRandomWord\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}]},{\"type\":\"function\",\"name\":\"reveal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"round\"},{\"type\":\"tuple\",\"name\":\"a\",\"components\":[{\"type\":\"bytes\",\"name\":\"val\"},{\"type\":\"uint256\",\"name\":\"bitlen\"}]}],\"outputs\":[]}]",
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

// GetCommitRevealValues is a free data retrieval call binding the contract method 0xc5623103.
//
// Solidity: function getCommitRevealValues(uint256 _round, uint256 _index) view returns(((bytes,uint256),(bytes,uint256),address))
func (_Crrrng *CrrrngCaller) GetCommitRevealValues(opts *bind.CallOpts, _round *big.Int, _index *big.Int) (Struct1, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getCommitRevealValues", _round, _index)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetCommitRevealValues is a free data retrieval call binding the contract method 0xc5623103.
//
// Solidity: function getCommitRevealValues(uint256 _round, uint256 _index) view returns(((bytes,uint256),(bytes,uint256),address))
func (_Crrrng *CrrrngSession) GetCommitRevealValues(_round *big.Int, _index *big.Int) (Struct1, error) {
	return _Crrrng.Contract.GetCommitRevealValues(&_Crrrng.CallOpts, _round, _index)
}

// GetCommitRevealValues is a free data retrieval call binding the contract method 0xc5623103.
//
// Solidity: function getCommitRevealValues(uint256 _round, uint256 _index) view returns(((bytes,uint256),(bytes,uint256),address))
func (_Crrrng *CrrrngCallerSession) GetCommitRevealValues(_round *big.Int, _index *big.Int) (Struct1, error) {
	return _Crrrng.Contract.GetCommitRevealValues(&_Crrrng.CallOpts, _round, _index)
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

// GetUserInfosAtRound is a free data retrieval call binding the contract method 0xdf006e33.
//
// Solidity: function getUserInfosAtRound(address _owner, uint256 _round) view returns((uint256,bool,bool))
func (_Crrrng *CrrrngCaller) GetUserInfosAtRound(opts *bind.CallOpts, _owner common.Address, _round *big.Int) (Struct3, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getUserInfosAtRound", _owner, _round)

	if err != nil {
		return *new(Struct3), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct3)).(*Struct3)

	return out0, err

}

// GetUserInfosAtRound is a free data retrieval call binding the contract method 0xdf006e33.
//
// Solidity: function getUserInfosAtRound(address _owner, uint256 _round) view returns((uint256,bool,bool))
func (_Crrrng *CrrrngSession) GetUserInfosAtRound(_owner common.Address, _round *big.Int) (Struct3, error) {
	return _Crrrng.Contract.GetUserInfosAtRound(&_Crrrng.CallOpts, _owner, _round)
}

// GetUserInfosAtRound is a free data retrieval call binding the contract method 0xdf006e33.
//
// Solidity: function getUserInfosAtRound(address _owner, uint256 _round) view returns((uint256,bool,bool))
func (_Crrrng *CrrrngCallerSession) GetUserInfosAtRound(_owner common.Address, _round *big.Int) (Struct3, error) {
	return _Crrrng.Contract.GetUserInfosAtRound(&_Crrrng.CallOpts, _owner, _round)
}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,uint256,address,bytes,bytes,(bytes,uint256),uint8,bool,bool))
func (_Crrrng *CrrrngCaller) GetValuesAtRound(opts *bind.CallOpts, _round *big.Int) (Struct2, error) {
	var out []interface{}
	err := _Crrrng.contract.Call(opts, &out, "getValuesAtRound", _round)

	if err != nil {
		return *new(Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct2)).(*Struct2)

	return out0, err

}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,uint256,address,bytes,bytes,(bytes,uint256),uint8,bool,bool))
func (_Crrrng *CrrrngSession) GetValuesAtRound(_round *big.Int) (Struct2, error) {
	return _Crrrng.Contract.GetValuesAtRound(&_Crrrng.CallOpts, _round)
}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,uint256,address,bytes,bytes,(bytes,uint256),uint8,bool,bool))
func (_Crrrng *CrrrngCallerSession) GetValuesAtRound(_round *big.Int) (Struct2, error) {
	return _Crrrng.Contract.GetValuesAtRound(&_Crrrng.CallOpts, _round)
}

// CalculateOmega is a paid mutator transaction binding the contract method 0x371b28db.
//
// Solidity: function calculateOmega(uint256 round) returns()
func (_Crrrng *CrrrngTransactor) CalculateOmega(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "calculateOmega", round)
}

// CalculateOmega is a paid mutator transaction binding the contract method 0x371b28db.
//
// Solidity: function calculateOmega(uint256 round) returns()
func (_Crrrng *CrrrngSession) CalculateOmega(round *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.CalculateOmega(&_Crrrng.TransactOpts, round)
}

// CalculateOmega is a paid mutator transaction binding the contract method 0x371b28db.
//
// Solidity: function calculateOmega(uint256 round) returns()
func (_Crrrng *CrrrngTransactorSession) CalculateOmega(round *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.CalculateOmega(&_Crrrng.TransactOpts, round)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrng *CrrrngTransactor) Commit(opts *bind.TransactOpts, round *big.Int, c Struct0) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "commit", round, c)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrng *CrrrngSession) Commit(round *big.Int, c Struct0) (*types.Transaction, error) {
	return _Crrrng.Contract.Commit(&_Crrrng.TransactOpts, round, c)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrng *CrrrngTransactorSession) Commit(round *big.Int, c Struct0) (*types.Transaction, error) {
	return _Crrrng.Contract.Commit(&_Crrrng.TransactOpts, round, c)
}

// Initialize is a paid mutator transaction binding the contract method 0xcfcde6a2.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y, bytes bigNumTwoPowerOfDelta, uint256 delta) returns()
func (_Crrrng *CrrrngTransactor) Initialize(opts *bind.TransactOpts, v []Struct0, x Struct0, y Struct0, bigNumTwoPowerOfDelta []byte, delta *big.Int) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "initialize", v, x, y, bigNumTwoPowerOfDelta, delta)
}

// Initialize is a paid mutator transaction binding the contract method 0xcfcde6a2.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y, bytes bigNumTwoPowerOfDelta, uint256 delta) returns()
func (_Crrrng *CrrrngSession) Initialize(v []Struct0, x Struct0, y Struct0, bigNumTwoPowerOfDelta []byte, delta *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.Initialize(&_Crrrng.TransactOpts, v, x, y, bigNumTwoPowerOfDelta, delta)
}

// Initialize is a paid mutator transaction binding the contract method 0xcfcde6a2.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y, bytes bigNumTwoPowerOfDelta, uint256 delta) returns()
func (_Crrrng *CrrrngTransactorSession) Initialize(v []Struct0, x Struct0, y Struct0, bigNumTwoPowerOfDelta []byte, delta *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.Initialize(&_Crrrng.TransactOpts, v, x, y, bigNumTwoPowerOfDelta, delta)
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

// Recover is a paid mutator transaction binding the contract method 0x9ad0c6a6.
//
// Solidity: function recover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y, bytes bigNumTwoPowerOfDelta, uint256 delta) returns()
func (_Crrrng *CrrrngTransactor) Recover(opts *bind.TransactOpts, round *big.Int, v []Struct0, x Struct0, y Struct0, bigNumTwoPowerOfDelta []byte, delta *big.Int) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "recover", round, v, x, y, bigNumTwoPowerOfDelta, delta)
}

// Recover is a paid mutator transaction binding the contract method 0x9ad0c6a6.
//
// Solidity: function recover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y, bytes bigNumTwoPowerOfDelta, uint256 delta) returns()
func (_Crrrng *CrrrngSession) Recover(round *big.Int, v []Struct0, x Struct0, y Struct0, bigNumTwoPowerOfDelta []byte, delta *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.Recover(&_Crrrng.TransactOpts, round, v, x, y, bigNumTwoPowerOfDelta, delta)
}

// Recover is a paid mutator transaction binding the contract method 0x9ad0c6a6.
//
// Solidity: function recover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y, bytes bigNumTwoPowerOfDelta, uint256 delta) returns()
func (_Crrrng *CrrrngTransactorSession) Recover(round *big.Int, v []Struct0, x Struct0, y Struct0, bigNumTwoPowerOfDelta []byte, delta *big.Int) (*types.Transaction, error) {
	return _Crrrng.Contract.Recover(&_Crrrng.TransactOpts, round, v, x, y, bigNumTwoPowerOfDelta, delta)
}

// RequestRandomWord is a paid mutator transaction binding the contract method 0x0a01fb9e.
//
// Solidity: function requestRandomWord() returns(uint256)
func (_Crrrng *CrrrngTransactor) RequestRandomWord(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "requestRandomWord")
}

// RequestRandomWord is a paid mutator transaction binding the contract method 0x0a01fb9e.
//
// Solidity: function requestRandomWord() returns(uint256)
func (_Crrrng *CrrrngSession) RequestRandomWord() (*types.Transaction, error) {
	return _Crrrng.Contract.RequestRandomWord(&_Crrrng.TransactOpts)
}

// RequestRandomWord is a paid mutator transaction binding the contract method 0x0a01fb9e.
//
// Solidity: function requestRandomWord() returns(uint256)
func (_Crrrng *CrrrngTransactorSession) RequestRandomWord() (*types.Transaction, error) {
	return _Crrrng.Contract.RequestRandomWord(&_Crrrng.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0x170a3a76.
//
// Solidity: function reveal(uint256 round, (bytes,uint256) a) returns()
func (_Crrrng *CrrrngTransactor) Reveal(opts *bind.TransactOpts, round *big.Int, a Struct0) (*types.Transaction, error) {
	return _Crrrng.contract.Transact(opts, "reveal", round, a)
}

// Reveal is a paid mutator transaction binding the contract method 0x170a3a76.
//
// Solidity: function reveal(uint256 round, (bytes,uint256) a) returns()
func (_Crrrng *CrrrngSession) Reveal(round *big.Int, a Struct0) (*types.Transaction, error) {
	return _Crrrng.Contract.Reveal(&_Crrrng.TransactOpts, round, a)
}

// Reveal is a paid mutator transaction binding the contract method 0x170a3a76.
//
// Solidity: function reveal(uint256 round, (bytes,uint256) a) returns()
func (_Crrrng *CrrrngTransactorSession) Reveal(round *big.Int, a Struct0) (*types.Transaction, error) {
	return _Crrrng.Contract.Reveal(&_Crrrng.TransactOpts, round, a)
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

// CrrrngRevealAIterator is returned from FilterRevealA and is used to iterate over the raw logs and unpacked data for RevealA events raised by the Crrrng contract.
type CrrrngRevealAIterator struct {
	Event *CrrrngRevealA // Event containing the contract specifics and raw log

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
func (it *CrrrngRevealAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngRevealA)
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
		it.Event = new(CrrrngRevealA)
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
func (it *CrrrngRevealAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngRevealAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngRevealA represents a RevealA event raised by the Crrrng contract.
type CrrrngRevealA struct {
	RevealLeftCount *big.Int
	AVal            []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRevealA is a free log retrieval operation binding the contract event 0xe34b79a7cd99d71a223e65f8beb51ff3e4e3d556129abf2216b3207fac7d17c2.
//
// Solidity: event RevealA(uint256 revealLeftCount, bytes aVal)
func (_Crrrng *CrrrngFilterer) FilterRevealA(opts *bind.FilterOpts) (*CrrrngRevealAIterator, error) {

	logs, sub, err := _Crrrng.contract.FilterLogs(opts, "RevealA")
	if err != nil {
		return nil, err
	}
	return &CrrrngRevealAIterator{contract: _Crrrng.contract, event: "RevealA", logs: logs, sub: sub}, nil
}

// WatchRevealA is a free log subscription operation binding the contract event 0xe34b79a7cd99d71a223e65f8beb51ff3e4e3d556129abf2216b3207fac7d17c2.
//
// Solidity: event RevealA(uint256 revealLeftCount, bytes aVal)
func (_Crrrng *CrrrngFilterer) WatchRevealA(opts *bind.WatchOpts, sink chan<- *CrrrngRevealA) (event.Subscription, error) {

	logs, sub, err := _Crrrng.contract.WatchLogs(opts, "RevealA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngRevealA)
				if err := _Crrrng.contract.UnpackLog(event, "RevealA", log); err != nil {
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

// ParseRevealA is a log parse operation binding the contract event 0xe34b79a7cd99d71a223e65f8beb51ff3e4e3d556129abf2216b3207fac7d17c2.
//
// Solidity: event RevealA(uint256 revealLeftCount, bytes aVal)
func (_Crrrng *CrrrngFilterer) ParseRevealA(log types.Log) (*CrrrngRevealA, error) {
	event := new(CrrrngRevealA)
	if err := _Crrrng.contract.UnpackLog(event, "RevealA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
