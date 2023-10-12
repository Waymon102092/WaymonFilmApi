package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PassWordCost = 12 //密码加密难度
)

type Admin struct {
	gorm.Model
	RoleId   int64  `json:"role_id" gorm:"not null;index"`
	UserName string `json:"user_name" gorm:"user_name"`
	NickName string `json:"nick_name" gorm:"nick_name"`
	Password string `json:"password" gorm:"password"`
	Logins   int    `json:"logins" gorm:"default:0"`
	LastTime int64  `json:"last_time" gorm:"last_time"`
	LastIp   string `json:"last_ip" gorm:"last_ip"`
	Time     int64  `json:"time" gorm:"time"`
	Status   int    `json:"status" gorm:"status"`
}

func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	admin.Password = string(bytes)
	return nil
}

func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	return err == nil
}
