package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	ParentId int64  `json:"parent_id" gorm:"not null;index"`
	Title    string `json:"title" gorm:"title"`
	Icon     string `json:"icon" gorm:"icon"`
	Index    string `json:"index" gorm:"index"`
	Route    string `json:"route" gorm:"route"`
	Time     int64  `json:"time" gorm:"time"`
	Sort     int    `json:"sort" gorm:"sort"`
	Status   int    `json:"status" gorm:"status"`
}
