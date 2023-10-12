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

type MovieResponse struct {
	MovieList []MovieList `json:"movieList"`
}

type MovieComingResponse struct {
	Coming []MovieList `json:"coming"`
}

type MovieList struct {
	Id              int     `json:"id"`
	HaspromotionTag bool    `json:"haspromotionTag"`
	Img             string  `json:"img"`
	Version         string  `json:"version"`
	Nm              string  `json:"nm"`
	PreShow         bool    `json:"preShow"`
	Sc              float64 `json:"sc"`
	GlobalReleased  bool    `json:"globalReleased"`
	Wish            int     `json:"wish"`
	Star            string  `json:"star"`
	Rt              string  `json:"rt"`
	ShowInfo        string  `json:"showInfo"`
	Showst          int     `json:"showst"`
	Wishst          int     `json:"wishst"`
	ComingTitle     string  `json:"comingTitle"`
}

type MovieService struct {
	CityId int `json:"city_id" form:"city_id"`
	model.BaseLimit
}

func (service *MovieService) MovieHot(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	baseUri := viper.GetString("netStart.baseUri")
	hot := viper.GetString("netStart.hot")
	uri := fmt.Sprintf("%s%s?ci=%d&limit=%d&offset=%d", baseUri, hot, service.CityId, service.Size, service.Page)
	result, err := waymon.GetByte(uri)
	if err != nil {
		code = e.GetJsonError
	}
	movieResponse := MovieResponse{}
	err = json.Unmarshal(result, &movieResponse)
	if err != nil {
		code = e.UnmarshalError
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   movieResponse.MovieList,
	}
}

func (service *MovieService) MovieComing(ctx context.Context) res.Response {
	code := e.Success
	if service.Page == 0 {
		service.Page = 1
	}
	if service.Size == 0 {
		service.Size = 10
	}
	baseUri := viper.GetString("netStart.baseUri")
	come := viper.GetString("netStart.come")
	uri := fmt.Sprintf("%s%s?ci=%d&limit=%d", baseUri, come, service.CityId, service.Size)
	result, err := waymon.GetByte(uri)
	if err != nil {
		code = e.GetJsonError
	}
	movieComingResponse := MovieComingResponse{}
	err = json.Unmarshal(result, &movieComingResponse)
	if err != nil {
		code = e.UnmarshalError
	}
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   movieComingResponse.Coming,
	}
}
