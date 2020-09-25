// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package foundation

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

// FoundationABI is the input ABI used to generate the binding from.
const FoundationABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"finishVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_info\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minPar\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPar\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"nominees\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nominee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"LogActivateNominee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"council\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"}],\"name\":\"LogNewCouncil\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nominee\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"LogNewNominees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"council\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"name\":\"LogVoteNominee\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"safeWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"StartVote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"foundation\",\"type\":\"address\"}],\"name\":\"transferToFoundation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Councils\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"GetNomineeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"nominee\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GetVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"introduction\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Nominees\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NotCouncils\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FoundationBin is the compiled bytecode used for deploying new contracts.
var FoundationBin = "0x608060405260008060146101000a81548160ff021916908315150217905550604051620024683803806200246883398181016040526200004391908101906200057e565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555089600490805190602001906200009b929190620002c4565b5088600181905550876002819055508660059080519060200190620000c2929190620002c4565b5085600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600681905550836007819055508260088190555060008090505b8251811015620002b357620001336200034b565b60405180606001604052808584815181106200014b57fe5b6020026020010151815260200160008152602001848152509050600981908060018154018082558091505090600182039060005260206000209060030201600090919290919091506000820151816000019080519060200190620001b19291906200036c565b50602082015181600101556040820151816002019080519060200190620001da9291906200036c565b50505050600980549050600a858481518110620001f357fe5b60200260200101516040516200020a91906200079c565b9081526020016040518091039020819055507fc8bb688d7129db8989edffeb4cd66831c19134ee4436dea9821ce51275586f638483815181106200024a57fe5b6020026020010151846001600a8887815181106200026457fe5b60200260200101516040516200027b91906200079c565b908152602001604051809103902054036040516200029c93929190620007b5565b60405180910390a15080806001019150506200011f565b50505050505050505050506200099c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200030757805160ff191683800117855562000338565b8280016001018555821562000338579182015b82811115620003375782518255916020019190600101906200031a565b5b509050620003479190620003f3565b5090565b60405180606001604052806060815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003af57805160ff1916838001178555620003e0565b82800160010185558215620003e0579182015b82811115620003df578251825591602001919060010190620003c2565b5b509050620003ef9190620003f3565b5090565b6200041891905b8082111562000414576000816000905550600101620003fa565b5090565b90565b6000815190506200042c8162000968565b92915050565b600082601f8301126200044457600080fd5b81516200045b62000455826200082e565b62000800565b9150818183526020840193506020810190508360005b83811015620004a557815186016200048a8882620004af565b84526020840193506020830192505060018101905062000471565b5050505092915050565b600082601f830112620004c157600080fd5b8151620004d8620004d28262000857565b62000800565b91508082526020830160208301858383011115620004f557600080fd5b6200050283828462000921565b50505092915050565b600082601f8301126200051d57600080fd5b8151620005346200052e8262000884565b62000800565b915080825260208301602083018583830111156200055157600080fd5b6200055e83828462000921565b50505092915050565b600081519050620005788162000982565b92915050565b6000806000806000806000806000806101408b8d0312156200059f57600080fd5b60008b015167ffffffffffffffff811115620005ba57600080fd5b620005c88d828e016200050b565b9a50506020620005db8d828e0162000567565b9950506040620005ee8d828e0162000567565b98505060608b015167ffffffffffffffff8111156200060c57600080fd5b6200061a8d828e016200050b565b97505060806200062d8d828e016200041b565b96505060a0620006408d828e0162000567565b95505060c0620006538d828e0162000567565b94505060e0620006668d828e0162000567565b9350506101008b015167ffffffffffffffff8111156200068557600080fd5b620006938d828e0162000432565b9250506101208b015167ffffffffffffffff811115620006b257600080fd5b620006c08d828e016200050b565b9150509295989b9194979a5092959850565b6000620006df82620008bc565b620006eb8185620008c7565b9350620006fd81856020860162000921565b620007088162000957565b840191505092915050565b60006200072082620008b1565b6200072c8185620008c7565b93506200073e81856020860162000921565b620007498162000957565b840191505092915050565b60006200076182620008b1565b6200076d8185620008d8565b93506200077f81856020860162000921565b80840191505092915050565b620007968162000917565b82525050565b6000620007aa828462000754565b915081905092915050565b60006060820190508181036000830152620007d1818662000713565b90508181036020830152620007e78185620006d2565b9050620007f860408301846200078b565b949350505050565b6000604051905081810181811067ffffffffffffffff821117156200082457600080fd5b8060405250919050565b600067ffffffffffffffff8211156200084657600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200086f57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff8211156200089c57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000620008f082620008f7565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b838110156200094157808201518184015260208101905062000924565b8381111562000951576000848401525b50505050565b6000601f19601f8301169050919050565b6200097381620008e3565b81146200097f57600080fd5b50565b6200098d8162000917565b81146200099957600080fd5b50565b611abc80620009ac6000396000f3fe6080604052600436106100dd5760003560e01c80636dc78b741161007f578063bef4876b11610059578063bef4876b1461028f578063d34dd1f0146102ba578063e15511c1146102e3578063fc85ab6414610320576100dd565b80636dc78b74146101fe578063829ece58146102275780638da5cb5b14610264576100dd565b80632aebcbb6116100bb5780632aebcbb61461018957806333c2d628146101a05780633ccfd60b146101bc57806356b6693f146101d3576100dd565b806306fdde03146100e25780631017cb571461010d578063237f02c01461014c575b600080fd5b3480156100ee57600080fd5b506100f761035f565b6040516101049190611766565b60405180910390f35b34801561011957600080fd5b50610134600480360361012f919081019061138d565b6103fd565b60405161014393929190611888565b60405180910390f35b34801561015857600080fd5b50610173600480360361016e9190810190611422565b61063f565b60405161018091906116c9565b60405180910390f35b34801561019557600080fd5b5061019e61067b565b005b6101ba60048036036101b59190810190611422565b610728565b005b3480156101c857600080fd5b506101d16107a6565b005b3480156101df57600080fd5b506101e861087f565b6040516101f59190611766565b60405180910390f35b34801561020a57600080fd5b5061022560048036036102209190810190611328565b61091d565b005b34801561023357600080fd5b5061024e60048036036102499190810190611422565b6109f7565b60405161025b91906116c9565b60405180910390f35b34801561027057600080fd5b50610279610a33565b60405161028691906116c9565b60405180910390f35b34801561029b57600080fd5b506102a4610a58565b6040516102b1919061174b565b60405180910390f35b3480156102c657600080fd5b506102e160048036036102dc9190810190611351565b610a6b565b005b3480156102ef57600080fd5b5061030a600480360361030591908101906113ce565b610b85565b604051610317919061186d565b60405180910390f35b34801561032c57600080fd5b5061034760048036036103429190810190611422565b610beb565b60405161035693929190611788565b60405180910390f35b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b505050505081565b60006060600080600a8560405161041491906116b2565b90815260200160405180910390205411610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a906117ed565b60405180910390fd5b61046b6111cf565b60096001600a8760405161047f91906116b2565b908152602001604051809103902054038154811061049957fe5b9060005260206000209060030201604051806060016040529081600082018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561054b5780601f106105205761010080835404028352916020019161054b565b820191906000526020600020905b81548152906001019060200180831161052e57829003601f168201915b5050505050815260200160018201548152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105f75780601f106105cc576101008083540402835291602001916105f7565b820191906000526020600020905b8154815290600101906020018083116105da57829003601f168201915b5050505050815250509050806020015181604001516001600a8860405161061e91906116b2565b90815260200160405180910390205403819150935093509350509193909250565b600c818154811061064c57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461070b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107029061184d565b60405180910390fd5b6001600060146101000a81548160ff021916908315150217905550565b60001515600060149054906101000a900460ff1615151461074857600080fd5b600154421015801561075b575060025442105b61079a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610791906117cd565b60405180910390fd5b6107a381610d52565b50565b60011515600060149054906101000a900460ff161515146107c657600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461081f57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f1935050505015801561087c573d6000803e3d6000fd5b50565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109155780601f106108ea57610100808354040283529160200191610915565b820191906000526020600020905b8154815290600101906020018083116108f857829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461097657600080fd5b60011515600060149054906101000a900460ff1615151461099657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f193505050501580156109f3573d6000803e3d6000fd5b5050565b600d8181548110610a0457fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060149054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac457600080fd5b60011515600060149054906101000a900460ff16151514610ae457600080fd5b6000809050610afe6006548361116790919063ffffffff16565b90508273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610b46573d6000803e3d6000fd5b507f2ba5dd9653174138784e5d96e6414a4792d581bf8b5622f9a78f2e76d7ee5d1a8382604051610b78929190611722565b60405180910390a1505050565b6000600b83604051610b9791906116b2565b908152602001604051809103902060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60098181548110610bf857fe5b9060005260206000209060030201600091509050806000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ca45780601f10610c7957610100808354040283529160200191610ca4565b820191906000526020600020905b815481529060010190602001808311610c8757829003601f168201915b505050505090806001015490806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d485780601f10610d1d57610100808354040283529160200191610d48565b820191906000526020600020905b815481529060010190602001808311610d2b57829003601f168201915b5050505050905083565b610d5a6111cf565b60098281548110610d6757fe5b9060005260206000209060030201604051806060016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e195780601f10610dee57610100808354040283529160200191610e19565b820191906000526020600020905b815481529060010190602001808311610dfc57829003601f168201915b5050505050815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ec55780601f10610e9a57610100808354040283529160200191610ec5565b820191906000526020600020905b815481529060010190602001808311610ea857829003601f168201915b50505050508152505090506000610ee76006543461119690919063ffffffff16565b90506007548110158015610f5e5750600b8260000151604051610f0a919061169b565b908152602001604051809103902060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600854038111155b610f9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f949061180d565b60405180910390fd5b60085481600b8460000151604051610fb5919061169b565b908152602001604051809103902060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054011115611043576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103a9061182d565b60405180910390fd5b80826020018181510191508181525050816009848154811061106157fe5b9060005260206000209060030201600082015181600001908051906020019061108b9291906111f0565b506020820151816001015560408201518160020190805190602001906110b29291906111f0565b5090505080600b83600001516040516110cb919061169b565b908152602001604051809103902060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055507f7e30f4b6a1054cfb4ad4be8640616baab2d7c975f2c968279e706c461882e27a3383600001518360405161115a939291906116e4565b60405180910390a1505050565b6000808284029050600084148061118657508284828161118357fe5b04145b61118c57fe5b8091505092915050565b60008082116111a157fe5b60008284816111ac57fe5b0490508284816111b857fe5b068184020184146111c557fe5b8091505092915050565b60405180606001604052806060815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061123157805160ff191683800117855561125f565b8280016001018555821561125f579182015b8281111561125e578251825591602001919060010190611243565b5b50905061126c9190611270565b5090565b61129291905b8082111561128e576000816000905550600101611276565b5090565b90565b6000813590506112a481611a34565b92915050565b6000813590506112b981611a4b565b92915050565b600082601f8301126112d057600080fd5b81356112e36112de826118f3565b6118c6565b915080825260208301602083018583830111156112ff57600080fd5b61130a8382846119e1565b50505092915050565b60008135905061132281611a62565b92915050565b60006020828403121561133a57600080fd5b6000611348848285016112aa565b91505092915050565b6000806040838503121561136457600080fd5b6000611372858286016112aa565b925050602061138385828601611313565b9150509250929050565b60006020828403121561139f57600080fd5b600082013567ffffffffffffffff8111156113b957600080fd5b6113c5848285016112bf565b91505092915050565b600080604083850312156113e157600080fd5b600083013567ffffffffffffffff8111156113fb57600080fd5b611407858286016112bf565b925050602061141885828601611295565b9150509250929050565b60006020828403121561143457600080fd5b600061144284828501611313565b91505092915050565b611454816119ab565b82525050565b61146381611951565b82525050565b61147281611975565b82525050565b60006114838261192a565b61148d8185611935565b935061149d8185602086016119f0565b6114a681611a23565b840191505092915050565b60006114bc8261192a565b6114c68185611946565b93506114d68185602086016119f0565b80840191505092915050565b60006114ed8261191f565b6114f78185611935565b93506115078185602086016119f0565b61151081611a23565b840191505092915050565b60006115268261191f565b6115308185611946565b93506115408185602086016119f0565b80840191505092915050565b6000611559600b83611935565b91507f6f7574206f662074696d650000000000000000000000000000000000000000006000830152602082019050919050565b6000611599600883611935565b91507f6e6f74206e616d650000000000000000000000000000000000000000000000006000830152602082019050919050565b60006115d9600c83611935565b91507f6f7574206f662072616e676500000000000000000000000000000000000000006000830152602082019050919050565b6000611619601083611935565b91507f6d6f7265207468656e206d6178506172000000000000000000000000000000006000830152602082019050919050565b6000611659600e83611935565b91507f746f2061646472657373206572720000000000000000000000000000000000006000830152602082019050919050565b611695816119a1565b82525050565b60006116a7828461151b565b915081905092915050565b60006116be82846114b1565b915081905092915050565b60006020820190506116de600083018461145a565b92915050565b60006060820190506116f9600083018661144b565b818103602083015261170b81856114e2565b905061171a604083018461168c565b949350505050565b6000604082019050611737600083018561144b565b611744602083018461168c565b9392505050565b60006020820190506117606000830184611469565b92915050565b6000602082019050818103600083015261178081846114e2565b905092915050565b600060608201905081810360008301526117a281866114e2565b90506117b1602083018561168c565b81810360408301526117c381846114e2565b9050949350505050565b600060208201905081810360008301526117e68161154c565b9050919050565b600060208201905081810360008301526118068161158c565b9050919050565b60006020820190508181036000830152611826816115cc565b9050919050565b600060208201905081810360008301526118468161160c565b9050919050565b600060208201905081810360008301526118668161164c565b9050919050565b6000602082019050611882600083018461168c565b92915050565b600060608201905061189d600083018661168c565b81810360208301526118af8185611478565b90506118be604083018461168c565b949350505050565b6000604051905081810181811067ffffffffffffffff821117156118e957600080fd5b8060405250919050565b600067ffffffffffffffff82111561190a57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061195c82611981565b9050919050565b600061196e82611981565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006119b6826119bd565b9050919050565b60006119c8826119cf565b9050919050565b60006119da82611981565b9050919050565b82818337600083830152505050565b60005b83811015611a0e5780820151818401526020810190506119f3565b83811115611a1d576000848401525b50505050565b6000601f19601f8301169050919050565b611a3d81611951565b8114611a4857600080fd5b50565b611a5481611963565b8114611a5f57600080fd5b50565b611a6b816119a1565b8114611a7657600080fd5b5056fea365627a7a7231582069036d4bb50648a59b969a51f89d146cee559fd3a0eb2599f255cd521d774ba76c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployFoundation deploys a new Ethereum contract, binding an instance of Foundation to it.
func DeployFoundation(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _begin *big.Int, _end *big.Int, _info string, _to common.Address, _fare *big.Int, _minPar *big.Int, _maxPar *big.Int, nominees []string, source string) (common.Address, *types.Transaction, *Foundation, error) {
	parsed, err := abi.JSON(strings.NewReader(FoundationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FoundationBin), backend, _name, _begin, _end, _info, _to, _fare, _minPar, _maxPar, nominees, source)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Foundation{FoundationCaller: FoundationCaller{contract: contract}, FoundationTransactor: FoundationTransactor{contract: contract}, FoundationFilterer: FoundationFilterer{contract: contract}}, nil
}

