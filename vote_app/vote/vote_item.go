package vote

import "github.com/jinzhu/gorm"

/******sql******
CREATE TABLE `vote_item` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subject_id` int(10) NOT NULL DEFAULT '0' COMMENT '主题id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '选项名称',
  `desc` varchar(20) NOT NULL DEFAULT '' COMMENT '选项描述',
  `img` varchar(20) NOT NULL DEFAULT '' COMMENT '选项图片',
  `address` varchar(30) NOT NULL DEFAULT '' COMMENT '链上地址',
  `vote_number` int(10) NOT NULL DEFAULT '0' COMMENT '选项得票数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='投票选项'
******sql******/
// VoteItem 投票选项
type VoteItem struct {
	ID         int    `gorm:"primary_key" json:"-"`
	SubjectID  int    `json:"subject_id"`  // 主题id
	Name       string `json:"name"`        // 选项名称
	Desc       string `json:"desc"`        // 选项描述
	Img        string `json:"img"`         // 选项图片
	Address    string `json:"address"`     // 链上地址
	VoteNumber int    `json:"vote_number"` // 选项得票数
}

func (v *VoteItem) TableName() string {
	return "vote_item"
}

func (v *VoteItem) Create(db *gorm.DB) error {
	return db.Table(v.TableName()).Create(v).Error
}

func (v *VoteItem) Update(db *gorm.DB) error {
	return db.Table(v.TableName()).Where("id=?", v.ID).Save(v).Error
}

func GetVoteItemByIdAndSid(db *gorm.DB, subId, id int) *VoteItem {
	var record VoteItem
	db.Table(record.TableName()).Where("subject_id=? and id=?", subId, id).Scan(&record)
	return &record
}

type ItemList []VoteItem

func GetVoteItemById(db *gorm.DB, subId int) ItemList {
	var list ItemList
	var record VoteItem
	db.Table(record.TableName()).Where("subject_id=?", subId).Scan(&list)
	return list
}

func GetVoteItemList(db *gorm.DB, subId []int) ItemList {
	var list ItemList
	var record VoteItem
	db.Table(record.TableName()).Where("subject_id in (?)", subId).Scan(&list)
	return list
}
