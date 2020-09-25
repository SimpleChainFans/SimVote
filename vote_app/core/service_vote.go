package core

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gitlab.dataqin.com/sipc/vote_app/ethereum"
	"gitlab.dataqin.com/sipc/vote_app/ethereum/foundation"
	"gitlab.dataqin.com/sipc/vote_app/user"
	"gitlab.dataqin.com/sipc/vote_app/utils"
	"gitlab.dataqin.com/sipc/vote_app/vote"
	"math/big"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

/**
 * @Classname vote_service
 * @Author Johnathan
 * @Date 2020/8/10 14:10
 * @Created by Goalnd 2020
 */

type ServiceVote interface {
	GetVoteSubjectList(status, kind, uid, limit int, subHash string) []SubjectList
	GetVoteSubjectInfo(uid int, subHash string) (*SubjectVote, error)
	GetVoteDetail(subHash string) []DetailRecord
	CreateNewOrdinaryVote(title, desc, sel, start, end, amount, host, appid, callBackURL, appkey string, kind, mulChoice, uid int) (*PayOrders, error)
	CreateNewCouncilVote(title, desc, sel, start, end, ticketValue, callBackURL, focus, appid string, repeat, max, min, isShow, uid int) (*CreateCouncilPayOrders, error)
	DeployCouncilContract(sign, hash, address string, nonce int) bool
	StartCouncilInput(subHash string, uid, itemId, number int) (*StartCouncilInfo, int)
	ReceiveNotifyService(version, appid, out, msg, signType, sign, appkey string, status int) error
	ReceiveCouncilNotifyService(version, appid, out, msg, signType, sign, appkey string, status int) error
	StartVote(uid int, subHash string, itemId []string) int
	StartCouncilVote(uid, itemId, number int, subHash, sign string) int
	TransferToFoundation(hash, private, to string) (bool, string)
}

type serviceVote struct {
	privateHexKey string
	address       string
	db            *gorm.DB
	ethClient     *ethereum.EthRpcClient
}

func NewServiceVote(db *gorm.DB, node, privateHexKey, address string) ServiceVote {
	ethClient := ethereum.NewEthRpcClient(node)
	return &serviceVote{
		privateHexKey: privateHexKey,
		address:       address,
		db:            db,
		ethClient:     ethClient,
	}
}

type SubjectList struct {
	//ID        int       `json:"id"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UserID    int       `json:"user_id"`    // 发起投票用户
	TypeID    int8      `json:"type_id"`    // 模板类型1：简易投票 2：长文本类型 3：图片类型 4：基金会类型
	Title     string    `json:"title"`      // 标题
	Desc      string    `json:"desc"`       // 主题介绍
	StartTime time.Time `json:"start_time"` // 开始时间
	EndTime   time.Time `json:"end_time"`   // 截止时间
	Status    int8      `json:"status"`     // 状态：0：部署中 1：部署完成进行中 2：已截止
	Sum       int       `json:"sum"`        //参加人数
	IsStart   int       `json:"isStart"`    //开启状态0未开启1开启
	IsExpried int       `json:"isExpried"`  //有无截止时间0无1有
	Name      string    `json:"name"`       //发起人
	Avatar    string    `json:"avatar"`     //头像
}

