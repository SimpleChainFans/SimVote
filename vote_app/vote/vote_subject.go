package vote

import (
	"github.com/jinzhu/gorm"
	"time"
)

/******sql******
CREATE TABLE `vote_subject` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '发起投票用户',
  `type_id` tinyint(3) NOT NULL DEFAULT '0' COMMENT '模板类型1：简易投票 2：长文本类型 3：图片类型 4：基金会类型',
  `title` varchar(20) NOT NULL DEFAULT '' COMMENT '标题',
  `desc` varchar(30) NOT NULL DEFAULT '' COMMENT '主题介绍',
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '截止时间',
  `multiple_choice` tinyint(3) NOT NULL DEFAULT '0' COMMENT '允许多选，0：不允许多选1：允许多选',
  `repeat_choice` tinyint(3) NOT NULL DEFAULT '0' COMMENT '允许重复投票，0：不允许 1：允许',
  `anonymous_choice` tinyint(3) NOT NULL DEFAULT '0' COMMENT '允许匿名投票，0：不允许 1：允许',
  `constract_address` varchar(30) NOT NULL DEFAULT '' COMMENT '合约地址',
  `min_limit` int(10) NOT NULL DEFAULT '0' COMMENT '一个用户最低起投票数',
  `max_limit` int(10) NOT NULL DEFAULT '0' COMMENT '一个用户最高投票数',
  `each_ticket_value` int(10) NOT NULL DEFAULT '0' COMMENT '每张投票价值n个sipc',
  `show_owner` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否展示发起投票用户 0：不展示 1：展示',
  `ticket_owner` tinyint(3) NOT NULL DEFAULT '0' COMMENT '投票所得sipc处理方式 0：全部转入基金会 1：全部退还给用户 2：未入选退还给用户',
  `found_address` varchar(30) NOT NULL DEFAULT '' COMMENT '基金会地址',
  `status` tinyint(10) NOT NULL DEFAULT '0' COMMENT '状态：0：部署中 1：部署完成进行中 2：已截止',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='投票主题'
******sql******/
// VoteSubject 投票主题
type VoteSubject struct {
	ID               int       `gorm:"primary_key" json:"-"`
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
	TxHash           string    `gorm:"column:txHash" json:"txHash"`
	SelectHash       string    `gorm:"column:selectHash" json:"selectHash"`
	TicketHash       string    `gorm:"column:ticketHash" json:"ticketHash"`
	MinLimit         int       `json:"min_limit"`                                // 一个用户最低起投票数
	MaxLimit         int       `json:"max_limit"`                                // 一个用户最高投票数
	EachTicketValue  float64   `json:"each_ticket_value"`                        // 每张投票价值n个sipc
	ShowOwner        int8      `json:"show_owner"`                               // 是否展示发起投票用户 0：不展示 1：展示
	TicketOwner      int8      `json:"ticket_owner"`                             // 投票所得sipc处理方式 0：全部转入基金会 1：全部退还给用户 2：未入选退还给用户
	FoundAddress     string    `json:"found_address"`                            // 基金会地址
	IsPay            int       `json:"is_pay"`                                   //支付状态-1:超时取消0:待支付1:已支付
	Status           int8      `json:"status"`                                   // 状态：0：部署中 1：部署完成进行中 2：已截止
	IsExpried        int       `json:"is_expried"`                               //有无截止时间0无1有
	VerifyResult     int       `gorm:"column:verifyResult" json:"verify_result"` //合约验证结果0未验证1成功2失败
	VerifyMsg        string    `gorm:"column:verifyMsg" json:"verify_msg"`       //错误原因
	Hash             string    `json:"hash"`
}

func (v *VoteSubject) TableName() string {
	return "vote_subject"
}

//支付状态
const (
	CANCELPAY  = -1 //支付取消
	WAITPAY    = 0  //待支付
	PAYSUCCESS = 1  //支付成功
)

const (
	DEPLOYWAIT = 0 //部署中
	DEPLOYING  = 1 //部署成功投票进行中
	STOP       = 2 //截止
)

