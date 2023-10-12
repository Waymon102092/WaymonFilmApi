package serializer

import (
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
)

type Banner struct {
	Id     int64  `json:"id"`
	Type   int    `json:"type"`
	Img    string `json:"img"`
	Param  string `json:"param"`
	Sort   int    `json:"sort"`
	Time   string `json:"time"`
	Status int    `json:"status"`
}

func BuildBanner(banner model.Banner) Banner {
	return Banner{
		Id:     int64(banner.ID),
		Type:   banner.Type,
		Img:    banner.Img,
		Param:  banner.Param,
		Sort:   banner.Sort,
		Time:   waymon.OnTime(banner.Time),
		Status: banner.Status,
	}
}

func BuildBanners(items []model.Banner) (banners []Banner) {
	for _, item := range items {
		banner := BuildBanner(item)
		banners = append(banners, banner)
	}
	return
}
