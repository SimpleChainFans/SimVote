package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.dataqin.com/sipc/vote_app/utils"
	"math/big"
)

func (ethClient *EthRpcClient) SendRawTransaction(rawTransactionData string) (transactionHash string, callErr error) {
	transactionHash = ""
	callErr = ethClient.Client.Call(&transactionHash, "eth_sendRawTransaction", rawTransactionData)
	return
}

type RpcTransaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

func (ethClient *EthRpcClient) GetTransactionByHash(txHash string) (transaction *RpcTransaction, callErr error) {
	transaction = &RpcTransaction{}
	callErr = ethClient.Client.Call(transaction, "eth_getTransactionByHash", txHash)
	return
}

type RpcTransactionReceipt struct {
	TransactionHash   string        `json:"transactionHash"`
	TransactionIndex  string        `json:"transactionIndex"`
	BlockNumber       string        `json:"blockNumber"`
	BlockHash         string        `json:"blockHash"`
	CumulativeGasUsed string        `json:"cumulativeGasUsed"`
	ContractAddress   string        `json:"contractAddress"`
	Logs              []interface{} `json:"logs"`
	logsBloom         string        `json:"logsBloom"`
	Status            string        `json:"status"`
}

// 获取回执
func (ethClient *EthRpcClient) GetTransactionReceipt(txId string) (receipt *RpcTransactionReceipt, callErr error) {
	receipt = &RpcTransactionReceipt{}
	callErr = ethClient.Client.Call(receipt, "eth_getTransactionReceipt", txId)
	return
}

//获取块数
func (ethClient *EthRpcClient) GetBlockNumber() (blockNum string, callErr error) {
	callErr = ethClient.Client.Call(&blockNum, "eth_blockNumber")
	return
}

//获取地址余额
func (ethClient *EthRpcClient) GetBalance(address string) (amount *big.Int, callErr error) {
	var balance string
	callErr = ethClient.Client.Call(&balance, "eth_getBalance", address, "latest")
	amount = utils.HexToBigInt(balance)
	return
}

func (ethClient *EthRpcClient) SignedAndSendTransaction(to string, privateKeyStr string, amount int64) (string, error) {
	toAddress := common.HexToAddress(to)
	client := ethclient.NewClient(ethClient.Client)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	privateKey, err := StringToPrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("publicKey to ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(
		nonce, toAddress, big.NewInt(amount*1e+18),
		uint64(21000), gasPrice, nil,
	)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//signTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil

	//b, err := rlp.EncodeToBytes(signTx)
	//if err != nil {
	//	return "", err
	//}
	//return hex.EncodeToString(b), nil
}

func (ethClient *EthRpcClient) SendTransaction() {

}

func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
