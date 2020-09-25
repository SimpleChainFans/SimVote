package cmd

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

/**
 * @Classname router
 * @Author Johnathan
 * @Date 2020/8/10 9:57
 * @Created by Goalnd 2020
 */
func NewRouters(server *server, e *echo.Echo, jwtConfig *middleware.JWTConfig) {
	union := e.Group("api/v1/union")
	RegisterUnion(server, union)

	vote := e.Group("api/v1/vote")
	RegisterVote(server, vote)
	voteAuth := e.Group("api/v1/vote")
	voteAuth.Use(middleware.JWTWithConfig(*jwtConfig))
	voteAuth.Use(LoginVerify())
	RegisterVoteAuth(server, voteAuth)
}

func RegisterUnion(this *server, union *echo.Group) {
	union.GET("", this.UnionGet)
	union.POST("/verify", this.UnionCheck)
	union.POST("/login", this.UnionLogin)
}

func RegisterVote(this *server, vote *echo.Group) {
	vote.GET("/fee", this.GetServiceCharge)              //发起投票费用
	vote.GET("/cron/status", this.VoteSubjectStatusCron) //发起投票状态定时任务
	vote.GET("/cron/send", this.SendVoteCron)            //投票状态定时任务
	vote.GET("/cron/deploy", this.DeployContract)        //合约部署
	vote.GET("/cron/expired", this.ExpiredVote)          //过期状态定时任务
	vote.GET("/cron/timeout", this.PayTimeOut)           //订单超时
	vote.GET("/cron/repeat", this.RepeatStartVote)       //补选项
	vote.GET("/cron/verify", this.VerifyContract)        //合约验证
	vote.GET("/cron/hash", this.AddVoteHash)             //补hash
	vote.POST("/receive", this.ReceiveNotify)            //接收通知
	vote.POST("/info", this.GetVoteInfo)                 //投票详情
	vote.POST("/list", this.GetVoteList)                 //投票市场
	vote.POST("/detail", this.VoteDetail)                //投票明细
	/*----------------------理事会----------------------*/
	//vote.POST("/council/receive", this.ReceiveCouncilNotify)        //理事会投票接收通知
	vote.POST("/council/startInput", this.StartCouncilVoteInput) //理事会投票input
	vote.POST("/council/transfer", this.FoundationReturn)        //合约资金划转
	vote.POST("/council/deploy", this.DeployCouncilContract)     //部署理事会投票
	vote.POST("/council/start", this.StartCouncilVote)           //理事会投票
}

func RegisterVoteAuth(this *server, vote *echo.Group) {
	/*----------------------理事会投票----------------------*/
	vote.POST("/simple", this.CreateOrdinaryVote)                //发起投票
	vote.POST("/start", this.StartVote)                          //投票
	vote.POST("/council/vote", this.CreateCouncilVote)           //发起理事会投票
	vote.POST("/council/startInput", this.StartCouncilVoteInput) //投票appInfo
}

var validate *validator.Validate

// 验证单例
func ValidatorInstance() *validator.Validate {
	validate = validator.New()
	/*	if validate == nil {

		validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
			return utils.IsPhone(fl.Field().String())
		})
		// 验证码规则
		validate.RegisterValidation("code", func(fl validator.FieldLevel) bool {
			return utils.IsValidationCode(fl.Field().String())
		})
		validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			return utils.IsPassword(fl.Field().String())
		})
		validate.RegisterValidation("oldPassword", func(fl validator.FieldLevel) bool {
			return utils.IsOldPasswordRule(fl.Field().String())
		})
		validate.RegisterValidation("address", func(fl validator.FieldLevel) bool {
			return utils.IsCryptoCurrencyAddress(fl.Field().String())
		})
		validate.RegisterValidation("bank_card", func(fl validator.FieldLevel) bool {
			return utils.IsBankCard(fl.Field().String())
		})
	}*/
	return validate
}
