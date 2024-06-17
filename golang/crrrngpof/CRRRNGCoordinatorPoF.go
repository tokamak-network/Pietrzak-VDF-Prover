// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crrrngpof

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

// VDFCRRNGPoFCommitValue is an auto generated low-level Go binding around an user-defined struct.
type VDFCRRNGPoFCommitValue struct {
	Commit          BigNumber
	OperatorAddress common.Address
}

// VDFCRRNGPoFFulfillStatus is an auto generated low-level Go binding around an user-defined struct.
type VDFCRRNGPoFFulfillStatus struct {
	Executed  bool
	Succeeded bool
}

// VDFCRRNGPoFOperatorStatusAtRound is an auto generated low-level Go binding around an user-defined struct.
type VDFCRRNGPoFOperatorStatusAtRound struct {
	CommitIndex *big.Int
	Committed   bool
}

// VDFCRRNGPoFValueAtRound is an auto generated low-level Go binding around an user-defined struct.
type VDFCRRNGPoFValueAtRound struct {
	StartTime     *big.Int
	RequestedTime *big.Int
	CommitCounts  *big.Int
	Consumer      common.Address
	CommitsString []byte
	Omega         BigNumber
	Stage         uint8
	IsCompleted   bool
	IsVerified    bool
}

// CrrrngpofMetaData contains all meta data concerning the Crrrngpof contract.
var CrrrngpofMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"disputePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"avgL2GasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"avgL1GasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premiumPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penaltyPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flatFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyLeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySucceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BigNumbers__ShouldNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CRRNGCoordinator_InsufficientDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CRRNGCoordinator_NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotEndedOrStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionInvalidAtThisStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCommittedParticipant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughOperators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughParticipated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFulfilledExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStartedRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVerifiedAtTOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OmegaAlreadyCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OmegaNotCompleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreviousRoundNotRecovered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecovNotMatchX\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SendFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ShouldNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillInCommitStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmittedSameOmega\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooEarlyToRefund\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TwoOrMoreCommittedPleaseRecover\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commitCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"commitVal\",\"type\":\"bytes\"}],\"name\":\"CommitC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hashedOmega\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"leader\",\"type\":\"address\"}],\"name\":\"FulfillRandomness\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"omega\",\"type\":\"bytes\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"}],\"name\":\"calculateDirectFundingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"disputeLeadershipAtRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber[]\",\"name\":\"v\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"x\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"y\",\"type\":\"tuple\"}],\"name\":\"disputeRecover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateDirectFundingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"fulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"fulfillRandomnessOfFailed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getCommitValue\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"commit\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"internalType\":\"structVDFCRRNGPoF.CommitValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getCommittedOperatorsAtRound\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getConsumerAtRound\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getCostAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getDisputeEndTimeAndLeaderAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getDisputeEndTimeOfOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDisputePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getFulfillStatusAtRound\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"succeeded\",\"type\":\"bool\"}],\"internalType\":\"structVDFCRRNGPoF.FulfillStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSetUpValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getUserStatusAtRound\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"commitIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"committed\",\"type\":\"bool\"}],\"internalType\":\"structVDFCRRNGPoF.OperatorStatusAtRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getValuesAtRound\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitCounts\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"commitsString\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"omega\",\"type\":\"tuple\"},{\"internalType\":\"enumVDFCRRNGPoF.Stages\",\"name\":\"stage\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"}],\"internalType\":\"structVDFCRRNGPoF.ValueAtRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber[]\",\"name\":\"v\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"x\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"y\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFulfilledRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRecoveredRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"operatorWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"reRequestRandomWordAtRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bitlen\",\"type\":\"uint256\"}],\"internalType\":\"structBigNumber\",\"name\":\"y\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"refundAtRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"}],\"name\":\"requestRandomWordDirectFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"disputePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"avgL2GasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"avgL1GasUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premiumPercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flatFee\",\"type\":\"uint256\"}],\"name\":\"setSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrrrngpofABI is the input ABI used to generate the binding from.
// Deprecated: Use CrrrngpofMetaData.ABI instead.
var CrrrngpofABI = CrrrngpofMetaData.ABI

// Crrrngpof is an auto generated Go binding around an Ethereum contract.
type Crrrngpof struct {
	CrrrngpofCaller     // Read-only binding to the contract
	CrrrngpofTransactor // Write-only binding to the contract
	CrrrngpofFilterer   // Log filterer for contract events
}

// CrrrngpofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrrrngpofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrrrngpofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrrrngpofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrrrngpofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrrrngpofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrrrngpofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrrrngpofSession struct {
	Contract     *Crrrngpof        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrrrngpofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrrrngpofCallerSession struct {
	Contract *CrrrngpofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrrrngpofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrrrngpofTransactorSession struct {
	Contract     *CrrrngpofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrrrngpofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrrrngpofRaw struct {
	Contract *Crrrngpof // Generic contract binding to access the raw methods on
}

// CrrrngpofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrrrngpofCallerRaw struct {
	Contract *CrrrngpofCaller // Generic read-only contract binding to access the raw methods on
}

// CrrrngpofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrrrngpofTransactorRaw struct {
	Contract *CrrrngpofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrrrngpof creates a new instance of Crrrngpof, bound to a specific deployed contract.
func NewCrrrngpof(address common.Address, backend bind.ContractBackend) (*Crrrngpof, error) {
	contract, err := bindCrrrngpof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crrrngpof{CrrrngpofCaller: CrrrngpofCaller{contract: contract}, CrrrngpofTransactor: CrrrngpofTransactor{contract: contract}, CrrrngpofFilterer: CrrrngpofFilterer{contract: contract}}, nil
}

// NewCrrrngpofCaller creates a new read-only instance of Crrrngpof, bound to a specific deployed contract.
func NewCrrrngpofCaller(address common.Address, caller bind.ContractCaller) (*CrrrngpofCaller, error) {
	contract, err := bindCrrrngpof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrrrngpofCaller{contract: contract}, nil
}

// NewCrrrngpofTransactor creates a new write-only instance of Crrrngpof, bound to a specific deployed contract.
func NewCrrrngpofTransactor(address common.Address, transactor bind.ContractTransactor) (*CrrrngpofTransactor, error) {
	contract, err := bindCrrrngpof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrrrngpofTransactor{contract: contract}, nil
}

// NewCrrrngpofFilterer creates a new log filterer instance of Crrrngpof, bound to a specific deployed contract.
func NewCrrrngpofFilterer(address common.Address, filterer bind.ContractFilterer) (*CrrrngpofFilterer, error) {
	contract, err := bindCrrrngpof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrrrngpofFilterer{contract: contract}, nil
}

// bindCrrrngpof binds a generic wrapper to an already deployed contract.
func bindCrrrngpof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrrrngpofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crrrngpof *CrrrngpofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crrrngpof.Contract.CrrrngpofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crrrngpof *CrrrngpofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrngpof.Contract.CrrrngpofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crrrngpof *CrrrngpofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crrrngpof.Contract.CrrrngpofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crrrngpof *CrrrngpofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crrrngpof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crrrngpof *CrrrngpofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrngpof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crrrngpof *CrrrngpofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crrrngpof.Contract.contract.Transact(opts, method, params...)
}