// Foundation is an auto generated Go binding around an Ethereum contract.
type Foundation struct {
	FoundationCaller     // Read-only binding to the contract
	FoundationTransactor // Write-only binding to the contract
	FoundationFilterer   // Log filterer for contract events
}

// FoundationCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoundationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoundationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoundationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoundationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoundationSession struct {
	Contract     *Foundation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoundationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoundationCallerSession struct {
	Contract *FoundationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FoundationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoundationTransactorSession struct {
	Contract     *FoundationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FoundationRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoundationRaw struct {
	Contract *Foundation // Generic contract binding to access the raw methods on
}

// FoundationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoundationCallerRaw struct {
	Contract *FoundationCaller // Generic read-only contract binding to access the raw methods on
}

// FoundationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoundationTransactorRaw struct {
	Contract *FoundationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoundation creates a new instance of Foundation, bound to a specific deployed contract.
func NewFoundation(address common.Address, backend bind.ContractBackend) (*Foundation, error) {
	contract, err := bindFoundation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foundation{FoundationCaller: FoundationCaller{contract: contract}, FoundationTransactor: FoundationTransactor{contract: contract}, FoundationFilterer: FoundationFilterer{contract: contract}}, nil
}

// NewFoundationCaller creates a new read-only instance of Foundation, bound to a specific deployed contract.
func NewFoundationCaller(address common.Address, caller bind.ContractCaller) (*FoundationCaller, error) {
	contract, err := bindFoundation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoundationCaller{contract: contract}, nil
}

// NewFoundationTransactor creates a new write-only instance of Foundation, bound to a specific deployed contract.
func NewFoundationTransactor(address common.Address, transactor bind.ContractTransactor) (*FoundationTransactor, error) {
	contract, err := bindFoundation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoundationTransactor{contract: contract}, nil
}

// NewFoundationFilterer creates a new log filterer instance of Foundation, bound to a specific deployed contract.
func NewFoundationFilterer(address common.Address, filterer bind.ContractFilterer) (*FoundationFilterer, error) {
	contract, err := bindFoundation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoundationFilterer{contract: contract}, nil
}

// bindFoundation binds a generic wrapper to an already deployed contract.
func bindFoundation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FoundationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foundation *FoundationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Foundation.Contract.FoundationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foundation *FoundationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.Contract.FoundationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foundation *FoundationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foundation.Contract.FoundationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foundation *FoundationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Foundation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foundation *FoundationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foundation *FoundationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foundation.Contract.contract.Transact(opts, method, params...)
}

