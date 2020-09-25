package ethereum

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/sirupsen/logrus"
	"gitlab.dataqin.com/sipc/vote_app/ethereum/foundation"
	"math/big"
	"strings"
	"testing"
)

/**
 * @Classname foundation
 * @Author Johnathan
 * @Date 2020/8/19 14:40
 * @Created by Goalnd 2020
 */

type VoteInfo struct {
	Name  string   `json:"name"`
	Info  string   `json:"info"`
	Begin *big.Int `json:"begin"`
	End   *big.Int `json:"end"`
}

const privateKey = "1d279d0f8fd2dd991b2ab74708d139a1fcfccf712b5b43af44c3d2eef0e37179"
const address = "0xe8c8a5373d3541c27f3a369c5e860d18f5b256c6"

//const privateKey = "f4fb142a51dd7edda51003b6b5e1943bcc9b3641c5610001bb75fb0add558469"
//const address = "0xb9df4b612ccf188907d11b7fb7d24eb1f0bc14a8"
const SubprivateKey = "ad9ea04f3d01c0758240f35280b35c84886adca94435cd385d179047ef11784e"
const SubAddress = "0xfc53af2db0712e02663ed7530de1f52e45c34df3"
const FoundationContract = "0x32bF6faB840AEBaBE6c3293694d9952c7abfa1Fa"
const Focus = "0xec44cfa70aedf56a9b65e4f7d092567671d3b618"

func TestEthRpcClient_DeployoundationContract(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	client := ethClient.GetClient()
	name := "选择一理事会？"
	info := `选择一个`
	begin := big.NewInt(1598248862)
	end := big.NewInt(1598716800)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := struct {
		Name  string `json:"name"`
		Begin int    `json:"begin"`
		End   int    `json:"end"`
		Info  string `json:"info"`
		To    string `json:"to"`
	}{
		name,
		int(begin.Int64()),
		int(end.Int64()),
		info,
		Focus,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	to := common.BytesToAddress([]byte(nil))
	msg := ethereum.CallMsg{GasPrice: gasPrice, Value: big.NewInt(0), Gas: 1000000, From: common.HexToAddress(address), To: &to, Data: bytes}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		t.Error(err.Error())
		return
	}
	gasLimit = 8000000
	auth, err := ethClient.TransactOpts(privateKey, client, gasLimit, big.NewInt(0))
	if err != nil {
		t.Error(err.Error())
		return
	}
	peoples := []string{"张学政", "韩梦", "曹可磊", "闫琳琳", "谢磊", "刘哲", "高旭", "朱卿赫", "向文祥"}
	addr, tx, _, err := foundation.DeployFoundation(auth, client, name, begin, end, info, common.HexToAddress(Focus), big.NewInt(1000000000000000000),
		big.NewInt(1), big.NewInt(10), peoples, "test")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(addr.Hex())
	t.Log(tx.Hash().Hex())
}