func (this *serviceVote) GetVoteSubjectList(status, kind, uid, limit int, subHash string) []SubjectList {
	results := make([]SubjectList, 0)
	var list vote.SubjectList
	var subject *vote.VoteSubject
	var startAt string
	switch kind {
	case 1: //我发起的
		if len(subHash) != 0 && subHash != "" {
			subject, _ = vote.GetVoteSubjectByHash(this.db, subHash)
			startAt = subject.StartTime.Format("2006-01-02 15:04:05")
		}
		list = vote.GetVoteSubjectList(this.db, status, uid, limit, startAt)
	case 2: //参与的
		suIdList := vote.GetSubjectIdByUser(this.db, uid)
		if len(suIdList) == 0 {
			break
		}
		if len(subHash) != 0 && subHash != "" {
			subject, _ = vote.GetVoteSubjectByHash(this.db, subHash)
			startAt = subject.StartTime.Format("2006-01-02 15:04:05")
		}
		sIds := suIdList.SubjectIds()
		list = vote.GetVoteSubjectListByIds(this.db, status, limit, startAt, sIds)
	default: //市场
		if len(subHash) != 0 && subHash != "" {
			subject, _ = vote.GetVoteSubjectByHash(this.db, subHash)
			startAt = subject.StartTime.Format("2006-01-02 15:04:05")
		}
		list = vote.GetVoteSubjectList(this.db, status, 0, limit, startAt)
	}
	//sidList := list.SubjectIds()
	//itemList := vote.GetVoteItemList(this.db, sidList)
	uidList := list.SubjectUIds()
	userList := user.GetUserList(this.db, uidList)
	for _, v := range list {
		totalList := vote.GetSubjectTotalNum(this.db, v.ID)
		sum := len(totalList)
		record := SubjectList{
			//ID:        v.ID,
			Hash:      v.Hash,
			CreatedAt: v.CreatedAt,
			UserID:    v.UserID,
			TypeID:    v.TypeID,
			Title:     v.Title,
			Desc:      v.Desc,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
			Status:    v.Status,
			IsExpried: v.IsExpried,
			Sum:       sum,
		}
		if v.StartTime.Before(time.Now()) {
			record.IsStart = 1
		}
		for _, u := range userList {
			if u.ID == v.UserID {
				record.Name = u.Nickname
				record.Avatar = u.Avatar
				break
			}
		}
		results = append(results, record)
	}
	return results
}

type DetailRecord struct {
	Name     string    `json:"name"`
	CreateAt time.Time `json:"createAt"`
	Select   string    `json:"select"`
	Avatar   string    `json:"avatar"`
	Status   int       `json:"status"`
}

func (this *serviceVote) GetVoteDetail(subHash string) []DetailRecord {
	subject, _ := vote.GetVoteSubjectByHash(this.db, subHash)
	items := vote.GetVoteItemById(this.db, subject.ID)
	records := vote.GetAllRecordList(this.db, subject.ID)
	uids := records.RecordUIds()
	users := user.GetUserList(this.db, uids)
	list := make([]DetailRecord, 0)
	for _, user := range users {
		for _, record := range records {
			if user.ID == record.UserID {
				for _, item := range items {
					if record.ItemID == item.ID {
						detail := DetailRecord{user.Nickname, record.CreatedAt, item.Name, user.Avatar, record.Status}
						list = append(list, detail)
					}
				}
			}
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreateAt.After(list[j].CreateAt)
	})
	return list
}

type SubjectVote struct {
	//ID               int       `json:"id"`
	Hash             string    `json:"hash"`
	CreatedAt        time.Time `json:"created_at"`        // 创建时间
	UserID           int       `json:"user_id"`           // 发起投票用户
	TypeID           int8      `json:"type_id"`           // 模板类型1：简易投票 2：长文本类型 3：图片类型 4：基金会类型
	Title            string    `json:"title"`             // 标题
	Desc             string    `json:"desc"`              // 主题介绍
	StartTime        time.Time `json:"start_time"`        // 开始时间
	EndTime          time.Time `json:"end_time"`          // 截止时间
	MultipleChoice   int8      `json:"multiple_choice"`   // 允许多选，0：不允许多选1：允许多选
	RepeatChoice     int8      `json:"repeat_choice"`     // 允许重复投票，0：不允许 1：允许
	AnonymousChoice  int8      `json:"anonymous_choice"`  // 允许匿名投票，0：不允许 1：允许
	ConstractAddress string    `json:"constract_address"` // 合约地址
	MinLimit         int       `json:"min_limit"`         // 一个用户最低起投票数
	MaxLimit         int       `json:"max_limit"`         // 一个用户最高投票数
	EachTicketValue  float64   `json:"each_ticket_value"` // 每张投票价值n个sipc
	ShowOwner        int8      `json:"show_owner"`        // 是否展示发起投票用户 0：不展示 1：展示
	TicketOwner      int8      `json:"ticket_owner"`      // 投票所得sipc处理方式 0：全部转入基金会 1：全部退还给用户 2：未入选退还给用户
	FoundAddress     string    `json:"found_address"`     // 基金会地址
	Status           int8      `json:"status"`            // 状态：0：部署中 1：部署完成进行中 2：已截止
	SelectInfo       string    `json:"select_info"`       //选项信息(json字符串)
	IsStart          int       `json:"isStart"`           //是否开启0未开启1开启
	Name             string    `json:"name"`              //发起人
	Avatar           string    `json:"avatar"`            //头像
	IsVote           int       `json:"isVote"`            //是否投票过
	IsPay            int       `json:"isPay"`
	IsExpried        int       `json:"isExpried"` //有无截止时间0无1有
}