// Councils is a free data retrieval call binding the contract method 0x237f02c0.
//
// Solidity: function Councils(uint256 ) constant returns(address)
func (_Foundation *FoundationCaller) Councils(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Foundation.contract.Call(opts, out, "Councils", arg0)
	return *ret0, err
}

// Councils is a free data retrieval call binding the contract method 0x237f02c0.
//
// Solidity: function Councils(uint256 ) constant returns(address)
func (_Foundation *FoundationSession) Councils(arg0 *big.Int) (common.Address, error) {
	return _Foundation.Contract.Councils(&_Foundation.CallOpts, arg0)
}

// Councils is a free data retrieval call binding the contract method 0x237f02c0.
//
// Solidity: function Councils(uint256 ) constant returns(address)
func (_Foundation *FoundationCallerSession) Councils(arg0 *big.Int) (common.Address, error) {
	return _Foundation.Contract.Councils(&_Foundation.CallOpts, arg0)
}

// GetNomineeInfo is a free data retrieval call binding the contract method 0x1017cb57.
//
// Solidity: function GetNomineeInfo(string name) constant returns(uint256 vote, string source, uint256 num)
func (_Foundation *FoundationCaller) GetNomineeInfo(opts *bind.CallOpts, name string) (struct {
	Vote   *big.Int
	Source string
	Num    *big.Int
}, error) {
	ret := new(struct {
		Vote   *big.Int
		Source string
		Num    *big.Int
	})
	out := ret
	err := _Foundation.contract.Call(opts, out, "GetNomineeInfo", name)
	return *ret, err
}

