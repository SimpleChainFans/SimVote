package cmd

import (
	"github.com/labstack/echo"
	"gitlab.dataqin.com/sipc/vote_app/core"
	"net/http"
	"strings"
)

/**
 * @Classname vote
 * @Author Johnathan
 * @Date 2020/8/11 9:19
 * @Created by Goalnd 2020
 */
/*---------------------------------------------------------------------发起投票---------------------------------------------------------------------*/
func (this *server) GetServiceCharge(c echo.Context) error {
	vo := struct {
		Fee string `json:"fee"`
	}{this.conf.VoteAmount}
	return this.ResponseSuccess(c, vo)
}

type CreateOrdinaryVoteArgs struct {
	Title          string `json:"title" form:"title" validate:"required"`
	Desc           string `json:"desc" form:"desc" validate:"required"`
	Type           int    `json:"type" form:"type" validate:"required"`
	Select         string `json:"select" form:"select" validate:"required"`
	StartAt        string `json:"start_at" form:"start_at"`
	EndAt          string `json:"end_at" form:"end_at"`
	MultipleChoice int    `json:"multiple_choice" form:"multiple_choice"`
	/*	RepeatChoice   int    `json:"repeat_choice"`
		MaxLimit       int    `json:"max_limit"`
		MinLimit       int    `json:"min_limit"`
		TicketVale     int    `json:"ticket_vale"`
		IsShow         int    `json:"is_show"`
		Ticket         int    `json:"ticket"`
		Address        string `json:"address"`*/
}

