package vote

import (
	"github.com/jinzhu/gorm"
	"time"
)

/******sql******
CREATE TABLE `vote_user_record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '投票时间',
  `subject_id` int(10) NOT NULL DEFAULT '0' COMMENT '主题id',
  `item_id` int(10) NOT NULL DEFAULT '0' COMMENT '选项id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '投票用户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户投票记录'
******sql******/
// VoteUserRecord 用户投票记录
type VoteUserRecord struct {
	ID        int       `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"created_at"`                       // 投票时间
	SubjectID int       `json:"subject_id"`                       // 主题id
	ItemID    int       `json:"item_id"`                          // 选项id
	UserID    int       `json:"user_id"`                          // 投票用户id
	VoteHash  string    `gorm:"column:voteHash"  json:"voteHash"` //投票hash
	Num       int       `json:"num"`                              //票数
	Status    int       `json:"status"`                           //投票状态
}

func (v *VoteUserRecord) TableName() string {
	return "vote_user_record"
}

const (
	VOTEWAIT    = 0 //待确定
	VOTESUCCESS = 1 //成功
	VOTEFAIL    = 2 //失败
)

func (v *VoteUserRecord) Create(db *gorm.DB) error {
	return db.Table(v.TableName()).Create(v).Error
}

func (v *VoteUserRecord) Update(db *gorm.DB) error {
	return db.Table(v.TableName()).Where("id=?", v.ID).Save(v).Error
}

func GetRecord(db *gorm.DB, sId, uId int) (*VoteUserRecord, bool) {
	var record VoteUserRecord
	res := db.Table(record.TableName()).Where("subject_id=? and user_id=?", sId, uId).Scan(&record).RecordNotFound()
	return &record, res
}

type VoteUserList []VoteUserRecord

func GetWaitRecordList(db *gorm.DB) VoteUserList {
	var list VoteUserList
	var record VoteUserRecord
	db.Table(record.TableName()).Where("status=0").Scan(&list)
	return list
}

func GetVoteCount(db *gorm.DB, sId, uId, itemId int) VoteUserList {
	var record VoteUserRecord
	var list VoteUserList
	db.Table(record.TableName()).Where("subject_id=? and user_id=? and item_id=? and status=1", sId, uId, itemId).Scan(&list)
	return list
}

func GetFailRecordList(db *gorm.DB, begin, end time.Time) VoteUserList {
	var list VoteUserList
	var record VoteUserRecord
	db.Table(record.TableName()).Where("status!=1 and created_at between ? and ?", begin, end).Scan(&list)
	return list
}

func GetAllRecordList(db *gorm.DB, sid int) VoteUserList {
	var list VoteUserList
	var record VoteUserRecord
	db.Table(record.TableName()).Where("subject_id=? and status!=2", sid).Order("created_at desc").Scan(&list)
	return list
}

func GetSubjectIdByUser(db *gorm.DB, uid int) VoteUserList {
	var list VoteUserList
	var record VoteUserRecord
	db.Table(record.TableName()).Select("subject_id").Where("user_id=?", uid).Group("subject_id").Scan(&list)
	return list
}

func GetReturnUser(db *gorm.DB, sid int, itemId []int) VoteUserList {
	var list VoteUserList
	var record VoteUserRecord
	db.Table(record.TableName()).Where("subject_id=? and item_id in (?)", sid, itemId).Scan(&list)
	return list
}

//人数统计
func GetCountNum(db *gorm.DB, sid int) int {
	var record VoteUserRecord
	var count int
	db.Table(record.TableName()).Where("subject_id=?", sid).Count(&count)
	return count
}

func GetRecordByUserId(db *gorm.DB, uid, sid int) bool {
	var record VoteUserRecord
	return db.Table(record.TableName()).Where("user_id=? and subject_id=?", uid, sid).Scan(&record).RecordNotFound()
}

func (this *VoteUserList) SubjectIds() []int {
	if len(*this) == 0 {
		return nil
	}
	subjects := make([]int, 0, len(*this))
	for _, v := range *this {
		subjects = append(subjects, v.SubjectID)
	}
	return subjects
}

type TotalVo struct {
	UserId int `gorm:"column:user_id" json:"user_id"`
}

type TotalList []TotalVo

func GetSubjectTotalNum(db *gorm.DB, sid int) TotalList {
	var list TotalList
	var record VoteUserRecord
	db.Table(record.TableName()).Select("user_id").Where("subject_id=? and status=1", sid).Group("user_id").Scan(&list)
	return list
}

func (this *VoteUserList) RecordUIds() []int {
	if len(*this) == 0 {
		return nil
	}
	users := make([]int, 0, len(*this))
	for _, v := range *this {
		users = append(users, v.UserID)
	}
	return users
}
