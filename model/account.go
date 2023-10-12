package model

import "gorm.io/gorm"

type Account struct { //账户
	gorm.Model
	Type     int    `json:"type" gorm:"default:1"` // 1代理 2票商 3员工
	MemberId int64  `json:"member_id" gorm:"not null;index"`
	Tag      int    `json:"tag" gorm:"default:1"` //只支持支付宝
	Name     string `json:"name" gorm:"name"`
	Card     string `json:"card" gorm:"card"`
	UserName string `json:"user_name" gorm:"user_name"`
	Time     int64  `json:"time" gorm:"time"`
	Status   int    `json:"status" gorm:"status"`
}
