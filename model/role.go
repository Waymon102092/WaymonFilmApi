package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Title  string `json:"title" gorm:"title"`
	Time   int64  `json:"time" gorm:"time"`
	Sort   int    `json:"sort" gorm:"sort"`
	Status int    `json:"status" gorm:"status"`
}
