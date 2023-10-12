package model

import "gorm.io/gorm"

type Help struct {
	gorm.Model
	Type    int    `json:"type" gorm:"default:1"` // 1用户端 2票商 3代理端 4员工端
	Title   string `json:"title" gorm:"title"`
	Content string `json:"content" gorm:"content"`
	Time    int64  `json:"time" gorm:"time"`
	Sort    int    `json:"sort" gorm:"sort"`
	Status  int    `json:"status" gorm:"status"`
}
