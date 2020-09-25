package cmd

import (
	"github.com/labstack/echo"
	"gitlab.dataqin.com/sipc/vote_app/core"
)

/**
 * @Classname cron
 * @Author Johnathan
 * @Date 2020/8/12 16:13
 * @Created by Goalnd 2020
 */
func (this *server) VoteSubjectStatusCron(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.VoteStatusUpdate()
	return this.ResponseSuccess(c, "Success")
}

func (this *server) SendVoteCron(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.SendVoteCron()
	return this.ResponseSuccess(c, "Success")
}

func (this *server) DeployContract(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.DeployContractCron()
	return this.ResponseSuccess(c, "Success")
}

func (this *server) ExpiredVote(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.ExpiredVoteCron()
	return this.ResponseSuccess(c, "Success")
}

func (this *server) PayTimeOut(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.PayTimeOutCron()
	return this.ResponseSuccess(c, "Success")
}

func (this *server) RepeatStartVote(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.RepeatStartVote()
	return this.ResponseSuccess(c, "Success")
}

func (this *server) VerifyContract(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.VerfiyVoteContract(this.conf.VerifyURL, this.conf.ContractVersion)
	return this.ResponseSuccess(c, "Success")
}

func (this *server) AddVoteHash(c echo.Context) error {
	if c.Request().Host != "localhost:7688" {
		// 无权限访问
		return this.NoPermissionError(c)
	}
	vs := core.NewServiceVoteCron(this.db, this.conf.EthRpcHost, this.conf.VotePrivateKey, this.conf.VoteAddress)
	vs.AddHash()
	return this.ResponseSuccess(c, "Success")
}
