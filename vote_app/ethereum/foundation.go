package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"gitlab.dataqin.com/sipc/vote_app/ethereum/foundation"
	"math/big"
)

/**
 * @Classname foundation
 * @Author Johnathan
 * @Date 2020/8/19 14:42
 * @Created by Goalnd 2020
 */

func (ethClient *EthRpcClient) Foundation(contractAddress string) (*foundation.Foundation, error) {
	client := ethClient.GetClient()
	contractAddr := common.HexToAddress(contractAddress)
	crossInstance, err := foundation.NewFoundation(contractAddr, client)
	if err != nil {
		return nil, err
	}
	return crossInstance, nil
}

// 添加选项
/*func (ethClient *EthRpcClient) AddCouncilNominees(contracAddress, privateHexKey string, nominees []string) (string, error) {
	foundationInstance, err := ethClient.Foundation(contracAddress)
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
	tx, err := foundationInstance.AddNormalNominees(auth, nominees, "理事会选举")
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}*/

// 规则设置
/*func (ethClient *EthRpcClient) SetTicket(contracAddress, privateHexKey string, fare, min, max int64) (string, error) {
	foundationInstance, err := ethClient.Foundation(contracAddress)
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
	tx, err := foundationInstance.SetTicket(auth, big.NewInt(fare), big.NewInt(min), big.NewInt(max))
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}
*/
//结束投票
func (ethClient *EthRpcClient) FinishFoundation(contracAddress, privateHexKey string) (string, error) {
	foundationInstance, err := ethClient.Foundation(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	client := ethClient.GetClient()
	auth, err := ethClient.TransactOpts(
		privateHexKey,
		client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
	)
	tx, err := foundationInstance.FinishVote(auth)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// 是否结束
func (ethClient *EthRpcClient) IsFinish(contracAddress string) (bool, error) {
	foundationInstance, err := ethClient.Foundation(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}
	return foundationInstance.Finished(nil)
}

// 为投票人转账
func (ethClient *EthRpcClient) SafeWithdrawal(contracAddress, privateHexKey, voter string, num int64) (string, error) {
	foundationInstance, err := ethClient.Foundation(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	client := ethClient.GetClient()
	auth, err := ethClient.TransactOpts(
		privateHexKey,
		client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
	)
	tx, err := foundationInstance.SafeWithdrawal(auth, common.HexToAddress(voter), big.NewInt(num))
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// 合约账户提取
func (ethClient *EthRpcClient) Withdrawal(contracAddress, privateHexKey string) (string, error) {
	foundationInstance, err := ethClient.Foundation(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	client := ethClient.GetClient()
	auth, err := ethClient.TransactOpts(
		privateHexKey,
		client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
	)
	tx, err := foundationInstance.Withdraw(auth)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// 合约所有钱转指定账户
func (ethClient *EthRpcClient) TransferToFoundation(contracAddress, privateHexKey, address string) (string, error) {
	foundationInstance, err := ethClient.Foundation(contracAddress)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	client := ethClient.GetClient()
	auth, err := ethClient.TransactOpts(
		privateHexKey,
		client, uint64(5000000), big.NewInt(0), // 300000 gasLimit 0.03 eth
	)
	tx, err := foundationInstance.TransferToFoundation(auth, common.HexToAddress(address))
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return tx.Hash().Hex(), nil
}
