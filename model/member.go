package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Tag       int    `json:"tag" gorm:"default:1"` // 1小程序 2公众号 3抖音 4支付宝
	MediaId   int64  `json:"media_id" gorm:"media_id"`
	PromoteId int64  `json:"promote_id" gorm:"promote_id"`
	ParentId  int64  `json:"parent_id" gorm:"index"`
	StaffId   int64  `json:"staff_id" gorm:"index"`
	Tel       string `json:"tel" gorm:"column:tel"`
	NickName  string `json:"nick_name" gorm:"column:nick_name"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url"`
	Code      string `json:"code" gorm:"not null;index"`
	Wechat    string `json:"wechat" gorm:"wechat"`
	WeOpenId  string `json:"we_open_id" gorm:"index"`
	OpenId    string `json:"open_id" gorm:"index"`
	UnionId   string `json:"union_id" gorm:"not null;index"`
	TTOpenId  string `json:"tt_open_id" gorm:"index"`
	TTUnionId string `json:"tt_union_id" gorm:"index"`
	Gender    string `json:"gender" gorm:"column:gender"`
	Province  string `json:"province" gorm:"column:province"`
	City      string `json:"city" gorm:"column:city"`
	District  string `json:"district" gorm:"district"`
	Money     int    `json:"money" gorm:"default:0"`
	CodeCover string `json:"code_cover" gorm:"code_cover"` //二维码图片 2 员工权限
	Time      int64  `json:"time" gorm:"column:time"`
	State     int    `json:"state" gorm:"default:0"`  // 1认证
	Status    int    `json:"status" gorm:"default:0"` // 0 用户 1 代理 2 员工 3 票商
}
