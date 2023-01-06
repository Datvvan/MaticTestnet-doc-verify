// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity_card\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash_docs\",\"type\":\"string\"}],\"name\":\"addContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"identity_card\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"company\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c39806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806327edd0611461003b57806384270c111461006b575b600080fd5b61005560048036038101906100509190610632565b610087565b6040516100629190610729565b60405180910390f35b6100856004803603810190610080919061074b565b610426565b005b606060005b6000805490508110156103e65760008082815481106100ae576100ad610822565b5b90600052602060002090600402016040518060800160405290816000820180546100d790610880565b80601f016020809104026020016040519081016040528092919081815260200182805461010390610880565b80156101505780601f1061012557610100808354040283529160200191610150565b820191906000526020600020905b81548152906001019060200180831161013357829003601f168201915b5050505050815260200160018201805461016990610880565b80601f016020809104026020016040519081016040528092919081815260200182805461019590610880565b80156101e25780601f106101b7576101008083540402835291602001916101e2565b820191906000526020600020905b8154815290600101906020018083116101c557829003601f168201915b505050505081526020016002820180546101fb90610880565b80601f016020809104026020016040519081016040528092919081815260200182805461022790610880565b80156102745780601f1061024957610100808354040283529160200191610274565b820191906000526020600020905b81548152906001019060200180831161025757829003601f168201915b5050505050815260200160038201805461028d90610880565b80601f01602080910402602001604051908101604052809291908181526020018280546102b990610880565b80156103065780601f106102db57610100808354040283529160200191610306565b820191906000526020600020905b8154815290600101906020018083116102e957829003601f168201915b50505050508152505090508460405160200161032291906108ed565b60405160208183030381529060405280519060200120816040015160405160200161034d91906108ed565b604051602081830303815290604052805190602001201480156103bf57508360405160200161037c91906108ed565b6040516020818303038152906040528051906020012081602001516040516020016103a791906108ed565b60405160208183030381529060405280519060200120145b156103d257806060015192505050610420565b5080806103de9061093d565b91505061008c565b6040518060400160405280600981526020017f4e6f7420666f756e6400000000000000000000000000000000000000000000008152509150505b92915050565b600060405180608001604052808681526020018581526020018481526020018381525090506000819080600181540180825580915050600190039060005260206000209060040201600090919091909150600082015181600001908161048c9190610b31565b5060208201518160010190816104a29190610b31565b5060408201518160020190816104b89190610b31565b5060608201518160030190816104ce9190610b31565b5050505050505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61053f826104f6565b810181811067ffffffffffffffff8211171561055e5761055d610507565b5b80604052505050565b60006105716104d8565b905061057d8282610536565b919050565b600067ffffffffffffffff82111561059d5761059c610507565b5b6105a6826104f6565b9050602081019050919050565b82818337600083830152505050565b60006105d56105d084610582565b610567565b9050828152602081018484840111156105f1576105f06104f1565b5b6105fc8482856105b3565b509392505050565b600082601f830112610619576106186104ec565b5b81356106298482602086016105c2565b91505092915050565b60008060408385031215610649576106486104e2565b5b600083013567ffffffffffffffff811115610667576106666104e7565b5b61067385828601610604565b925050602083013567ffffffffffffffff811115610694576106936104e7565b5b6106a085828601610604565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106e45780820151818401526020810190506106c9565b60008484015250505050565b60006106fb826106aa565b61070581856106b5565b93506107158185602086016106c6565b61071e816104f6565b840191505092915050565b6000602082019050818103600083015261074381846106f0565b905092915050565b60008060008060808587031215610765576107646104e2565b5b600085013567ffffffffffffffff811115610783576107826104e7565b5b61078f87828801610604565b945050602085013567ffffffffffffffff8111156107b0576107af6104e7565b5b6107bc87828801610604565b935050604085013567ffffffffffffffff8111156107dd576107dc6104e7565b5b6107e987828801610604565b925050606085013567ffffffffffffffff81111561080a576108096104e7565b5b61081687828801610604565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061089857607f821691505b6020821081036108ab576108aa610851565b5b50919050565b600081905092915050565b60006108c7826106aa565b6108d181856108b1565b93506108e18185602086016106c6565b80840191505092915050565b60006108f982846108bc565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b600061094882610933565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361097a57610979610904565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026109e77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826109aa565b6109f186836109aa565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610a2e610a29610a2484610933565b610a09565b610933565b9050919050565b6000819050919050565b610a4883610a13565b610a5c610a5482610a35565b8484546109b7565b825550505050565b600090565b610a71610a64565b610a7c818484610a3f565b505050565b5b81811015610aa057610a95600082610a69565b600181019050610a82565b5050565b601f821115610ae557610ab681610985565b610abf8461099a565b81016020851015610ace578190505b610ae2610ada8561099a565b830182610a81565b50505b505050565b600082821c905092915050565b6000610b0860001984600802610aea565b1980831691505092915050565b6000610b218383610af7565b9150826002028217905092915050565b610b3a826106aa565b67ffffffffffffffff811115610b5357610b52610507565b5b610b5d8254610880565b610b68828285610aa4565b600060209050601f831160018114610b9b5760008415610b89578287015190505b610b938582610b15565b865550610bfb565b601f198416610ba986610985565b60005b82811015610bd157848901518255600182019150602085019450602081019050610bac565b86831015610bee5784890151610bea601f891682610af7565b8355505b6001600288020188555050505b50505050505056fea26469706673582212207bd2b59c5b2fc1ba6752b4e9fb7933265a1e48a944c7984ecc151cdb5173bc6a64736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetContract is a free data retrieval call binding the contract method 0x27edd061.
//
// Solidity: function getContract(string identity_card, string company) view returns(string)
func (_Api *ApiCaller) GetContract(opts *bind.CallOpts, identity_card string, company string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getContract", identity_card, company)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x27edd061.
//
// Solidity: function getContract(string identity_card, string company) view returns(string)
func (_Api *ApiSession) GetContract(identity_card string, company string) (string, error) {
	return _Api.Contract.GetContract(&_Api.CallOpts, identity_card, company)
}

// GetContract is a free data retrieval call binding the contract method 0x27edd061.
//
// Solidity: function getContract(string identity_card, string company) view returns(string)
func (_Api *ApiCallerSession) GetContract(identity_card string, company string) (string, error) {
	return _Api.Contract.GetContract(&_Api.CallOpts, identity_card, company)
}

// AddContract is a paid mutator transaction binding the contract method 0x84270c11.
//
// Solidity: function addContract(string name, string company, string identity_card, string hash_docs) returns()
func (_Api *ApiTransactor) AddContract(opts *bind.TransactOpts, name string, company string, identity_card string, hash_docs string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addContract", name, company, identity_card, hash_docs)
}

// AddContract is a paid mutator transaction binding the contract method 0x84270c11.
//
// Solidity: function addContract(string name, string company, string identity_card, string hash_docs) returns()
func (_Api *ApiSession) AddContract(name string, company string, identity_card string, hash_docs string) (*types.Transaction, error) {
	return _Api.Contract.AddContract(&_Api.TransactOpts, name, company, identity_card, hash_docs)
}

// AddContract is a paid mutator transaction binding the contract method 0x84270c11.
//
// Solidity: function addContract(string name, string company, string identity_card, string hash_docs) returns()
func (_Api *ApiTransactorSession) AddContract(name string, company string, identity_card string, hash_docs string) (*types.Transaction, error) {
	return _Api.Contract.AddContract(&_Api.TransactOpts, name, company, identity_card, hash_docs)
}
