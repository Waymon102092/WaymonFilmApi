package services

import (
	"Waymon_api/dao"
	"Waymon_api/model"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/log"
	"Waymon_api/pkg/res"
	"context"
	"go.uber.org/zap"
	"time"
)

type AccountService struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Card     string `json:"card" form:"card"`
	Tag      int    `json:"tag" form:"tag"`
}

func (service *AccountService) AccountInfo(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["member_id"] = memberId
	accountDao := dao.NewAccountDao()
	account, err := accountDao.AccountInfo(condition)
	if err != nil {
		code = e.AccountInfoError
		zap.S().Error("AccountInfoError" + err.Error())
		log.WaymonLogger.Error("AccountInfoError" + err.Error())
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   account.ID,
	}
}

func (service *AccountService) AccountAdd(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	account := &model.Account{
		MemberId: memberId,
		Name:     service.Name,
		Card:     service.Card,
		UserName: service.Username,
		Tag:      service.Tag,
		Time:     time.Now().Unix(),
		Status:   0,
	}
	accountDao := dao.NewAccountDao()
	err := accountDao.AccountAdd(account)
	if err != nil {
		code = e.AccountAddError
		zap.S().Error("AccountAddError" + err.Error())
		log.WaymonLogger.Error("AccountAddError" + err.Error())
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   nil,
	}
}
