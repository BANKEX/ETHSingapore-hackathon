// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"name\":\"newBlocks\",\"type\":\"bytes\"},{\"name\":\"protectedBlockNumber\",\"type\":\"uint256\"},{\"name\":\"protectedBlockHash\",\"type\":\"address\"}],\"name\":\"submitBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"name\":\"newBlocks\",\"type\":\"bytes\"},{\"name\":\"protectedBlockNumber\",\"type\":\"uint256\"},{\"name\":\"protectedBlockHash\",\"type\":\"address\"},{\"name\":\"rsv\",\"type\":\"bytes\"}],\"name\":\"submitBlocksSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blocksLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"length\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BlocksSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3610daa806100cf6000396000f30060806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806347097aea146100935780636680c54814610144578063715018a61461023b5780638ce0b5a2146102525780638da5cb5b1461027d5780638f32d59b146102d4578063f25b3f9914610303578063f2fde38b14610370575b600080fd5b34801561009f57600080fd5b5061012e60048036038101908080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103b3565b6040518082815260200191505060405180910390f35b34801561015057600080fd5b5061022560048036038101908080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506103dd565b6040518082815260200191505060405180910390f35b34801561024757600080fd5b506102506105dd565b005b34801561025e57600080fd5b506102676106af565b6040518082815260200191505060405180910390f35b34801561028957600080fd5b506102926106bc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102e057600080fd5b506102e96106e5565b604051808215151515815260200191505060405180910390f35b34801561030f57600080fd5b5061032e6004803603810190808035906020019092919050505061073c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561037c57600080fd5b506103b1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061077f565b005b60006103bd6106e5565b15156103c857600080fd5b6103d48585858561079e565b50949350505050565b6000806000878787876040516020018085815260200184805190602001908083835b60208310151561042457805182526020820191506020810190506020830392506103ff565b6001836020036101000a0380198251168184511680821785525050505050509050018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014019450505050506040516020818303038152906040526040518082805190602001908083835b6020831015156104d757805182526020820191506020810190506020830392506104b2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915061050f82610a3e565b905061051b8185610af9565b73ffffffffffffffffffffffffffffffffffffffff166105396106bc565b73ffffffffffffffffffffffffffffffffffffffff161415156105c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f496e76616c6964207369676e617475726500000000000000000000000000000081525060200191505060405180910390fd5b6105d08888888861079e565b9250505095945050505050565b6105e56106e5565b15156105f057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600180549050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600060018281548110151561074d57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6107876106e5565b151561079257600080fd5b61079b81610bf1565b50565b600080600080600080601489518115156107b457fe5b0494506001805490508a141515610833576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f496e76616c69642066726f6d496e64657800000000000000000000000000000081525060200191505060405180910390fd5b60008a14806108a657508673ffffffffffffffffffffffffffffffffffffffff1660018981548110151561086357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b151561091a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f57726f6e672070726f74656374656420626c6f636b206e756d6265720000000081525060200191505060405180910390fd5b6109328a600180549050610ceb90919063ffffffff16565b9350610947858b610d0c90919063ffffffff16565b6001816109549190610d2d565b508392505b848310156109e6576014830260200190506c01000000000000000000000000818a0151049150816001848c0181548110151561099157fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508280600101935050610959565b84841015610a2c576001805490507ff32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99426040518082815260200191505060405180910390a25b83850395505050505050949350505050565b60008160405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c0182600019166000191681526020019150506040516020818303038152906040526040518082805190602001908083835b602083101515610ac55780518252602082019150602081019050602083039250610aa0565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209050919050565b60008060008060418551141515610b135760009350610be8565b6020850151925060408501519150606085015160001a9050601b8160ff161015610b3e57601b810190505b601b8160ff1614158015610b565750601c8160ff1614155b15610b645760009350610be8565b600186828585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015610bdb573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610c2d57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080838311151515610cfd57600080fd5b82840390508091505092915050565b6000808284019050838110151515610d2357600080fd5b8091505092915050565b815481835581811115610d5457818360005260206000209182019101610d539190610d59565b5b505050565b610d7b91905b80821115610d77576000816000905550600101610d5f565b5090565b905600a165627a7a723058201687e0e47f25f617374be008a495eafa0aa15aed0fd7fbfdaecbc7ac2f37f3d90029`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_Store *StoreCaller) Blocks(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "blocks", i)
	return *ret0, err
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_Store *StoreSession) Blocks(i *big.Int) (common.Address, error) {
	return _Store.Contract.Blocks(&_Store.CallOpts, i)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(i uint256) constant returns(address)
func (_Store *StoreCallerSession) Blocks(i *big.Int) (common.Address, error) {
	return _Store.Contract.Blocks(&_Store.CallOpts, i)
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_Store *StoreCaller) BlocksLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "blocksLength")
	return *ret0, err
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_Store *StoreSession) BlocksLength() (*big.Int, error) {
	return _Store.Contract.BlocksLength(&_Store.CallOpts)
}

