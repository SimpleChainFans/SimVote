package user

import (
	"github.com/jinzhu/gorm"
	"time"
)

/******sql******
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(30) NOT NULL DEFAULT '' COMMENT '头像',
  `sipc_open_id` varchar(20) NOT NULL DEFAULT '' COMMENT 'sipc的联合id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sipc_open_id_idx` (`sipc_open_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='用户表'
******sql******/
// User 用户表
type User struct {
	ID         int       `gorm:"primary_key" json:"-"`       // 用户id
	CreatedAt  time.Time `json:"created_at"`                 // 创建时间
	Nickname   string    `json:"nickname"`                   // 昵称
	Avatar     string    `json:"avatar"`                     // 头像
	SipcOpenID string    `gorm:"unique" json:"sipc_open_id"` // sipc的联合id
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Create(db *gorm.DB) error {
	return db.Table(u.TableName()).Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Table(u.TableName()).Where("id=?", u.ID).Save(u).Error
}

func GetUserById(db *gorm.DB, id int) *User {
	var user User
	db.Table(user.TableName()).Where("id=?", id).Scan(&user)
	return &user
}

func GetUserBySipcId(db *gorm.DB, sipcId string) (*User, bool) {
	var user User
	res := db.Table(user.TableName()).Where("sipc_open_id=?", sipcId).Scan(&user).RecordNotFound()
	return &user, res
}

type UserList []User

func GetUserList(db *gorm.DB, uid []int) UserList {
	var list UserList
	var record User
	db.Table(record.TableName()).Where("id in (?)", uid).Scan(&list)
	return list
}
