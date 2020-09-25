// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vote

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

// VoteABI is the input ABI used to generate the binding from.
const VoteABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_info\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"nominees\",\"type\":\"string[]\"}],\"name\":\"AddVoteInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Nominees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"SendVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"introduction\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VoteBin is the compiled bytecode used for deploying new contracts.
var VoteBin = "0x608060405260008060146101000a81548160ff0219169083151502179055503480156200002b57600080fd5b5060405162001085380380620010858339818101604052620000519190810190620001fd565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360039080519060200190620000a9929190620000db565b5082600181905550816002819055508060049080519060200190620000d0929190620000db565b505050505062000350565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200011e57805160ff19168380011785556200014f565b828001600101855582156200014f579182015b828111156200014e57825182559160200191906001019062000131565b5b5090506200015e919062000162565b5090565b6200018791905b808211156200018357600081600090555060010162000169565b5090565b90565b600082601f8301126200019c57600080fd5b8151620001b3620001ad82620002c9565b6200029b565b91508082526020830160208301858383011115620001d057600080fd5b620001dd83828462000300565b50505092915050565b600081519050620001f78162000336565b92915050565b600080600080608085870312156200021457600080fd5b600085015167ffffffffffffffff8111156200022f57600080fd5b6200023d878288016200018a565b94505060206200025087828801620001e6565b93505060406200026387828801620001e6565b925050606085015167ffffffffffffffff8111156200028157600080fd5b6200028f878288016200018a565b91505092959194509250565b6000604051905081810181811067ffffffffffffffff82111715620002bf57600080fd5b8060405250919050565b600067ffffffffffffffff821115620002e157600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b60005b838110156200032057808201518184015260208101905062000303565b8381111562000330576000848401525b50505050565b6200034181620002f6565b81146200034d57600080fd5b50565b610d2580620003606000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b146100ef578063b6cc8a931461010d578063bef4876b14610129578063fc85ab641461014757610088565b806306fdde031461008d5780632aebcbb6146100ab57806356b6693f146100b557806370c268de146100d3575b600080fd5b610095610179565b6040516100a29190610af3565b60405180910390f35b6100b3610217565b005b6100bd61028d565b6040516100ca9190610af3565b60405180910390f35b6100ed60048036036100e8919081019061096d565b61032b565b005b6100f761044d565b6040516101049190610abd565b60405180910390f35b610127600480360361012291908101906109ae565b610472565b005b61013161059d565b60405161013e9190610ad8565b60405180910390f35b610161600480360361015c91908101906109ae565b6105b0565b60405161017093929190610b55565b60405180910390f35b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561020f5780601f106101e45761010080835404028352916020019161020f565b820191906000526020600020905b8154815290600101906020018083116101f257829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461027057600080fd5b6001600060146101000a81548160ff021916908315150217905550565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103235780601f106102f857610100808354040283529160200191610323565b820191906000526020600020905b81548152906001019060200180831161030657829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461038457600080fd5b6000600580549050905060008090505b8251811015610448576103a56107ca565b60405180606001604052808385018152602001600081526020018584815181106103cb57fe5b6020026020010151815250905060058190806001815401808255809150509060018203906000526020600020906003020160009091929091909150600082015181600001556020820151816001015560408201518160020190805190602001906104369291906107eb565b50505050508080600101915050610394565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60001515600060149054906101000a900460ff1615151461049257600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104eb57600080fd5b600154421180156104fd575060025442105b61053c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053390610b15565b60405180910390fd5b6000811015801561055257506005805490508111155b610591576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058890610b35565b60405180910390fd5b61059a8161067f565b50565b600060149054906101000a900460ff1681565b600581815481106105bd57fe5b9060005260206000209060030201600091509050806000015490806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106755780601f1061064a57610100808354040283529160200191610675565b820191906000526020600020905b81548152906001019060200180831161065857829003601f168201915b5050505050905083565b6106876107ca565b6005828154811061069457fe5b90600052602060002090600302016040518060600160405290816000820154815260200160018201548152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561075a5780601f1061072f5761010080835404028352916020019161075a565b820191906000526020600020905b81548152906001019060200180831161073d57829003601f168201915b50505050508152505090506001816020018181510191508181525050806005838154811061078457fe5b9060005260206000209060030201600082015181600001556020820151816001015560408201518160020190805190602001906107c29291906107eb565b509050505050565b60405180606001604052806000815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061082c57805160ff191683800117855561085a565b8280016001018555821561085a579182015b8281111561085957825182559160200191906001019061083e565b5b509050610867919061086b565b5090565b61088d91905b80821115610889576000816000905550600101610871565b5090565b90565b600082601f8301126108a157600080fd5b81356108b46108af82610bc0565b610b93565b9150818183526020840193506020810190508360005b838110156108fa57813586016108e08882610904565b8452602084019350602083019250506001810190506108ca565b5050505092915050565b600082601f83011261091557600080fd5b813561092861092382610be8565b610b93565b9150808252602083016020830185838301111561094457600080fd5b61094f838284610c78565b50505092915050565b60008135905061096781610ccb565b92915050565b60006020828403121561097f57600080fd5b600082013567ffffffffffffffff81111561099957600080fd5b6109a584828501610890565b91505092915050565b6000602082840312156109c057600080fd5b60006109ce84828501610958565b91505092915050565b6109e081610c30565b82525050565b6109ef81610c42565b82525050565b6000610a0082610c14565b610a0a8185610c1f565b9350610a1a818560208601610c87565b610a2381610cba565b840191505092915050565b6000610a3b600b83610c1f565b91507f6f7574206f662074696d650000000000000000000000000000000000000000006000830152602082019050919050565b6000610a7b601083610c1f565b91507f756e636f7272656374206e756d626572000000000000000000000000000000006000830152602082019050919050565b610ab781610c6e565b82525050565b6000602082019050610ad260008301846109d7565b92915050565b6000602082019050610aed60008301846109e6565b92915050565b60006020820190508181036000830152610b0d81846109f5565b905092915050565b60006020820190508181036000830152610b2e81610a2e565b9050919050565b60006020820190508181036000830152610b4e81610a6e565b9050919050565b6000606082019050610b6a6000830186610aae565b610b776020830185610aae565b8181036040830152610b8981846109f5565b9050949350505050565b6000604051905081810181811067ffffffffffffffff82111715610bb657600080fd5b8060405250919050565b600067ffffffffffffffff821115610bd757600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610bff57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000610c3b82610c4e565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610ca5578082015181840152602081019050610c8a565b83811115610cb4576000848401525b50505050565b6000601f19601f8301169050919050565b610cd481610c6e565b8114610cdf57600080fd5b5056fea365627a7a72315820fcd5d9a6bbf9f81581bc57f64dfdf22bc8544fab93f78f4276e83f677ef2361d6c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _begin *big.Int, _end *big.Int, _info string) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VoteBin), backend, _name, _begin, _end, _info)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// Nominees is a free data retrieval call binding the contract method 0xfc85ab64.
//
// Solidity: function Nominees(uint256 ) constant returns(uint256 id, uint256 vote, string source)
func (_Vote *VoteCaller) Nominees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Vote   *big.Int
	Source string
}, error) {
	ret := new(struct {
		Id     *big.Int
		Vote   *big.Int
		Source string
	})
	out := ret
	err := _Vote.contract.Call(opts, out, "Nominees", arg0)
	return *ret, err
}