//发起普通投票生成支付订单
func (this *server) CreateOrdinaryVote(c echo.Context) error {
	userId := this.GetUserId(c) //获取用户id
	req := new(CreateOrdinaryVoteArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	//refererURL := c.Request().Header.Get("Referer")
	refererURL := this.conf.IndexURL + "result/confirming"
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	order, err := vs.CreateNewOrdinaryVote(req.Title, req.Desc, req.Select, req.StartAt, req.EndAt, this.conf.VoteAmount, this.conf.LocalApi, this.conf.AppId, refererURL, this.conf.AppKey, req.Type, req.MultipleChoice, int(userId))
	if err != nil {
		c.Logger().Error(err)
		if err.Error() == "out range" {
			return this.InvalidParametersError(c)
		}
		return this.InternalServiceError(c)
	}
	return this.ResponseSuccess(c, order)
}

//理事会投票
type CreateCouncilVoteArgs struct {
	Title        string `json:"title" form:"title" validate:"required"`
	Desc         string `json:"desc" form:"desc" validate:"required"`
	Select       string `json:"select" form:"select" validate:"required"` //候选人信息
	StartAt      string `json:"start_at" form:"start_at"`
	EndAt        string `json:"end_at" form:"end_at"`
	RepeatChoice int    `json:"repeat_choice" form:"repeat_choice"`
	MaxLimit     int    `json:"max_limit" form:"max_limit" validate:"required"`       // 一个用户最搞起投票数
	MinLimit     int    `json:"min_limit" form:"min_limit" validate:"required"`       // 一个用户最低起投票数
	TicketValue  string `json:"ticket_value" form:"ticket_value" validate:"required"` // 每张投票价值n个sipc
	IsShow       int    `json:"is_show" form:"is_show"`                               // 匿名投票 0：不展示 1：展示
}

type CouncilPayOrdersV0 struct {
	*core.PayOrders
	Input string `json:"input"`
}

func (this *server) CreateCouncilVote(c echo.Context) error {
	userId := this.GetUserId(c) //获取用户id
	req := new(CreateCouncilVoteArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	//refererURL := c.Request().Header.Get("Referer")
	refererURL := this.conf.IndexURL + "result/confirming"
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	result, err := vs.CreateNewCouncilVote(req.Title, req.Desc, req.Select, req.StartAt, req.EndAt, req.TicketValue, refererURL, this.conf.VoteAddress, this.conf.AppId, req.RepeatChoice,
		req.MaxLimit, req.MinLimit, req.IsShow, int(userId))
	if err != nil {
		c.Logger().Error(err)
		return this.InternalServiceError(c)
	}
	return this.ResponseSuccess(c, result)

}

type DeployCouncilContractArgs struct {
	Hash    string `json:"hash" form:"hash" validate:"required"`
	TxSign  string `json:"txSign" form:"txSign" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Nonce   int    `json:"nonce" form:"nonce"`
}

//部署合约
func (this *server) DeployCouncilContract(c echo.Context) error {
	req := new(DeployCouncilContractArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	res := vs.DeployCouncilContract(req.TxSign, req.Hash, req.Address, req.Nonce)
	if !res {
		return this.InternalServiceError(c)
	}
	return this.ResponseSuccess(c, "Success")
}

/*---------------------------------------------------------------------通知接收---------------------------------------------------------------------*/

type ReceiveNotifyArgs struct {
	Version    string `json:"version" form:"version" validate:"required"`
	Appid      string `json:"appId" form:"appId" validate:"required"`
	OutTradeNo string `json:"outTradeNo" form:"outTradeNo" validate:"required"`
	Msg        string `json:"msg" form:"msg"` //放置交易hash
	Status     int    `json:"status" form:"status" validate:"required"`
	SignType   string `json:"signType" form:"signType" validate:"required"`
	Sign       string `json:"sign" form:"sign" validate:"required"`
}

//接收通知
func (this *server) ReceiveNotify(c echo.Context) error {
	req := new(ReceiveNotifyArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	err := vs.ReceiveNotifyService(req.Version, req.Appid, req.OutTradeNo, req.Msg, req.SignType, req.Sign, this.conf.AppKey, req.Status)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, struct {
			Status int    `json:"status"`
			Msg    string `json:"msg"`
		}{0, err.Error()})
	}
	return c.JSON(http.StatusOK, struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	}{1, "Success"})
}

//接收通知
func (this *server) ReceiveCouncilNotify(c echo.Context) error {
	req := new(ReceiveNotifyArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	err := vs.ReceiveCouncilNotifyService(req.Version, req.Appid, req.OutTradeNo, req.Msg, req.SignType, req.Sign, this.conf.AppKey, req.Status)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusOK, struct {
			Status int    `json:"status"`
			Msg    string `json:"msg"`
		}{0, err.Error()})
	}
	return c.JSON(http.StatusOK, struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	}{1, "Success"})
}

/*---------------------------------------------------------------------进行投票---------------------------------------------------------------------*/

type StartVoteArgs struct {
	Hash   string `json:"hash" form:"hash" validate:"required"`
	ItemId string `json:"item_id" form:"item_id" validate:"required"`
}

func (this *server) StartVote(c echo.Context) error {
	userId := this.GetUserId(c) //获取用户id
	req := new(StartVoteArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	items := strings.Split(req.ItemId, ",")
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	code := vs.StartVote(int(userId), req.Hash, items)
	switch code {
	case 1:
		return this.InternalServiceError(c)
	case 4:
		return ResponseOutput(c, code, "不允许重复投票", nil)
	case 5:
		return ResponseOutput(c, code, "未查询到该选项", nil)
	case 6:
		return ResponseOutput(c, code, "投票时间未开启", nil)
	case 8:
		return ResponseOutput(c, code, "等待区块确认", nil)
	case 9:
		return ResponseOutput(c, code, "已达到最大投票人数限制", nil)
	}
	return ResponseOutput(c, code, "投票成功!", nil)
}

type StartCouncilVoteArgs struct {
	Hash   string `json:"hash" form:"hash" validate:"required"`
	ItemId int    `json:"item_id" form:"item_id" validate:"required"`
	Number int    `json:"number" form:"number"  validate:"required"`
	TxSign string `json:"txSign" form:"txSign" validate:"required"`
	Uid    int    `json:"uid" form:"uid"  validate:"required"`
}

func (this *server) StartCouncilVote(c echo.Context) error {
	req := new(StartCouncilVoteArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	code := vs.StartCouncilVote(req.Uid, req.ItemId, req.Number, req.Hash, req.TxSign)
	switch code {
	case 1:
		return this.InternalServiceError(c)
	case 4:
		return ResponseOutput(c, code, "不允许重复投票", nil)
	case 5:
		return ResponseOutput(c, code, "未查询到该选项", nil)
	case 6:
		return ResponseOutput(c, code, "投票时间未开启", nil)
	case 8:
		return ResponseOutput(c, code, "等待区块确认", nil)
	}
	return this.ResponseSuccess(c, "Success")
}

type StartCouncilVoteInputArgs struct {
	Hash   string `json:"hash" form:"hash" validate:"required"`
	ItemId int    `json:"item_id" form:"item_id" validate:"required"`
	Number int    `json:"number" form:"number"  validate:"required"`
}

func (this *server) StartCouncilVoteInput(c echo.Context) error {
	userId := this.GetUserId(c) //获取用户id
	req := new(StartCouncilVoteInputArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	refererURL := this.conf.IndexURL + "view-detail/" + req.Hash
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	result, code := vs.StartCouncilInput(req.Hash, int(userId), req.ItemId, req.Number)
	switch code {
	case -1:
		return this.InvalidParametersError(c)
	case 1:
		return this.InternalServiceError(c)
	case 4:
		return ResponseOutput(c, code, "不允许重复投票", nil)
	case 5:
		return ResponseOutput(c, code, "未查询到该选项", nil)
	case 6:
		return ResponseOutput(c, code, "投票时间未开启", nil)
	case 8:
		return ResponseOutput(c, code, "等待区块确认", nil)
	case 9:
		return ResponseOutput(c, code, "已达到最大投票人数限制", nil)
	}
	result.CallbackURL = refererURL
	return this.ResponseSuccess(c, result)
}

/*---------------------------------------------------------------------投票信息---------------------------------------------------------------------*/

type GetVoteListArgs struct {
	Hash   string `json:"hash" form:"hash"`
	UserId int    `json:"user_id" form:"user_id"`
	Status int    `json:"status" form:"status"`
	Type   int    `json:"type"  form:"type"` // 0.市场 1.我发起的 2.我参与的
	Limit  int    `json:"limit" form:"limit" default:"10"`
}

//投票市场
func (this *server) GetVoteList(c echo.Context) error {
	req := new(GetVoteListArgs)
	var err error
	if err = c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	list := vs.GetVoteSubjectList(req.Status, req.Type, req.UserId, req.Limit, req.Hash)
	return this.ResponseSuccess(c, list)
}

type GetVoteInfoArgs struct {
	Hash   string `json:"hash" form:"hash" validate:"required"`
	UserId int    `json:"user_id" form:"user_id"`
}

//投票详情
func (this *server) GetVoteInfo(c echo.Context) error {
	req := new(GetVoteInfoArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	result, err := vs.GetVoteSubjectInfo(req.UserId, req.Hash)
	if err != nil {
		c.Logger().Error(err)
		return this.InternalServiceError(c)
	}
	if len(result.Hash) == 0 {
		return ResponseOutput(c, 7, "未查询该数据", nil)
	}
	return this.ResponseSuccess(c, result)
}

type VoteDetailArgs struct {
	Hash string `json:"hash" form:"hash" validate:"required" `
}

//投票明细
func (this *server) VoteDetail(c echo.Context) error {
	req := new(VoteDetailArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	list := vs.GetVoteDetail(req.Hash)
	return this.ResponseSuccess(c, list)
}

/*---------------------------------------------------------------------理事会资金返回---------------------------------------------------------------------*/
//合约金额提取
type FoundationReturnArgs struct {
	PrivateKey string `json:"privateKey" form:"privateKey" validate:"required"`
	Hash       string `json:"hash" form:"hash" validate:"required"`
	To         string `json:"to" form:"to" validate:"required"`
	//TxSign string `json:"txSign" form:"txSign" validate:"required"`
}

func (this *server) FoundationReturn(c echo.Context) error {
	req := new(FoundationReturnArgs)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 参数验证
	if err := ValidatorInstance().Struct(req); err != nil {
		c.Logger().Error(err)
		return this.InvalidParametersError(c)
	}
	// 构建Service
	vs := core.NewServiceVote(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	res, tx := vs.TransferToFoundation(req.Hash, req.PrivateKey, req.To)
	if res {
		return this.ResponseSuccess(c, tx)
	} else {
		return ResponseOutput(c, 2, "操作失败", tx)
	}
}
