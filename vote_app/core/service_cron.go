package core

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"gitlab.dataqin.com/sipc/vote_app/ethereum"
	"gitlab.dataqin.com/sipc/vote_app/utils"
	"gitlab.dataqin.com/sipc/vote_app/vote"
	"io/ioutil"
	"math/big"
	"time"
)

/**
 * @Classname service_cron
 * @Author Johnathan
 * @Date 2020/8/11 16:27
 * @Created by Goalnd 2020
 */
type ServiceVoteCron interface {
	VoteStatusUpdate()
	SendVoteCron()
	DeployContractCron()
	ExpiredVoteCron()
	PayTimeOutCron()
	RepeatStartVote()
	VerfiyVoteContract(url, ver string)
	AddHash()
}

type serviceVoteCron struct {
	privateHexKey string
	address       string
	db            *gorm.DB
	ethClient     *ethereum.EthRpcClient
}

func NewServiceVoteCron(db *gorm.DB, node, privateHexKey, address string) ServiceVoteCron {
	ethClient := ethereum.NewEthRpcClient(node)
	return &serviceVoteCron{
		privateHexKey: privateHexKey,
		address:       address,
		db:            db,
		ethClient:     ethClient,
	}
}

/*---------------------------------------------------------------------定时任务---------------------------------------------------------------------*/

func (this *serviceVoteCron) VoteStatusUpdate() {
	list := vote.GetWaitVoteSubjectList(this.db)
	for _, v := range list {
		if v.TxHash == "" || len(v.TxHash) == 0 { //还未接收到通知的投票
			continue
		}
		result, err := this.ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(v.TxHash))
		if err != nil {
			logrus.Warn("TransactionReceipt[TxHash]:", err.Error())
			continue
		}
		if result != nil {
			if result.Status == 1 {
				if v.TypeID == 4 {
					v.Status = vote.DEPLOYING
					v.Update(this.db)
					continue
				}
				if v.SelectHash != "" {
					recepit, err := this.ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(v.SelectHash))
					if err != nil {
						logrus.Warn("TransactionReceipt[SelectHash]:", err.Error())
						break
					}
					if recepit != nil {
						if recepit.Status == 1 {
							v.Status = vote.DEPLOYING
							v.Update(this.db)
						} else { //重新添加选项
							err := addSelect(this.db, this.ethClient, this.privateHexKey, &v)
							if err != nil {
								logrus.Warn("addSelect:", err.Error())
								break
							}
						}
					}
				} else { //添加选项
					err := addSelect(this.db, this.ethClient, this.privateHexKey, &v)
					if err != nil {
						logrus.Error(err.Error())
						break
					}
				}
			}
		}
	}
}

func addSelect(db *gorm.DB, ethClinet *ethereum.EthRpcClient, private string, v *vote.VoteSubject) error {
	nominees := make([]string, 0)
	//添加投票选项
	items := vote.GetVoteItemById(db, v.ID)
	for _, val := range items {
		nominees = append(nominees, val.Name)
	}
	tx, err := ethClinet.AddVoteSelect(v.ConstractAddress, private, nominees)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	v.SelectHash = tx
	v.Update(db)
	return nil
}

func (this *serviceVoteCron) SendVoteCron() {
	list := vote.GetWaitRecordList(this.db)
	for _, v := range list {
		result, _ := this.ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(v.VoteHash))
		if result != nil {
			if result.Status == 0 {
				v.Status = vote.VOTEFAIL
			}
			if result.Status == 1 {
				item := vote.GetVoteItemByIdAndSid(this.db, v.SubjectID, v.ItemID)
				v.Status = vote.VOTESUCCESS
				item.VoteNumber += v.Num
				item.Update(this.db)
			}
			v.Update(this.db)

		}
	}
}

func (this *serviceVoteCron) DeployContractCron() {
	list := vote.GetPaySuccess(this.db)
	for _, v := range list {
		start := v.StartTime.Unix()
		end := v.EndTime.Unix()
		//部署合约
		contractAddress, tx, err := this.ethClient.DeployStoreContract(this.privateHexKey, this.address, v.Title, v.Desc, big.NewInt(0), big.NewInt(start), big.NewInt(end))
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		v.ConstractAddress = contractAddress
		v.TxHash = tx
		err = v.Update(this.db)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
	}
}

