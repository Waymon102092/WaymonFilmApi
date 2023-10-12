package services

import (
	"Waymon_api/dao"
	"Waymon_api/model"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/log"
	"Waymon_api/pkg/res"
	"Waymon_api/serializer"
	"context"
	"go.uber.org/zap"
)

type ConfigService struct {
	Id     int64  `json:"id" form:"id"`
	About  string `json:"about" form:"about"`
	Proxy  string `json:"proxy" form:"proxy"`
	Policy string `json:"policy" form:"policy"`
	Status int    `json:"status" form:"status"`
}

func (service *ConfigService) ConfigInfo(ctx context.Context) res.Response {
	code := e.Success
	configDao := dao.NewConfigDao()
	config, err := configDao.ConfigInfo()
	if err != nil {
		code = e.ConfigInfoError
		zap.S().Error("InfoError" + err.Error())
		log.WaymonLogger.Error("InfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildConfig(config),
	}
}

func (service *ConfigService) ConfigEdit(ctx context.Context) res.Response {
	code := e.Success
	config := &model.Config{}
	if service.About != "" {
		config.About = service.About
	}
	if service.Proxy != "" {
		config.Proxy = service.Proxy
	}
	if service.Policy != "" {
		config.Policy = service.Policy
	}
	configDao := dao.NewConfigDao()
	err := configDao.ConfigEdit(service.Id, config)
	if err != nil {
		code = e.ConfigEditError
		zap.S().Error("ConfigEditError" + err.Error())
		log.WaymonLogger.Error("ConfigEditError" + err.Error())
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