// 部署投票合约
func TestEthRpcClient_DeployStoreFoundationContract(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.3.148:9545")
	name := "选择一理事会？"
	info := `选择一个`
	begin := big.NewInt(1599805642)
	end := big.NewInt(1599905642)
	fare := big.NewInt(1e18)
	min := big.NewInt(1)
	max := big.NewInt(5)
	//peoples := []string{"张学政", "韩梦", "曹可磊", "闫琳琳", "谢磊", "刘哲", "高旭", "朱卿赫", "向文祥"}
	peoples := []string{"ff", "gg"}
	private, _ := crypto.HexToECDSA(privateKey)

	client := ethClient.GetClient()
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err.Error())
		return
	}
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	if err != nil {
		t.Error(err.Error())
		return
	}
	input, err := parsed.Pack("", name, begin, end, info, common.HexToAddress(Focus), fare, min, max, peoples, "理事会投票")
	if err != nil {
		t.Error(err.Error())
		return
	}
	input = append(common.FromHex(foundation.FoundationBin), input...)
	str := `608060405260008060146101000a81548160ff021916908315150217905550604051620024683803806200246883398181016040526200004391908101906200057e565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555089600490805190602001906200009b929190620002c4565b5088600181905550876002819055508660059080519060200190620000c2929190620002c4565b5085600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600681905550836007819055508260088190555060008090505b8251811015620002b357620001336200034b565b60405180606001604052808584815181106200014b57fe5b6020026020010151815260200160008152602001848152509050600981908060018154018082558091505090600182039060005260206000209060030201600090919290919091506000820151816000019080519060200190620001b19291906200036c565b50602082015181600101556040820151816002019080519060200190620001da9291906200036c565b50505050600980549050600a858481518110620001f357fe5b60200260200101516040516200020a91906200079c565b9081526020016040518091039020819055507fc8bb688d7129db8989edffeb4cd66831c19134ee4436dea9821ce51275586f638483815181106200024a57fe5b6020026020010151846001600a8887815181106200026457fe5b60200260200101516040516200027b91906200079c565b908152602001604051809103902054036040516200029c93929190620007b5565b60405180910390a15080806001019150506200011f565b50505050505050505050506200099c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200030757805160ff191683800117855562000338565b8280016001018555821562000338579182015b82811115620003375782518255916020019190600101906200031a565b5b509050620003479190620003f3565b5090565b60405180606001604052806060815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003af57805160ff1916838001178555620003e0565b82800160010185558215620003e0579182015b82811115620003df578251825591602001919060010190620003c2565b5b509050620003ef9190620003f3565b5090565b6200041891905b8082111562000414576000816000905550600101620003fa565b5090565b90565b6000815190506200042c8162000968565b92915050565b600082601f8301126200044457600080fd5b81516200045b62000455826200082e565b62000800565b9150818183526020840193506020810190508360005b83811015620004a557815186016200048a8882620004af565b84526020840193506020830192505060018101905062000471565b5050505092915050565b600082601f830112620004c157600080fd5b8151620004d8620004d28262000857565b62000800565b91508082526020830160208301858383011115620004f557600080fd5b6200050283828462000921565b50505092915050565b600082601f8301126200051d57600080fd5b8151620005346200052e8262000884565b62000800565b915080825260208301602083018583830111156200055157600080fd5b6200055e83828462000921565b50505092915050565b600081519050620005788162000982565b92915050565b6000806000806000806000806000806101408b8d0312156200059f57600080fd5b60008b015167ffffffffffffffff811115620005ba57600080fd5b620005c88d828e016200050b565b9a50506020620005db8d828e0162000567565b9950506040620005ee8d828e0162000567565b98505060608b015167ffffffffffffffff8111156200060c57600080fd5b6200061a8d828e016200050b565b97505060806200062d8d828e016200041b565b96505060a0620006408d828e0162000567565b95505060c0620006538d828e0162000567565b94505060e0620006668d828e0162000567565b9350506101008b015167ffffffffffffffff8111156200068557600080fd5b620006938d828e0162000432565b9250506101208b015167ffffffffffffffff811115620006b257600080fd5b620006c08d828e016200050b565b9150509295989b9194979a5092959850565b6000620006df82620008bc565b620006eb8185620008c7565b9350620006fd81856020860162000921565b620007088162000957565b840191505092915050565b60006200072082620008b1565b6200072c8185620008c7565b93506200073e81856020860162000921565b620007498162000957565b840191505092915050565b60006200076182620008b1565b6200076d8185620008d8565b93506200077f81856020860162000921565b80840191505092915050565b620007968162000917565b82525050565b6000620007aa828462000754565b915081905092915050565b60006060820190508181036000830152620007d1818662000713565b90508181036020830152620007e78185620006d2565b9050620007f860408301846200078b565b949350505050565b6000604051905081810181811067ffffffffffffffff821117156200082457600080fd5b8060405250919050565b600067ffffffffffffffff8211156200084657600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200086f57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff8211156200089c57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000620008f082620008f7565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b838110156200094157808201518184015260208101905062000924565b8381111562000951576000848401525b50505050565b6000601f19601f8301169050919050565b6200097381620008e3565b81146200097f57600080fd5b50565b6200098d8162000917565b81146200099957600080fd5b50565b611abc80620009ac6000396000f3fe6080604052600436106100dd5760003560e01c80636dc78b741161007f578063bef4876b11610059578063bef4876b1461028f578063d34dd1f0146102ba578063e15511c1146102e3578063fc85ab6414610320576100dd565b80636dc78b74146101fe578063829ece58146102275780638da5cb5b14610264576100dd565b80632aebcbb6116100bb5780632aebcbb61461018957806333c2d628146101a05780633ccfd60b146101bc57806356b6693f146101d3576100dd565b806306fdde03146100e25780631017cb571461010d578063237f02c01461014c575b600080fd5b3480156100ee57600080fd5b506100f761035f565b6040516101049190611766565b60405180910390f35b34801561011957600080fd5b50610134600480360361012f919081019061138d565b6103fd565b60405161014393929190611888565b60405180910390f35b34801561015857600080fd5b50610173600480360361016e9190810190611422565b61063f565b60405161018091906116c9565b60405180910390f35b34801561019557600080fd5b5061019e61067b565b005b6101ba60048036036101b59190810190611422565b610728565b005b3480156101c857600080fd5b506101d16107a6565b005b3480156101df57600080fd5b506101e861087f565b6040516101f59190611766565b60405180910390f35b34801561020a57600080fd5b5061022560048036036102209190810190611328565b61091d565b005b34801561023357600080fd5b5061024e60048036036102499190810190611422565b6109f7565b60405161025b91906116c9565b60405180910390f35b34801561027057600080fd5b50610279610a33565b60405161028691906116c9565b60405180910390f35b34801561029b57600080fd5b506102a4610a58565b6040516102b1919061174b565b60405180910390f35b3480156102c657600080fd5b506102e160048036036102dc9190810190611351565b610a6b565b005b3480156102ef57600080fd5b5061030a600480360361030591908101906113ce565b610b85565b604051610317919061186d565b60405180910390f35b34801561032c57600080fd5b5061034760048036036103429190810190611422565b610beb565b60405161035693929190611788565b60405180910390f35b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b505050505081565b60006060600080600a8560405161041491906116b2565b90815260200160405180910390205411610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045a906117ed565b60405180910390fd5b61046b6111cf565b60096001600a8760405161047f91906116b2565b908152602001604051809103902054038154811061049957fe5b9060005260206000209060030201604051806060016040529081600082018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561054b5780601f106105205761010080835404028352916020019161054b565b820191906000526020600020905b81548152906001019060200180831161052e57829003601f168201915b5050505050815260200160018201548152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105f75780601f106105cc576101008083540402835291602001916105f7565b820191906000526020600020905b8154815290600101906020018083116105da57829003601f168201915b5050505050815250509050806020015181604001516001600a8860405161061e91906116b2565b90815260200160405180910390205403819150935093509350509193909250565b600c818154811061064c57fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461070b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107029061184d565b60405180910390fd5b6001600060146101000a81548160ff021916908315150217905550565b60001515600060149054906101000a900460ff1615151461074857600080fd5b600154421015801561075b575060025442105b61079a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610791906117cd565b60405180910390fd5b6107a381610d52565b50565b60011515600060149054906101000a900460ff161515146107c657600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461081f57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f1935050505015801561087c573d6000803e3d6000fd5b50565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109155780601f106108ea57610100808354040283529160200191610915565b820191906000526020600020905b8154815290600101906020018083116108f857829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461097657600080fd5b60011515600060149054906101000a900460ff1615151461099657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f193505050501580156109f3573d6000803e3d6000fd5b5050565b600d8181548110610a0457fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060149054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac457600080fd5b60011515600060149054906101000a900460ff16151514610ae457600080fd5b6000809050610afe6006548361116790919063ffffffff16565b90508273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610b46573d6000803e3d6000fd5b507f2ba5dd9653174138784e5d96e6414a4792d581bf8b5622f9a78f2e76d7ee5d1a8382604051610b78929190611722565b60405180910390a1505050565b6000600b83604051610b9791906116b2565b908152602001604051809103902060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60098181548110610bf857fe5b9060005260206000209060030201600091509050806000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ca45780601f10610c7957610100808354040283529160200191610ca4565b820191906000526020600020905b815481529060010190602001808311610c8757829003601f168201915b505050505090806001015490806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d485780601f10610d1d57610100808354040283529160200191610d48565b820191906000526020600020905b815481529060010190602001808311610d2b57829003601f168201915b5050505050905083565b610d5a6111cf565b60098281548110610d6757fe5b9060005260206000209060030201604051806060016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e195780601f10610dee57610100808354040283529160200191610e19565b820191906000526020600020905b815481529060010190602001808311610dfc57829003601f168201915b5050505050815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ec55780601f10610e9a57610100808354040283529160200191610ec5565b820191906000526020600020905b815481529060010190602001808311610ea857829003601f168201915b50505050508152505090506000610ee76006543461119690919063ffffffff16565b90506007548110158015610f5e5750600b8260000151604051610f0a919061169b565b908152602001604051809103902060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600854038111155b610f9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f949061180d565b60405180910390fd5b60085481600b8460000151604051610fb5919061169b565b908152602001604051809103902060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054011115611043576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103a9061182d565b60405180910390fd5b80826020018181510191508181525050816009848154811061106157fe5b9060005260206000209060030201600082015181600001908051906020019061108b9291906111f0565b506020820151816001015560408201518160020190805190602001906110b29291906111f0565b5090505080600b83600001516040516110cb919061169b565b908152602001604051809103902060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055507f7e30f4b6a1054cfb4ad4be8640616baab2d7c975f2c968279e706c461882e27a3383600001518360405161115a939291906116e4565b60405180910390a1505050565b6000808284029050600084148061118657508284828161118357fe5b04145b61118c57fe5b8091505092915050565b60008082116111a157fe5b60008284816111ac57fe5b0490508284816111b857fe5b068184020184146111c557fe5b8091505092915050565b60405180606001604052806060815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061123157805160ff191683800117855561125f565b8280016001018555821561125f579182015b8281111561125e578251825591602001919060010190611243565b5b50905061126c9190611270565b5090565b61129291905b8082111561128e576000816000905550600101611276565b5090565b90565b6000813590506112a481611a34565b92915050565b6000813590506112b981611a4b565b92915050565b600082601f8301126112d057600080fd5b81356112e36112de826118f3565b6118c6565b915080825260208301602083018583830111156112ff57600080fd5b61130a8382846119e1565b50505092915050565b60008135905061132281611a62565b92915050565b60006020828403121561133a57600080fd5b6000611348848285016112aa565b91505092915050565b6000806040838503121561136457600080fd5b6000611372858286016112aa565b925050602061138385828601611313565b9150509250929050565b60006020828403121561139f57600080fd5b600082013567ffffffffffffffff8111156113b957600080fd5b6113c5848285016112bf565b91505092915050565b600080604083850312156113e157600080fd5b600083013567ffffffffffffffff8111156113fb57600080fd5b611407858286016112bf565b925050602061141885828601611295565b9150509250929050565b60006020828403121561143457600080fd5b600061144284828501611313565b91505092915050565b611454816119ab565b82525050565b61146381611951565b82525050565b61147281611975565b82525050565b60006114838261192a565b61148d8185611935565b935061149d8185602086016119f0565b6114a681611a23565b840191505092915050565b60006114bc8261192a565b6114c68185611946565b93506114d68185602086016119f0565b80840191505092915050565b60006114ed8261191f565b6114f78185611935565b93506115078185602086016119f0565b61151081611a23565b840191505092915050565b60006115268261191f565b6115308185611946565b93506115408185602086016119f0565b80840191505092915050565b6000611559600b83611935565b91507f6f7574206f662074696d650000000000000000000000000000000000000000006000830152602082019050919050565b6000611599600883611935565b91507f6e6f74206e616d650000000000000000000000000000000000000000000000006000830152602082019050919050565b60006115d9600c83611935565b91507f6f7574206f662072616e676500000000000000000000000000000000000000006000830152602082019050919050565b6000611619601083611935565b91507f6d6f7265207468656e206d6178506172000000000000000000000000000000006000830152602082019050919050565b6000611659600e83611935565b91507f746f2061646472657373206572720000000000000000000000000000000000006000830152602082019050919050565b611695816119a1565b82525050565b60006116a7828461151b565b915081905092915050565b60006116be82846114b1565b915081905092915050565b60006020820190506116de600083018461145a565b92915050565b60006060820190506116f9600083018661144b565b818103602083015261170b81856114e2565b905061171a604083018461168c565b949350505050565b6000604082019050611737600083018561144b565b611744602083018461168c565b9392505050565b60006020820190506117606000830184611469565b92915050565b6000602082019050818103600083015261178081846114e2565b905092915050565b600060608201905081810360008301526117a281866114e2565b90506117b1602083018561168c565b81810360408301526117c381846114e2565b9050949350505050565b600060208201905081810360008301526117e68161154c565b9050919050565b600060208201905081810360008301526118068161158c565b9050919050565b60006020820190508181036000830152611826816115cc565b9050919050565b600060208201905081810360008301526118468161160c565b9050919050565b600060208201905081810360008301526118668161164c565b9050919050565b6000602082019050611882600083018461168c565b92915050565b600060608201905061189d600083018661168c565b81810360208301526118af8185611478565b90506118be604083018461168c565b949350505050565b6000604051905081810181811067ffffffffffffffff821117156118e957600080fd5b8060405250919050565b600067ffffffffffffffff82111561190a57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061195c82611981565b9050919050565b600061196e82611981565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006119b6826119bd565b9050919050565b60006119c8826119cf565b9050919050565b60006119da82611981565b9050919050565b82818337600083830152505050565b60005b83811015611a0e5780820151818401526020810190506119f3565b83811115611a1d576000848401525b50505050565b6000601f19601f8301169050919050565b611a3d81611951565b8114611a4857600080fd5b50565b611a5481611963565b8114611a5f57600080fd5b50565b611a6b816119a1565b8114611a7657600080fd5b5056fea365627a7a7231582069036d4bb50648a59b969a51f89d146cee559fd3a0eb2599f255cd521d774ba76c6578706572696d656e74616cf564736f6c634300050d00400000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000005f559420000000000000000000000000000000000000000000000000000000005f6020200000000000000000000000000000000000000000000000000000000000000180000000000000000000000000ec44cfa70aedf56a9b65e4f7d092567671d3b6180000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000360000000000000000000000000000000000000000000000000000000000000000fe79086e4ba8be4bc9ae6b58be8af9500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006e6b58be8af9500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000009e5bca0e5ada6e694bf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006e99fa9e6a2a600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009e69bb9e58fafe7a38a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006e58898e593b20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000fe79086e4ba8be4bc9ae68a95e7a5a80000000000000000000000000000000000`
	input, _ = hex.DecodeString(str)
	//t.Log(hex.EncodeToString(input))
	toAddr := common.BytesToAddress([]byte(nil))
	value := big.NewInt(0)
	msg := ethereum.CallMsg{GasPrice: gasPrice, Value: value, Gas: gasPrice.Uint64(), From: common.HexToAddress(address), To: &toAddr, Data: input}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		t.Error(err.Error())
		return
	}

	res := hex.EncodeToString(input)
	gasLimit = 8000000
	t.Log(res)
	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(address), nil)
	t.Log(nonce)
	if err != nil {
		t.Fatal(err)
	}
	tx := types.NewContractCreation(
		nonce, value,
		gasLimit, gasPrice, input,
	)
	t.Log(tx.Hash().Hex())
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	if err != nil {
		t.Fatal(err)
	}

	bf, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		t.Fatal(err)
	}
	rawTx := hexutil.Encode(bf)
	t.Log(bf) // 签名结果

	// 跨链交易转发
	txHash, err := ethClient.SendRawTransaction(rawTx)
	//err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	contract := crypto.CreateAddress(common.HexToAddress(address), signedTx.Nonce())
	t.Log(txHash, "-----", contract.Hex())
	/*0x06233629550929c77726759c4fe077e48b208137fb01ce692ce68336a801aaad ----- 0x5CD70f90fB5DB2546CCd4DA5C2396c9460178e29*/
}

