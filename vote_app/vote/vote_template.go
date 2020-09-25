package vote

/******sql******
CREATE TABLE `vote_template` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(20) NOT NULL DEFAULT '' COMMENT '模板名称',
  `type_id` int(1) NOT NULL DEFAULT '0' COMMENT '模板类型1：简易投票 2：长文本类型 3：图片类型 4：基金会类型',
  `img` varchar(50) NOT NULL DEFAULT '' COMMENT '模板封面',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='选举模板'
******sql******/
// VoteTemplate 选举模板
type VoteTemplate struct {
	ID     int    `gorm:"primary_key" json:"-"`
	Title  string `json:"title"`   // 模板名称
	TypeID int    `json:"type_id"` // 模板类型1：简易投票 2：长文本类型 3：图片类型 4：基金会类型
	Img    string `json:"img"`     // 模板封面
}
