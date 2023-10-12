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

type ReportCategoryService struct {
	ReportCategoryId int64  `json:"report_category_id" form:"report_category_id"`
	Title            string `json:"title" form:"title"`
	Sort             int    `json:"sort" form:"sort"`
	Status           int    `json:"status" form:"status"`
	model.BaseLimit
}

func (service *ReportCategoryService) ReportCategoryList(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	if service.ReportCategoryId > 0 {
		condition["id"] = service.ReportCategoryId
	}
	likeCondition := ""
	if service.Title != "" {
		likeCondition += fmt.Sprintf("title LIKE '%s'", "%"+service.Title+"%")
	}
	reportCategoryDao := dao.NewReportCategoryDao()
	reportCategories, count, err := reportCategoryDao.ReportCategoryList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.ReportCategoryListError
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
			ItemList: serializer.BuildReportCategories(reportCategories),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *ReportCategoryService) ReportCategoryAdd(ctx context.Context) res.Response {
	code := e.Success
	reportCategory := &model.ReportCategory{
		Title:  service.Title,
		Sort:   1,
		Time:   time.Now().Unix(),
		Status: 1,
	}
	reportCategoryDao := dao.NewReportCategoryDao()
	err := reportCategoryDao.ReportCategoryAdd(reportCategory)
	if err != nil {
		code = e.ReportCategoryAddError
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

func (service *ReportCategoryService) ReportCategoryInfo(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	if service.ReportCategoryId > 0 {
		condition["reportCategory_id"] = service.ReportCategoryId
	}
	reportCategoryDao := dao.NewReportCategoryDao()
	reportCategory, err := reportCategoryDao.ReportCategoryInfo(condition)
	if err != nil {
		code = e.ReportCategoryInfoError
		zap.S().Error("ReportCategoryInfoError" + err.Error())
		log.WaymonLogger.Error("ReportCategoryInfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildReportCategory(reportCategory),
	}
}

func (service *ReportCategoryService) ReportCategoryEdit(ctx context.Context) res.Response {
	code := e.Success
	reportCategory := &model.ReportCategory{}
	if service.Title != "" {
		reportCategory.Title = service.Title
	}
	if service.Sort > 0 {
		reportCategory.Sort = service.Sort
	}
	if service.Status > -1 {
		reportCategory.Status = service.Status
	}
	reportCategoryDao := dao.NewReportCategoryDao()
	err := reportCategoryDao.ReportCategoryEdit(service.ReportCategoryId, reportCategory)
	if err != nil {
		code = e.ReportCategoryEditError
		zap.S().Error("ReportCategoryEditError" + err.Error())
		log.WaymonLogger.Error("ReportCategoryEditError" + err.Error())
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

func (service *ReportCategoryService) ReportCategoryStatus(ctx context.Context) res.Response {
	code := e.Success
	reportCategoryDao := dao.NewReportCategoryDao()
	err := reportCategoryDao.ReportCategoryStatus(service.ReportCategoryId, service.Status)
	if err != nil {
		code = e.ReportCategoryStatusError
		zap.S().Error("ReportCategoryStatusError" + err.Error())
		log.WaymonLogger.Error("ReportCategoryStatusError" + err.Error())
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

func (service *ReportCategoryService) ReportCategoryCount(ctx context.Context) res.Response {
	code := e.Success
	reportCategoryDao := dao.NewReportCategoryDao()
	count, err := reportCategoryDao.ReportCategoryCount()
	if err != nil {
		code = e.ReportCategoryCountError
		zap.S().Error("ReportCategoryCountError" + err.Error())
		log.WaymonLogger.Error("ReportCategoryCountError" + err.Error())
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

func (service *ReportCategoryService) ReportCategoryDelete(ctx context.Context) res.Response {
	code := e.Success
	reportCategoryDao := dao.NewReportCategoryDao()
	err := reportCategoryDao.ReportCategoryDelete(service.ReportCategoryId)
	if err != nil {
		code = e.ReportCategoryDeleteError
		zap.S().Error("ReportCategoryDeleteError" + err.Error())
		log.WaymonLogger.Error("ReportCategoryDeleteError" + err.Error())
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