//添加候选人(没有投票限制使用AddNormalNominees)
func TestEthRpcClient_AddNormalNominees(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	private, _ := crypto.HexToECDSA(privateKey)
	client := ethClient.GetClient()
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	if err != nil {
		t.Error(err.Error())
		return
	}
	addrList := []common.Address{
		common.HexToAddress("0x95bebb32e53a954464dfd8fb45da3ed1ac645d08"),
		common.HexToAddress("0xb9df4b612ccf188907d11b7fb7d24eb1f0bc14a8"),
	}
	//focus := common.HexToAddress(Focus)
	input, err := parsed.Pack("AddNormalNominees", addrList, "理事会投票")
	if err != nil {
		t.Error(err.Error())
		return
	}
	res := hex.EncodeToString(input)
	t.Log(res)
	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(address), nil)
	t.Log(nonce)
	if err != nil {
		t.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err.Error())
		return
	}
	//合约地址(投票部署的合约地址)
	toAddr := common.HexToAddress(FoundationContract)

	value := big.NewInt(0)
	msg := ethereum.CallMsg{GasPrice: gasPrice, Value: value, From: common.HexToAddress(address), To: &toAddr, Data: input}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//gasLimit = 300000
	tx := types.NewTransaction(
		nonce, toAddr, value,
		gasLimit, gasPrice, input,
	)
	t.Log(tx.Hash().Hex())
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	if err != nil {
		t.Fatal(err)
	}

	bf, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		t.Fatal(err)
	}
	rawTx := hexutil.Encode(bf)
	t.Log(bf) // 签名结果

	// 交易转发
	txHash, err := ethClient.SendRawTransaction(rawTx)
	//err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(txHash)
}