type SelectArgs struct {
	Id   int    `json:"id"`
	Hash string `json:"hash"`
	//SubjectID  int    `json:"subject_id"`  // 主题id
	Name       string `json:"name"`        // 选项名称
	Desc       string `json:"desc"`        // 选项描述
	Img        string `json:"img"`         // 选项图片
	Address    string `json:"address"`     // 链上地址
	VoteNumber int    `json:"vote_number"` // 选项得票数
	Proportion string `json:"proportion"`  //占比
}

func (this *serviceVote) GetVoteSubjectInfo(uid int, subHash string) (*SubjectVote, error) {
	subject, res := vote.GetVoteSubjectByHash(this.db, subHash)
	if res {
		logrus.Warn("subject not found")
		return nil, errors.New("not found subject")
	}
	//获取用户
	u := user.GetUserById(this.db, subject.UserID)
	//获取选项
	items := vote.GetVoteItemById(this.db, subject.ID)
	var sum int
	list := make([]SelectArgs, 0)
	for _, v := range items {
		sum += v.VoteNumber
	}
	isVote := 0
	//检查是否投票
	if subject.RepeatChoice != 1 {
		if uid != 0 {
			if !vote.GetRecordByUserId(this.db, uid, subject.ID) {
				isVote = 1
			}
		}
	}

	s, _ := decimal.NewFromString(fmt.Sprintf("%d", sum))
	for _, v := range items {
		vNum, _ := decimal.NewFromString(fmt.Sprintf("%d", v.VoteNumber))
		var pro string
		if sum != 0 {
			pro = vNum.Div(s).String()
		} else {
			pro = "0"
		}
		result := SelectArgs{
			Id: v.ID,
			//SubjectID:  v.SubjectID,
			Hash:       subject.Hash,
			Name:       v.Name,
			Desc:       v.Desc,
			Img:        v.Img,
			Address:    v.Address,
			VoteNumber: v.VoteNumber,
			Proportion: pro,
		}
		list = append(list, result)
	}
	bytes, err := json.Marshal(list)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	vo := SubjectVote{
		//ID:               subject.ID,
		Hash:             subject.Hash,
		CreatedAt:        subject.CreatedAt,
		UserID:           subject.UserID,
		TypeID:           subject.TypeID,
		Title:            subject.Title,
		Desc:             subject.Desc,
		StartTime:        subject.StartTime,
		EndTime:          subject.EndTime,
		MultipleChoice:   subject.MultipleChoice,
		RepeatChoice:     subject.RepeatChoice,
		AnonymousChoice:  subject.AnonymousChoice,
		ConstractAddress: subject.ConstractAddress,
		MinLimit:         subject.MinLimit,
		MaxLimit:         subject.MaxLimit,
		EachTicketValue:  subject.EachTicketValue,
		ShowOwner:        subject.ShowOwner,
		TicketOwner:      subject.TicketOwner,
		FoundAddress:     subject.FoundAddress,
		Status:           subject.Status,
		SelectInfo:       string(bytes),
		Name:             u.Nickname,
		Avatar:           u.Avatar,
		IsVote:           isVote,
		IsPay:            subject.IsPay,
		IsExpried:        subject.IsExpried,
	}
	if subject.StartTime.Before(time.Now()) {
		vo.IsStart = 1
	}
	return &vo, nil
}

type PayOrders struct {
	Version     string `json:"version"` //接口版本
	AppId       string `json:"appId"`
	OutTradeNo  string `json:"outTradeNo"`
	OrderTime   int64  `json:"orderTime"`
	Detail      string `json:"detail"`
	Amount      string `json:"amount"`
	NotifyURL   string `json:"notifyURL"`
	CallBackURL string `json:"callbackURL"`
	Wallet      string `json:"wallet"` //钱包类型0中心化1去中心化2两者都支持
	//Address     string `json:"address"` //去中心化地址
	Asset    string `json:"asset"`
	SignType string `json:"signType"`
	Sign     string `json:"sign"`
}

