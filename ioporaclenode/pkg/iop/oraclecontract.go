// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iop

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

// OracleContractIBOSSignature is an auto generated low-level Go binding around an user-defined struct.
type OracleContractIBOSSignature struct {
	S [2]*big.Int
	R [4]*big.Int
}

// OracleContractMetaData contains all meta data concerning the OracleContract contract.
var OracleContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distKeyContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowThreshold\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowMultisig\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowIBOS\",\"type\":\"bool\"}],\"name\":\"AllowedModesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldM\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newM\",\"type\":\"uint256\"}],\"name\":\"MultisigRequiredChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ValidationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ValidationResponse\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"VALIDATOR_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"allowIBOS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"allowMultisig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"allowThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"multisigRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedNonce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"setMultisigRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"tss\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ms\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ibos\",\"type\":\"bool\"}],\"name\":\"setAllowedModes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateBlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"identityBytes\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"S\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"R\",\"type\":\"uint256[4]\"}],\"internalType\":\"structOracleContract.IBOSSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitBlockValidationIBOS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"identityBytes\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"S\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"R\",\"type\":\"uint256[4]\"}],\"internalType\":\"structOracleContract.IBOSSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitTransactionValidationIBOS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"tssSig\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes[]\",\"name\":\"msSigs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitBlockValidationHybrid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"tssSig\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes[]\",\"name\":\"msSigs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"submitTransactionValidationHybrid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seed\",\"type\":\"uint256\"}],\"name\":\"isValidationFeeReceiver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitBlockValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitTransactionValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findBlockValidationResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findTransactionValidationResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// OracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleContractMetaData.ABI instead.
var OracleContractABI = OracleContractMetaData.ABI

// OracleContract is an auto generated Go binding around an Ethereum contract.
type OracleContract struct {
	OracleContractCaller     // Read-only binding to the contract
	OracleContractTransactor // Write-only binding to the contract
	OracleContractFilterer   // Log filterer for contract events
}

// OracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleContractSession struct {
	Contract     *OracleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleContractCallerSession struct {
	Contract *OracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleContractTransactorSession struct {
	Contract     *OracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleContractRaw struct {
	Contract *OracleContract // Generic contract binding to access the raw methods on
}

// OracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleContractCallerRaw struct {
	Contract *OracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleContractTransactorRaw struct {
	Contract *OracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleContract creates a new instance of OracleContract, bound to a specific deployed contract.
func NewOracleContract(address common.Address, backend bind.ContractBackend) (*OracleContract, error) {
	contract, err := bindOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// NewOracleContractCaller creates a new read-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractCaller(address common.Address, caller bind.ContractCaller) (*OracleContractCaller, error) {
	contract, err := bindOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractCaller{contract: contract}, nil
}

// NewOracleContractTransactor creates a new write-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleContractTransactor, error) {
	contract, err := bindOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransactor{contract: contract}, nil
}

// NewOracleContractFilterer creates a new log filterer instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleContractFilterer, error) {
	contract, err := bindOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleContractFilterer{contract: contract}, nil
}

// bindOracleContract binds a generic wrapper to an already deployed contract.
func bindOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.OracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transact(opts, method, params...)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) TOTALFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "TOTAL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) VALIDATORFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "VALIDATOR_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_OracleContract *OracleContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_OracleContract *OracleContractSession) Admin() (common.Address, error) {
	return _OracleContract.Contract.Admin(&_OracleContract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_OracleContract *OracleContractCallerSession) Admin() (common.Address, error) {
	return _OracleContract.Contract.Admin(&_OracleContract.CallOpts)
}

// AllowIBOS is a free data retrieval call binding the contract method 0x1d41b6e9.
//
// Solidity: function allowIBOS() view returns(bool)
func (_OracleContract *OracleContractCaller) AllowIBOS(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "allowIBOS")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowIBOS is a free data retrieval call binding the contract method 0x1d41b6e9.
//
// Solidity: function allowIBOS() view returns(bool)
func (_OracleContract *OracleContractSession) AllowIBOS() (bool, error) {
	return _OracleContract.Contract.AllowIBOS(&_OracleContract.CallOpts)
}

// AllowIBOS is a free data retrieval call binding the contract method 0x1d41b6e9.
//
// Solidity: function allowIBOS() view returns(bool)
func (_OracleContract *OracleContractCallerSession) AllowIBOS() (bool, error) {
	return _OracleContract.Contract.AllowIBOS(&_OracleContract.CallOpts)
}

// AllowMultisig is a free data retrieval call binding the contract method 0xbacc4e54.
//
// Solidity: function allowMultisig() view returns(bool)
func (_OracleContract *OracleContractCaller) AllowMultisig(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "allowMultisig")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowMultisig is a free data retrieval call binding the contract method 0xbacc4e54.
//
// Solidity: function allowMultisig() view returns(bool)
func (_OracleContract *OracleContractSession) AllowMultisig() (bool, error) {
	return _OracleContract.Contract.AllowMultisig(&_OracleContract.CallOpts)
}

// AllowMultisig is a free data retrieval call binding the contract method 0xbacc4e54.
//
// Solidity: function allowMultisig() view returns(bool)
func (_OracleContract *OracleContractCallerSession) AllowMultisig() (bool, error) {
	return _OracleContract.Contract.AllowMultisig(&_OracleContract.CallOpts)
}

// AllowThreshold is a free data retrieval call binding the contract method 0xff7b16d2.
//
// Solidity: function allowThreshold() view returns(bool)
func (_OracleContract *OracleContractCaller) AllowThreshold(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "allowThreshold")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowThreshold is a free data retrieval call binding the contract method 0xff7b16d2.
//
// Solidity: function allowThreshold() view returns(bool)
func (_OracleContract *OracleContractSession) AllowThreshold() (bool, error) {
	return _OracleContract.Contract.AllowThreshold(&_OracleContract.CallOpts)
}

// AllowThreshold is a free data retrieval call binding the contract method 0xff7b16d2.
//
// Solidity: function allowThreshold() view returns(bool)
func (_OracleContract *OracleContractCallerSession) AllowThreshold() (bool, error) {
	return _OracleContract.Contract.AllowThreshold(&_OracleContract.CallOpts)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) FindBlockValidationResult(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findBlockValidationResult", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) FindBlockValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) FindBlockValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) FindTransactionValidationResult(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findTransactionValidationResult", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) FindTransactionValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindTransactionValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindTransactionValidationResult is a free data retrieval call binding the contract method 0x43434590.
//
// Solidity: function findTransactionValidationResult(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) FindTransactionValidationResult(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.FindTransactionValidationResult(&_OracleContract.CallOpts, _hash)
}

// IsValidationFeeReceiver is a free data retrieval call binding the contract method 0xf5ba1d3d.
//
// Solidity: function isValidationFeeReceiver(uint256 _stake, uint256 _seed) pure returns(bool)
func (_OracleContract *OracleContractCaller) IsValidationFeeReceiver(opts *bind.CallOpts, _stake *big.Int, _seed *big.Int) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "isValidationFeeReceiver", _stake, _seed)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidationFeeReceiver is a free data retrieval call binding the contract method 0xf5ba1d3d.
//
// Solidity: function isValidationFeeReceiver(uint256 _stake, uint256 _seed) pure returns(bool)
func (_OracleContract *OracleContractSession) IsValidationFeeReceiver(_stake *big.Int, _seed *big.Int) (bool, error) {
	return _OracleContract.Contract.IsValidationFeeReceiver(&_OracleContract.CallOpts, _stake, _seed)
}