func (v *VoteSubject) Create(db *gorm.DB) error {
	return db.Table(v.TableName()).Create(v).Error
}

func (v *VoteSubject) Update(db *gorm.DB) error {
	return db.Table(v.TableName()).Where("id=?", v.ID).Save(v).Error
}

func GetVoteSubjectById(db *gorm.DB, id int) (*VoteSubject, bool) {
	var record VoteSubject
	res := db.Table(record.TableName()).Where("id=?", id).Scan(&record).RecordNotFound()
	return &record, res
}

func GetVoteSubjectByHash(db *gorm.DB, hash string) (*VoteSubject, bool) {
	var record VoteSubject
	res := db.Table(record.TableName()).Where("hash=?", hash).Scan(&record).RecordNotFound()
	return &record, res
}

type SubjectList []VoteSubject

func GetPaySuccess(db *gorm.DB) SubjectList {
	var list SubjectList
	var record VoteSubject
	db.Table(record.TableName()).Where("is_pay=1 and constract_address='' and type_id!=4").Scan(&list)
	return list
}

func GetNotHashSubjectList(db *gorm.DB) SubjectList {
	var list SubjectList
	var record VoteSubject
	db.Table(record.TableName()).Where("hash=''").Scan(&list)
	return list
}

func GetVoteingSubjectList(db *gorm.DB) SubjectList {
	var list SubjectList
	var record VoteSubject
	db.Table(record.TableName()).Where("status=1").Scan(&list)
	return list
}

func GetVotePayStatus(db *gorm.DB) SubjectList {
	var list SubjectList
	var record VoteSubject
	db.Table(record.TableName()).Where("is_pay=0").Scan(&list)
	return list
}

func GetWaitVoteSubjectList(db *gorm.DB) SubjectList {
	var list SubjectList
	var record VoteSubject
	db.Table(record.TableName()).Where("status=0 and is_pay=1").Scan(&list)
	return list
}

func GetVoteSubjectList(db *gorm.DB, status, uid, limit int, startAt string) SubjectList {
	var list SubjectList
	var record VoteSubject
	d := db.Table(record.TableName())
	if status == 0 { //进行中
		d = d.Where("status!=2")
	} else { //已完成
		d = d.Where("status = 2")
	}
	if uid != 0 {
		d = d.Where("user_id=?", uid)
	}
	if startAt != "" && len(startAt) != 0 {
		d = d.Where("start_time<?", startAt)
	}
	//if sid != 0 {
	//	d = d.Where("id>?", sid)
	//}
	d.Where("is_pay=1").Order("start_time desc,id asc").Limit(limit).Scan(&list)
	return list
}

func GetVoteSubjectListByIds(db *gorm.DB, status, limit int, startAt string, ids []int) SubjectList {
	var list SubjectList
	var record VoteSubject
	d := db.Table(record.TableName()).Where("id in (?)", ids)
	if status == 0 { //进行中
		d = d.Where("status!=2")
	} else { //已完成
		d = d.Where("status = 2")
	}
	if startAt != "" && len(startAt) != 0 {
		d = d.Where("start_time<?", startAt)
	}
	//if sid != 0 {
	//	d = d.Where("id>?", sid)
	//}
	d.Order("start_time desc,id asc").Limit(limit).Scan(&list)
	return list
}

func GetUnVerfiyVote(db *gorm.DB) SubjectList {
	var list SubjectList
	var record VoteSubject
	db.Table(record.TableName()).Where("verifyResult=0 and constract_address!=''").Scan(&list)
	return list
}

func (this *SubjectList) SubjectIds() []int {
	if len(*this) == 0 {
		return nil
	}
	hashes := make([]int, 0, len(*this))
	for _, v := range *this {
		hashes = append(hashes, v.ID)
	}
	return hashes
}

func (this *SubjectList) SubjectUIds() []int {
	if len(*this) == 0 {
		return nil
	}
	users := make([]int, 0, len(*this))
	for _, v := range *this {
		users = append(users, v.UserID)
	}
	return users
}
