package core

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"gitlab.dataqin.com/sipc/vote_app/user"
	"gitlab.dataqin.com/sipc/vote_app/utils"
)

/**
 * @Classname service_user
 * @Author Johnathan
 * @Date 2020/8/10 14:19
 * @Created by Goalnd 2020
 */
type ServiceUser interface {
	//UnionLoginValidator(did, vc, loginId, ciText, name, avatar, sipcId, url string) error
	UnionLoginValidator(did, vc, loginId, ciText, url string) (error, bool)
	UnionLogin(string) *user.User
	CreateUnionSign(version, appid, loginId, detail, signType, appkey string) string
}

type serviceUser struct {
	db *gorm.DB
}

func NewServiceUser(db *gorm.DB) ServiceUser {
	return &serviceUser{db: db}
}

type PostResp struct {
	Code   int    `json:"status"`
	Msg    string `json:"message"`
	Result verify `json:"result"`
}

type verify struct {
	Verified bool   `json:"verified"`
	UserId   string `json:"userId"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

//func (this *serviceUser) UnionLoginValidator(did, vc, loginId, ciText, name, avatar, sipcId, url string) error {
func (this *serviceUser) UnionLoginValidator(did, vc, loginId, ciText, url string) (error, bool) {
	req, err := json.Marshal(struct {
		Did        string `json:"did"`
		Vc         string `json:"vc"`
		LoginId    string `json:"login_id"`
		CipherText string `json:"cipher_text"`
	}{did, vc, loginId, ciText})
	if err != nil {
		logrus.Error(err.Error())
		return err, false
	}
	resByte, err := utils.PostRequest(url, req)
	if err != nil {
		logrus.Error(err.Error())
		return err, false
	}
	var result PostResp
	err = json.Unmarshal(resByte, &result)
	if err != nil {
		logrus.Error(err.Error())
		return err, false
	}
	fmt.Println(result.Result)
	/*	err = json.Unmarshal([]byte(result.Result), &verifyResult)
		if err != nil {
			logrus.Error(err.Error())
			return err, false
		}*/
	if !result.Result.Verified {
		v := user.UnionVerify{
			LoginId:    loginId,
			Did:        did,
			CipherText: ciText,
			SipcId:     result.Result.UserId,
			Status:     user.FAIL,
			IsUse:      0,
		}
		err = v.Create(this.db)
		if err != nil {
			logrus.Error(err.Error())
			return err, false
		}
		return nil, false
	}
	u, res := user.GetUserBySipcId(this.db, result.Result.UserId)
	if res { //没有用户创建
		u := user.User{
			SipcOpenID: result.Result.UserId,
			Avatar:     result.Result.Avatar,
			Nickname:   result.Result.NickName}
		err = u.Create(this.db)
		if err != nil {
			logrus.Error(err.Error())
			return err, false
		}
	}
	if u.Avatar != result.Result.Avatar || u.Nickname != result.Result.NickName {
		if u.Avatar != result.Result.Avatar {
			u.Avatar = result.Result.Avatar
		}
		if u.Nickname != result.Result.NickName {
			u.Nickname = result.Result.NickName
		}
		u.Update(this.db)
	}

	v := user.UnionVerify{
		LoginId:    loginId,
		Did:        did,
		CipherText: ciText,
		SipcId:     result.Result.UserId,
		Status:     user.SUCCESS,
		IsUse:      1,
	}
	err = v.Create(this.db)
	if err != nil {
		logrus.Error(err.Error())
		return err, false
	}
	return nil, true
}

func (this *serviceUser) UnionLogin(loginId string) *user.User {
	result, res := user.GetOneByLoginId(this.db, loginId)
	if res {
		logrus.Warn("loginId used")
		return nil
	}
	if result.Status == 0 { //失败
		logrus.Warn("Login fail")
		return nil
	}
	//	成功-->jwt token构建
	result.IsUse = 0
	err := result.Update(this.db)
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}
	u, res := user.GetUserBySipcId(this.db, result.SipcId)
	if res {
		logrus.Warn("No data")
		return nil
	}
	return u
}

func (this *serviceUser) CreateUnionSign(version, appid, loginId, detail, signType, appkey string) string {
	params := make(map[string]string)
	params["version"] = version
	params["appId"] = appid
	params["loginId"] = loginId
	params["detail"] = detail
	params["signType"] = signType
	strParams := utils.Ksort(params)
	strParams += fmt.Sprintf("&appkey=%s", appkey)
	sign, _ := signCreate(params, appkey, signType)
	return sign
}
