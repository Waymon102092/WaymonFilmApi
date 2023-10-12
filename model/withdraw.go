package model

import "gorm.io/gorm"

type Withdraw struct {
	gorm.Model
	Type     int   `json:"type" form:"default:1"` // 1
	MemberId int64 `json:"member_id" gorm:"not null;index"`
	AmountId int64 `json:"amount_id" gorm:"not null;index"`
	Money    int   `json:"money" gorm:"money"`
	Price    int   `json:"price" gorm:"price"`
	Time     int64 `json:"time" gorm:"time"`
	Status   int   `json:"status" gorm:"status"` //0 ing 1success -1 fail
}
