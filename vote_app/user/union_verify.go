package user

import "github.com/jinzhu/gorm"

/**
 * @Classname union_verify
 * @Author Johnathan
 * @Date 2020/8/10 14:39
 * @Created by Goalnd 2020
 */
type UnionVerify struct {
	Id         uint   `gorm:"column:id" json:"id"`
	LoginId    string `gorm:"column:loginId" json:"loginId"`
	Did        string `gorm:"column:did" json:"did"`
	CipherText string `gorm:"column:cipherText" json:"cipherText"`
	SipcId     string `gorm:"column:sipcId" json:"sipcId"`
	Status     int    `gorm:"column:status" json:"status"`
	IsUse      int    `gorm:"column:isUse" json:"isUse"`
}

const (
	FAIL    = 0
	SUCCESS = 1
)

func (u *UnionVerify) TableName() string {
	return "union_verify"
}

func (u *UnionVerify) Create(db *gorm.DB) error {
	return db.Table(u.TableName()).Create(u).Error
}

func (u *UnionVerify) Update(db *gorm.DB) error {
	return db.Table(u.TableName()).Save(u).Error
}

func GetOneByLoginId(db *gorm.DB, loginId string) (*UnionVerify, bool) {
	var result UnionVerify
	res := db.Table(result.TableName()).Where("loginId=? and isUse=1", loginId).Scan(&result).RecordNotFound()
	return &result, res
}