// CalculateDirectFundingPrice is a free data retrieval call binding the contract method 0x2b5cb59a.
//
// Solidity: function calculateDirectFundingPrice(uint32 _callbackGasLimit) view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) CalculateDirectFundingPrice(opts *bind.CallOpts, _callbackGasLimit uint32) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "calculateDirectFundingPrice", _callbackGasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDirectFundingPrice is a free data retrieval call binding the contract method 0x2b5cb59a.
//
// Solidity: function calculateDirectFundingPrice(uint32 _callbackGasLimit) view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) CalculateDirectFundingPrice(_callbackGasLimit uint32) (*big.Int, error) {
	return _Crrrngpof.Contract.CalculateDirectFundingPrice(&_Crrrngpof.CallOpts, _callbackGasLimit)
}

// CalculateDirectFundingPrice is a free data retrieval call binding the contract method 0x2b5cb59a.
//
// Solidity: function calculateDirectFundingPrice(uint32 _callbackGasLimit) view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) CalculateDirectFundingPrice(_callbackGasLimit uint32) (*big.Int, error) {
	return _Crrrngpof.Contract.CalculateDirectFundingPrice(&_Crrrngpof.CallOpts, _callbackGasLimit)
}

// EstimateDirectFundingPrice is a free data retrieval call binding the contract method 0x6644cb44.
//
// Solidity: function estimateDirectFundingPrice(uint32 _callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) EstimateDirectFundingPrice(opts *bind.CallOpts, _callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "estimateDirectFundingPrice", _callbackGasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateDirectFundingPrice is a free data retrieval call binding the contract method 0x6644cb44.
//
// Solidity: function estimateDirectFundingPrice(uint32 _callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) EstimateDirectFundingPrice(_callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Crrrngpof.Contract.EstimateDirectFundingPrice(&_Crrrngpof.CallOpts, _callbackGasLimit, gasPrice)
}

// EstimateDirectFundingPrice is a free data retrieval call binding the contract method 0x6644cb44.
//
// Solidity: function estimateDirectFundingPrice(uint32 _callbackGasLimit, uint256 gasPrice) view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) EstimateDirectFundingPrice(_callbackGasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Crrrngpof.Contract.EstimateDirectFundingPrice(&_Crrrngpof.CallOpts, _callbackGasLimit, gasPrice)
}

// GetCommitValue is a free data retrieval call binding the contract method 0x20403ed6.
//
// Solidity: function getCommitValue(uint256 _round, uint256 _index) view returns(((bytes,uint256),address))
func (_Crrrngpof *CrrrngpofCaller) GetCommitValue(opts *bind.CallOpts, _round *big.Int, _index *big.Int) (VDFCRRNGPoFCommitValue, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getCommitValue", _round, _index)

	if err != nil {
		return *new(VDFCRRNGPoFCommitValue), err
	}

	out0 := *abi.ConvertType(out[0], new(VDFCRRNGPoFCommitValue)).(*VDFCRRNGPoFCommitValue)

	return out0, err

}

// GetCommitValue is a free data retrieval call binding the contract method 0x20403ed6.
//
// Solidity: function getCommitValue(uint256 _round, uint256 _index) view returns(((bytes,uint256),address))
func (_Crrrngpof *CrrrngpofSession) GetCommitValue(_round *big.Int, _index *big.Int) (VDFCRRNGPoFCommitValue, error) {
	return _Crrrngpof.Contract.GetCommitValue(&_Crrrngpof.CallOpts, _round, _index)
}

// GetCommitValue is a free data retrieval call binding the contract method 0x20403ed6.
//
// Solidity: function getCommitValue(uint256 _round, uint256 _index) view returns(((bytes,uint256),address))
func (_Crrrngpof *CrrrngpofCallerSession) GetCommitValue(_round *big.Int, _index *big.Int) (VDFCRRNGPoFCommitValue, error) {
	return _Crrrngpof.Contract.GetCommitValue(&_Crrrngpof.CallOpts, _round, _index)
}

// GetCommittedOperatorsAtRound is a free data retrieval call binding the contract method 0xfd1f1f7f.
//
// Solidity: function getCommittedOperatorsAtRound(uint256 round) view returns(address[])
func (_Crrrngpof *CrrrngpofCaller) GetCommittedOperatorsAtRound(opts *bind.CallOpts, round *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getCommittedOperatorsAtRound", round)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCommittedOperatorsAtRound is a free data retrieval call binding the contract method 0xfd1f1f7f.
//
// Solidity: function getCommittedOperatorsAtRound(uint256 round) view returns(address[])
func (_Crrrngpof *CrrrngpofSession) GetCommittedOperatorsAtRound(round *big.Int) ([]common.Address, error) {
	return _Crrrngpof.Contract.GetCommittedOperatorsAtRound(&_Crrrngpof.CallOpts, round)
}

// GetCommittedOperatorsAtRound is a free data retrieval call binding the contract method 0xfd1f1f7f.
//
// Solidity: function getCommittedOperatorsAtRound(uint256 round) view returns(address[])
func (_Crrrngpof *CrrrngpofCallerSession) GetCommittedOperatorsAtRound(round *big.Int) ([]common.Address, error) {
	return _Crrrngpof.Contract.GetCommittedOperatorsAtRound(&_Crrrngpof.CallOpts, round)
}

// GetConsumerAtRound is a free data retrieval call binding the contract method 0x8c0822d3.
//
// Solidity: function getConsumerAtRound(uint256 round) view returns(address)
func (_Crrrngpof *CrrrngpofCaller) GetConsumerAtRound(opts *bind.CallOpts, round *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getConsumerAtRound", round)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConsumerAtRound is a free data retrieval call binding the contract method 0x8c0822d3.
//
// Solidity: function getConsumerAtRound(uint256 round) view returns(address)
func (_Crrrngpof *CrrrngpofSession) GetConsumerAtRound(round *big.Int) (common.Address, error) {
	return _Crrrngpof.Contract.GetConsumerAtRound(&_Crrrngpof.CallOpts, round)
}

// GetConsumerAtRound is a free data retrieval call binding the contract method 0x8c0822d3.
//
// Solidity: function getConsumerAtRound(uint256 round) view returns(address)
func (_Crrrngpof *CrrrngpofCallerSession) GetConsumerAtRound(round *big.Int) (common.Address, error) {
	return _Crrrngpof.Contract.GetConsumerAtRound(&_Crrrngpof.CallOpts, round)
}

// GetCostAtRound is a free data retrieval call binding the contract method 0x705ce484.
//
// Solidity: function getCostAtRound(uint256 round) view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) GetCostAtRound(opts *bind.CallOpts, round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getCostAtRound", round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCostAtRound is a free data retrieval call binding the contract method 0x705ce484.
//
// Solidity: function getCostAtRound(uint256 round) view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) GetCostAtRound(round *big.Int) (*big.Int, error) {
	return _Crrrngpof.Contract.GetCostAtRound(&_Crrrngpof.CallOpts, round)
}

// GetCostAtRound is a free data retrieval call binding the contract method 0x705ce484.
//
// Solidity: function getCostAtRound(uint256 round) view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetCostAtRound(round *big.Int) (*big.Int, error) {
	return _Crrrngpof.Contract.GetCostAtRound(&_Crrrngpof.CallOpts, round)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address operator) view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) GetDepositAmount(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getDepositAmount", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address operator) view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) GetDepositAmount(operator common.Address) (*big.Int, error) {
	return _Crrrngpof.Contract.GetDepositAmount(&_Crrrngpof.CallOpts, operator)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address operator) view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetDepositAmount(operator common.Address) (*big.Int, error) {
	return _Crrrngpof.Contract.GetDepositAmount(&_Crrrngpof.CallOpts, operator)
}

