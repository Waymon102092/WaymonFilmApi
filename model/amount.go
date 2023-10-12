package model

import "gorm.io/gorm"

type Amount struct {
	gorm.Model
	Type     int   `json:"type" gorm:"default:1"` // 1 代理  2 票商  3 员工
	MemberId int64 `json:"member_id" gorm:"not null;index"`
	OrderId  int64 `json:"order_id" gorm:"not null;index"`
	TradeId  int64 `json:"trade_id" gorm:"not null;index"`
	Tag      int   `json:"tag" form:"tag"` //1 + 2 -
	Money    int   `json:"money" gorm:"money"`
	Price    int   `json:"price" gorm:"price"`
	Time     int64 `json:"time" gorm:"time"`
	Status   int   `json:"status" gorm:"status"`
}