func (this *serviceVoteCron) ExpiredVoteCron() {
	list := vote.GetVoteingSubjectList(this.db)
	for _, v := range list {
		if v.EndTime.Before(time.Now()) {
			v.Status = vote.STOP
			v.Update(this.db)
			if v.TypeID == 4 {
				this.ethClient.FinishFoundation(v.ConstractAddress, this.privateHexKey)
			} else { //simple
				this.ethClient.FinishVote(v.ConstractAddress, this.privateHexKey)
			}
		}
	}

}

func (this *serviceVoteCron) PayTimeOutCron() {
	list := vote.GetVotePayStatus(this.db)
	h, _ := time.ParseDuration("-90m")
	lastTime := time.Now().Add(h)
	for _, v := range list {
		if v.CreatedAt.Before(lastTime) {
			v.IsPay = -1
			v.Update(this.db)
		}
	}
}

func (this *serviceVoteCron) RepeatStartVote() {
	end := time.Now()
	h, _ := time.ParseDuration("-1h")
	begin := end.Add(h)
	list := vote.GetFailRecordList(this.db, begin, end)
	for _, v := range list {
		result, _ := this.ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(v.VoteHash))
		if result != nil {
			if result.BlockHash.Hex() != "" {
				if result.Status == 0 {
					//	重新投票
					subject, _ := vote.GetVoteSubjectById(this.db, v.SubjectID)
					if subject.TypeID == 4 {
						continue
					}
					//判断是否过期
					if time.Now().After(subject.EndTime) {
						v.Status = vote.VOTESUCCESS
						v.Update(this.db)
						continue
					}
					items := vote.GetVoteItemById(this.db, subject.ID)
					index := -1
					for i, val := range items {
						if val.ID == v.ItemID {
							index = i
						}
					}
					tx, err := this.ethClient.SendVote(subject.ConstractAddress, this.privateHexKey, index)
					if err != nil {
						logrus.Error(err.Error())
						continue
					}
					v.VoteHash = tx
					v.Update(this.db)
				}
				if result.Status == 1 {
					v.Status = vote.VOTESUCCESS
					v.Update(this.db)
				}
			}

		}
	}
}

// 合约验证
func (this *serviceVoteCron) VerfiyVoteContract(url, ver string) {
	initURL := url + "/init"
	list := vote.GetUnVerfiyVote(this.db)
	for _, v := range list {
		//验证初始化
		initReq := struct {
			Address  string `json:"address"`
			Version  string `json:"version"`
			Compiler string `json:"compiler"`
		}{
			Address:  v.ConstractAddress,
			Version:  ver,
			Compiler: "Solidity(Single File)",
		}
		reqByte, _ := json.Marshal(&initReq)
		_, err := utils.PostRequest(initURL, reqByte)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		//	验证合约
		var file string
		if v.TypeID == 4 {
			b, err := ioutil.ReadFile("./conf/contracts/foundation.sol")
			if err != nil {
				logrus.Error(err.Error())
				continue
			}
			file = string(b)
		} else {
			b, err := ioutil.ReadFile("./conf/contracts/vote.sol")
			if err != nil {
				logrus.Error(err.Error())
				continue
			}
			file = string(b)
		}
		logrus.Info(file[0])
		req := struct {
			Address    string `json:"address"`
			SourceCode string `json:"sourceCode"`
			Optimizer  string `json:"optimizer"`
			Libraries  string `json:"libraries"`
		}{
			Address:    v.ConstractAddress,
			SourceCode: file,
			//Optimizer:  "{\"enabled\":false,\"runs\":200}",
			//Libraries:  "{}",
		}
		reqByte, _ = json.Marshal(&req)
		result, err := utils.PostRequest(url, reqByte)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		resp := make(map[string]interface{})
		err = json.Unmarshal(result, &resp)
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		if resp["err"] != nil { //验证失败
			v.VerifyResult = 2
			v.VerifyMsg = resp["message"].(string)
		} else {
			v.VerifyResult = 1
		}
		v.Update(this.db)
	}
}

func (this *serviceVoteCron) AddHash() {
	list := vote.GetNotHashSubjectList(this.db)
	for _, v := range list {
		hash := utils.RandStringBytes(6)
		for {
			//哈希验重
			_, result := vote.GetVoteSubjectByHash(this.db, hash)
			if !result {
				hash = utils.RandStringBytes(6)
			} else {
				break
			}
		}
		v.Hash = hash
		v.Update(this.db)
	}
	return
}