// GetDisputeEndTimeAndLeaderAtRound is a free data retrieval call binding the contract method 0x26fc6d61.
//
// Solidity: function getDisputeEndTimeAndLeaderAtRound(uint256 round) view returns(uint256, address)
func (_Crrrngpof *CrrrngpofCaller) GetDisputeEndTimeAndLeaderAtRound(opts *bind.CallOpts, round *big.Int) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getDisputeEndTimeAndLeaderAtRound", round)

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
func (_Crrrngpof *CrrrngpofSession) GetDisputeEndTimeAndLeaderAtRound(round *big.Int) (*big.Int, common.Address, error) {
	return _Crrrngpof.Contract.GetDisputeEndTimeAndLeaderAtRound(&_Crrrngpof.CallOpts, round)
}

// GetDisputeEndTimeAndLeaderAtRound is a free data retrieval call binding the contract method 0x26fc6d61.
//
// Solidity: function getDisputeEndTimeAndLeaderAtRound(uint256 round) view returns(uint256, address)
func (_Crrrngpof *CrrrngpofCallerSession) GetDisputeEndTimeAndLeaderAtRound(round *big.Int) (*big.Int, common.Address, error) {
	return _Crrrngpof.Contract.GetDisputeEndTimeAndLeaderAtRound(&_Crrrngpof.CallOpts, round)
}

// GetDisputeEndTimeOfOperator is a free data retrieval call binding the contract method 0x42dfbb15.
//
// Solidity: function getDisputeEndTimeOfOperator(address operator) view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) GetDisputeEndTimeOfOperator(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getDisputeEndTimeOfOperator", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeEndTimeOfOperator is a free data retrieval call binding the contract method 0x42dfbb15.
//
// Solidity: function getDisputeEndTimeOfOperator(address operator) view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) GetDisputeEndTimeOfOperator(operator common.Address) (*big.Int, error) {
	return _Crrrngpof.Contract.GetDisputeEndTimeOfOperator(&_Crrrngpof.CallOpts, operator)
}

// GetDisputeEndTimeOfOperator is a free data retrieval call binding the contract method 0x42dfbb15.
//
// Solidity: function getDisputeEndTimeOfOperator(address operator) view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetDisputeEndTimeOfOperator(operator common.Address) (*big.Int, error) {
	return _Crrrngpof.Contract.GetDisputeEndTimeOfOperator(&_Crrrngpof.CallOpts, operator)
}

// GetDisputePeriod is a free data retrieval call binding the contract method 0x5aea0ec4.
//
// Solidity: function getDisputePeriod() view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) GetDisputePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getDisputePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputePeriod is a free data retrieval call binding the contract method 0x5aea0ec4.
//
// Solidity: function getDisputePeriod() view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) GetDisputePeriod() (*big.Int, error) {
	return _Crrrngpof.Contract.GetDisputePeriod(&_Crrrngpof.CallOpts)
}

// GetDisputePeriod is a free data retrieval call binding the contract method 0x5aea0ec4.
//
// Solidity: function getDisputePeriod() view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetDisputePeriod() (*big.Int, error) {
	return _Crrrngpof.Contract.GetDisputePeriod(&_Crrrngpof.CallOpts)
}

// GetFeeSettings is a free data retrieval call binding the contract method 0x2b38400e.
//
// Solidity: function getFeeSettings() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Crrrngpof *CrrrngpofCaller) GetFeeSettings(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getFeeSettings")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetFeeSettings is a free data retrieval call binding the contract method 0x2b38400e.
//
// Solidity: function getFeeSettings() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Crrrngpof *CrrrngpofSession) GetFeeSettings() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Crrrngpof.Contract.GetFeeSettings(&_Crrrngpof.CallOpts)
}

// GetFeeSettings is a free data retrieval call binding the contract method 0x2b38400e.
//
// Solidity: function getFeeSettings() view returns(uint256, uint256, uint256, uint256, uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetFeeSettings() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Crrrngpof.Contract.GetFeeSettings(&_Crrrngpof.CallOpts)
}

