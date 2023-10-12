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
	"go.uber.org/zap"
	"strings"
	"time"
)

type ReportService struct {
	ReportId         int64  `json:"report_id" form:"report_id"`
	ReportCategoryId int64  `json:"report_category_id" form:"report_category_id"`
	MemberId         int64  `json:"member_id" form:"member_id"`
	Content          string `json:"content" form:"content"`
	Imgs             string `json:"imgs" form:"imgs"`
	Sort             int    `json:"sort" form:"sort"`
	Time             int64  `json:"time" form:"time"`
	Status           int    `json:"status" form:"status"`
	model.BaseLimit
}

func (service *ReportService) ReportInfo(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	if service.ReportId > 0 {
		condition["id"] = service.ReportId
	}
	reportDao := dao.NewReportDao()
	report, err := reportDao.ReportInfo(condition)
	if err != nil {
		code = e.ReportInfoError
		zap.S().Error("ReportInfoError" + err.Error())
		log.WaymonLogger.Error("ReportInfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildReport(report, &model.Member{}, model.ReportCategory{}),
	}
}

func (service *ReportService) ReportList(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	likeCondition := ""
	reportDao := dao.NewReportDao()
	reports, count, err := reportDao.ReportList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.ReportListError
		zap.S().Error("ReportListError" + err.Error())
		log.WaymonLogger.Error("ReportListError" + err.Error())
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
			ItemList: serializer.BuildReports(reports),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *ReportService) ReportAdd(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	imgs := make([]string, 0)
	if service.Imgs != "" {
		imgs = strings.Split(service.Imgs, ",")
	}
	report := &model.Report{
		MemberId:         memberId,
		ReportCategoryId: service.ReportCategoryId,
		Content:          service.Content,
		Time:             time.Now().Unix(),
		Status:           0,
	}
	reportDao := dao.NewReportDao()
	err := reportDao.ReportAdd(report, imgs)
	if err != nil {
		code = e.ReportAddError
		zap.S().Error("ReportAddError" + err.Error())
		log.WaymonLogger.Error("ReportAddError" + err.Error())
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

func (service *ReportService) ReportEdit(ctx context.Context) res.Response {
	code := e.Success
	report := &model.Report{}
	reportDao := dao.NewReportDao()
	err := reportDao.ReportEdit(service.ReportId, report)
	if err != nil {
		code = e.ReportEditError
		zap.S().Error("ReportEditError" + err.Error())
		log.WaymonLogger.Error("ReportEditError" + err.Error())
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

func (service *ReportService) ReportStatus(ctx context.Context) res.Response {
	code := e.Success
	reportDao := dao.NewReportDao()
	err := reportDao.ReportStatus(service.ReportId, service.Status)
	if err != nil {
		code = e.ReportStatusError
		zap.S().Error("ReportStatusError" + err.Error())
		log.WaymonLogger.Error("ReportStatusError" + err.Error())
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

func (service *ReportService) ReportCount(ctx context.Context) res.Response {
	code := e.Success
	reportDao := dao.NewReportDao()
	count, err := reportDao.ReportCount()
	if err != nil {
		code = e.ReportCountError
		zap.S().Error("ReportCountError" + err.Error())
		log.WaymonLogger.Error("ReportCountError" + err.Error())
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

func (service *ReportService) ReportDelete(ctx context.Context) res.Response {
	code := e.Success
	reportDao := dao.NewReportDao()
	err := reportDao.ReportDelete(service.ReportId)
	if err != nil {
		code = e.ReportDeleteError
		zap.S().Error("ReportDeleteError" + err.Error())
		log.WaymonLogger.Error("ReportDeleteError" + err.Error())
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