// IsValidationFeeReceiver is a free data retrieval call binding the contract method 0xf5ba1d3d.
//
// Solidity: function isValidationFeeReceiver(uint256 _stake, uint256 _seed) pure returns(bool)
func (_OracleContract *OracleContractCallerSession) IsValidationFeeReceiver(_stake *big.Int, _seed *big.Int) (bool, error) {
	return _OracleContract.Contract.IsValidationFeeReceiver(&_OracleContract.CallOpts, _stake, _seed)
}

// MultisigRequired is a free data retrieval call binding the contract method 0x7241784c.
//
// Solidity: function multisigRequired() view returns(uint256)
func (_OracleContract *OracleContractCaller) MultisigRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "multisigRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultisigRequired is a free data retrieval call binding the contract method 0x7241784c.
//
// Solidity: function multisigRequired() view returns(uint256)
func (_OracleContract *OracleContractSession) MultisigRequired() (*big.Int, error) {
	return _OracleContract.Contract.MultisigRequired(&_OracleContract.CallOpts)
}

// MultisigRequired is a free data retrieval call binding the contract method 0x7241784c.
//
// Solidity: function multisigRequired() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) MultisigRequired() (*big.Int, error) {
	return _OracleContract.Contract.MultisigRequired(&_OracleContract.CallOpts)
}

// UsedNonce is a free data retrieval call binding the contract method 0x9723fb6d.
//
// Solidity: function usedNonce(uint256 ) view returns(bool)
func (_OracleContract *OracleContractCaller) UsedNonce(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "usedNonce", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNonce is a free data retrieval call binding the contract method 0x9723fb6d.
//
// Solidity: function usedNonce(uint256 ) view returns(bool)
func (_OracleContract *OracleContractSession) UsedNonce(arg0 *big.Int) (bool, error) {
	return _OracleContract.Contract.UsedNonce(&_OracleContract.CallOpts, arg0)
}

// UsedNonce is a free data retrieval call binding the contract method 0x9723fb6d.
//
// Solidity: function usedNonce(uint256 ) view returns(bool)
func (_OracleContract *OracleContractCallerSession) UsedNonce(arg0 *big.Int) (bool, error) {
	return _OracleContract.Contract.UsedNonce(&_OracleContract.CallOpts, arg0)
}

// SetAllowedModes is a paid mutator transaction binding the contract method 0xe6f25ccc.
//
// Solidity: function setAllowedModes(bool tss, bool ms, bool ibos) returns()
func (_OracleContract *OracleContractTransactor) SetAllowedModes(opts *bind.TransactOpts, tss bool, ms bool, ibos bool) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "setAllowedModes", tss, ms, ibos)
}

// SetAllowedModes is a paid mutator transaction binding the contract method 0xe6f25ccc.
//
// Solidity: function setAllowedModes(bool tss, bool ms, bool ibos) returns()
func (_OracleContract *OracleContractSession) SetAllowedModes(tss bool, ms bool, ibos bool) (*types.Transaction, error) {
	return _OracleContract.Contract.SetAllowedModes(&_OracleContract.TransactOpts, tss, ms, ibos)
}

// SetAllowedModes is a paid mutator transaction binding the contract method 0xe6f25ccc.
//
// Solidity: function setAllowedModes(bool tss, bool ms, bool ibos) returns()
func (_OracleContract *OracleContractTransactorSession) SetAllowedModes(tss bool, ms bool, ibos bool) (*types.Transaction, error) {
	return _OracleContract.Contract.SetAllowedModes(&_OracleContract.TransactOpts, tss, ms, ibos)
}

// SetMultisigRequired is a paid mutator transaction binding the contract method 0x3205772a.
//
// Solidity: function setMultisigRequired(uint256 m) returns()
func (_OracleContract *OracleContractTransactor) SetMultisigRequired(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "setMultisigRequired", m)
}

// SetMultisigRequired is a paid mutator transaction binding the contract method 0x3205772a.
//
// Solidity: function setMultisigRequired(uint256 m) returns()
func (_OracleContract *OracleContractSession) SetMultisigRequired(m *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SetMultisigRequired(&_OracleContract.TransactOpts, m)
}