// GetNomineeInfo is a free data retrieval call binding the contract method 0x1017cb57.
//
// Solidity: function GetNomineeInfo(string name) constant returns(uint256 vote, string source, uint256 num)
func (_Foundation *FoundationSession) GetNomineeInfo(name string) (struct {
	Vote   *big.Int
	Source string
	Num    *big.Int
}, error) {
	return _Foundation.Contract.GetNomineeInfo(&_Foundation.CallOpts, name)
}

// GetNomineeInfo is a free data retrieval call binding the contract method 0x1017cb57.
//
// Solidity: function GetNomineeInfo(string name) constant returns(uint256 vote, string source, uint256 num)
func (_Foundation *FoundationCallerSession) GetNomineeInfo(name string) (struct {
	Vote   *big.Int
	Source string
	Num    *big.Int
}, error) {
	return _Foundation.Contract.GetNomineeInfo(&_Foundation.CallOpts, name)
}

// GetVote is a free data retrieval call binding the contract method 0xe15511c1.
//
// Solidity: function GetVote(string nominee, address voter) constant returns(uint256)
func (_Foundation *FoundationCaller) GetVote(opts *bind.CallOpts, nominee string, voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Foundation.contract.Call(opts, out, "GetVote", nominee, voter)
	return *ret0, err
}

