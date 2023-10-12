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

type FansService struct {
	Tip       int   `json:"tip" form:"tip"`
	Tag       int   `json:"tag" form:"tag"`
	TimeStamp int64 `json:"timeStamp" form:"timeStamp"`
	model.BaseLimit
}

func (service *FansService) FansList(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	if service.Tip == 1 { //代理
		condition["parent_id"] = memberId
	} else { //员工
		condition["staff_id"] = memberId
	}
	starTime, endTime := waymon.MonthDuration(service.TimeStamp)
	likeCondition := fmt.Sprintf("time > %d and time < %d", starTime, endTime)
	fansDao := dao.NewFansDao()
	fans, count, err := fansDao.FansList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.MemberListError
		zap.S().Error("MemberListError" + err.Error())
		log.WaymonLogger.Error("MemberListError" + err.Error())
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
			ItemList: serializer.BuildFans(fans),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *FansService) FansCount(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	if service.Tip == 1 { //代理
		condition["parent_id"] = memberId
	} else { //员工
		condition["staff_id"] = memberId
	}
	var startTime, endTime int64
	if service.Tag == 1 {
		startTime, endTime = waymon.GetTimeStamp(1)
	}
	if service.Tag == 2 {
		startTime, endTime = waymon.GetTimeStamp(5)
	}
	if service.Tag == 3 {
		startTime, endTime = waymon.GetTimeStamp(8)
	}
	fansDao := dao.NewFansDao()
	count, err := fansDao.FansCount(condition, startTime, endTime)
	if err != nil {
		code = e.MemberCountError
		zap.S().Error("MemberCountError" + err.Error())
		log.WaymonLogger.Error("MemberCountError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   count,
	}
}

func (service *FansService) FansOrder(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	condition["member.parent_id"] = memberId
	condition["order.status"] = 3
	likeCondition := ""
	var startTime, endTime int64
	if service.Tag == 1 {
		startTime, endTime = waymon.GetTimeStamp(1)
		likeCondition += fmt.Sprintf("order.time > %d and order.time < %d", startTime, endTime)
	}
	if service.Tag == 2 {
		startTime, endTime = waymon.GetTimeStamp(5)
		likeCondition += fmt.Sprintf("order.time > %d and order.time < %d", startTime, endTime)
	}
	if service.Tag == 3 {
		startTime, endTime = waymon.GetTimeStamp(8)
		likeCondition += fmt.Sprintf("order.time > %d and order.time < %d", startTime, endTime)
	}
	fansDao := dao.NewFansDao()
	orders, count, err := fansDao.FansOrder(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.FansOrderError
		zap.S().Error("OrderCountError" + err.Error())
		log.WaymonLogger.Error("OrderCountError" + err.Error())
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
			ItemList: serializer.BuildFansOrders(orders),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *FansService) FansOrderCount(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["member.parent_id"] = memberId
	condition["order.status"] = 3
	likeCondition := ""
	var startTime, endTime int64
	if service.Tag == 1 {
		startTime, endTime = waymon.GetTimeStamp(1)
		likeCondition += fmt.Sprintf("order.time > %d and order.time < %d", startTime, endTime)
	}
	if service.Tag == 2 {
		startTime, endTime = waymon.GetTimeStamp(5)
		likeCondition += fmt.Sprintf("order.time > %d and order.time < %d", startTime, endTime)
	}
	if service.Tag == 3 {
		startTime, endTime = waymon.GetTimeStamp(8)
		likeCondition += fmt.Sprintf("order.time > %d and order.time < %d", startTime, endTime)
	}
	fansDao := dao.NewFansDao()
	count, err := fansDao.FansOrderCount(condition, likeCondition)
	if err != nil {
		code = e.FansOrderCountError
		zap.S().Error("OrderCountError" + err.Error())
		log.WaymonLogger.Error("OrderCountError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   count,
	}
}
