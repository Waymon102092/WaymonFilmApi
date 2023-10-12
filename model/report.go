package model

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Type             int    `json:"type" gorm:"default:1"`
	MemberId         int64  `json:"member_id" gorm:"not null;index"`
	ReportCategoryId int64  `json:"report_category_id" gorm:"not null;index"`
	Content          string `json:"content" gorm:"content"`
	Sort             int    `json:"sort" gorm:"sort"`
	Time             int64  `json:"time" gorm:"time"`
	Status           int    `json:"status" gorm:"status"`
}