// GetVote is a free data retrieval call binding the contract method 0xe15511c1.
//
// Solidity: function GetVote(string nominee, address voter) constant returns(uint256)
func (_Foundation *FoundationSession) GetVote(nominee string, voter common.Address) (*big.Int, error) {
	return _Foundation.Contract.GetVote(&_Foundation.CallOpts, nominee, voter)
}

// GetVote is a free data retrieval call binding the contract method 0xe15511c1.
//
// Solidity: function GetVote(string nominee, address voter) constant returns(uint256)
func (_Foundation *FoundationCallerSession) GetVote(nominee string, voter common.Address) (*big.Int, error) {
	return _Foundation.Contract.GetVote(&_Foundation.CallOpts, nominee, voter)
}

// Nominees is a free data retrieval call binding the contract method 0xfc85ab64.
//
// Solidity: function Nominees(uint256 ) constant returns(string name, uint256 vote, string source)
func (_Foundation *FoundationCaller) Nominees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name   string
	Vote   *big.Int
	Source string
}, error) {
	ret := new(struct {
		Name   string
		Vote   *big.Int
		Source string
	})
	out := ret
	err := _Foundation.contract.Call(opts, out, "Nominees", arg0)
	return *ret, err
}

// Nominees is a free data retrieval call binding the contract method 0xfc85ab64.
//
// Solidity: function Nominees(uint256 ) constant returns(string name, uint256 vote, string source)
func (_Foundation *FoundationSession) Nominees(arg0 *big.Int) (struct {
	Name   string
	Vote   *big.Int
	Source string
}, error) {
	return _Foundation.Contract.Nominees(&_Foundation.CallOpts, arg0)
}

// Nominees is a free data retrieval call binding the contract method 0xfc85ab64.
//
// Solidity: function Nominees(uint256 ) constant returns(string name, uint256 vote, string source)
func (_Foundation *FoundationCallerSession) Nominees(arg0 *big.Int) (struct {
	Name   string
	Vote   *big.Int
	Source string
}, error) {
	return _Foundation.Contract.Nominees(&_Foundation.CallOpts, arg0)
}

// NotCouncils is a free data retrieval call binding the contract method 0x829ece58.
//
// Solidity: function NotCouncils(uint256 ) constant returns(address)
func (_Foundation *FoundationCaller) NotCouncils(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Foundation.contract.Call(opts, out, "NotCouncils", arg0)
	return *ret0, err
}

// NotCouncils is a free data retrieval call binding the contract method 0x829ece58.
//
// Solidity: function NotCouncils(uint256 ) constant returns(address)
func (_Foundation *FoundationSession) NotCouncils(arg0 *big.Int) (common.Address, error) {
	return _Foundation.Contract.NotCouncils(&_Foundation.CallOpts, arg0)
}

// NotCouncils is a free data retrieval call binding the contract method 0x829ece58.
//
// Solidity: function NotCouncils(uint256 ) constant returns(address)
func (_Foundation *FoundationCallerSession) NotCouncils(arg0 *big.Int) (common.Address, error) {
	return _Foundation.Contract.NotCouncils(&_Foundation.CallOpts, arg0)
}

// Finished is a free data retrieval call binding the contract method 0xbef4876b.
//
// Solidity: function finished() constant returns(bool)
func (_Foundation *FoundationCaller) Finished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Foundation.contract.Call(opts, out, "finished")
	return *ret0, err
}

// Finished is a free data retrieval call binding the contract method 0xbef4876b.
//
// Solidity: function finished() constant returns(bool)
func (_Foundation *FoundationSession) Finished() (bool, error) {
	return _Foundation.Contract.Finished(&_Foundation.CallOpts)
}

// Finished is a free data retrieval call binding the contract method 0xbef4876b.
//
// Solidity: function finished() constant returns(bool)
func (_Foundation *FoundationCallerSession) Finished() (bool, error) {
	return _Foundation.Contract.Finished(&_Foundation.CallOpts)
}

// Introduction is a free data retrieval call binding the contract method 0x56b6693f.
//
// Solidity: function introduction() constant returns(string)
func (_Foundation *FoundationCaller) Introduction(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Foundation.contract.Call(opts, out, "introduction")
	return *ret0, err
}

// Introduction is a free data retrieval call binding the contract method 0x56b6693f.
//
// Solidity: function introduction() constant returns(string)
func (_Foundation *FoundationSession) Introduction() (string, error) {
	return _Foundation.Contract.Introduction(&_Foundation.CallOpts)
}

