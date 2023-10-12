package model

import "gorm.io/gorm"

type Promote struct {
	gorm.Model
	MediaId int64  `json:"media_id" gorm:"media_id"`
	Title   string `json:"title" gorm:"title"`
	Time    int64  `json:"time" gorm:"time"`
	Sort    int    `json:"sort" gorm:"sort"`
	Status  int    `json:"status" gorm:"status"`
}
