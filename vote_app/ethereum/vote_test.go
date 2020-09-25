package ethereum

import (
	"math/big"
	"testing"
)

/**
 * @Classname vote_test
 * @Author Johnathan
 * @Date 2020/8/12 9:59
 * @Created by Goalnd 2020
 */

const privateHexKey = `1519f22661a26794a9dd7265014dec671eedd6211353b13b5c812b5640fa0064`
const contracAddress = "0x08c7aD6852B848A3824f1C4036CD17c674D721a2"

func TestDeployVote(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	info := `zxz`
	address, txHash, err := ethClient.DeployStoreContract(privateHexKey, "0x95bebb32e53a954464dfd8fb45da3ed1ac645d08", "test!", info, big.NewInt(0), big.NewInt(1597230000), big.NewInt(1597514400))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(address, txHash)
}

func TestVoteTransactor_AddVoteInfo(t *testing.T) {
	ethClient := NewEthRpcClient("http://192.168.4.121:8555")
	selects := []string{
		"A", "B", "C", "D",
	}
	tx, err := ethClient.AddVoteSelect(contracAddress, privateHexKey, selects)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx) //0x0dec72ef7fdde5daa93375a9392a3603e84dd95bb0211cfd18e8f7751404aec2
}

func TestVoteSession_SendVote(t *testing.T) {
	ethClient := NewEthRpcClient("http://47.91.16.204:8545")
	tx, err := ethClient.SendVote(contracAddress, privateHexKey, 0)
	if err != nil {
		t.Fatal(err)
	}
	/*	voteInstance, err := ethClient.Vote(contracAddress)
		if err != nil {
			t.Fatal(err.Error())
		}
		client := ethClient.GetClient()
		auth, err := ethClient.TransactOpts(
			privateHexKey,
			client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
		)
		if err != nil {
			t.Fatal(err)
		}
		transaction, err := voteInstance.SendVote(auth, big.NewInt(0))
		if err != nil {
			t.Fatal(err)
		}*/
	t.Log(tx) //0x2d85a86d62b726925910e303002f133895191de7dad4ed1cb1eb4b53a495a4e3
}
