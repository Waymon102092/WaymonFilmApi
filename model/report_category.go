package model

import "gorm.io/gorm"

type ReportCategory struct {
	gorm.Model
	Title  string `json:"title" gorm:"title"`
	Sort   int    `json:"sort" gorm:"sort"`
	Time   int64  `json:"time" gorm:"time"`
	Status int    `json:"status" gorm:"status"`
}
