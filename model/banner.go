package model

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Type   int    `json:"type" gorm:"not null;index"`
	Img    string `json:"img" gorm:"img"`
	Param  string `json:"param" gorm:"param"`
	Sort   int    `json:"sort" gorm:"default:1"`
	Time   int64  `json:"time" gorm:"time"`
	Status int    `json:"status" gorm:"status"`
}
