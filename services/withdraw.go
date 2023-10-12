package services

import (
	"Waymon_api/dao"
	"Waymon_api/model"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/log"
	"Waymon_api/pkg/res"
	"Waymon_api/pkg/waymon"
	"Waymon_api/serializer"
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type WithdrawService struct {
	MemberId   int64 `json:"member_id" form:"member_id"`
	WithdrawId int64 `json:"withdraw_id" form:"withdraw_id"`
	TimeStamp  int64 `json:"timeStamp" form:"timeStamp"`
	Type       int   `json:"type" form:"type"`
	Tip        int   `json:"tip" form:"tip"`
	Tag        int   `json:"tag" form:"tag"`
	Money      int   `json:"money" form:"money"`
	Status     int   `json:"status" form:"status"`
	model.BaseLimit
}

func (service *WithdrawService) WithdrawInfo(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	if service.WithdrawId > 0 {
		condition["id"] = service.WithdrawId
	}
	withdrawDao := dao.NewWithdrawDao()
	withdraw, err := withdrawDao.WithdrawInfo(condition)
	if err != nil {
		code = e.WithdrawInfoError
		zap.S().Error("WithdrawInfoError" + err.Error())
		log.WaymonLogger.Error("WithdrawInfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildWithdraw(withdraw),
	}
}

func (service *WithdrawService) WithdrawList(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	condition["member_id"] = memberId
	likeCondition := ""
	if service.TimeStamp > 0 {
		starTime, endTime := waymon.MonthDuration(service.TimeStamp)
		likeCondition += fmt.Sprintf("time > %d and time < %d", starTime, endTime)
	}
	withdrawDao := dao.NewWithdrawDao()
	withdraws, count, err := withdrawDao.WithdrawList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.WithdrawListError
		zap.S().Error("WithdrawListError" + err.Error())
		log.WaymonLogger.Error("WithdrawListError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: res.Responses{
			ItemList: serializer.BuildWithdraws(withdraws),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *WithdrawService) WithdrawAdd(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	withdraw := &model.Withdraw{
		MemberId: memberId,
		Money:    service.Money,
		Type:     service.Type,
		Time:     time.Now().Unix(),
		Status:   0,
	}
	withdrawDao := dao.NewWithdrawDao()
	err := withdrawDao.WithdrawAdd(withdraw)
	if err != nil {
		code = e.WithdrawAddError
		zap.S().Error("WithdrawAddError" + err.Error())
		log.WaymonLogger.Error("WithdrawAddError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   nil,
	}
}

func (service *WithdrawService) WithdrawEdit(ctx context.Context) res.Response {
	code := e.Success
	withdraw := &model.Withdraw{}
	withdrawDao := dao.NewWithdrawDao()
	err := withdrawDao.WithdrawEdit(service.WithdrawId, withdraw)
	if err != nil {
		code = e.WithdrawEditError
		zap.S().Error("WithdrawEditError" + err.Error())
		log.WaymonLogger.Error("WithdrawEditError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   nil,
	}
}

func (service *WithdrawService) WithdrawStatus(ctx context.Context) res.Response {
	code := e.Success
	withdrawDao := dao.NewWithdrawDao()
	err := withdrawDao.WithdrawStatus(service.WithdrawId, service.Status)
	if err != nil {
		code = e.WithdrawStatusError
		zap.S().Error("WithdrawStatusError" + err.Error())
		log.WaymonLogger.Error("WithdrawStatusError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   nil,
	}
}

func (service *WithdrawService) WithdrawMoney(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["member_id"] = memberId
	if service.Status > 0 {
		condition["status"] = service.Status
	}
	likeCondition := ""
	if service.Tag == 1 {
		startTime, endTime := waymon.GetTimeStamp(1)
		likeCondition += fmt.Sprintf("time > %d and time < %d", startTime, endTime)
	}
	if service.Tag == 2 {
		startTime, endTime := waymon.GetTimeStamp(5)
		likeCondition += fmt.Sprintf("time > %d and time < %d", startTime, endTime)
	}
	if service.Tag == 3 {
		startTime, endTime := waymon.GetTimeStamp(8)
		likeCondition += fmt.Sprintf("time > %d and time < %d", startTime, endTime)
	}
	withdrawDao := dao.NewWithdrawDao()
	money, err := withdrawDao.WithdrawMoney(condition, likeCondition)
	if err != nil {
		code = e.AmountStatusError
		zap.S().Error("AmountStatusError" + err.Error())
		log.WaymonLogger.Error("AmountStatusError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   money,
	}
}
