package services

import (
	"Waymon_api/dao"
	"Waymon_api/model"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/log"
	"Waymon_api/pkg/res"
	"Waymon_api/pkg/waymon"
	"Waymon_api/pkg/wechat/mini"
	"Waymon_api/serializer"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type PosterService struct {
	model.BaseLimit
}

func (service *PosterService) PosterList(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	condition["status"] = 1
	posterDao := dao.NewPosterDao()
	posters, count, err := posterDao.PosterList(condition, "", service.BaseLimit)
	if err != nil {
		code = e.PosterListError
		zap.S().Error("PosterListError" + err.Error())
		log.WaymonLogger.Error("PosterListError" + err.Error())
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
			ItemList: serializer.BuildPosters(posters),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *PosterService) PosterCode(ctx context.Context, memberId int64) res.Response {
	code := e.Success
	page := "pages/index/index"
	scene := fmt.Sprintf("%d", memberId)
	img, err := mini.CreateCode(page, scene)
	if err != nil {
		code = e.PosterError
		zap.S().Error("MemberInfoError" + err.Error())
		log.WaymonLogger.Error("MemberInfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   img,
	}
}