// Introduction is a free data retrieval call binding the contract method 0x56b6693f.
//
// Solidity: function introduction() constant returns(string)
func (_Foundation *FoundationCallerSession) Introduction() (string, error) {
	return _Foundation.Contract.Introduction(&_Foundation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Foundation *FoundationCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Foundation.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Foundation *FoundationSession) Name() (string, error) {
	return _Foundation.Contract.Name(&_Foundation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Foundation *FoundationCallerSession) Name() (string, error) {
	return _Foundation.Contract.Name(&_Foundation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Foundation *FoundationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Foundation.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Foundation *FoundationSession) Owner() (common.Address, error) {
	return _Foundation.Contract.Owner(&_Foundation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Foundation *FoundationCallerSession) Owner() (common.Address, error) {
	return _Foundation.Contract.Owner(&_Foundation.CallOpts)
}

// StartVote is a paid mutator transaction binding the contract method 0x33c2d628.
//
// Solidity: function StartVote(uint256 num) returns()
func (_Foundation *FoundationTransactor) StartVote(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "StartVote", num)
}

// StartVote is a paid mutator transaction binding the contract method 0x33c2d628.
//
// Solidity: function StartVote(uint256 num) returns()
func (_Foundation *FoundationSession) StartVote(num *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.StartVote(&_Foundation.TransactOpts, num)
}

// StartVote is a paid mutator transaction binding the contract method 0x33c2d628.
//
// Solidity: function StartVote(uint256 num) returns()
func (_Foundation *FoundationTransactorSession) StartVote(num *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.StartVote(&_Foundation.TransactOpts, num)
}

// FinishVote is a paid mutator transaction binding the contract method 0x2aebcbb6.
//
// Solidity: function finishVote() returns()
func (_Foundation *FoundationTransactor) FinishVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "finishVote")
}

// FinishVote is a paid mutator transaction binding the contract method 0x2aebcbb6.
//
// Solidity: function finishVote() returns()
func (_Foundation *FoundationSession) FinishVote() (*types.Transaction, error) {
	return _Foundation.Contract.FinishVote(&_Foundation.TransactOpts)
}

// FinishVote is a paid mutator transaction binding the contract method 0x2aebcbb6.
//
// Solidity: function finishVote() returns()
func (_Foundation *FoundationTransactorSession) FinishVote() (*types.Transaction, error) {
	return _Foundation.Contract.FinishVote(&_Foundation.TransactOpts)
}

// SafeWithdrawal is a paid mutator transaction binding the contract method 0xd34dd1f0.
//
// Solidity: function safeWithdrawal(address voter, uint256 num) returns()
func (_Foundation *FoundationTransactor) SafeWithdrawal(opts *bind.TransactOpts, voter common.Address, num *big.Int) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "safeWithdrawal", voter, num)
}

// SafeWithdrawal is a paid mutator transaction binding the contract method 0xd34dd1f0.
//
// Solidity: function safeWithdrawal(address voter, uint256 num) returns()
func (_Foundation *FoundationSession) SafeWithdrawal(voter common.Address, num *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.SafeWithdrawal(&_Foundation.TransactOpts, voter, num)
}

// SafeWithdrawal is a paid mutator transaction binding the contract method 0xd34dd1f0.
//
// Solidity: function safeWithdrawal(address voter, uint256 num) returns()
func (_Foundation *FoundationTransactorSession) SafeWithdrawal(voter common.Address, num *big.Int) (*types.Transaction, error) {
	return _Foundation.Contract.SafeWithdrawal(&_Foundation.TransactOpts, voter, num)
}

// TransferToFoundation is a paid mutator transaction binding the contract method 0x6dc78b74.
//
// Solidity: function transferToFoundation(address foundation) returns()
func (_Foundation *FoundationTransactor) TransferToFoundation(opts *bind.TransactOpts, foundation common.Address) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "transferToFoundation", foundation)
}

// TransferToFoundation is a paid mutator transaction binding the contract method 0x6dc78b74.
//
// Solidity: function transferToFoundation(address foundation) returns()
func (_Foundation *FoundationSession) TransferToFoundation(foundation common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.TransferToFoundation(&_Foundation.TransactOpts, foundation)
}

