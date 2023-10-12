package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Title  string `json:"title" gorm:"title"`
	Url    string `json:"url" gorm:"url"`
	Time   int64  `json:"time" gorm:"time"`
	Status int    `json:"status" gorm:"default:1"`
}
