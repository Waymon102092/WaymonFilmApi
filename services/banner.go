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
)

type BannerService struct {
	BannerId int64  `json:"banner_id" form:"banner_id"`
	Type     int    `json:"type" form:"type"`
	Img      string `json:"img" form:"img"`
	Param    string `json:"param" form:"param"`
	Sort     int    `json:"sort" form:"sort"`
	Status   int    `json:"status" form:"status"`
	model.BaseLimit
}

func (service *BannerService) BannerList(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	likeCondition := ""
	bannerDao := dao.NewBannerDao()
	banners, count, err := bannerDao.BannerList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.BannerListError
		zap.S().Error("BannerListError" + err.Error())
		log.WaymonLogger.Error("BannerListError" + err.Error())
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
			ItemList: serializer.BuildBanners(banners),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}