/*func TestSetTicket(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	client := ethClient.GetClient()
	foundationInstance, err := foundation.NewFoundation(common.HexToAddress(FoundationContract), client)
	if err != nil {
		t.Error(err.Error())
		return
	}
	gasLimit := uint64(300000)
	auth, err := ethClient.TransactOpts(privateKey, client, gasLimit, big.NewInt(0))
	if err != nil {
		t.Error(err.Error())
		return
	}
	fare := big.NewInt(1)
	min := big.NewInt(1)
	max := big.NewInt(10)
	//focus := common.HexToAddress(Focus)
	tx, err := foundationInstance.SetTicket(auth, fare, min, max)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(tx.Hash().Hex())
}*/

// 设置规则
func TestEthRpcClient_SetTicket(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	private, _ := crypto.HexToECDSA(privateKey)
	client := ethClient.GetClient()
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	if err != nil {
		t.Error(err.Error())
		return
	}
	fare := big.NewInt(1e18)
	min := big.NewInt(1)
	max := big.NewInt(10)
	focus := common.HexToAddress(Focus)
	input, err := parsed.Pack("SetTicket", fare, min, max, focus)
	if err != nil {
		t.Error(err.Error())
		return
	}
	res := hex.EncodeToString(input)
	t.Log(res)
	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(address), nil)
	t.Log(nonce)
	if err != nil {
		t.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err.Error())
		return
	}
	//合约地址(投票部署的合约地址)
	toAddr := common.HexToAddress(FoundationContract)

	value := big.NewInt(0)
	msg := ethereum.CallMsg{GasPrice: gasPrice, Value: value, From: common.HexToAddress(address), To: &toAddr, Data: input}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//gasLimit = 300000
	tx := types.NewTransaction(
		nonce, toAddr, value,
		gasLimit, gasPrice, input,
	)
	t.Log(tx.Hash().Hex())
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	if err != nil {
		t.Fatal(err)
	}

	bf, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		t.Fatal(err)
	}
	rawTx := hexutil.Encode(bf)
	t.Log(bf) // 签名结果

	// 交易转发
	txHash, err := ethClient.SendRawTransaction(rawTx)
	//err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(txHash)
}

