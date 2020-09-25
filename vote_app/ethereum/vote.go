package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"gitlab.dataqin.com/sipc/vote_app/ethereum/vote"
	"math/big"
)

/**
 * @Classname vote
 * @Author Johnathan
 * @Date 2020/8/12 14:21
 * @Created by Goalnd 2020
 */
// 从私钥椭圆地址获取公钥
func (ethClient *EthRpcClient) CommonAddressFromPrivate(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress, nil
}

func (ethClient *EthRpcClient) TransactOpts(
	privateHexKey string, client *ethclient.Client, gasLimit uint64,
	value *big.Int,
) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateHexKey)
	if err != nil {
		return nil, err
	}
	fromAddress, err := ethClient.CommonAddressFromPrivate(privateKey)
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value       // in wei
	auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func (ethClient *EthRpcClient) Vote(contractAddress string) (*vote.Vote, error) {
	client := ethClient.GetClient()
	contractAddr := common.HexToAddress(contractAddress)
	crossInstance, err := vote.NewVote(contractAddr, client)
	if err != nil {
		return nil, err
	}
	return crossInstance, nil
}

// 部署Vote合约
func (ethClient *EthRpcClient) DeployStoreContract(privateHexKey, address, name, info string, value, begin, to *big.Int) (string, string, error) {
	client := ethClient.GetClient()
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", "", err
	}
	data := struct {
		Name  string `json:"name"`
		Begin int    `json:"begin"`
		End   int    `json:"end"`
		Info  string `json:"info"`
	}{
		name,
		int(begin.Int64()),
		int(to.Int64()),
		info,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", "", err
	}
	//abiParsed, err := abi.JSON(strings.NewReader(VoteABI))
	//if err != nil {
	//	return "", "", err
	//}
	//abiParsed.Pack("")
	//str:=string(bytes)
	//resutl:=common.Hex2Bytes(str)
	t := common.BytesToAddress([]byte(nil))
	msg := ethereum.CallMsg{GasPrice: gasPrice, Value: value, Gas: 1000000, From: common.HexToAddress(address), To: &t, Data: bytes}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return "", "", err
	}
	gasLimit = 8000000
	auth, err := ethClient.TransactOpts(privateHexKey, client, gasLimit, value)
	if err != nil {
		return "", "", err
	}
	//auth.GasPrice=big.NewInt(1e16)
	addr, tx, _, err := vote.DeployVote(auth, client, name, begin, to, info)
	if err != nil {
		return "", "", err
	}
	return addr.Hex(), tx.Hash().Hex(), nil
}

// 添加选项
func (ethClient *EthRpcClient) AddVoteSelect(contracAddress, privateHexKey string, nominees []string) (string, error) {
	voteInstance, err := ethClient.Vote(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	client := ethClient.GetClient()
	auth, err := ethClient.TransactOpts(
		privateHexKey,
		client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
	)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	tx, err := voteInstance.AddVoteInfo(auth, nominees)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// 投票
func (ethClient *EthRpcClient) SendVote(contracAddress, privateHexKey string, id int) (string, error) {
	voteInstance, err := ethClient.Vote(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	client := ethClient.GetClient()
	auth, err := ethClient.TransactOpts(
		privateHexKey,
		client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
	)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	tx, err := voteInstance.SendVote(auth, big.NewInt(int64(id)))
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}

type VoteNominess struct {
	Id     *big.Int
	Vote   *big.Int
	Source string
}

// 获取投票信息
func (ethClient *EthRpcClient) VoteInfo(contracAddress string, num int64) (*VoteNominess, error) {
	voteInstance, err := ethClient.Vote(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	result, err := voteInstance.Nominees(nil, big.NewInt(num))
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	res := VoteNominess{
		Id:     result.Id,
		Vote:   result.Vote,
		Source: result.Source,
	}
	return &res, nil
}

// 结束投票
func (ethClient *EthRpcClient) FinishVote(contracAddress, privateHexKey string) (string, error) {
	voteInstance, err := ethClient.Vote(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	client := ethClient.GetClient()
	auth, err := ethClient.TransactOpts(
		privateHexKey,
		client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
	)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	tx, err := voteInstance.FinishVote(auth)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}