// BlocksLength is a free data retrieval call binding the contract method 0x8ce0b5a2.
//
// Solidity: function blocksLength() constant returns(uint256)
func (_Store *StoreCallerSession) BlocksLength() (*big.Int, error) {
	return _Store.Contract.BlocksLength(&_Store.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Store *StoreCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Store *StoreSession) IsOwner() (bool, error) {
	return _Store.Contract.IsOwner(&_Store.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Store *StoreCallerSession) IsOwner() (bool, error) {
	return _Store.Contract.IsOwner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x47097aea.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address) returns(uint256)
func (_Store *StoreTransactor) SubmitBlocks(opts *bind.TransactOpts, fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "submitBlocks", fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x47097aea.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address) returns(uint256)
func (_Store *StoreSession) SubmitBlocks(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocks(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x47097aea.
//
// Solidity: function submitBlocks(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address) returns(uint256)
func (_Store *StoreTransactorSession) SubmitBlocks(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocks(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x6680c548.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address, rsv bytes) returns(uint256)
func (_Store *StoreTransactor) SubmitBlocksSigned(opts *bind.TransactOpts, fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address, rsv []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "submitBlocksSigned", fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash, rsv)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x6680c548.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address, rsv bytes) returns(uint256)
func (_Store *StoreSession) SubmitBlocksSigned(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address, rsv []byte) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocksSigned(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash, rsv)
}

// SubmitBlocksSigned is a paid mutator transaction binding the contract method 0x6680c548.
//
// Solidity: function submitBlocksSigned(fromIndex uint256, newBlocks bytes, protectedBlockNumber uint256, protectedBlockHash address, rsv bytes) returns(uint256)
func (_Store *StoreTransactorSession) SubmitBlocksSigned(fromIndex *big.Int, newBlocks []byte, protectedBlockNumber *big.Int, protectedBlockHash common.Address, rsv []byte) (*types.Transaction, error) {
	return _Store.Contract.SubmitBlocksSigned(&_Store.TransactOpts, fromIndex, newBlocks, protectedBlockNumber, protectedBlockHash, rsv)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Store *StoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// StoreBlocksSubmittedIterator is returned from FilterBlocksSubmitted and is used to iterate over the raw logs and unpacked data for BlocksSubmitted events raised by the Store contract.
type StoreBlocksSubmittedIterator struct {
	Event *StoreBlocksSubmitted // Event containing the contract specifics and raw log

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
func (it *StoreBlocksSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBlocksSubmitted)
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
		it.Event = new(StoreBlocksSubmitted)
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
func (it *StoreBlocksSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBlocksSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBlocksSubmitted represents a BlocksSubmitted event raised by the Store contract.
type StoreBlocksSubmitted struct {
	Length *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlocksSubmitted is a free log retrieval operation binding the contract event 0xf32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99.
//
// Solidity: e BlocksSubmitted(length indexed uint256, time uint256)
func (_Store *StoreFilterer) FilterBlocksSubmitted(opts *bind.FilterOpts, length []*big.Int) (*StoreBlocksSubmittedIterator, error) {

	var lengthRule []interface{}
	for _, lengthItem := range length {
		lengthRule = append(lengthRule, lengthItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "BlocksSubmitted", lengthRule)
	if err != nil {
		return nil, err
	}
	return &StoreBlocksSubmittedIterator{contract: _Store.contract, event: "BlocksSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlocksSubmitted is a free log subscription operation binding the contract event 0xf32c68e7736e0f3f51cf7e6d33003550534f6ce10665ed8430cd92d66b0bbb99.
//
// Solidity: e BlocksSubmitted(length indexed uint256, time uint256)
func (_Store *StoreFilterer) WatchBlocksSubmitted(opts *bind.WatchOpts, sink chan<- *StoreBlocksSubmitted, length []*big.Int) (event.Subscription, error) {

	var lengthRule []interface{}
	for _, lengthItem := range length {
		lengthRule = append(lengthRule, lengthItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "BlocksSubmitted", lengthRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBlocksSubmitted)
				if err := _Store.contract.UnpackLog(event, "BlocksSubmitted", log); err != nil {
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

// StoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Store contract.
type StoreOwnershipTransferredIterator struct {
	Event *StoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOwnershipTransferred)
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
		it.Event = new(StoreOwnershipTransferred)
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
func (it *StoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOwnershipTransferred represents a OwnershipTransferred event raised by the Store contract.
type StoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Store *StoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StoreOwnershipTransferredIterator{contract: _Store.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Store *StoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOwnershipTransferred)
				if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