func TestStartVoteFoundation(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	client := ethClient.GetClient()
	foundationInstance, err := foundation.NewFoundation(common.HexToAddress(FoundationContract), client)
	if err != nil {
		t.Error(err.Error())
		return
	}
	gasLimit := uint64(300000)
	auth, err := ethClient.TransactOpts(SubprivateKey, client, gasLimit, big.NewInt(1e18))
	if err != nil {
		t.Error(err.Error())
		return
	}
	tx, err := foundationInstance.StartVote(auth, big.NewInt(0))
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(tx.Hash().Hex())
}

//理事会投票
func TestStartFoundationVote(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.3.148:9545")
	private, _ := crypto.HexToECDSA(SubprivateKey)
	client := ethClient.GetClient()
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	if err != nil {
		t.Error(err.Error())
		return
	}
	//合约地址(投票部署的合约地址)
	toAddr := common.HexToAddress(FoundationContract)
	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(SubAddress), nil)
	t.Log(nonce)
	if err != nil {
		t.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err.Error())
		return
	}
	num := big.NewInt(0) //候选人序号
	input, err := parsed.Pack("StartVote", num)
	if err != nil {
		t.Error(err.Error())
		return
	}
	res := hex.EncodeToString(input)
	t.Log(res)
	str := `33c2d6280000000000000000000000000000000000000000000000000000000000000000`
	input, _ = hex.DecodeString(str)
	value := big.NewInt(2e18)
	msg := ethereum.CallMsg{GasPrice: gasPrice, Value: value, From: common.HexToAddress(SubAddress), To: &toAddr, Data: input}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//gasLimit = uint64(300000)
	tx := types.NewTransaction(
		nonce, toAddr, value,
		gasLimit, gasPrice, input,
	)
	t.Log(tx.Hash().Hex())
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	if err != nil {
		t.Fatal(err)
	}

	bf, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		t.Fatal(err)
	}
	rawTx := hexutil.Encode(bf)
	t.Log(bf) // 签名结果

	// 交易转发
	txHash, err := ethClient.SendRawTransaction(rawTx)
	//err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(txHash)
}

