// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package decert

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
)

// DecertCertificate is an auto generated low-level Go binding around an user-defined struct.
type DecertCertificate struct {
	Issuer    common.Address
	Recipient common.Address
	CertHash  [32]byte
	Link      string
	IssuedAt  *big.Int
}

// DecertRevokedStatus is an auto generated low-level Go binding around an user-defined struct.
type DecertRevokedStatus struct {
	IsRevoked bool
	Reason    string
	RevokedAt *big.Int
}

// DecertMetaData contains all meta data concerning the Decert contract.
var DecertMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issuedAt\",\"type\":\"uint256\"}],\"internalType\":\"structDecert.Certificate[]\",\"name\":\"_certificate\",\"type\":\"tuple[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"}],\"name\":\"certData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issuedAt\",\"type\":\"uint256\"}],\"internalType\":\"structDecert.Certificate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"changeBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_certHash\",\"type\":\"bytes32\"}],\"name\":\"hashToID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenID\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revokedAt\",\"type\":\"uint256\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_revokedTokenID\",\"type\":\"uint256\"}],\"name\":\"revokedStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRevoked\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"revokedAt\",\"type\":\"uint256\"}],\"internalType\":\"structDecert.RevokedStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newBatchSize\",\"type\":\"uint256\"}],\"name\":\"setBatchSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DecertABI is the input ABI used to generate the binding from.
// Deprecated: Use DecertMetaData.ABI instead.
var DecertABI = DecertMetaData.ABI

// Decert is an auto generated Go binding around an Ethereum contract.
type Decert struct {
	DecertCaller     // Read-only binding to the contract
	DecertTransactor // Write-only binding to the contract
	DecertFilterer   // Log filterer for contract events
}

// DecertCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecertSession struct {
	Contract     *Decert           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecertCallerSession struct {
	Contract *DecertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DecertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecertTransactorSession struct {
	Contract     *DecertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecertRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecertRaw struct {
	Contract *Decert // Generic contract binding to access the raw methods on
}

// DecertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecertCallerRaw struct {
	Contract *DecertCaller // Generic read-only contract binding to access the raw methods on
}

// DecertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecertTransactorRaw struct {
	Contract *DecertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecert creates a new instance of Decert, bound to a specific deployed contract.
func NewDecert(address common.Address, backend bind.ContractBackend) (*Decert, error) {
	contract, err := bindDecert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Decert{DecertCaller: DecertCaller{contract: contract}, DecertTransactor: DecertTransactor{contract: contract}, DecertFilterer: DecertFilterer{contract: contract}}, nil
}

// NewDecertCaller creates a new read-only instance of Decert, bound to a specific deployed contract.
func NewDecertCaller(address common.Address, caller bind.ContractCaller) (*DecertCaller, error) {
	contract, err := bindDecert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecertCaller{contract: contract}, nil
}

// NewDecertTransactor creates a new write-only instance of Decert, bound to a specific deployed contract.
func NewDecertTransactor(address common.Address, transactor bind.ContractTransactor) (*DecertTransactor, error) {
	contract, err := bindDecert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecertTransactor{contract: contract}, nil
}

// NewDecertFilterer creates a new log filterer instance of Decert, bound to a specific deployed contract.
func NewDecertFilterer(address common.Address, filterer bind.ContractFilterer) (*DecertFilterer, error) {
	contract, err := bindDecert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecertFilterer{contract: contract}, nil
}

// bindDecert binds a generic wrapper to an already deployed contract.
func bindDecert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DecertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Decert *DecertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Decert.Contract.DecertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Decert *DecertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Decert.Contract.DecertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Decert *DecertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Decert.Contract.DecertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Decert *DecertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Decert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Decert *DecertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Decert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Decert *DecertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Decert.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Decert *DecertCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Decert *DecertSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Decert.Contract.BalanceOf(&_Decert.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Decert *DecertCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Decert.Contract.BalanceOf(&_Decert.CallOpts, owner)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_Decert *DecertCaller) BatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "batchSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_Decert *DecertSession) BatchSize() (*big.Int, error) {
	return _Decert.Contract.BatchSize(&_Decert.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_Decert *DecertCallerSession) BatchSize() (*big.Int, error) {
	return _Decert.Contract.BatchSize(&_Decert.CallOpts)
}

// CertData is a free data retrieval call binding the contract method 0x1520e2ee.
//
// Solidity: function certData(uint256 _tokenID) view returns((address,address,bytes32,string,uint256))
func (_Decert *DecertCaller) CertData(opts *bind.CallOpts, _tokenID *big.Int) (DecertCertificate, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "certData", _tokenID)

	if err != nil {
		return *new(DecertCertificate), err
	}

	out0 := *abi.ConvertType(out[0], new(DecertCertificate)).(*DecertCertificate)

	return out0, err

}

