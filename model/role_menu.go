package model

import "gorm.io/gorm"

type RoleMenu struct {
	gorm.Model
	RoleId int64 `json:"role_id" gorm:"not null;index"`
	MenuId int64 `json:"menu_id" gorm:"not null;index"`
	Time   int64 `json:"time" gorm:"time"`
	Sort   int   `json:"sort" gorm:"sort"`
	Status int   `json:"status" gorm:"status"`
}