// 获取投票信息
func TestGetNominessInfo(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.3.148:9545")
	client := ethClient.GetClient()
	foundationInstance, err := foundation.NewFoundation(common.HexToAddress(FoundationContract), client)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//opts := bind.CallOpts{Pending: true, From: common.HexToAddress(SubAddress), BlockNumber: nil, Context: context.Background()}
	res, err := foundationInstance.GetNomineeInfo(nil, "奥巴马")
	t.Log("Num", res.Num)
	t.Log("Source", res.Source)
	t.Log("Vote", res.Vote)
}
func TestGetNominess(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.3.148:9545")
	client := ethClient.GetClient()
	foundationInstance, err := foundation.NewFoundation(common.HexToAddress(FoundationContract), client)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//opts := bind.CallOpts{Pending: true, From: common.HexToAddress(SubAddress), BlockNumber: nil, Context: context.Background()}
	res, err := foundationInstance.Nominees(nil, big.NewInt(0))
	t.Log("Name", res.Name)
	t.Log("Source", res.Source)
	t.Log("Vote", res.Vote)
}

func TestGetVote(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	//client := ethClient.GetClient()
	foundationInstance, err := ethClient.Foundation(FoundationContract)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//opts := bind.CallOpts{Pending: true, From: common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"), BlockNumber: nil, Context: context.Background()}
	res, err := foundationInstance.GetVote(nil, "", common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"))
	t.Log("num", res)
}

//划转到基金会
func TestFlowToFoundation(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	client := ethClient.GetClient()
	foundationInstance, err := ethClient.Foundation(FoundationContract)
	if err != nil {
		t.Error(err.Error())
		return
	}
	gasLimit := uint64(300000)
	auth, err := ethClient.TransactOpts(privateKey, client, gasLimit, big.NewInt(0))
	if err != nil {
		t.Error(err.Error())
		return
	}
	//opts := bind.CallOpts{Pending: true, From: common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"), BlockNumber: nil, Context: context.Background()}
	res, err := foundationInstance.TransferToFoundation(auth, common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"))
	t.Log(res.Hash().Hex())
}

func TestFlowFoundation(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	private, _ := crypto.HexToECDSA(privateKey)

	client := ethClient.GetClient()
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err.Error())
		return
	}
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	if err != nil {
		t.Error(err.Error())
		return
	}
	input, err := parsed.Pack("transferToFoundation", common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	str := `6dc78b74000000000000000000000000fc53af2db0712e02663ed7530de1f52e45c34df3`
	input, _ = hex.DecodeString(str)
	//t.Log(hex.EncodeToString(input))
	res := hex.EncodeToString(input)
	t.Log(res)
	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(address), nil)
	t.Log(nonce)
	if err != nil {
		t.Fatal(err)
	}
	//合约地址(投票部署的合约地址)
	toAddr := common.HexToAddress(FoundationContract)

	value := big.NewInt(0)
	msg := ethereum.CallMsg{GasPrice: gasPrice, Value: value, From: common.HexToAddress(address), To: &toAddr, Data: input}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		t.Error(err.Error())
		return
	}
	//gasLimit = 300000
	tx := types.NewTransaction(
		nonce, toAddr, value,
		gasLimit, gasPrice, input,
	)
	t.Log(tx.Hash().Hex())
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, private)
	if err != nil {
		t.Fatal(err)
	}

	bf, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		t.Fatal(err)
	}
	rawTx := hexutil.Encode(bf)
	t.Log(bf) // 签名结果

	// 交易转发
	txHash, err := ethClient.SendRawTransaction(rawTx)
	//err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(txHash)
}