// SetMultisigRequired is a paid mutator transaction binding the contract method 0x3205772a.
//
// Solidity: function setMultisigRequired(uint256 m) returns()
func (_OracleContract *OracleContractTransactorSession) SetMultisigRequired(m *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SetMultisigRequired(&_OracleContract.TransactOpts, m)
}

// SubmitBlockValidationHybrid is a paid mutator transaction binding the contract method 0xa47e5c9c.
//
// Solidity: function submitBlockValidationHybrid(bytes32 _hash, bool _result, uint8 mode, uint256[2] tssSig, bytes[] msSigs, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactor) SubmitBlockValidationHybrid(opts *bind.TransactOpts, _hash [32]byte, _result bool, mode uint8, tssSig [2]*big.Int, msSigs [][]byte, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitBlockValidationHybrid", _hash, _result, mode, tssSig, msSigs, nonce)
}

// SubmitBlockValidationHybrid is a paid mutator transaction binding the contract method 0xa47e5c9c.
//
// Solidity: function submitBlockValidationHybrid(bytes32 _hash, bool _result, uint8 mode, uint256[2] tssSig, bytes[] msSigs, uint256 nonce) returns()
func (_OracleContract *OracleContractSession) SubmitBlockValidationHybrid(_hash [32]byte, _result bool, mode uint8, tssSig [2]*big.Int, msSigs [][]byte, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationHybrid(&_OracleContract.TransactOpts, _hash, _result, mode, tssSig, msSigs, nonce)
}

// SubmitBlockValidationHybrid is a paid mutator transaction binding the contract method 0xa47e5c9c.
//
// Solidity: function submitBlockValidationHybrid(bytes32 _hash, bool _result, uint8 mode, uint256[2] tssSig, bytes[] msSigs, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitBlockValidationHybrid(_hash [32]byte, _result bool, mode uint8, tssSig [2]*big.Int, msSigs [][]byte, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationHybrid(&_OracleContract.TransactOpts, _hash, _result, mode, tssSig, msSigs, nonce)
}

// SubmitBlockValidationIBOS is a paid mutator transaction binding the contract method 0x96e456d1.
//
// Solidity: function submitBlockValidationIBOS(bytes32 _hash, bool _result, bytes[] identityBytes, (uint256[2],uint256[4])[] signatures, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactor) SubmitBlockValidationIBOS(opts *bind.TransactOpts, _hash [32]byte, _result bool, identityBytes [][]byte, signatures []OracleContractIBOSSignature, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitBlockValidationIBOS", _hash, _result, identityBytes, signatures, nonce)
}

// SubmitBlockValidationIBOS is a paid mutator transaction binding the contract method 0x96e456d1.
//
// Solidity: function submitBlockValidationIBOS(bytes32 _hash, bool _result, bytes[] identityBytes, (uint256[2],uint256[4])[] signatures, uint256 nonce) returns()
func (_OracleContract *OracleContractSession) SubmitBlockValidationIBOS(_hash [32]byte, _result bool, identityBytes [][]byte, signatures []OracleContractIBOSSignature, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationIBOS(&_OracleContract.TransactOpts, _hash, _result, identityBytes, signatures, nonce)
}

// SubmitBlockValidationIBOS is a paid mutator transaction binding the contract method 0x96e456d1.
//
// Solidity: function submitBlockValidationIBOS(bytes32 _hash, bool _result, bytes[] identityBytes, (uint256[2],uint256[4])[] signatures, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitBlockValidationIBOS(_hash [32]byte, _result bool, identityBytes [][]byte, signatures []OracleContractIBOSSignature, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationIBOS(&_OracleContract.TransactOpts, _hash, _result, identityBytes, signatures, nonce)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitBlockValidationResult(opts *bind.TransactOpts, _hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitBlockValidationResult", _hash, _result, _signature)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitBlockValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitBlockValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitTransactionValidationHybrid is a paid mutator transaction binding the contract method 0x4a2a8b0d.
//
// Solidity: function submitTransactionValidationHybrid(bytes32 _hash, bool _result, uint8 mode, uint256[2] tssSig, bytes[] msSigs, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationHybrid(opts *bind.TransactOpts, _hash [32]byte, _result bool, mode uint8, tssSig [2]*big.Int, msSigs [][]byte, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationHybrid", _hash, _result, mode, tssSig, msSigs, nonce)
}

