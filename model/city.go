package model

import "gorm.io/gorm"

type City struct {
	gorm.Model
	CityId     int    `json:"city_id" gorm:"not null;index"`
	CityName   string `json:"city_name" gorm:"city_name"`
	CityLetter string `json:"city_letter" gorm:"city_letter"`
	Time       int64  `json:"time" gorm:"time"`
	Status     int    `json:"status" gorm:"status"`
}