func TestBack(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.3.146:8545")
	client := ethClient.GetClient()
	foundationInstance, err := ethClient.Foundation(FoundationContract)
	if err != nil {
		t.Error(err.Error())
		return
	}
	gasLimit := uint64(300000)
	auth, err := ethClient.TransactOpts(privateKey, client, gasLimit, big.NewInt(0))
	if err != nil {
		t.Error(err.Error())
		return
	}
	//opts := bind.CallOpts{Pending: true, From: common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"), BlockNumber: nil, Context: context.Background()}
	res, err := foundationInstance.SafeWithdrawal(auth, common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"), big.NewInt(3))
	t.Log(res.Hash().Hex())
}

func TestEthRpcClient_FinishVote(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.3.146:8545")
	client := ethClient.GetClient()
	foundationInstance, err := ethClient.Foundation(FoundationContract)
	if err != nil {
		t.Error(err.Error())
		return
	}
	gasLimit := uint64(300000)
	auth, err := ethClient.TransactOpts("25926cd431628bb7ba2bc0a397a9a66af802da80f83a8d1654c7d2a335dba863", client, gasLimit, big.NewInt(0))
	if err != nil {
		t.Error(err.Error())
		return
	}
	res, err := foundationInstance.FinishVote(auth)
	t.Log(res.Hash().Hex())
}

func TestEthRpcClient_FinishStatus(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.3.148:9545")
	foundationInstance, err := ethClient.Foundation(FoundationContract)
	if err != nil {
		t.Error(err.Error())
		return
	}
	res, err := foundationInstance.Finished(nil)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	t.Log(res)
}