// SubmitTransactionValidationHybrid is a paid mutator transaction binding the contract method 0x4a2a8b0d.
//
// Solidity: function submitTransactionValidationHybrid(bytes32 _hash, bool _result, uint8 mode, uint256[2] tssSig, bytes[] msSigs, uint256 nonce) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationHybrid(_hash [32]byte, _result bool, mode uint8, tssSig [2]*big.Int, msSigs [][]byte, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationHybrid(&_OracleContract.TransactOpts, _hash, _result, mode, tssSig, msSigs, nonce)
}

// SubmitTransactionValidationHybrid is a paid mutator transaction binding the contract method 0x4a2a8b0d.
//
// Solidity: function submitTransactionValidationHybrid(bytes32 _hash, bool _result, uint8 mode, uint256[2] tssSig, bytes[] msSigs, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationHybrid(_hash [32]byte, _result bool, mode uint8, tssSig [2]*big.Int, msSigs [][]byte, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationHybrid(&_OracleContract.TransactOpts, _hash, _result, mode, tssSig, msSigs, nonce)
}

// SubmitTransactionValidationIBOS is a paid mutator transaction binding the contract method 0xb6fdebe7.
//
// Solidity: function submitTransactionValidationIBOS(bytes32 _hash, bool _result, bytes[] identityBytes, (uint256[2],uint256[4])[] signatures, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationIBOS(opts *bind.TransactOpts, _hash [32]byte, _result bool, identityBytes [][]byte, signatures []OracleContractIBOSSignature, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationIBOS", _hash, _result, identityBytes, signatures, nonce)
}

// SubmitTransactionValidationIBOS is a paid mutator transaction binding the contract method 0xb6fdebe7.
//
// Solidity: function submitTransactionValidationIBOS(bytes32 _hash, bool _result, bytes[] identityBytes, (uint256[2],uint256[4])[] signatures, uint256 nonce) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationIBOS(_hash [32]byte, _result bool, identityBytes [][]byte, signatures []OracleContractIBOSSignature, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationIBOS(&_OracleContract.TransactOpts, _hash, _result, identityBytes, signatures, nonce)
}