var VoteType = map[int]string{
	1: "简易投票",
	2: "长文本类型投票",
	3: "图片类型投票",
	4: "基金会类型投票",
}

func (this *serviceVote) CreateNewOrdinaryVote(title, desc, sel, start, end, amount, host, appid, callBackURL, appkey string, kind, mulChoice, uid int) (*PayOrders, error) {
	sess := this.db.Begin()
	var startAt, endAt time.Time
	if len(start) == 0 {
		startAt = time.Now()
	} else {
		startAt = utils.String2LocalTime(start)
	}
	ad, _ := time.ParseDuration("24h")
	flagAt := startAt.Add(ad * 30)
	isExpried := 0
	if len(end) == 0 {
		//30天时间
		endAt = startAt.Add(ad * 30)
	} else {
		endAt = utils.String2LocalTime(end)
	}
	isExpried = 1
	if flagAt.Unix() < endAt.Unix() {
		logrus.Warn("time out range")
		return nil, errors.New("out range")
	}
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
	subject := vote.VoteSubject{
		UserID: uid, TypeID: int8(kind), Title: title, Desc: desc, StartTime: startAt, EndTime: endAt, MultipleChoice: int8(mulChoice), IsExpried: isExpried, Hash: hash,
	}
	err := subject.Create(sess)
	if err != nil {
		logrus.Error(err.Error())
		sess.Rollback()
		return nil, err
	}
	//添加选项
	var selectList []vote.VoteItem
	err = json.Unmarshal([]byte(sel), &selectList)
	if err != nil {
		logrus.Error(err.Error())
		sess.Rollback()
		return nil, err
	}
	for _, v := range selectList {
		v.SubjectID = subject.ID
		err = v.Create(sess)
		if err != nil {
			logrus.Error(err.Error())
			sess.Rollback()
			return nil, err
		}
	}
	notify := url.QueryEscape(host + "/api/v1/vote/receive")
	order, err := createOrder(appid, callBackURL, notify, "2", amount, appkey, kind)
	//投票订单关联
	info := vote.OrderVote{
		VoteId:     uint(subject.ID),
		OutTradeNo: order.OutTradeNo,
		OrderTime:  fmt.Sprintf("%d", order.OrderTime),
		Detail:     order.Detail,
		Amount:     order.Amount,
		SignType:   order.SignType,
		Sign:       order.Sign,
	}
	err = info.Create(sess)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	sess.Commit()
	return order, err
}

type CreateCouncilPayOrders struct {
	Hash        string `json:"hash"`    //投票hash
	Version     string `json:"version"` //接口版本
	AppId       string `json:"appId"`
	Detail      string `json:"detail"`
	Amount      string `json:"amount"`
	Wallet      string `json:"wallet"` //钱包类型0中心化1去中心化2两者都支持
	Asset       string `json:"asset"`
	CallBackURL string `json:"callbackURL"`
	Input       string `json:"input"`
	GasPrice    string `json:"gasPrice"`
}

