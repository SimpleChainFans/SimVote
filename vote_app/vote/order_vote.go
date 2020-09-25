package vote

import "github.com/jinzhu/gorm"

/**
 * @Classname order_vote
 * @Author Johnathan
 * @Date 2020/8/11 15:07
 * @Created by Goalnd 2020
 */
type OrderVote struct {
	Id         uint   `gorm:"column:id" json:"id"`
	OutTradeNo string `gorm:"column:outTradeNo" json:"outTradeNo"`
	VoteId     uint   `gorm:"column:voteId" json:"voteId"`
	OrderTime  string `gorm:"column:orderTime" json:"orderTime"`
	Detail     string `gorm:"column:detail" json:"detail"`
	Amount     string `gorm:"column:amount" json:"amount"`
	SignType   string `gorm:"column:signType" json:"signType"`
	Sign       string `gorm:"column:sign" json:"sign"`
}

func (o *OrderVote) TableName() string {
	return "order_vote"
}

func (o *OrderVote) Create(db *gorm.DB) error {
	return db.Table(o.TableName()).Create(o).Error
}

func GetOrderByNo(db *gorm.DB, no string) (*OrderVote, bool) {
	var record OrderVote
	res := db.Table(record.TableName()).Where("outTradeNo=?", no).Scan(&record).RecordNotFound()
	return &record, res
}
