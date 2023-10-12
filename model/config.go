package model

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	About       string  `json:"about" gorm:"type:text"`
	Proxy       string  `json:"proxy" gorm:"type:text"`
	Policy      string  `json:"policy" gorm:"type:text"`
	DiscountPre float64 `json:"discount_pre" gorm:"default:0"`
	AgentPre    float64 `json:"agent_pre" gorm:"default:0"`
	StaffPre    float64 `json:"staff_pre" gorm:"default:0"`
	PartnerPre  float64 `json:"partner_pre" gorm:"default:0"` //平台比率
	StopStatus  int     `json:"stop_status" gorm:"default:1"`
	Time        int64   `json:"time" gorm:"time"`
	Status      int     `json:"status" gorm:"default:1"`
}
