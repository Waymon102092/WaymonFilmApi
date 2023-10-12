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
)

type AmountService struct {
	MemberId  int64 `json:"member_id" form:"member_id"`
	AmountId  int64 `json:"amount_id" form:"amount_id"`
	TimeStamp int64 `json:"timeStamp" form:"timeStamp"`
	Type      int   `json:"type" form:"type"`
	Money     int   `json:"money" form:"money"`
	Status    int   `json:"status" form:"status"`
	Tag       int   `json:"tag" form:"tag"`
	model.BaseLimit
}

func (service *AmountService) AmountList(ctx context.Context, memberId int64) res.Response {
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
	amountDao := dao.NewAmountDao()
	amounts, count, err := amountDao.AmountList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.AmountListError
		zap.S().Error("AmountListError" + err.Error())
		log.WaymonLogger.Error("AmountListError" + err.Error())
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
			ItemList: serializer.BuildAmounts(amounts),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *AmountService) AmountMoney(ctx context.Context, memberId int64) res.Response {
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
	amountDao := dao.NewAmountDao()
	money, err := amountDao.AmountMoney(condition, likeCondition)
	if err != nil {
		code = e.AmountCountError
		zap.S().Error("AmountCountError" + err.Error())
		log.WaymonLogger.Error("AmountCountError" + err.Error())
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

func (service *AmountService) AmountSettle(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	amountDao := dao.NewAmountDao()
	money, err := amountDao.AmountSettle(memberId)
	if err != nil {
		code = e.AmountCountError
		zap.S().Error("AmountCountError" + err.Error())
		log.WaymonLogger.Error("AmountCountError" + err.Error())
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

func (service *AmountService) AmountAccumulate(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	amountDao := dao.NewAmountDao()
	money, err := amountDao.AmountAccumulate(memberId)
	if err != nil {
		code = e.AmountCountError
		zap.S().Error("AmountCountError" + err.Error())
		log.WaymonLogger.Error("AmountCountError" + err.Error())
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
