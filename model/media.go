package model

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	PlatId int64  `json:"plat_id" gorm:"plat_id"`
	Title  string `json:"title" gorm:"title"`
	Sort   int    `json:"sort" gorm:"sort"`
	Time   int64  `json:"time" gorm:"time"`
	Status int    `json:"status" gorm:"status"`
}