//创建理事会投票
func (this *serviceVote) CreateNewCouncilVote(title, desc, sel, start, end, ticketValue, callBackURL, focus, appid string, repeat, max, min, isShow, uid int) (*CreateCouncilPayOrders, error) {
	fare, err := strconv.ParseFloat(ticketValue, 64)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	sess := this.db.Begin()
	var startAt, endAt time.Time
	if len(start) == 0 {
		startAt = time.Now()
	} else {
		startAt = utils.String2LocalTime(start)
	}
	ad, _ := time.ParseDuration("24h")
	flagAt := startAt.Add(ad * 30)
	isExpried := 0
	if len(end) == 0 {
		//30天时间
		endAt = startAt.Add(ad * 30)
	} else {
		endAt = utils.String2LocalTime(end)
	}
	isExpried = 1
	if flagAt.Unix() < endAt.Unix() {
		logrus.Warn("time out range")
		return nil, errors.New("out range")
	}
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
	subject := vote.VoteSubject{
		Hash:            hash,
		UserID:          uid,
		TypeID:          4,
		Title:           title,
		Desc:            desc,
		StartTime:       startAt,
		EndTime:         endAt,
		RepeatChoice:    int8(repeat),
		MaxLimit:        max,
		MinLimit:        min,
		EachTicketValue: fare,
		ShowOwner:       int8(isShow),
		IsExpried:       isExpried,
	}
	err = subject.Create(sess)
	if err != nil {
		logrus.Error(err.Error())
		sess.Rollback()
		return nil, err
	}
	//添加选项
	var selectList []vote.VoteItem
	err = json.Unmarshal([]byte(sel), &selectList)
	if err != nil {
		logrus.Error(err.Error())
		sess.Rollback()
		return nil, err
	}
	for _, v := range selectList {
		v.SubjectID = subject.ID
		err = v.Create(sess)
		if err != nil {
			logrus.Error(err.Error())
			sess.Rollback()
			return nil, err
		}
	}
	sess.Commit()
	gasPrice, err := this.ethClient.GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	result := CreateCouncilPayOrders{
		Version:     "v1",
		Hash:        hash,
		AppId:       appid,
		Detail:      "部署理事会投票合约",
		Amount:      "0",
		Asset:       "SIPC",
		Wallet:      "1",
		CallBackURL: callBackURL,
		GasPrice:    gasPrice.String(),
	}
	nominees := make([]string, 0)
	for _, v := range selectList {
		nominees = append(nominees, v.Name)
	}
	//start = fmt.Sprintf("%d", startAt.Unix())
	//end = fmt.Sprintf("%d", endAt.Unix())
	input := getVoteInput(title, desc, focus, startAt, endAt, int64(max), int64(min), fare, nominees)
	result.Input = input
	return &result, err
}

func getVoteInput(title, desc, focus string, start, end time.Time, max, min int64, fare float64, nominees []string) string {
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	ticketValue := utils.FloatToBigInt(fare)
	//input, err := parsed.Pack("", title, big.NewInt(startAt.Unix()), big.NewInt(endAt.Unix()), desc, common.HexToAddress(focus))
	input, err := parsed.Pack("", title, big.NewInt(start.Unix()), big.NewInt(end.Unix()), desc, common.HexToAddress(focus),
		ticketValue, big.NewInt(min), big.NewInt(max), nominees, "理事会投票")
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}
	input = append(common.FromHex(foundation.FoundationBin), input...)
	return hex.EncodeToString(input)
}

func (this *serviceVote) DeployCouncilContract(sign, hash, address string, nonce int) bool {
	vote, res := vote.GetVoteSubjectByHash(this.db, hash)
	if res {
		return false
	}
	// 交易转发
	txHash, err := this.ethClient.SendRawTransaction(sign)
	if err != nil {
		logrus.Error("SendRawTransaction:", err.Error())
		return false
	}
	contract := crypto.CreateAddress(common.HexToAddress(address), uint64(nonce))
	vote.TxHash = txHash
	vote.ConstractAddress = contract.Hex()
	vote.IsPay = 1
	vote.Update(this.db)
	return true
}

//创建订单
func createOrder(appid, call, notify, wallet, amount, appkey string, kind int) (*PayOrders, error) {
	//生成投票支付订单
	snowFlake, err := utils.NewSnowflake(2)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	orderTime := time.Now().UnixNano() / 1e6
	callBack := url.QueryEscape(call)
	order := PayOrders{
		Version:     "v1",
		AppId:       appid,
		OutTradeNo:  snowFlake.OrderNum(),
		OrderTime:   orderTime,
		Detail:      VoteType[kind],
		Amount:      amount,
		NotifyURL:   notify,
		CallBackURL: callBack,
		Wallet:      wallet,
		//Address:     address,
		Asset:    "SIPC",
		SignType: "SHA-256",
	}
	order.Sign = createSign(order.Version, order.AppId, order.OutTradeNo, order.Detail, order.Amount, "SIPC", order.SignType, appkey, order.OrderTime)
	return &order, nil
}

