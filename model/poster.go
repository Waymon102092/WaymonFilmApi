package model

import "gorm.io/gorm"

type Poster struct {
	gorm.Model
	Img    string `json:"img" gorm:"img"`
	Sort   int    `json:"sort" gorm:"sort"`
	Time   int64  `json:"time" gorm:"time"`
	Status int    `json:"status" gorm:"status"`
}
