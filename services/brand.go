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
	"time"
)

type BrandService struct {
	BrandId int64  `json:"brand_id" form:"brand_id"`
	Title   string `json:"title" form:"title"`
	Sort    int    `json:"sort" form:"sort"`
	Status  int    `json:"status" form:"status"`
	model.BaseLimit
}

func (service *BrandService) BrandInfo(ctx context.Context) res.Response {
	code := e.Success
	condition := make(map[string]interface{})
	if service.BrandId > 0 {
		condition["id"] = service.BrandId
	}
	brandDao := dao.NewBrandDao()
	brand, err := brandDao.BrandInfo(condition)
	if err != nil {
		code = e.BrandInfoError
		zap.S().Error("BrandInfoError" + err.Error())
		log.WaymonLogger.Error("BrandInfoError" + err.Error())
		return res.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   nil,
		}
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildBrand(brand),
	}
}

func (service *BrandService) BrandList(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	condition := make(map[string]interface{})
	if service.BrandId > 0 {
		condition["brand_id"] = service.BrandId
	}
	likeCondition := ""
	brandDao := dao.NewBrandDao()
	brands, count, err := brandDao.BrandList(condition, likeCondition, service.BaseLimit)
	if err != nil {
		code = e.BrandListError
		zap.S().Error("BrandListError" + err.Error())
		log.WaymonLogger.Error("BrandListError" + err.Error())
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
			ItemList: serializer.BuildBrands(brands),
			Total:    waymon.PageCount(count, service.Size),
		},
	}
}

func (service *BrandService) BrandAdd(ctx context.Context) res.Response {
	code := e.Success
	brand := &model.Brand{
		Title:  service.Title,
		Sort:   service.Sort,
		Time:   time.Now().Unix(),
		Status: service.Status,
	}
	brandDao := dao.NewBrandDao()
	err := brandDao.BrandAdd(brand)
	if err != nil {
		code = e.BrandAddError
		zap.S().Error("BrandAddError" + err.Error())
		log.WaymonLogger.Error("BrandAddError" + err.Error())
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

func (service *BrandService) BrandEdit(ctx context.Context) res.Response {
	code := e.Success
	brand := &model.Brand{}
	if service.Title != "" {
		brand.Title = service.Title
	}
	if service.Sort > 0 {
		brand.Sort = service.Sort
	}
	if service.Status > -1 {
		brand.Status = service.Status
	}
	brandDao := dao.NewBrandDao()
	err := brandDao.BrandEdit(service.BrandId, brand)
	if err != nil {
		code = e.BrandEditError
		zap.S().Error("BrandEditError" + err.Error())
		log.WaymonLogger.Error("BrandEditError" + err.Error())
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

func (service *BrandService) BrandStatus(ctx context.Context) res.Response {
	code := e.Success
	brandDao := dao.NewBrandDao()
	err := brandDao.BrandStatus(service.BrandId, service.Status)
	if err != nil {
		code = e.BrandStatusError
		zap.S().Error("BrandStatusError" + err.Error())
		log.WaymonLogger.Error("BrandStatusError" + err.Error())
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

func (service *BrandService) BrandCount(ctx context.Context) res.Response {
	code := e.Success
	brandDao := dao.NewBrandDao()
	count, err := brandDao.BrandCount(1)
	if err != nil {
		code = e.BrandCountError
		zap.S().Error("BrandCountError" + err.Error())
		log.WaymonLogger.Error("BrandCountError" + err.Error())
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

func (service *BrandService) BrandDelete(ctx context.Context) res.Response {
	code := e.Success
	brandDao := dao.NewBrandDao()
	err := brandDao.BrandDelete(service.BrandId)
	if err != nil {
		code = e.BrandDeleteError
		zap.S().Error("BrandDeleteError" + err.Error())
		log.WaymonLogger.Error("BrandDeleteError" + err.Error())
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