// GetFulfillStatusAtRound is a free data retrieval call binding the contract method 0x199e424d.
//
// Solidity: function getFulfillStatusAtRound(uint256 round) view returns((bool,bool))
func (_Crrrngpof *CrrrngpofCaller) GetFulfillStatusAtRound(opts *bind.CallOpts, round *big.Int) (VDFCRRNGPoFFulfillStatus, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getFulfillStatusAtRound", round)

	if err != nil {
		return *new(VDFCRRNGPoFFulfillStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(VDFCRRNGPoFFulfillStatus)).(*VDFCRRNGPoFFulfillStatus)

	return out0, err

}

// GetFulfillStatusAtRound is a free data retrieval call binding the contract method 0x199e424d.
//
// Solidity: function getFulfillStatusAtRound(uint256 round) view returns((bool,bool))
func (_Crrrngpof *CrrrngpofSession) GetFulfillStatusAtRound(round *big.Int) (VDFCRRNGPoFFulfillStatus, error) {
	return _Crrrngpof.Contract.GetFulfillStatusAtRound(&_Crrrngpof.CallOpts, round)
}

// GetFulfillStatusAtRound is a free data retrieval call binding the contract method 0x199e424d.
//
// Solidity: function getFulfillStatusAtRound(uint256 round) view returns((bool,bool))
func (_Crrrngpof *CrrrngpofCallerSession) GetFulfillStatusAtRound(round *big.Int) (VDFCRRNGPoFFulfillStatus, error) {
	return _Crrrngpof.Contract.GetFulfillStatusAtRound(&_Crrrngpof.CallOpts, round)
}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xde4303bc.
//
// Solidity: function getMinimumDepositAmount() view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) GetMinimumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getMinimumDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xde4303bc.
//
// Solidity: function getMinimumDepositAmount() view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) GetMinimumDepositAmount() (*big.Int, error) {
	return _Crrrngpof.Contract.GetMinimumDepositAmount(&_Crrrngpof.CallOpts)
}

// GetMinimumDepositAmount is a free data retrieval call binding the contract method 0xde4303bc.
//
// Solidity: function getMinimumDepositAmount() view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetMinimumDepositAmount() (*big.Int, error) {
	return _Crrrngpof.Contract.GetMinimumDepositAmount(&_Crrrngpof.CallOpts)
}

// GetNextRound is a free data retrieval call binding the contract method 0x8011d1ed.
//
// Solidity: function getNextRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) GetNextRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getNextRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextRound is a free data retrieval call binding the contract method 0x8011d1ed.
//
// Solidity: function getNextRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) GetNextRound() (*big.Int, error) {
	return _Crrrngpof.Contract.GetNextRound(&_Crrrngpof.CallOpts)
}

// GetNextRound is a free data retrieval call binding the contract method 0x8011d1ed.
//
// Solidity: function getNextRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetNextRound() (*big.Int, error) {
	return _Crrrngpof.Contract.GetNextRound(&_Crrrngpof.CallOpts)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x1a4a7c61.
//
// Solidity: function getOperatorCount() view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) GetOperatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getOperatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorCount is a free data retrieval call binding the contract method 0x1a4a7c61.
//
// Solidity: function getOperatorCount() view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) GetOperatorCount() (*big.Int, error) {
	return _Crrrngpof.Contract.GetOperatorCount(&_Crrrngpof.CallOpts)
}

// GetOperatorCount is a free data retrieval call binding the contract method 0x1a4a7c61.
//
// Solidity: function getOperatorCount() view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) GetOperatorCount() (*big.Int, error) {
	return _Crrrngpof.Contract.GetOperatorCount(&_Crrrngpof.CallOpts)
}

// GetSetUpValues is a free data retrieval call binding the contract method 0x4d541fd3.
//
// Solidity: function getSetUpValues() pure returns(uint256, uint256, uint256, uint256, bytes, bytes, bytes)
func (_Crrrngpof *CrrrngpofCaller) GetSetUpValues(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, []byte, []byte, []byte, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getSetUpValues")

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
func (_Crrrngpof *CrrrngpofSession) GetSetUpValues() (*big.Int, *big.Int, *big.Int, *big.Int, []byte, []byte, []byte, error) {
	return _Crrrngpof.Contract.GetSetUpValues(&_Crrrngpof.CallOpts)
}

// GetSetUpValues is a free data retrieval call binding the contract method 0x4d541fd3.
//
// Solidity: function getSetUpValues() pure returns(uint256, uint256, uint256, uint256, bytes, bytes, bytes)
func (_Crrrngpof *CrrrngpofCallerSession) GetSetUpValues() (*big.Int, *big.Int, *big.Int, *big.Int, []byte, []byte, []byte, error) {
	return _Crrrngpof.Contract.GetSetUpValues(&_Crrrngpof.CallOpts)
}

// GetUserStatusAtRound is a free data retrieval call binding the contract method 0xff235219.
//
// Solidity: function getUserStatusAtRound(address _operator, uint256 _round) view returns((uint256,bool))
func (_Crrrngpof *CrrrngpofCaller) GetUserStatusAtRound(opts *bind.CallOpts, _operator common.Address, _round *big.Int) (VDFCRRNGPoFOperatorStatusAtRound, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getUserStatusAtRound", _operator, _round)

	if err != nil {
		return *new(VDFCRRNGPoFOperatorStatusAtRound), err
	}

	out0 := *abi.ConvertType(out[0], new(VDFCRRNGPoFOperatorStatusAtRound)).(*VDFCRRNGPoFOperatorStatusAtRound)

	return out0, err

}

// GetUserStatusAtRound is a free data retrieval call binding the contract method 0xff235219.
//
// Solidity: function getUserStatusAtRound(address _operator, uint256 _round) view returns((uint256,bool))
func (_Crrrngpof *CrrrngpofSession) GetUserStatusAtRound(_operator common.Address, _round *big.Int) (VDFCRRNGPoFOperatorStatusAtRound, error) {
	return _Crrrngpof.Contract.GetUserStatusAtRound(&_Crrrngpof.CallOpts, _operator, _round)
}

// GetUserStatusAtRound is a free data retrieval call binding the contract method 0xff235219.
//
// Solidity: function getUserStatusAtRound(address _operator, uint256 _round) view returns((uint256,bool))
func (_Crrrngpof *CrrrngpofCallerSession) GetUserStatusAtRound(_operator common.Address, _round *big.Int) (VDFCRRNGPoFOperatorStatusAtRound, error) {
	return _Crrrngpof.Contract.GetUserStatusAtRound(&_Crrrngpof.CallOpts, _operator, _round)
}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,uint256,address,bytes,(bytes,uint256),uint8,bool,bool))
func (_Crrrngpof *CrrrngpofCaller) GetValuesAtRound(opts *bind.CallOpts, _round *big.Int) (VDFCRRNGPoFValueAtRound, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "getValuesAtRound", _round)

	if err != nil {
		return *new(VDFCRRNGPoFValueAtRound), err
	}

	out0 := *abi.ConvertType(out[0], new(VDFCRRNGPoFValueAtRound)).(*VDFCRRNGPoFValueAtRound)

	return out0, err

}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,uint256,address,bytes,(bytes,uint256),uint8,bool,bool))
func (_Crrrngpof *CrrrngpofSession) GetValuesAtRound(_round *big.Int) (VDFCRRNGPoFValueAtRound, error) {
	return _Crrrngpof.Contract.GetValuesAtRound(&_Crrrngpof.CallOpts, _round)
}