// TransferToFoundation is a paid mutator transaction binding the contract method 0x6dc78b74.
//
// Solidity: function transferToFoundation(address foundation) returns()
func (_Foundation *FoundationTransactorSession) TransferToFoundation(foundation common.Address) (*types.Transaction, error) {
	return _Foundation.Contract.TransferToFoundation(&_Foundation.TransactOpts, foundation)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Foundation *FoundationTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foundation.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Foundation *FoundationSession) Withdraw() (*types.Transaction, error) {
	return _Foundation.Contract.Withdraw(&_Foundation.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Foundation *FoundationTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Foundation.Contract.Withdraw(&_Foundation.TransactOpts)
}

// FoundationLogActivateNomineeIterator is returned from FilterLogActivateNominee and is used to iterate over the raw logs and unpacked data for LogActivateNominee events raised by the Foundation contract.
type FoundationLogActivateNomineeIterator struct {
	Event *FoundationLogActivateNominee // Event containing the contract specifics and raw log

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
func (it *FoundationLogActivateNomineeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationLogActivateNominee)
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
		it.Event = new(FoundationLogActivateNominee)
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
func (it *FoundationLogActivateNomineeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationLogActivateNomineeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationLogActivateNominee represents a LogActivateNominee event raised by the Foundation contract.
type FoundationLogActivateNominee struct {
	Nominee common.Address
	Num     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogActivateNominee is a free log retrieval operation binding the contract event 0x8aa9fdf82d1dac84a65e21709e2337b363c8be7b601643d74de10715494072dc.
//
// Solidity: event LogActivateNominee(address nominee, uint256 num)
func (_Foundation *FoundationFilterer) FilterLogActivateNominee(opts *bind.FilterOpts) (*FoundationLogActivateNomineeIterator, error) {

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "LogActivateNominee")
	if err != nil {
		return nil, err
	}
	return &FoundationLogActivateNomineeIterator{contract: _Foundation.contract, event: "LogActivateNominee", logs: logs, sub: sub}, nil
}

// WatchLogActivateNominee is a free log subscription operation binding the contract event 0x8aa9fdf82d1dac84a65e21709e2337b363c8be7b601643d74de10715494072dc.
//
// Solidity: event LogActivateNominee(address nominee, uint256 num)
func (_Foundation *FoundationFilterer) WatchLogActivateNominee(opts *bind.WatchOpts, sink chan<- *FoundationLogActivateNominee) (event.Subscription, error) {

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "LogActivateNominee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationLogActivateNominee)
				if err := _Foundation.contract.UnpackLog(event, "LogActivateNominee", log); err != nil {
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

// ParseLogActivateNominee is a log parse operation binding the contract event 0x8aa9fdf82d1dac84a65e21709e2337b363c8be7b601643d74de10715494072dc.
//
// Solidity: event LogActivateNominee(address nominee, uint256 num)
func (_Foundation *FoundationFilterer) ParseLogActivateNominee(log types.Log) (*FoundationLogActivateNominee, error) {
	event := new(FoundationLogActivateNominee)
	if err := _Foundation.contract.UnpackLog(event, "LogActivateNominee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FoundationLogNewCouncilIterator is returned from FilterLogNewCouncil and is used to iterate over the raw logs and unpacked data for LogNewCouncil events raised by the Foundation contract.
type FoundationLogNewCouncilIterator struct {
	Event *FoundationLogNewCouncil // Event containing the contract specifics and raw log

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
func (it *FoundationLogNewCouncilIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationLogNewCouncil)
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
		it.Event = new(FoundationLogNewCouncil)
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
func (it *FoundationLogNewCouncilIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationLogNewCouncilIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationLogNewCouncil represents a LogNewCouncil event raised by the Foundation contract.
type FoundationLogNewCouncil struct {
	Council common.Address
	Source  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogNewCouncil is a free log retrieval operation binding the contract event 0x1454128cbc71b0d4692d0b6f26b553290f380a07c0930a2fcc64887113e418d8.
//
// Solidity: event LogNewCouncil(address council, string source)
func (_Foundation *FoundationFilterer) FilterLogNewCouncil(opts *bind.FilterOpts) (*FoundationLogNewCouncilIterator, error) {

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "LogNewCouncil")
	if err != nil {
		return nil, err
	}
	return &FoundationLogNewCouncilIterator{contract: _Foundation.contract, event: "LogNewCouncil", logs: logs, sub: sub}, nil
}

// WatchLogNewCouncil is a free log subscription operation binding the contract event 0x1454128cbc71b0d4692d0b6f26b553290f380a07c0930a2fcc64887113e418d8.
//
// Solidity: event LogNewCouncil(address council, string source)
func (_Foundation *FoundationFilterer) WatchLogNewCouncil(opts *bind.WatchOpts, sink chan<- *FoundationLogNewCouncil) (event.Subscription, error) {

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "LogNewCouncil")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationLogNewCouncil)
				if err := _Foundation.contract.UnpackLog(event, "LogNewCouncil", log); err != nil {
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

// ParseLogNewCouncil is a log parse operation binding the contract event 0x1454128cbc71b0d4692d0b6f26b553290f380a07c0930a2fcc64887113e418d8.
//
// Solidity: event LogNewCouncil(address council, string source)
func (_Foundation *FoundationFilterer) ParseLogNewCouncil(log types.Log) (*FoundationLogNewCouncil, error) {
	event := new(FoundationLogNewCouncil)
	if err := _Foundation.contract.UnpackLog(event, "LogNewCouncil", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FoundationLogNewNomineesIterator is returned from FilterLogNewNominees and is used to iterate over the raw logs and unpacked data for LogNewNominees events raised by the Foundation contract.
type FoundationLogNewNomineesIterator struct {
	Event *FoundationLogNewNominees // Event containing the contract specifics and raw log

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
func (it *FoundationLogNewNomineesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationLogNewNominees)
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
		it.Event = new(FoundationLogNewNominees)
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
func (it *FoundationLogNewNomineesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationLogNewNomineesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationLogNewNominees represents a LogNewNominees event raised by the Foundation contract.
type FoundationLogNewNominees struct {
	Nominee string
	From    string
	Num     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogNewNominees is a free log retrieval operation binding the contract event 0xc8bb688d7129db8989edffeb4cd66831c19134ee4436dea9821ce51275586f63.
//
// Solidity: event LogNewNominees(string nominee, string from, uint256 num)
func (_Foundation *FoundationFilterer) FilterLogNewNominees(opts *bind.FilterOpts) (*FoundationLogNewNomineesIterator, error) {

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "LogNewNominees")
	if err != nil {
		return nil, err
	}
	return &FoundationLogNewNomineesIterator{contract: _Foundation.contract, event: "LogNewNominees", logs: logs, sub: sub}, nil
}

// WatchLogNewNominees is a free log subscription operation binding the contract event 0xc8bb688d7129db8989edffeb4cd66831c19134ee4436dea9821ce51275586f63.
//
// Solidity: event LogNewNominees(string nominee, string from, uint256 num)
func (_Foundation *FoundationFilterer) WatchLogNewNominees(opts *bind.WatchOpts, sink chan<- *FoundationLogNewNominees) (event.Subscription, error) {

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "LogNewNominees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationLogNewNominees)
				if err := _Foundation.contract.UnpackLog(event, "LogNewNominees", log); err != nil {
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

// ParseLogNewNominees is a log parse operation binding the contract event 0xc8bb688d7129db8989edffeb4cd66831c19134ee4436dea9821ce51275586f63.
//
// Solidity: event LogNewNominees(string nominee, string from, uint256 num)
func (_Foundation *FoundationFilterer) ParseLogNewNominees(log types.Log) (*FoundationLogNewNominees, error) {
	event := new(FoundationLogNewNominees)
	if err := _Foundation.contract.UnpackLog(event, "LogNewNominees", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FoundationLogTransferIterator is returned from FilterLogTransfer and is used to iterate over the raw logs and unpacked data for LogTransfer events raised by the Foundation contract.
type FoundationLogTransferIterator struct {
	Event *FoundationLogTransfer // Event containing the contract specifics and raw log

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
func (it *FoundationLogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationLogTransfer)
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
		it.Event = new(FoundationLogTransfer)
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
func (it *FoundationLogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationLogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationLogTransfer represents a LogTransfer event raised by the Foundation contract.
type FoundationLogTransfer struct {
	Voter common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogTransfer is a free log retrieval operation binding the contract event 0x2ba5dd9653174138784e5d96e6414a4792d581bf8b5622f9a78f2e76d7ee5d1a.
//
// Solidity: event LogTransfer(address voter, uint256 value)
func (_Foundation *FoundationFilterer) FilterLogTransfer(opts *bind.FilterOpts) (*FoundationLogTransferIterator, error) {

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "LogTransfer")
	if err != nil {
		return nil, err
	}
	return &FoundationLogTransferIterator{contract: _Foundation.contract, event: "LogTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTransfer is a free log subscription operation binding the contract event 0x2ba5dd9653174138784e5d96e6414a4792d581bf8b5622f9a78f2e76d7ee5d1a.
//
// Solidity: event LogTransfer(address voter, uint256 value)
func (_Foundation *FoundationFilterer) WatchLogTransfer(opts *bind.WatchOpts, sink chan<- *FoundationLogTransfer) (event.Subscription, error) {

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "LogTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationLogTransfer)
				if err := _Foundation.contract.UnpackLog(event, "LogTransfer", log); err != nil {
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

// ParseLogTransfer is a log parse operation binding the contract event 0x2ba5dd9653174138784e5d96e6414a4792d581bf8b5622f9a78f2e76d7ee5d1a.
//
// Solidity: event LogTransfer(address voter, uint256 value)
func (_Foundation *FoundationFilterer) ParseLogTransfer(log types.Log) (*FoundationLogTransfer, error) {
	event := new(FoundationLogTransfer)
	if err := _Foundation.contract.UnpackLog(event, "LogTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FoundationLogVoteNomineeIterator is returned from FilterLogVoteNominee and is used to iterate over the raw logs and unpacked data for LogVoteNominee events raised by the Foundation contract.
type FoundationLogVoteNomineeIterator struct {
	Event *FoundationLogVoteNominee // Event containing the contract specifics and raw log

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
func (it *FoundationLogVoteNomineeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoundationLogVoteNominee)
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
		it.Event = new(FoundationLogVoteNominee)
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
func (it *FoundationLogVoteNomineeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoundationLogVoteNomineeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoundationLogVoteNominee represents a LogVoteNominee event raised by the Foundation contract.
type FoundationLogVoteNominee struct {
	Voter   common.Address
	Council string
	Votes   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogVoteNominee is a free log retrieval operation binding the contract event 0x7e30f4b6a1054cfb4ad4be8640616baab2d7c975f2c968279e706c461882e27a.
//
// Solidity: event LogVoteNominee(address voter, string council, uint256 votes)
func (_Foundation *FoundationFilterer) FilterLogVoteNominee(opts *bind.FilterOpts) (*FoundationLogVoteNomineeIterator, error) {

	logs, sub, err := _Foundation.contract.FilterLogs(opts, "LogVoteNominee")
	if err != nil {
		return nil, err
	}
	return &FoundationLogVoteNomineeIterator{contract: _Foundation.contract, event: "LogVoteNominee", logs: logs, sub: sub}, nil
}

// WatchLogVoteNominee is a free log subscription operation binding the contract event 0x7e30f4b6a1054cfb4ad4be8640616baab2d7c975f2c968279e706c461882e27a.
//
// Solidity: event LogVoteNominee(address voter, string council, uint256 votes)
func (_Foundation *FoundationFilterer) WatchLogVoteNominee(opts *bind.WatchOpts, sink chan<- *FoundationLogVoteNominee) (event.Subscription, error) {

	logs, sub, err := _Foundation.contract.WatchLogs(opts, "LogVoteNominee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoundationLogVoteNominee)
				if err := _Foundation.contract.UnpackLog(event, "LogVoteNominee", log); err != nil {
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

// ParseLogVoteNominee is a log parse operation binding the contract event 0x7e30f4b6a1054cfb4ad4be8640616baab2d7c975f2c968279e706c461882e27a.
//
// Solidity: event LogVoteNominee(address voter, string council, uint256 votes)
func (_Foundation *FoundationFilterer) ParseLogVoteNominee(log types.Log) (*FoundationLogVoteNominee, error) {
	event := new(FoundationLogVoteNominee)
	if err := _Foundation.contract.UnpackLog(event, "LogVoteNominee", log); err != nil {
		return nil, err
	}
	return event, nil
}
