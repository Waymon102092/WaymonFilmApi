package model

import "gorm.io/gorm"

type Hot struct {
	gorm.Model
	City   string `json:"city" gorm:"city"`
	Time   int64  `json:"time" gorm:"time"`
	Sort   int    `json:"sort" gorm:"sort"`
	Status int    `json:"status" gorm:"status"`
}