// CertData is a free data retrieval call binding the contract method 0x1520e2ee.
//
// Solidity: function certData(uint256 _tokenID) view returns((address,address,bytes32,string,uint256))
func (_Decert *DecertSession) CertData(_tokenID *big.Int) (DecertCertificate, error) {
	return _Decert.Contract.CertData(&_Decert.CallOpts, _tokenID)
}

// CertData is a free data retrieval call binding the contract method 0x1520e2ee.
//
// Solidity: function certData(uint256 _tokenID) view returns((address,address,bytes32,string,uint256))
func (_Decert *DecertCallerSession) CertData(_tokenID *big.Int) (DecertCertificate, error) {
	return _Decert.Contract.CertData(&_Decert.CallOpts, _tokenID)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Decert *DecertCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Decert *DecertSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Decert.Contract.GetApproved(&_Decert.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Decert *DecertCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Decert.Contract.GetApproved(&_Decert.CallOpts, tokenId)
}

// HashToID is a free data retrieval call binding the contract method 0xdeef1e82.
//
// Solidity: function hashToID(bytes32 _certHash) view returns(uint256)
func (_Decert *DecertCaller) HashToID(opts *bind.CallOpts, _certHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "hashToID", _certHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashToID is a free data retrieval call binding the contract method 0xdeef1e82.
//
// Solidity: function hashToID(bytes32 _certHash) view returns(uint256)
func (_Decert *DecertSession) HashToID(_certHash [32]byte) (*big.Int, error) {
	return _Decert.Contract.HashToID(&_Decert.CallOpts, _certHash)
}

// HashToID is a free data retrieval call binding the contract method 0xdeef1e82.
//
// Solidity: function hashToID(bytes32 _certHash) view returns(uint256)
func (_Decert *DecertCallerSession) HashToID(_certHash [32]byte) (*big.Int, error) {
	return _Decert.Contract.HashToID(&_Decert.CallOpts, _certHash)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Decert *DecertCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Decert *DecertSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Decert.Contract.IsApprovedForAll(&_Decert.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Decert *DecertCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Decert.Contract.IsApprovedForAll(&_Decert.CallOpts, owner, operator)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_Decert *DecertCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "issuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_Decert *DecertSession) Issuer() (common.Address, error) {
	return _Decert.Contract.Issuer(&_Decert.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_Decert *DecertCallerSession) Issuer() (common.Address, error) {
	return _Decert.Contract.Issuer(&_Decert.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Decert *DecertCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Decert *DecertSession) Name() (string, error) {
	return _Decert.Contract.Name(&_Decert.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Decert *DecertCallerSession) Name() (string, error) {
	return _Decert.Contract.Name(&_Decert.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Decert *DecertCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Decert *DecertSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Decert.Contract.OwnerOf(&_Decert.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Decert *DecertCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Decert.Contract.OwnerOf(&_Decert.CallOpts, tokenId)
}

// RevokedStatus is a free data retrieval call binding the contract method 0xd2bc1f85.
//
// Solidity: function revokedStatus(uint256 _revokedTokenID) view returns((bool,string,uint256))
func (_Decert *DecertCaller) RevokedStatus(opts *bind.CallOpts, _revokedTokenID *big.Int) (DecertRevokedStatus, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "revokedStatus", _revokedTokenID)

	if err != nil {
		return *new(DecertRevokedStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DecertRevokedStatus)).(*DecertRevokedStatus)

	return out0, err

}

// RevokedStatus is a free data retrieval call binding the contract method 0xd2bc1f85.
//
// Solidity: function revokedStatus(uint256 _revokedTokenID) view returns((bool,string,uint256))
func (_Decert *DecertSession) RevokedStatus(_revokedTokenID *big.Int) (DecertRevokedStatus, error) {
	return _Decert.Contract.RevokedStatus(&_Decert.CallOpts, _revokedTokenID)
}

// RevokedStatus is a free data retrieval call binding the contract method 0xd2bc1f85.
//
// Solidity: function revokedStatus(uint256 _revokedTokenID) view returns((bool,string,uint256))
func (_Decert *DecertCallerSession) RevokedStatus(_revokedTokenID *big.Int) (DecertRevokedStatus, error) {
	return _Decert.Contract.RevokedStatus(&_Decert.CallOpts, _revokedTokenID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Decert *DecertCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Decert *DecertSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Decert.Contract.SupportsInterface(&_Decert.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Decert *DecertCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Decert.Contract.SupportsInterface(&_Decert.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Decert *DecertCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Decert *DecertSession) Symbol() (string, error) {
	return _Decert.Contract.Symbol(&_Decert.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Decert *DecertCallerSession) Symbol() (string, error) {
	return _Decert.Contract.Symbol(&_Decert.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Decert *DecertCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Decert *DecertSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Decert.Contract.TokenByIndex(&_Decert.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Decert *DecertCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Decert.Contract.TokenByIndex(&_Decert.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Decert *DecertCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Decert *DecertSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Decert.Contract.TokenOfOwnerByIndex(&_Decert.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Decert *DecertCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Decert.Contract.TokenOfOwnerByIndex(&_Decert.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Decert *DecertCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Decert *DecertSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Decert.Contract.TokenURI(&_Decert.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Decert *DecertCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Decert.Contract.TokenURI(&_Decert.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Decert *DecertCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Decert.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Decert *DecertSession) TotalSupply() (*big.Int, error) {
	return _Decert.Contract.TotalSupply(&_Decert.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Decert *DecertCallerSession) TotalSupply() (*big.Int, error) {
	return _Decert.Contract.TotalSupply(&_Decert.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Decert *DecertTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Decert *DecertSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.Approve(&_Decert.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Decert *DecertTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.Approve(&_Decert.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x30ead00d.
//
// Solidity: function batchMint((address,address,bytes32,string,uint256)[] _certificate) returns()
func (_Decert *DecertTransactor) BatchMint(opts *bind.TransactOpts, _certificate []DecertCertificate) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "batchMint", _certificate)
}

// BatchMint is a paid mutator transaction binding the contract method 0x30ead00d.
//
// Solidity: function batchMint((address,address,bytes32,string,uint256)[] _certificate) returns()
func (_Decert *DecertSession) BatchMint(_certificate []DecertCertificate) (*types.Transaction, error) {
	return _Decert.Contract.BatchMint(&_Decert.TransactOpts, _certificate)
}

// BatchMint is a paid mutator transaction binding the contract method 0x30ead00d.
//
// Solidity: function batchMint((address,address,bytes32,string,uint256)[] _certificate) returns()
func (_Decert *DecertTransactorSession) BatchMint(_certificate []DecertCertificate) (*types.Transaction, error) {
	return _Decert.Contract.BatchMint(&_Decert.TransactOpts, _certificate)
}

// ChangeBaseURI is a paid mutator transaction binding the contract method 0x39a0c6f9.
//
// Solidity: function changeBaseURI(string newBaseURI) returns()
func (_Decert *DecertTransactor) ChangeBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "changeBaseURI", newBaseURI)
}

// ChangeBaseURI is a paid mutator transaction binding the contract method 0x39a0c6f9.
//
// Solidity: function changeBaseURI(string newBaseURI) returns()
func (_Decert *DecertSession) ChangeBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Decert.Contract.ChangeBaseURI(&_Decert.TransactOpts, newBaseURI)
}

// ChangeBaseURI is a paid mutator transaction binding the contract method 0x39a0c6f9.
//
// Solidity: function changeBaseURI(string newBaseURI) returns()
func (_Decert *DecertTransactorSession) ChangeBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Decert.Contract.ChangeBaseURI(&_Decert.TransactOpts, newBaseURI)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xa35110cd.
//
// Solidity: function revokeCertificate(uint256[] _tokenID, string _reason, uint256 revokedAt) returns()
func (_Decert *DecertTransactor) RevokeCertificate(opts *bind.TransactOpts, _tokenID []*big.Int, _reason string, revokedAt *big.Int) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "revokeCertificate", _tokenID, _reason, revokedAt)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xa35110cd.
//
// Solidity: function revokeCertificate(uint256[] _tokenID, string _reason, uint256 revokedAt) returns()
func (_Decert *DecertSession) RevokeCertificate(_tokenID []*big.Int, _reason string, revokedAt *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.RevokeCertificate(&_Decert.TransactOpts, _tokenID, _reason, revokedAt)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xa35110cd.
//
// Solidity: function revokeCertificate(uint256[] _tokenID, string _reason, uint256 revokedAt) returns()
func (_Decert *DecertTransactorSession) RevokeCertificate(_tokenID []*big.Int, _reason string, revokedAt *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.RevokeCertificate(&_Decert.TransactOpts, _tokenID, _reason, revokedAt)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Decert *DecertTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Decert *DecertSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.SafeTransferFrom(&_Decert.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Decert *DecertTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.SafeTransferFrom(&_Decert.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Decert *DecertTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Decert *DecertSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Decert.Contract.SafeTransferFrom0(&_Decert.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Decert *DecertTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Decert.Contract.SafeTransferFrom0(&_Decert.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Decert *DecertTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Decert *DecertSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Decert.Contract.SetApprovalForAll(&_Decert.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Decert *DecertTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Decert.Contract.SetApprovalForAll(&_Decert.TransactOpts, operator, approved)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x576f35e3.
//
// Solidity: function setBatchSize(uint256 _newBatchSize) returns()
func (_Decert *DecertTransactor) SetBatchSize(opts *bind.TransactOpts, _newBatchSize *big.Int) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "setBatchSize", _newBatchSize)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x576f35e3.
//
// Solidity: function setBatchSize(uint256 _newBatchSize) returns()
func (_Decert *DecertSession) SetBatchSize(_newBatchSize *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.SetBatchSize(&_Decert.TransactOpts, _newBatchSize)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x576f35e3.
//
// Solidity: function setBatchSize(uint256 _newBatchSize) returns()
func (_Decert *DecertTransactorSession) SetBatchSize(_newBatchSize *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.SetBatchSize(&_Decert.TransactOpts, _newBatchSize)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Decert *DecertTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Decert *DecertSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.TransferFrom(&_Decert.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Decert *DecertTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Decert.Contract.TransferFrom(&_Decert.TransactOpts, from, to, tokenId)
}

// DecertApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Decert contract.
type DecertApprovalIterator struct {
	Event *DecertApproval // Event containing the contract specifics and raw log

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
func (it *DecertApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecertApproval)
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
		it.Event = new(DecertApproval)
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
func (it *DecertApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecertApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecertApproval represents a Approval event raised by the Decert contract.
type DecertApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Decert *DecertFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DecertApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Decert.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DecertApprovalIterator{contract: _Decert.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Decert *DecertFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DecertApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Decert.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecertApproval)
				if err := _Decert.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Decert *DecertFilterer) ParseApproval(log types.Log) (*DecertApproval, error) {
	event := new(DecertApproval)
	if err := _Decert.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecertApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Decert contract.
type DecertApprovalForAllIterator struct {
	Event *DecertApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DecertApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecertApprovalForAll)
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
		it.Event = new(DecertApprovalForAll)
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
func (it *DecertApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecertApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecertApprovalForAll represents a ApprovalForAll event raised by the Decert contract.
type DecertApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Decert *DecertFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DecertApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Decert.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DecertApprovalForAllIterator{contract: _Decert.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Decert *DecertFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DecertApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Decert.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecertApprovalForAll)
				if err := _Decert.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Decert *DecertFilterer) ParseApprovalForAll(log types.Log) (*DecertApprovalForAll, error) {
	event := new(DecertApprovalForAll)
	if err := _Decert.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecertTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Decert contract.
type DecertTransferIterator struct {
	Event *DecertTransfer // Event containing the contract specifics and raw log

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
func (it *DecertTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecertTransfer)
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
		it.Event = new(DecertTransfer)
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
func (it *DecertTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecertTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecertTransfer represents a Transfer event raised by the Decert contract.
type DecertTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Decert *DecertFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DecertTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Decert.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DecertTransferIterator{contract: _Decert.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Decert *DecertFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DecertTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Decert.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecertTransfer)
				if err := _Decert.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Decert *DecertFilterer) ParseTransfer(log types.Log) (*DecertTransfer, error) {
	event := new(DecertTransfer)
	if err := _Decert.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