// SubmitTransactionValidationIBOS is a paid mutator transaction binding the contract method 0xb6fdebe7.
//
// Solidity: function submitTransactionValidationIBOS(bytes32 _hash, bool _result, bytes[] identityBytes, (uint256[2],uint256[4])[] signatures, uint256 nonce) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationIBOS(_hash [32]byte, _result bool, identityBytes [][]byte, signatures []OracleContractIBOSSignature, nonce *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationIBOS(&_OracleContract.TransactOpts, _hash, _result, identityBytes, signatures, nonce)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationResult(opts *bind.TransactOpts, _hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationResult", _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateBlock(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateBlock", _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateTransaction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateTransaction", _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// OracleContractAllowedModesChangedIterator is returned from FilterAllowedModesChanged and is used to iterate over the raw logs and unpacked data for AllowedModesChanged events raised by the OracleContract contract.
type OracleContractAllowedModesChangedIterator struct {
	Event *OracleContractAllowedModesChanged // Event containing the contract specifics and raw log

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
func (it *OracleContractAllowedModesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractAllowedModesChanged)
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
		it.Event = new(OracleContractAllowedModesChanged)
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
func (it *OracleContractAllowedModesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractAllowedModesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractAllowedModesChanged represents a AllowedModesChanged event raised by the OracleContract contract.
type OracleContractAllowedModesChanged struct {
	AllowThreshold bool
	AllowMultisig  bool
	AllowIBOS      bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllowedModesChanged is a free log retrieval operation binding the contract event 0x5796377bd9a4625608bba8c12a69ccd67ec138ddd6ad2e4805beace042f4bf6b.
//
// Solidity: event AllowedModesChanged(bool allowThreshold, bool allowMultisig, bool allowIBOS)
func (_OracleContract *OracleContractFilterer) FilterAllowedModesChanged(opts *bind.FilterOpts) (*OracleContractAllowedModesChangedIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "AllowedModesChanged")
	if err != nil {
		return nil, err
	}
	return &OracleContractAllowedModesChangedIterator{contract: _OracleContract.contract, event: "AllowedModesChanged", logs: logs, sub: sub}, nil
}

// WatchAllowedModesChanged is a free log subscription operation binding the contract event 0x5796377bd9a4625608bba8c12a69ccd67ec138ddd6ad2e4805beace042f4bf6b.
//
// Solidity: event AllowedModesChanged(bool allowThreshold, bool allowMultisig, bool allowIBOS)
func (_OracleContract *OracleContractFilterer) WatchAllowedModesChanged(opts *bind.WatchOpts, sink chan<- *OracleContractAllowedModesChanged) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "AllowedModesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractAllowedModesChanged)
				if err := _OracleContract.contract.UnpackLog(event, "AllowedModesChanged", log); err != nil {
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

// ParseAllowedModesChanged is a log parse operation binding the contract event 0x5796377bd9a4625608bba8c12a69ccd67ec138ddd6ad2e4805beace042f4bf6b.
//
// Solidity: event AllowedModesChanged(bool allowThreshold, bool allowMultisig, bool allowIBOS)
func (_OracleContract *OracleContractFilterer) ParseAllowedModesChanged(log types.Log) (*OracleContractAllowedModesChanged, error) {
	event := new(OracleContractAllowedModesChanged)
	if err := _OracleContract.contract.UnpackLog(event, "AllowedModesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractMultisigRequiredChangedIterator is returned from FilterMultisigRequiredChanged and is used to iterate over the raw logs and unpacked data for MultisigRequiredChanged events raised by the OracleContract contract.
type OracleContractMultisigRequiredChangedIterator struct {
	Event *OracleContractMultisigRequiredChanged // Event containing the contract specifics and raw log

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
func (it *OracleContractMultisigRequiredChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractMultisigRequiredChanged)
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
		it.Event = new(OracleContractMultisigRequiredChanged)
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
func (it *OracleContractMultisigRequiredChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractMultisigRequiredChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractMultisigRequiredChanged represents a MultisigRequiredChanged event raised by the OracleContract contract.
type OracleContractMultisigRequiredChanged struct {
	OldM *big.Int
	NewM *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMultisigRequiredChanged is a free log retrieval operation binding the contract event 0x10855fc9fd9027e533ded2b6cb96fbfcdfd4b5d0d80292acf678b66781e9363c.
//
// Solidity: event MultisigRequiredChanged(uint256 oldM, uint256 newM)
func (_OracleContract *OracleContractFilterer) FilterMultisigRequiredChanged(opts *bind.FilterOpts) (*OracleContractMultisigRequiredChangedIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "MultisigRequiredChanged")
	if err != nil {
		return nil, err
	}
	return &OracleContractMultisigRequiredChangedIterator{contract: _OracleContract.contract, event: "MultisigRequiredChanged", logs: logs, sub: sub}, nil
}

// WatchMultisigRequiredChanged is a free log subscription operation binding the contract event 0x10855fc9fd9027e533ded2b6cb96fbfcdfd4b5d0d80292acf678b66781e9363c.
//
// Solidity: event MultisigRequiredChanged(uint256 oldM, uint256 newM)
func (_OracleContract *OracleContractFilterer) WatchMultisigRequiredChanged(opts *bind.WatchOpts, sink chan<- *OracleContractMultisigRequiredChanged) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "MultisigRequiredChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractMultisigRequiredChanged)
				if err := _OracleContract.contract.UnpackLog(event, "MultisigRequiredChanged", log); err != nil {
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

// ParseMultisigRequiredChanged is a log parse operation binding the contract event 0x10855fc9fd9027e533ded2b6cb96fbfcdfd4b5d0d80292acf678b66781e9363c.
//
// Solidity: event MultisigRequiredChanged(uint256 oldM, uint256 newM)
func (_OracleContract *OracleContractFilterer) ParseMultisigRequiredChanged(log types.Log) (*OracleContractMultisigRequiredChanged, error) {
	event := new(OracleContractMultisigRequiredChanged)
	if err := _OracleContract.contract.UnpackLog(event, "MultisigRequiredChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidationRequestIterator is returned from FilterValidationRequest and is used to iterate over the raw logs and unpacked data for ValidationRequest events raised by the OracleContract contract.
type OracleContractValidationRequestIterator struct {
	Event *OracleContractValidationRequest // Event containing the contract specifics and raw log

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
func (it *OracleContractValidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationRequest)
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
		it.Event = new(OracleContractValidationRequest)
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
func (it *OracleContractValidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationRequest represents a ValidationRequest event raised by the OracleContract contract.
type OracleContractValidationRequest struct {
	Typ  uint8
	From common.Address
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidationRequest is a free log retrieval operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) FilterValidationRequest(opts *bind.FilterOpts, from []common.Address) (*OracleContractValidationRequestIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationRequestIterator{contract: _OracleContract.contract, event: "ValidationRequest", logs: logs, sub: sub}, nil
}

// WatchValidationRequest is a free log subscription operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) WatchValidationRequest(opts *bind.WatchOpts, sink chan<- *OracleContractValidationRequest, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationRequest)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
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

// ParseValidationRequest is a log parse operation binding the contract event 0x3706933cbfd265e74e347f4c40263753cecc292080a1bfd0e9fd6ce994c08396.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash)
func (_OracleContract *OracleContractFilterer) ParseValidationRequest(log types.Log) (*OracleContractValidationRequest, error) {
	event := new(OracleContractValidationRequest)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidationResponseIterator is returned from FilterValidationResponse and is used to iterate over the raw logs and unpacked data for ValidationResponse events raised by the OracleContract contract.
type OracleContractValidationResponseIterator struct {
	Event *OracleContractValidationResponse // Event containing the contract specifics and raw log

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
func (it *OracleContractValidationResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationResponse)
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
		it.Event = new(OracleContractValidationResponse)
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
func (it *OracleContractValidationResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationResponse represents a ValidationResponse event raised by the OracleContract contract.
type OracleContractValidationResponse struct {
	Typ        uint8
	Aggregator common.Address
	Hash       [32]byte
	Valid      bool
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidationResponse is a free log retrieval operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) FilterValidationResponse(opts *bind.FilterOpts, aggregator []common.Address) (*OracleContractValidationResponseIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationResponse", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationResponseIterator{contract: _OracleContract.contract, event: "ValidationResponse", logs: logs, sub: sub}, nil
}

// WatchValidationResponse is a free log subscription operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) WatchValidationResponse(opts *bind.WatchOpts, sink chan<- *OracleContractValidationResponse, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationResponse", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationResponse)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
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

// ParseValidationResponse is a log parse operation binding the contract event 0x9739d27192db56a131c86c297677afccac89aa4f39cbc5bcd62a8d10ce559675.
//
// Solidity: event ValidationResponse(uint8 typ, address indexed aggregator, bytes32 hash, bool valid, uint256 fee)
func (_OracleContract *OracleContractFilterer) ParseValidationResponse(log types.Log) (*OracleContractValidationResponse, error) {
	event := new(OracleContractValidationResponse)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