// GetValuesAtRound is a free data retrieval call binding the contract method 0x7a498fad.
//
// Solidity: function getValuesAtRound(uint256 _round) view returns((uint256,uint256,uint256,address,bytes,(bytes,uint256),uint8,bool,bool))
func (_Crrrngpof *CrrrngpofCallerSession) GetValuesAtRound(_round *big.Int) (VDFCRRNGPoFValueAtRound, error) {
	return _Crrrngpof.Contract.GetValuesAtRound(&_Crrrngpof.CallOpts, _round)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Crrrngpof *CrrrngpofCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Crrrngpof *CrrrngpofSession) IsInitialized() (bool, error) {
	return _Crrrngpof.Contract.IsInitialized(&_Crrrngpof.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Crrrngpof *CrrrngpofCallerSession) IsInitialized() (bool, error) {
	return _Crrrngpof.Contract.IsInitialized(&_Crrrngpof.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_Crrrngpof *CrrrngpofCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_Crrrngpof *CrrrngpofSession) IsOperator(operator common.Address) (bool, error) {
	return _Crrrngpof.Contract.IsOperator(&_Crrrngpof.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_Crrrngpof *CrrrngpofCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _Crrrngpof.Contract.IsOperator(&_Crrrngpof.CallOpts, operator)
}

// LastFulfilledRound is a free data retrieval call binding the contract method 0x411fa71d.
//
// Solidity: function lastFulfilledRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) LastFulfilledRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "lastFulfilledRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFulfilledRound is a free data retrieval call binding the contract method 0x411fa71d.
//
// Solidity: function lastFulfilledRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) LastFulfilledRound() (*big.Int, error) {
	return _Crrrngpof.Contract.LastFulfilledRound(&_Crrrngpof.CallOpts)
}

// LastFulfilledRound is a free data retrieval call binding the contract method 0x411fa71d.
//
// Solidity: function lastFulfilledRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) LastFulfilledRound() (*big.Int, error) {
	return _Crrrngpof.Contract.LastFulfilledRound(&_Crrrngpof.CallOpts)
}

// LastRecoveredRound is a free data retrieval call binding the contract method 0xe68bb46f.
//
// Solidity: function lastRecoveredRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofCaller) LastRecoveredRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "lastRecoveredRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRecoveredRound is a free data retrieval call binding the contract method 0xe68bb46f.
//
// Solidity: function lastRecoveredRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofSession) LastRecoveredRound() (*big.Int, error) {
	return _Crrrngpof.Contract.LastRecoveredRound(&_Crrrngpof.CallOpts)
}

// LastRecoveredRound is a free data retrieval call binding the contract method 0xe68bb46f.
//
// Solidity: function lastRecoveredRound() view returns(uint256)
func (_Crrrngpof *CrrrngpofCallerSession) LastRecoveredRound() (*big.Int, error) {
	return _Crrrngpof.Contract.LastRecoveredRound(&_Crrrngpof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crrrngpof *CrrrngpofCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crrrngpof.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crrrngpof *CrrrngpofSession) Owner() (common.Address, error) {
	return _Crrrngpof.Contract.Owner(&_Crrrngpof.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crrrngpof *CrrrngpofCallerSession) Owner() (common.Address, error) {
	return _Crrrngpof.Contract.Owner(&_Crrrngpof.CallOpts)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrngpof *CrrrngpofTransactor) Commit(opts *bind.TransactOpts, round *big.Int, c BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "commit", round, c)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrngpof *CrrrngpofSession) Commit(round *big.Int, c BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.Commit(&_Crrrngpof.TransactOpts, round, c)
}

// Commit is a paid mutator transaction binding the contract method 0x8c8583b3.
//
// Solidity: function commit(uint256 round, (bytes,uint256) c) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) Commit(round *big.Int, c BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.Commit(&_Crrrngpof.TransactOpts, round, c)
}

// DisputeLeadershipAtRound is a paid mutator transaction binding the contract method 0xda69041b.
//
// Solidity: function disputeLeadershipAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactor) DisputeLeadershipAtRound(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "disputeLeadershipAtRound", round)
}

// DisputeLeadershipAtRound is a paid mutator transaction binding the contract method 0xda69041b.
//
// Solidity: function disputeLeadershipAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofSession) DisputeLeadershipAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.DisputeLeadershipAtRound(&_Crrrngpof.TransactOpts, round)
}

// DisputeLeadershipAtRound is a paid mutator transaction binding the contract method 0xda69041b.
//
// Solidity: function disputeLeadershipAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) DisputeLeadershipAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.DisputeLeadershipAtRound(&_Crrrngpof.TransactOpts, round)
}

// DisputeRecover is a paid mutator transaction binding the contract method 0x948e8468.
//
// Solidity: function disputeRecover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofTransactor) DisputeRecover(opts *bind.TransactOpts, round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "disputeRecover", round, v, x, y)
}

// DisputeRecover is a paid mutator transaction binding the contract method 0x948e8468.
//
// Solidity: function disputeRecover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofSession) DisputeRecover(round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.DisputeRecover(&_Crrrngpof.TransactOpts, round, v, x, y)
}

// DisputeRecover is a paid mutator transaction binding the contract method 0x948e8468.
//
// Solidity: function disputeRecover(uint256 round, (bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) DisputeRecover(round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.DisputeRecover(&_Crrrngpof.TransactOpts, round, v, x, y)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0x495bd94e.
//
// Solidity: function fulfillRandomness(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactor) FulfillRandomness(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "fulfillRandomness", round)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0x495bd94e.
//
// Solidity: function fulfillRandomness(uint256 round) returns()
func (_Crrrngpof *CrrrngpofSession) FulfillRandomness(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.FulfillRandomness(&_Crrrngpof.TransactOpts, round)
}

// FulfillRandomness is a paid mutator transaction binding the contract method 0x495bd94e.
//
// Solidity: function fulfillRandomness(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) FulfillRandomness(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.FulfillRandomness(&_Crrrngpof.TransactOpts, round)
}

// FulfillRandomnessOfFailed is a paid mutator transaction binding the contract method 0x27fe5cec.
//
// Solidity: function fulfillRandomnessOfFailed(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactor) FulfillRandomnessOfFailed(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "fulfillRandomnessOfFailed", round)
}