// Nominees is a free data retrieval call binding the contract method 0xfc85ab64.
//
// Solidity: function Nominees(uint256 ) constant returns(uint256 id, uint256 vote, string source)
func (_Vote *VoteSession) Nominees(arg0 *big.Int) (struct {
	Id     *big.Int
	Vote   *big.Int
	Source string
}, error) {
	return _Vote.Contract.Nominees(&_Vote.CallOpts, arg0)
}

// Nominees is a free data retrieval call binding the contract method 0xfc85ab64.
//
// Solidity: function Nominees(uint256 ) constant returns(uint256 id, uint256 vote, string source)
func (_Vote *VoteCallerSession) Nominees(arg0 *big.Int) (struct {
	Id     *big.Int
	Vote   *big.Int
	Source string
}, error) {
	return _Vote.Contract.Nominees(&_Vote.CallOpts, arg0)
}

// Finished is a free data retrieval call binding the contract method 0xbef4876b.
//
// Solidity: function finished() constant returns(bool)
func (_Vote *VoteCaller) Finished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vote.contract.Call(opts, out, "finished")
	return *ret0, err
}

// Finished is a free data retrieval call binding the contract method 0xbef4876b.
//
// Solidity: function finished() constant returns(bool)
func (_Vote *VoteSession) Finished() (bool, error) {
	return _Vote.Contract.Finished(&_Vote.CallOpts)
}

