package model

import "gorm.io/gorm"

type Apply struct {
	gorm.Model
	MemberId int64  `json:"member_id" gorm:"not null;index"`
	UserName string `json:"user_name" gorm:"user_name"`
	Cinema   string `json:"cinema" gorm:"cinema"`
	Express  string `json:"express" gorm:"express"`
	Name     string `json:"name" gorm:"name"`
	Time     int64  `json:"time" gorm:"time"`
	Status   int    `json:"status"`
}