// FulfillRandomnessOfFailed is a paid mutator transaction binding the contract method 0x27fe5cec.
//
// Solidity: function fulfillRandomnessOfFailed(uint256 round) returns()
func (_Crrrngpof *CrrrngpofSession) FulfillRandomnessOfFailed(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.FulfillRandomnessOfFailed(&_Crrrngpof.TransactOpts, round)
}

// FulfillRandomnessOfFailed is a paid mutator transaction binding the contract method 0x27fe5cec.
//
// Solidity: function fulfillRandomnessOfFailed(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) FulfillRandomnessOfFailed(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.FulfillRandomnessOfFailed(&_Crrrngpof.TransactOpts, round)
}

// Initialize is a paid mutator transaction binding the contract method 0x8efd09d1.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofTransactor) Initialize(opts *bind.TransactOpts, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "initialize", v, x, y)
}

// Initialize is a paid mutator transaction binding the contract method 0x8efd09d1.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofSession) Initialize(v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.Initialize(&_Crrrngpof.TransactOpts, v, x, y)
}

// Initialize is a paid mutator transaction binding the contract method 0x8efd09d1.
//
// Solidity: function initialize((bytes,uint256)[] v, (bytes,uint256) x, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) Initialize(v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.Initialize(&_Crrrngpof.TransactOpts, v, x, y)
}

// OperatorDeposit is a paid mutator transaction binding the contract method 0xf23fa3d3.
//
// Solidity: function operatorDeposit() payable returns()
func (_Crrrngpof *CrrrngpofTransactor) OperatorDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "operatorDeposit")
}

// OperatorDeposit is a paid mutator transaction binding the contract method 0xf23fa3d3.
//
// Solidity: function operatorDeposit() payable returns()
func (_Crrrngpof *CrrrngpofSession) OperatorDeposit() (*types.Transaction, error) {
	return _Crrrngpof.Contract.OperatorDeposit(&_Crrrngpof.TransactOpts)
}

// OperatorDeposit is a paid mutator transaction binding the contract method 0xf23fa3d3.
//
// Solidity: function operatorDeposit() payable returns()
func (_Crrrngpof *CrrrngpofTransactorSession) OperatorDeposit() (*types.Transaction, error) {
	return _Crrrngpof.Contract.OperatorDeposit(&_Crrrngpof.TransactOpts)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0x9238a02f.
//
// Solidity: function operatorWithdraw(uint256 amount) returns()
func (_Crrrngpof *CrrrngpofTransactor) OperatorWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "operatorWithdraw", amount)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0x9238a02f.
//
// Solidity: function operatorWithdraw(uint256 amount) returns()
func (_Crrrngpof *CrrrngpofSession) OperatorWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.OperatorWithdraw(&_Crrrngpof.TransactOpts, amount)
}

// OperatorWithdraw is a paid mutator transaction binding the contract method 0x9238a02f.
//
// Solidity: function operatorWithdraw(uint256 amount) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) OperatorWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.OperatorWithdraw(&_Crrrngpof.TransactOpts, amount)
}

// ReRequestRandomWordAtRound is a paid mutator transaction binding the contract method 0x3b246b1d.
//
// Solidity: function reRequestRandomWordAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactor) ReRequestRandomWordAtRound(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "reRequestRandomWordAtRound", round)
}

// ReRequestRandomWordAtRound is a paid mutator transaction binding the contract method 0x3b246b1d.
//
// Solidity: function reRequestRandomWordAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofSession) ReRequestRandomWordAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.ReRequestRandomWordAtRound(&_Crrrngpof.TransactOpts, round)
}

// ReRequestRandomWordAtRound is a paid mutator transaction binding the contract method 0x3b246b1d.
//
// Solidity: function reRequestRandomWordAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) ReRequestRandomWordAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.ReRequestRandomWordAtRound(&_Crrrngpof.TransactOpts, round)
}

// Recover is a paid mutator transaction binding the contract method 0x141348ad.
//
// Solidity: function recover(uint256 round, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofTransactor) Recover(opts *bind.TransactOpts, round *big.Int, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "recover", round, y)
}

// Recover is a paid mutator transaction binding the contract method 0x141348ad.
//
// Solidity: function recover(uint256 round, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofSession) Recover(round *big.Int, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.Recover(&_Crrrngpof.TransactOpts, round, y)
}

// Recover is a paid mutator transaction binding the contract method 0x141348ad.
//
// Solidity: function recover(uint256 round, (bytes,uint256) y) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) Recover(round *big.Int, y BigNumber) (*types.Transaction, error) {
	return _Crrrngpof.Contract.Recover(&_Crrrngpof.TransactOpts, round, y)
}

// RefundAtRound is a paid mutator transaction binding the contract method 0xf02cfbbf.
//
// Solidity: function refundAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactor) RefundAtRound(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "refundAtRound", round)
}

// RefundAtRound is a paid mutator transaction binding the contract method 0xf02cfbbf.
//
// Solidity: function refundAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofSession) RefundAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.RefundAtRound(&_Crrrngpof.TransactOpts, round)
}

// RefundAtRound is a paid mutator transaction binding the contract method 0xf02cfbbf.
//
// Solidity: function refundAtRound(uint256 round) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) RefundAtRound(round *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.RefundAtRound(&_Crrrngpof.TransactOpts, round)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crrrngpof *CrrrngpofTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crrrngpof *CrrrngpofSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crrrngpof.Contract.RenounceOwnership(&_Crrrngpof.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crrrngpof *CrrrngpofTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crrrngpof.Contract.RenounceOwnership(&_Crrrngpof.TransactOpts)
}

// RequestRandomWordDirectFunding is a paid mutator transaction binding the contract method 0x7072b977.
//
// Solidity: function requestRandomWordDirectFunding(uint32 callbackGasLimit) payable returns(uint256)
func (_Crrrngpof *CrrrngpofTransactor) RequestRandomWordDirectFunding(opts *bind.TransactOpts, callbackGasLimit uint32) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "requestRandomWordDirectFunding", callbackGasLimit)
}

// RequestRandomWordDirectFunding is a paid mutator transaction binding the contract method 0x7072b977.
//
// Solidity: function requestRandomWordDirectFunding(uint32 callbackGasLimit) payable returns(uint256)
func (_Crrrngpof *CrrrngpofSession) RequestRandomWordDirectFunding(callbackGasLimit uint32) (*types.Transaction, error) {
	return _Crrrngpof.Contract.RequestRandomWordDirectFunding(&_Crrrngpof.TransactOpts, callbackGasLimit)
}

