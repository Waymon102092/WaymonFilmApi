package services

import (
	"Waymon_api/model"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/res"
	"Waymon_api/pkg/waymon"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type CinemaResponse struct {
	CinemaId int    `json:"cinemaId"`
	Title    string `json:"title"`
	Price    struct {
		N string `json:"n"`
		Q string `json:"q"`
	} `json:"price"`
	Location string `json:"location"`
	Distance string `json:"distance"`
}

type CinemaService struct {
	Day        string  `json:"day" form:"day"`
	BrandId    int     `json:"brand_id" form:"brand_id"`
	CityId     int     `json:"city_id" form:"city_id"`
	DistrictId int     `json:"district_id" form:"district_id"`
	Lat        float64 `json:"lat" form:"lat"`
	Lng        float64 `json:"lng" form:"lng"`
	model.BaseLimit
}

func (service *CinemaService) CinemaList(ctx context.Context) res.Response {
	code := e.Success
	baseUri := viper.GetString("netStart.baseUri")
	cinema := viper.GetString("netStart.cinema")
	uri := fmt.Sprintf("%s%s?cityId=%d&lat=%f&lng=%f", baseUri, cinema, service.CityId, service.Lat, service.Lng)
	result, err := waymon.GetByte(uri)
	if err != nil {
		code = e.GetJsonError
	}
	cinemaResponse := CinemaResponse{}
	err = json.Unmarshal(result, &cinemaResponse)
	if err != nil {
		code = e.UnmarshalError
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   cinemaResponse,
	}
}
