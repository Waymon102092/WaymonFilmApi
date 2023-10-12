package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	MemberId int64  `json:"member_id" form:"not null;index"`
	Type     int    `json:"type" form:"type"` //1 文字 2trade 3明细
	Content  string `json:"content" form:"type:text"`
	Param    string `json:"param" form:"param"`
	Time     int64  `json:"time" form:"time"`
	Sort     int    `json:"sort" form:"sort"`
	Status   int    `json:"status" form:"status"`
}