// RequestRandomWordDirectFunding is a paid mutator transaction binding the contract method 0x7072b977.
//
// Solidity: function requestRandomWordDirectFunding(uint32 callbackGasLimit) payable returns(uint256)
func (_Crrrngpof *CrrrngpofTransactorSession) RequestRandomWordDirectFunding(callbackGasLimit uint32) (*types.Transaction, error) {
	return _Crrrngpof.Contract.RequestRandomWordDirectFunding(&_Crrrngpof.TransactOpts, callbackGasLimit)
}

// SetSettings is a paid mutator transaction binding the contract method 0xcc917f4e.
//
// Solidity: function setSettings(uint256 disputePeriod, uint256 minimumDepositAmount, uint256 avgL2GasUsed, uint256 avgL1GasUsed, uint256 premiumPercentage, uint256 flatFee) returns()
func (_Crrrngpof *CrrrngpofTransactor) SetSettings(opts *bind.TransactOpts, disputePeriod *big.Int, minimumDepositAmount *big.Int, avgL2GasUsed *big.Int, avgL1GasUsed *big.Int, premiumPercentage *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "setSettings", disputePeriod, minimumDepositAmount, avgL2GasUsed, avgL1GasUsed, premiumPercentage, flatFee)
}

// SetSettings is a paid mutator transaction binding the contract method 0xcc917f4e.
//
// Solidity: function setSettings(uint256 disputePeriod, uint256 minimumDepositAmount, uint256 avgL2GasUsed, uint256 avgL1GasUsed, uint256 premiumPercentage, uint256 flatFee) returns()
func (_Crrrngpof *CrrrngpofSession) SetSettings(disputePeriod *big.Int, minimumDepositAmount *big.Int, avgL2GasUsed *big.Int, avgL1GasUsed *big.Int, premiumPercentage *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.SetSettings(&_Crrrngpof.TransactOpts, disputePeriod, minimumDepositAmount, avgL2GasUsed, avgL1GasUsed, premiumPercentage, flatFee)
}