func createSign(version, appid, out, detail, amount, asset, signType, appkey string, orderTime int64) string {
	params := make(map[string]string)
	params["version"] = version
	params["appId"] = appid
	params["outTradeNo"] = out
	params["detail"] = detail
	params["amount"] = amount
	params["signType"] = signType
	params["asset"] = asset
	params["orderTime"] = fmt.Sprintf("%d", orderTime)
	sign, err := signCreate(params, appkey, signType)
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}
	return sign
}

// 生成签名
func signCreate(info map[string]string, appkey, signType string) (string, error) {
	//判断签名加密
	var sign string
	if signType == "MD5" { //MD5加密
		strParams := utils.Ksort(info)
		strParams += fmt.Sprintf("&appkey=%s", appkey)
		logrus.Info(strParams)
		m := md5.New()
		m.Write([]byte(strParams))
		code := m.Sum(nil)
		sign = strings.ToUpper(hex.EncodeToString(code))
	} else { //SHA-256加密
		strParams := utils.Ksort(info)
		strParams += fmt.Sprintf("&appkey=%s", appkey)
		logrus.Info(strParams)
		h := sha256.New()
		h.Write([]byte(strParams))
		sha := hex.EncodeToString(h.Sum(nil))
		sign = strings.ToUpper(sha)
	}
	return sign, nil
}

func (this *serviceVote) ReceiveNotifyService(version, appid, out, msg, signType, sign, appkey string, status int) error {
	params := make(map[string]string)
	params["version"] = version
	params["appId"] = appid
	params["outTradeNo"] = out
	params["msg"] = msg
	params["status"] = fmt.Sprintf("%d", status)
	params["signType"] = signType
	theSign, err := signCreate(params, appkey, signType)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	if theSign != sign {
		logrus.Warn("sign verify fail")
		return errors.New("sign verify fail")
	}
	//TODO：签名验证成功部署合约
	receipt, err := this.ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(msg))
	if receipt != nil {
		if receipt.Status == 0 {
			logrus.Warn("transaction fail")
			return errors.New("transaction fail")
		}
	}

	order, res := vote.GetOrderByNo(this.db, out)
	if res {
		logrus.Warn("not found order")
		return errors.New("not found order")
	}
	v, res := vote.GetVoteSubjectById(this.db, int(order.VoteId))
	if res {
		logrus.Warn("not found vote")
		return errors.New("not found vote")
	}
	start := v.StartTime.Unix()
	end := v.EndTime.Unix()
	v.IsPay = vote.PAYSUCCESS
	err = v.Update(this.db)
	//部署合约
	contractAddress, tx, err := this.ethClient.DeployStoreContract(this.privateHexKey, this.address, v.Title, v.Desc, big.NewInt(0), big.NewInt(start), big.NewInt(end))
	if err != nil {
		logrus.Error(err.Error())
		//v.Status = -1
		v.Update(this.db)
		return err
	}
	v.ConstractAddress = contractAddress
	v.TxHash = tx
	err = v.Update(this.db)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}

// 理事会通知接收
func (this *serviceVote) ReceiveCouncilNotifyService(version, appid, out, msg, signType, sign, appkey string, status int) error {
	params := make(map[string]string)
	params["version"] = version
	params["appId"] = appid
	params["outTradeNo"] = out
	params["msg"] = msg
	params["status"] = fmt.Sprintf("%d", status)
	params["signType"] = signType
	theSign, err := signCreate(params, appkey, signType)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	if theSign != sign {
		logrus.Warn("sign verify fail")
		return errors.New("sign verify fail")
	}
	/*验签通过更改状态*/
	receipt, err := this.ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(msg))
	if err != nil {
		logrus.Error("TransactionReceipt:", err.Error())
		return err
	}
	if receipt == nil {
		logrus.Warn("transaction not found")
		return errors.New("transaction not found")
	}
	if receipt.Status == 0 {
		logrus.Warn("transaction fail")
		return errors.New("transaction fail")
	}
	order, res := vote.GetOrderByNo(this.db, out)
	if res {
		logrus.Warn("not found order")
		return errors.New("not found order")
	}
	v, res := vote.GetVoteSubjectById(this.db, int(order.VoteId))
	if res {
		logrus.Warn("not found vote")
		return errors.New("not found vote")
	}
	v.IsPay = vote.PAYSUCCESS
	err = v.Update(this.db)
	return nil
}

