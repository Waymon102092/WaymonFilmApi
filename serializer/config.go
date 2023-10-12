package serializer

import (
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
)

type Config struct {
	Id     int64  `json:"id"`
	About  string `json:"about"`
	Proxy  string `json:"proxy"`
	Policy string `json:"policy"`
	Time   string `json:"time"`
	Status int    `json:"status"`
}

func BuildConfig(config model.Config) Config {
	return Config{
		Id:     int64(config.ID),
		About:  config.About,
		Proxy:  config.Proxy,
		Policy: config.Policy,
		Time:   waymon.OnTime(config.Time),
		Status: config.Status,
	}
}
