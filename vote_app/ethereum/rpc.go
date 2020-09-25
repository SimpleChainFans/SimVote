package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
)

type EthRpcClient struct {
	Client *rpc.Client
}

func NewEthRpcClient(ethHttpRpc string) *EthRpcClient {
	e := new(EthRpcClient)
	ethRpcClient, dialHttpErr := rpc.DialHTTP(ethHttpRpc)
	if dialHttpErr != nil {
		logrus.Fatal("EthRpcClient:", dialHttpErr)
	}
	e.Client = ethRpcClient
	return e
}

func (ethClient *EthRpcClient) GetClient() *ethclient.Client {
	return ethclient.NewClient(ethClient.Client)
}