func (this *serviceVote) StartVote(uid int, subHash string, itemId []string) int {
	subject, _ := vote.GetVoteSubjectByHash(this.db, subHash)
	if subject.Status == 0 {
		logrus.Warn("等待区块确认")
		return 8
	}
	if subject.StartTime.After(time.Now()) {
		logrus.Warn("投票时间未开启")
		return 6
	}
	if subject.RepeatChoice == 0 {
		record, res := vote.GetRecord(this.db, subject.ID, uid)
		if !res {
			if record.Status != 2 {
				logrus.Warn("重复投票")
				return 4
			}
		}
	}
	count := vote.GetCountNum(this.db, subject.ID)
	if count >= 100 {
		logrus.Warn("已达到最大投票人数限制")
		return 9
	}
	items := vote.GetVoteItemById(this.db, subject.ID)
	index := -1
	if len(itemId) == 1 { //单选
		for i, v := range items {
			id, _ := strconv.Atoi(itemId[0])
			if v.ID == id {
				index = i
				code := sendVote(this.db, this.ethClient, subject, index, id, subject.ID, uid, this.privateHexKey)
				if code != 0 {
					return code
				}
			}
		}
	} else { //多选
		for _, v := range itemId {
			index = -1
			id, _ := strconv.Atoi(v)
			for key, val := range items {
				if val.ID == id {
					index = key
					code := sendVote(this.db, this.ethClient, subject, index, id, subject.ID, uid, this.privateHexKey)
					if code != 0 {
						return code
					}
					break
				}
			}
		}
	}
	if index == -1 {
		logrus.Warn("未查询到该选项")
		return 5
	}
	return 0
}

type CouncilChoice struct {
	ItemId int `json:"itemId"`
	Num    int `json:"num"`
}

//理事会投票
func (this *serviceVote) StartCouncilVote(uid, itemId, number int, subHash, sign string) int {
	subject, _ := vote.GetVoteSubjectByHash(this.db, subHash)
	item := vote.GetVoteItemByIdAndSid(this.db, subject.ID, itemId)
	// 交易转发
	txHash, err := this.ethClient.SendRawTransaction(sign)
	//err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logrus.Error(err.Error())
		return 1
	}
	record := vote.VoteUserRecord{
		CreatedAt: time.Now(),
		SubjectID: subject.ID,
		ItemID:    itemId,
		UserID:    uid,
		Num:       number,
		VoteHash:  txHash,
	}
	receipt, _ := this.ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if receipt != nil {
		if receipt.Status == 1 {
			record.Status = 1
			item.VoteNumber += number
		}
	}
	item.Update(this.db)
	record.Create(this.db)
	return 0
}

type StartCouncilInfo struct {
	GasPrice        string  `json:"gasPrice"`
	Input           string  `json:"input"`
	Hash            string  `json:"hash"`
	Number          int     `json:"number"`
	Amount          float64 `json:"amount"`
	Asset           string  `json:"asset"`
	Detail          string  `json:"detail"`
	ItemId          int     `json:"item_id"`
	ContractAddress string  `json:"contractAddress"`
	TicketValue     float64 `json:"ticketValue"`
	Uid             int     `json:"uid"`
	CallbackURL     string  `json:"callbackURL"`
}

