package services

import (
	"Waymon_api/dao"
	"Waymon_api/model"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/res"
	"Waymon_api/pkg/waymon"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type DistrictResponse struct {
	District DistrictSubResponse `json:"district"`
}

type DistrictSubResponse struct {
	Name     string     `json:"name"`
	SubItems []District `json:"subItems"`
}

type District struct {
	Count int    `json:"count"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
}

type DistrictService struct {
	CityId int `json:"city_id" form:"city_id"`
}

func (service *DistrictService) DistrictFilter(ctx context.Context) res.Response {
	code := e.Success
	baseUri := viper.GetString("netStart.baseUri")
	filter := viper.GetString("netStart.filter_cinema")
	uri := fmt.Sprintf("%s%s?ci=%d", baseUri, filter, service.CityId)
	result, err := waymon.GetByte(uri)
	if err != nil {
		code = e.GetJsonError
	}
	districtResponse := DistrictResponse{}
	err = json.Unmarshal(result, &districtResponse)
	if err != nil {
		code = e.UnmarshalError
	}
	if len(districtResponse.District.SubItems) > 0 {
		wg := new(sync.WaitGroup)
		for _, item := range districtResponse.District.SubItems {
			wg.Add(1)
			districtDao := dao.NewDistrictDao()
			condition := make(map[string]interface{})
			condition["district_id"] = item.Id
			districtInfo, _ := districtDao.DistrictInfo(condition)
			if districtInfo.ID == 0 {
				district := &model.District{
					DistrictId:   item.Id,
					DistrictName: item.Name,
					Time:         time.Now().Unix(),
					Status:       1,
				}
				err = districtDao.DistrictAdd(district)
				if err != nil {
					continue
				}
			}
			wg.Done()
		}
		wg.Wait()
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   districtResponse.District.SubItems,
	}
}
