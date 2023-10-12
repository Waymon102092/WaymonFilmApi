package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TradeNo string `json:"trade_no" gorm:"trade_no"` //
	Time    int64  `json:"time" gorm:"time"`
	State   int    `json:"state" gorm:"state"`
	Status  int    `json:"status" gorm:"default:0"`
}