// Finished is a free data retrieval call binding the contract method 0xbef4876b.
//
// Solidity: function finished() constant returns(bool)
func (_Vote *VoteCallerSession) Finished() (bool, error) {
	return _Vote.Contract.Finished(&_Vote.CallOpts)
}

// Introduction is a free data retrieval call binding the contract method 0x56b6693f.
//
// Solidity: function introduction() constant returns(string)
func (_Vote *VoteCaller) Introduction(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Vote.contract.Call(opts, out, "introduction")
	return *ret0, err
}

// Introduction is a free data retrieval call binding the contract method 0x56b6693f.
//
// Solidity: function introduction() constant returns(string)
func (_Vote *VoteSession) Introduction() (string, error) {
	return _Vote.Contract.Introduction(&_Vote.CallOpts)
}

// Introduction is a free data retrieval call binding the contract method 0x56b6693f.
//
// Solidity: function introduction() constant returns(string)
func (_Vote *VoteCallerSession) Introduction() (string, error) {
	return _Vote.Contract.Introduction(&_Vote.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Vote *VoteCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Vote.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Vote *VoteSession) Name() (string, error) {
	return _Vote.Contract.Name(&_Vote.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Vote *VoteCallerSession) Name() (string, error) {
	return _Vote.Contract.Name(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vote *VoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vote.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vote *VoteSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vote *VoteCallerSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// AddVoteInfo is a paid mutator transaction binding the contract method 0x70c268de.
//
// Solidity: function AddVoteInfo(string[] nominees) returns()
func (_Vote *VoteTransactor) AddVoteInfo(opts *bind.TransactOpts, nominees []string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "AddVoteInfo", nominees)
}

// AddVoteInfo is a paid mutator transaction binding the contract method 0x70c268de.
//
// Solidity: function AddVoteInfo(string[] nominees) returns()
func (_Vote *VoteSession) AddVoteInfo(nominees []string) (*types.Transaction, error) {
	return _Vote.Contract.AddVoteInfo(&_Vote.TransactOpts, nominees)
}

// AddVoteInfo is a paid mutator transaction binding the contract method 0x70c268de.
//
// Solidity: function AddVoteInfo(string[] nominees) returns()
func (_Vote *VoteTransactorSession) AddVoteInfo(nominees []string) (*types.Transaction, error) {
	return _Vote.Contract.AddVoteInfo(&_Vote.TransactOpts, nominees)
}

// SendVote is a paid mutator transaction binding the contract method 0xb6cc8a93.
//
// Solidity: function SendVote(uint256 _id) returns()
func (_Vote *VoteTransactor) SendVote(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "SendVote", _id)
}

// SendVote is a paid mutator transaction binding the contract method 0xb6cc8a93.
//
// Solidity: function SendVote(uint256 _id) returns()
func (_Vote *VoteSession) SendVote(_id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SendVote(&_Vote.TransactOpts, _id)
}

// SendVote is a paid mutator transaction binding the contract method 0xb6cc8a93.
//
// Solidity: function SendVote(uint256 _id) returns()
func (_Vote *VoteTransactorSession) SendVote(_id *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.SendVote(&_Vote.TransactOpts, _id)
}

// FinishVote is a paid mutator transaction binding the contract method 0x2aebcbb6.
//
// Solidity: function finishVote() returns()
func (_Vote *VoteTransactor) FinishVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "finishVote")
}

// FinishVote is a paid mutator transaction binding the contract method 0x2aebcbb6.
//
// Solidity: function finishVote() returns()
func (_Vote *VoteSession) FinishVote() (*types.Transaction, error) {
	return _Vote.Contract.FinishVote(&_Vote.TransactOpts)
}

// FinishVote is a paid mutator transaction binding the contract method 0x2aebcbb6.
//
// Solidity: function finishVote() returns()
func (_Vote *VoteTransactorSession) FinishVote() (*types.Transaction, error) {
	return _Vote.Contract.FinishVote(&_Vote.TransactOpts)
}
