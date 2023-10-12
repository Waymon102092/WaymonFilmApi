package services

import (
	"Waymon_api/dao"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/log"
	"Waymon_api/pkg/res"
	"Waymon_api/serializer"
	"context"
	"go.uber.org/zap"
)

type CustomService struct {
	Type int64 `json:"type" form:"type"`
}

func (service *CustomService) CustomInfo(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	condition["type"] = service.Type
	condition["status"] = 1
	customDao := dao.NewCustomDao()
	custom, err := customDao.CustomInfo(condition)
	if err != nil {
		code = e.CustomInfoError
		zap.S().Error("CustomInfoError" + err.Error())
		log.WaymonLogger.Error("CustomInfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCustom(custom),
	}
}
