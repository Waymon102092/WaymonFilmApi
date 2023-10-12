package model

import "gorm.io/gorm"

type Custom struct {
	gorm.Model
	Type   int    `json:"type" gorm:"default:1"` // 1 用户  2 票商 3 抖音
	Img    string `json:"img" gorm:"img"`
	Time   int64  `json:"time" gorm:"time"`
	Sort   int    `json:"sort" gorm:"sort"`
	Status int    `json:"status" gorm:"status"`
}