func (this *serviceVote) StartCouncilInput(subHash string, uid, itemId, number int) (*StartCouncilInfo, int) {
	subject, _ := vote.GetVoteSubjectByHash(this.db, subHash)
	if subject.Status == 0 {
		logrus.Warn("等待区块确认")
		return nil, 8
	}
	if subject.StartTime.After(time.Now()) {
		logrus.Warn("投票时间未开启")
		return nil, 6
	}
	if subject.RepeatChoice == 0 { //不允许重复投票
		record, res := vote.GetRecord(this.db, subject.ID, uid)
		if !res {
			if record.Status != 2 {
				logrus.Warn("重复投票")
				return nil, 4
			}
		}
	} else {
		list := vote.GetVoteCount(this.db, subject.ID, uid, itemId)
		count := 0
		for _, v := range list {
			count += v.Num
		}
		if (count + number) > subject.MaxLimit {
			logrus.Warn("超出投票数量限制")
			return nil, 9
		}
	}
	items := vote.GetVoteItemById(this.db, subject.ID)
	num := -1
	for i, v := range items {
		if v.ID == itemId {
			num = i
			break
		}
	}
	if num == -1 {
		logrus.Warn("not found item")
		return nil, num
	}
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	if err != nil {
		logrus.Error(err.Error())
		return nil, 1
	}
	input, err := parsed.Pack("StartVote", big.NewInt(int64(num)))
	if err != nil {
		logrus.Error(err.Error())
		return nil, 1
	}
	gasPrice, err := this.ethClient.GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return nil, 1
	}
	amount, _ := decimal.NewFromFloat(subject.EachTicketValue).Mul(decimal.NewFromInt(int64(number))).Float64()
	result := StartCouncilInfo{
		GasPrice:        gasPrice.String(),
		Input:           hex.EncodeToString(input),
		Hash:            subHash,
		Amount:          amount,
		Detail:          "投票付款",
		Asset:           "SIPC",
		Number:          number,
		ItemId:          itemId,
		ContractAddress: subject.ConstractAddress,
		TicketValue:     subject.EachTicketValue,
		Uid:             uid,
	}
	return &result, 0
}

func sendVote(db *gorm.DB, ethClient *ethereum.EthRpcClient, subject *vote.VoteSubject, index, itemId, subId, uid int, privateHexKey string) int {
	tx, err := ethClient.SendVote(subject.ConstractAddress, privateHexKey, index)
	if err != nil {
		logrus.Error(err.Error())
		return 1
	}
	item := vote.GetVoteItemByIdAndSid(db, subId, itemId)
	u := user.GetUserById(db, uid)
	record := vote.VoteUserRecord{
		SubjectID: subId,
		ItemID:    item.ID,
		UserID:    u.ID,
		VoteHash:  tx,
		Status:    vote.VOTEWAIT,
		Num:       1,
	}
	record.Status = vote.VOTESUCCESS
	item.VoteNumber += 1
	err = item.Update(db)
	if err != nil {
		logrus.Error(err.Error())
		return 1
	}
	/*	result, _ := ethClient.GetClient().TransactionReceipt(context.Background(), common.HexToHash(tx))
		if result != nil {
			if result.BlockHash.Hex() != "" {
				if result.Status == 0 {
					record.Status = vote.VOTEFAIL
				}
				if result.Status == 1 {
					record.Status = vote.VOTESUCCESS
					item.VoteNumber += 1
					err = item.Update(db)
					if err != nil {
						logrus.Error(err.Error())
						return 1
					}
				}
			}
		}*/
	err = record.Create(db)
	if err != nil {
		logrus.Error(err.Error())
		return 1
	}
	return 0
}

/*func transferToFoundationInput(address string) string {
	parsed, err := abi.JSON(strings.NewReader(foundation.FoundationABI))
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}
	input, err := parsed.Pack("transferToFoundation", common.HexToAddress(address))
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}
	return hex.EncodeToString(input)
}*/

func (this *serviceVote) TransferToFoundation(hash, private, to string) (bool, string) {
	subject, res := vote.GetVoteSubjectByHash(this.db, hash)
	if res {
		logrus.Warn("vote not found")
		return false, "vote not found"
	}
	client := this.ethClient.GetClient()
	foundationInstance, err := this.ethClient.Foundation(subject.ConstractAddress)
	if err != nil {
		logrus.Error(err.Error())
		return false, err.Error()
	}
	res, err = foundationInstance.Finished(nil)
	if err != nil {
		logrus.Error(err.Error())
		return false, err.Error()
	}
	if res { //已结束
		gasLimit := uint64(300000)
		auth, err := this.ethClient.TransactOpts(private, client, gasLimit, big.NewInt(0))
		if err != nil {
			logrus.Error(err.Error())
			return false, err.Error()
		}
		//opts := bind.CallOpts{Pending: true, From: common.HexToAddress("0xfc53af2db0712e02663ed7530de1f52e45c34df3"), BlockNumber: nil, Context: context.Background()}
		result, err := foundationInstance.TransferToFoundation(auth, common.HexToAddress(to))
		if err != nil {
			logrus.Error(err.Error())
			return false, err.Error()
		}
		return true, result.Hash().Hex()
	} else { //未结束
		return false, "vote not over"
	}
}
