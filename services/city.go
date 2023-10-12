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

type CitiesResponse struct {
	Hot []string `json:"hot"`
	Cts []Cts    `json:"cts"`
}

type Cts struct {
	Id int    `json:"id"`
	Nm string `json:"nm"`
	Py string `json:"py"`
}

type CityResponse struct {
	Data City `json:"data"`
}

type City struct {
	Detail       string  `json:"detail"`
	ParentArea   int     `json:"parentArea"`
	CityPinyin   string  `json:"cityPinyin"`
	Lng          float64 `json:"lng"`
	IsForeign    bool    `json:"isForeign"`
	DpCityId     int     `json:"dpCityId"`
	Country      string  `json:"country"`
	IsOpen       bool    `json:"isOpen"`
	City         string  `json:"city"`
	Id           int     `json:"id"`
	OpenCityName string  `json:"openCityName"`
	OriginCityID int     `json:"originCityID"`
	Area         int     `json:"area"`
	AreaName     string  `json:"areaName"`
	Province     string  `json:"province"`
	District     string  `json:"district"`
	Lat          float64 `json:"lat"`
}

type CityService struct {
	Lat float64 `json:"lat" form:"lat"` //34.806964
	Lng float64 `json:"lng" form:"lng"` //113.543147
}

func (service *CityService) CityAdd(ctx context.Context) res.Response {
	code := e.Success
	baseUri := viper.GetString("netStart.baseUri")
	city := viper.GetString("netStart.cities")
	result, err := waymon.GetByte(baseUri + city)
	if err != nil {
		fmt.Println(err)
		code = e.GetJsonError
	}
	citiesResponse := CitiesResponse{}
	err = json.Unmarshal(result, &citiesResponse)
	if err != nil {
		fmt.Println(err)
		code = e.UnmarshalError
	}
	if len(citiesResponse.Hot) > 0 {
		wg := new(sync.WaitGroup)
		for _, hot := range citiesResponse.Hot {
			wg.Add(1)
			hotDao := dao.NewHotDao()
			condition := make(map[string]interface{})
			condition["city"] = hot
			hotInfo, _ := hotDao.HotInfo(condition)
			if hotInfo.ID == 0 {
				row := &model.Hot{
					City:   hot,
					Time:   time.Now().Unix(),
					Sort:   1,
					Status: 1,
				}
				err = hotDao.HotAdd(row)
				if err != nil {
					continue
				}
			}
			wg.Done()
		}
		wg.Wait()
	}
	if len(citiesResponse.Cts) > 0 {
		wg := new(sync.WaitGroup)
		for _, cts := range citiesResponse.Cts {
			wg.Add(1)
			cityDao := dao.NewCityDao()
			condition := make(map[string]interface{})
			condition["city_id"] = cts.Id
			cityInfo, _ := cityDao.CityInfo(condition)
			if cityInfo.ID == 0 {
				row := &model.City{
					CityId:     cts.Id,
					CityName:   cts.Nm,
					CityLetter: cts.Py,
					Time:       time.Now().Unix(),
					Status:     1,
				}
				err = cityDao.CityAdd(row)
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
		Data:   nil,
	}
}

func (service *CityService) CityCurrent(ctx context.Context) res.Response {
	code := e.Success
	baseUri := viper.GetString("netStart.baseUri")
	city := viper.GetString("netStart.city")
	uri := fmt.Sprintf("%s%s?lat=%f&lng=%f", baseUri, city, service.Lat, service.Lng)
	result, err := waymon.GetByte(uri)
	if err != nil {
		fmt.Println(err)
		code = e.GetJsonError
	}
	cityResponse := CityResponse{}
	err = json.Unmarshal(result, &cityResponse)
	if err != nil {
		fmt.Println(err)
		code = e.UnmarshalError
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   cityResponse.Data,
	}
}
