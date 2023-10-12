package model

import "gorm.io/gorm"

type ReportImg struct {
	gorm.Model
	ReportId int64  `json:"report_id" gorm:"not null;index"`
	Img      string `json:"img" gorm:"img"`
	Sort     int    `json:"sort" gorm:"sort"`
	Time     int64  `json:"time" gorm:"time"`
	Status   int    `json:"status" gorm:"status"`
}
