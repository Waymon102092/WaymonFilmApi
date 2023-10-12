package model

import "gorm.io/gorm"

type District struct {
	gorm.Model
	DistrictId   int    `json:"district_id" gorm:"not null;index"`
	DistrictName string `json:"district_name" gorm:"district_name"`
	Time         int64  `json:"time" gorm:"time"`
	Status       int    `json:"status" gorm:"status"`
}