// SetSettings is a paid mutator transaction binding the contract method 0xcc917f4e.
//
// Solidity: function setSettings(uint256 disputePeriod, uint256 minimumDepositAmount, uint256 avgL2GasUsed, uint256 avgL1GasUsed, uint256 premiumPercentage, uint256 flatFee) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) SetSettings(disputePeriod *big.Int, minimumDepositAmount *big.Int, avgL2GasUsed *big.Int, avgL1GasUsed *big.Int, premiumPercentage *big.Int, flatFee *big.Int) (*types.Transaction, error) {
	return _Crrrngpof.Contract.SetSettings(&_Crrrngpof.TransactOpts, disputePeriod, minimumDepositAmount, avgL2GasUsed, avgL1GasUsed, premiumPercentage, flatFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crrrngpof *CrrrngpofTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crrrngpof.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crrrngpof *CrrrngpofSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crrrngpof.Contract.TransferOwnership(&_Crrrngpof.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crrrngpof *CrrrngpofTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crrrngpof.Contract.TransferOwnership(&_Crrrngpof.TransactOpts, newOwner)
}

// CrrrngpofCommitCIterator is returned from FilterCommitC and is used to iterate over the raw logs and unpacked data for CommitC events raised by the Crrrngpof contract.
type CrrrngpofCommitCIterator struct {
	Event *CrrrngpofCommitC // Event containing the contract specifics and raw log

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
func (it *CrrrngpofCommitCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngpofCommitC)
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
		it.Event = new(CrrrngpofCommitC)
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
func (it *CrrrngpofCommitCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngpofCommitCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngpofCommitC represents a CommitC event raised by the Crrrngpof contract.
type CrrrngpofCommitC struct {
	CommitCount *big.Int
	CommitVal   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommitC is a free log retrieval operation binding the contract event 0xe2bd9c5fc4c79023527f33e9ed312ce0178f7adfc724f21a8bf76f43d93042c4.
//
// Solidity: event CommitC(uint256 commitCount, bytes commitVal)
func (_Crrrngpof *CrrrngpofFilterer) FilterCommitC(opts *bind.FilterOpts) (*CrrrngpofCommitCIterator, error) {

	logs, sub, err := _Crrrngpof.contract.FilterLogs(opts, "CommitC")
	if err != nil {
		return nil, err
	}
	return &CrrrngpofCommitCIterator{contract: _Crrrngpof.contract, event: "CommitC", logs: logs, sub: sub}, nil
}

// WatchCommitC is a free log subscription operation binding the contract event 0xe2bd9c5fc4c79023527f33e9ed312ce0178f7adfc724f21a8bf76f43d93042c4.
//
// Solidity: event CommitC(uint256 commitCount, bytes commitVal)
func (_Crrrngpof *CrrrngpofFilterer) WatchCommitC(opts *bind.WatchOpts, sink chan<- *CrrrngpofCommitC) (event.Subscription, error) {

	logs, sub, err := _Crrrngpof.contract.WatchLogs(opts, "CommitC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngpofCommitC)
				if err := _Crrrngpof.contract.UnpackLog(event, "CommitC", log); err != nil {
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
func (_Crrrngpof *CrrrngpofFilterer) ParseCommitC(log types.Log) (*CrrrngpofCommitC, error) {
	event := new(CrrrngpofCommitC)
	if err := _Crrrngpof.contract.UnpackLog(event, "CommitC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngpofFulfillRandomnessIterator is returned from FilterFulfillRandomness and is used to iterate over the raw logs and unpacked data for FulfillRandomness events raised by the Crrrngpof contract.
type CrrrngpofFulfillRandomnessIterator struct {
	Event *CrrrngpofFulfillRandomness // Event containing the contract specifics and raw log

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
func (it *CrrrngpofFulfillRandomnessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngpofFulfillRandomness)
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
		it.Event = new(CrrrngpofFulfillRandomness)
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
func (it *CrrrngpofFulfillRandomnessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngpofFulfillRandomnessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngpofFulfillRandomness represents a FulfillRandomness event raised by the Crrrngpof contract.
type CrrrngpofFulfillRandomness struct {
	Round       *big.Int
	HashedOmega *big.Int
	Success     bool
	Leader      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFulfillRandomness is a free log retrieval operation binding the contract event 0x81f6207cc5d4e30b3e281e9ce3531afc36baddb714b1a5847121dee9f6cbb9a1.
//
// Solidity: event FulfillRandomness(uint256 round, uint256 hashedOmega, bool success, address leader)
func (_Crrrngpof *CrrrngpofFilterer) FilterFulfillRandomness(opts *bind.FilterOpts) (*CrrrngpofFulfillRandomnessIterator, error) {

	logs, sub, err := _Crrrngpof.contract.FilterLogs(opts, "FulfillRandomness")
	if err != nil {
		return nil, err
	}
	return &CrrrngpofFulfillRandomnessIterator{contract: _Crrrngpof.contract, event: "FulfillRandomness", logs: logs, sub: sub}, nil
}

// WatchFulfillRandomness is a free log subscription operation binding the contract event 0x81f6207cc5d4e30b3e281e9ce3531afc36baddb714b1a5847121dee9f6cbb9a1.
//
// Solidity: event FulfillRandomness(uint256 round, uint256 hashedOmega, bool success, address leader)
func (_Crrrngpof *CrrrngpofFilterer) WatchFulfillRandomness(opts *bind.WatchOpts, sink chan<- *CrrrngpofFulfillRandomness) (event.Subscription, error) {

	logs, sub, err := _Crrrngpof.contract.WatchLogs(opts, "FulfillRandomness")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngpofFulfillRandomness)
				if err := _Crrrngpof.contract.UnpackLog(event, "FulfillRandomness", log); err != nil {
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

// ParseFulfillRandomness is a log parse operation binding the contract event 0x81f6207cc5d4e30b3e281e9ce3531afc36baddb714b1a5847121dee9f6cbb9a1.
//
// Solidity: event FulfillRandomness(uint256 round, uint256 hashedOmega, bool success, address leader)
func (_Crrrngpof *CrrrngpofFilterer) ParseFulfillRandomness(log types.Log) (*CrrrngpofFulfillRandomness, error) {
	event := new(CrrrngpofFulfillRandomness)
	if err := _Crrrngpof.contract.UnpackLog(event, "FulfillRandomness", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngpofOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crrrngpof contract.
type CrrrngpofOwnershipTransferredIterator struct {
	Event *CrrrngpofOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrrrngpofOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngpofOwnershipTransferred)
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
		it.Event = new(CrrrngpofOwnershipTransferred)
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
func (it *CrrrngpofOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngpofOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngpofOwnershipTransferred represents a OwnershipTransferred event raised by the Crrrngpof contract.
type CrrrngpofOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crrrngpof *CrrrngpofFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrrrngpofOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crrrngpof.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrrrngpofOwnershipTransferredIterator{contract: _Crrrngpof.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crrrngpof *CrrrngpofFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrrrngpofOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crrrngpof.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngpofOwnershipTransferred)
				if err := _Crrrngpof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Crrrngpof *CrrrngpofFilterer) ParseOwnershipTransferred(log types.Log) (*CrrrngpofOwnershipTransferred, error) {
	event := new(CrrrngpofOwnershipTransferred)
	if err := _Crrrngpof.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngpofRandomWordsRequestedIterator is returned from FilterRandomWordsRequested and is used to iterate over the raw logs and unpacked data for RandomWordsRequested events raised by the Crrrngpof contract.
type CrrrngpofRandomWordsRequestedIterator struct {
	Event *CrrrngpofRandomWordsRequested // Event containing the contract specifics and raw log

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
func (it *CrrrngpofRandomWordsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngpofRandomWordsRequested)
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
		it.Event = new(CrrrngpofRandomWordsRequested)
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
func (it *CrrrngpofRandomWordsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngpofRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngpofRandomWordsRequested represents a RandomWordsRequested event raised by the Crrrngpof contract.
type CrrrngpofRandomWordsRequested struct {
	Round  *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsRequested is a free log retrieval operation binding the contract event 0x5afc92f47019f6f981dbe48fba0d8bc10298b161505d2ead5ad4b975bfde77f4.
//
// Solidity: event RandomWordsRequested(uint256 round, address sender)
func (_Crrrngpof *CrrrngpofFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts) (*CrrrngpofRandomWordsRequestedIterator, error) {

	logs, sub, err := _Crrrngpof.contract.FilterLogs(opts, "RandomWordsRequested")
	if err != nil {
		return nil, err
	}
	return &CrrrngpofRandomWordsRequestedIterator{contract: _Crrrngpof.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

// WatchRandomWordsRequested is a free log subscription operation binding the contract event 0x5afc92f47019f6f981dbe48fba0d8bc10298b161505d2ead5ad4b975bfde77f4.
//
// Solidity: event RandomWordsRequested(uint256 round, address sender)
func (_Crrrngpof *CrrrngpofFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *CrrrngpofRandomWordsRequested) (event.Subscription, error) {

	logs, sub, err := _Crrrngpof.contract.WatchLogs(opts, "RandomWordsRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngpofRandomWordsRequested)
				if err := _Crrrngpof.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
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
func (_Crrrngpof *CrrrngpofFilterer) ParseRandomWordsRequested(log types.Log) (*CrrrngpofRandomWordsRequested, error) {
	event := new(CrrrngpofRandomWordsRequested)
	if err := _Crrrngpof.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrrrngpofRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the Crrrngpof contract.
type CrrrngpofRecoveredIterator struct {
	Event *CrrrngpofRecovered // Event containing the contract specifics and raw log

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
func (it *CrrrngpofRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrrrngpofRecovered)
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
		it.Event = new(CrrrngpofRecovered)
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
func (it *CrrrngpofRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrrrngpofRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrrrngpofRecovered represents a Recovered event raised by the Crrrngpof contract.
type CrrrngpofRecovered struct {
	Round     *big.Int
	Recoverer common.Address
	Omega     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x25899563c74bb3958cacda4bb21a8d47d17287b9008b87963ed1802f3fba71ec.
//
// Solidity: event Recovered(uint256 round, address recoverer, bytes omega)
func (_Crrrngpof *CrrrngpofFilterer) FilterRecovered(opts *bind.FilterOpts) (*CrrrngpofRecoveredIterator, error) {

	logs, sub, err := _Crrrngpof.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &CrrrngpofRecoveredIterator{contract: _Crrrngpof.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x25899563c74bb3958cacda4bb21a8d47d17287b9008b87963ed1802f3fba71ec.
//
// Solidity: event Recovered(uint256 round, address recoverer, bytes omega)
func (_Crrrngpof *CrrrngpofFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *CrrrngpofRecovered) (event.Subscription, error) {

	logs, sub, err := _Crrrngpof.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrrrngpofRecovered)
				if err := _Crrrngpof.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x25899563c74bb3958cacda4bb21a8d47d17287b9008b87963ed1802f3fba71ec.
//
// Solidity: event Recovered(uint256 round, address recoverer, bytes omega)
func (_Crrrngpof *CrrrngpofFilterer) ParseRecovered(log types.Log) (*CrrrngpofRecovered, error) {
	event := new(CrrrngpofRecovered)
	if err := _Crrrngpof.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
