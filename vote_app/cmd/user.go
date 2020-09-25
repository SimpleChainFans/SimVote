package cmd

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"gitlab.dataqin.com/sipc/vote_app/core"
	"time"
)

/**
 * @Classname uion_login
 * @Author Johnathan
 * @Date 2020/8/10 9:18
 * @Created by Goalnd 2020
 */
//联合登录
func (this *server) UnionGet(c echo.Context) error {
	refererURL := c.Request().Header.Get("Referer")
	u2 := uuid.NewV4()
	loginId := u2.String()
	result := struct {
		Version     string `json:"version"`
		LoginId     string `json:"loginId"`
		LoginUrl    string `json:"loginURL"`
		CallBackURL string `json:"callbackURL"`
		AppId       string `json:"appId"`
		Detail      string `json:"detail"`
		SignType    string `json:"signType"`
		Sign        string `json:"sign"`
	}{
		"v1",
		loginId,
		this.conf.LocalApi + "/api/v1/union/verify",
		fmt.Sprintf("%s?loginId=%s", refererURL, loginId),
		//refererURL,
		this.conf.AppId,
		"投票助手联合登录",
		"SHA-256",
		"",
	}
	// 构建Service
	us := core.NewServiceUser(this.db)
	result.Sign = us.CreateUnionSign(result.Version, this.conf.AppId, result.LoginId, result.Detail, result.SignType, this.conf.AppKey)
	return this.ResponseSuccess(c, result)
}

type UnionCheckArgs struct {
	Did        string `json:"did" form:"did" validate:"required"`
	Vc         string `json:"vc" form:"vc" validate:"required"`
	LoginId    string `json:"login_id" form:"login_id" validate:"required"`
	CipherText string `json:"cipher_text" form:"cipher_text" validate:"required"`
	/*	NikeName   string `json:"nike_name" form:"nike_name" validate:"required"`
		Avatar     string `json:"avatar" form:"avatar" validate:"required"`
		SipcId     string `json:"sipc_id" form:"sipc_id" validate:"required"`*/
}

//联合登录检查
func (this *server) UnionCheck(c echo.Context) error {
	req := new(UnionCheckArgs)
	// 参数绑定
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
	us := core.NewServiceUser(this.db)
	//err := us.UnionLoginValidator(req.Did, req.Vc, req.LoginId, req.CipherText, req.NikeName, req.Avatar, req.SipcId, this.conf.SipcApi+"v1/did/verify")
	err, result := us.UnionLoginValidator(req.Did, req.Vc, req.LoginId, req.CipherText, this.conf.SipcApi+"v1/did/verify")
	if err != nil {
		return this.InternalServiceError(c)
	}
	vo := struct {
		Result bool `json:"result"`
	}{result}
	return this.ResponseSuccess(c, vo)
}

type UnionLoginArgs struct {
	LoginId string `json:"login_id" form:"login_id" validate:"required"`
}

func (this *server) UnionLogin(c echo.Context) error {
	req := new(UnionLoginArgs)
	// 参数绑定
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
	us := core.NewServiceUser(this.db)
	user := us.UnionLogin(req.LoginId)
	if user == nil {
		return ResponseOutput(c, 3, "未查询到用户", nil)
	}
	// Set custom claims
	var expireAt int64
	expireAt = time.Now().Add(time.Hour * 12).Unix()
	claims := JwtCustomClaims{
		UserId:      uint(user.ID),
		LoginVerify: 1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(this.conf.Secret))
	if err != nil {
		return this.InternalServiceError(c)
	}
	vo := struct {
		UserId   int    `json:"userId"`
		SipcId   string `json:"sipc_id"`
		NikeName string `json:"nike_name"`
		Avatar   string `json:"avatar"`
		Token    string `json:"token"`
	}{
		user.ID, user.SipcOpenID, user.Nickname, user.Avatar, t,
	}
	return this.ResponseSuccess(c, vo)
}
