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

type HelpService struct {
	Type     int    `json:"type" form:"type"`
	HelpId   int64  `json:"help_id" form:"help_id"`
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	StarTime int64  `json:"star_time" form:"star_time"`
	EndTime  int64  `json:"end_time" form:"end_time"`
	Sort     int    `json:"sort" form:"sort"`
	Status   int    `json:"status" form:"status"`
	model.BaseLimit
}

func (service *HelpService) HelpInfo(ctx context.Context) res.Response {
	code := e.Success
	helpDao := dao.NewHelpDao()
	help, err := helpDao.HelpInfo(service.HelpId)
	if err != nil {
		code = e.HelpInfoError
		zap.S().Error("HelpInfoError" + err.Error())
		log.WaymonLogger.Error("HelpInfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildHelp(help),
	}
}

func (service *HelpService) HelpList(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	condition["status"] = 1
	if service.HelpId > 0 {
		condition["id"] = service.HelpId
	}
	if service.Type > 0 {
		condition["type"] = service.Type
	}
	likeCondition := ""
	if service.Title != "" {
		likeCondition += fmt.Sprintf("title LIKE '%s'", "%"+service.Title+"%")
	}
	if service.StarTime > 0 {
		likeCondition += fmt.Sprintf("time > %d", service.StarTime)
	}
	if service.EndTime > 0 {
		likeCondition += fmt.Sprintf("time < %d", service.EndTime)
	}
	helpDao := dao.NewHelpDao()
	helps, count, err := helpDao.HelpList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.HelpListError
		zap.S().Error("HelpListError" + err.Error())
		log.WaymonLogger.Error("HelpListError" + err.Error())
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
			ItemList: serializer.BuildHelps(helps),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *HelpService) HelpAdd(ctx context.Context) res.Response {
	code := e.Success
	help := &model.Help{
		Title:   service.Title,
		Content: service.Content,
		Time:    time.Now().Unix(),
		Sort:    service.Sort,
		Status:  service.Status,
	}
	helpDao := dao.NewHelpDao()
	err := helpDao.HelpAdd(help)
	if err != nil {
		code = e.HelpAddError
		zap.S().Error("HelpAddError" + err.Error())
		log.WaymonLogger.Error("HelpAddError" + err.Error())
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

func (service *HelpService) HelpEdit(ctx context.Context) res.Response {
	code := e.Success
	help := &model.Help{}
	if service.Title != "" {
		help.Title = service.Title
	}
	if service.Content != "" {
		help.Content = service.Content
	}
	helpDao := dao.NewHelpDao()
	err := helpDao.HelpEdit(service.HelpId, help)
	if err != nil {
		code = e.HelpEditError
		zap.S().Error("HelpEditError" + err.Error())
		log.WaymonLogger.Error("HelpEditError" + err.Error())
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

func (service *HelpService) HelpStatus(ctx context.Context) res.Response {
	code := e.Success
	helpDao := dao.NewHelpDao()
	err := helpDao.HelpStatus(service.HelpId, service.Status)
	if err != nil {
		code = e.HelpStatusError
		zap.S().Error("HelpStatusError" + err.Error())
		log.WaymonLogger.Error("HelpStatusError" + err.Error())
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

func (service *HelpService) HelpCount(ctx context.Context) res.Response {
	code := e.Success
	helpDao := dao.NewHelpDao()
	count, err := helpDao.HelpCount()
	if err != nil {
		code = e.HelpCountError
		zap.S().Error("HelpCountError" + err.Error())
		log.WaymonLogger.Error("HelpCountError" + err.Error())
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

func (service *HelpService) HelpDelete(ctx context.Context) res.Response {
	code := e.Success
	helpDao := dao.NewHelpDao()
	err := helpDao.HelpDelete(service.HelpId)
	if err != nil {
		code = e.HelpDeleteError
		zap.S().Error("HelpDeleteError" + err.Error())
		log.WaymonLogger.Error("HelpDeleteError" + err.Error())
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
