package serializer

import (
	"Waymon_api/dao"
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
	"fmt"
	"strconv"
)

type Fans struct {
	Id        int64  `json:"id"`
	Tel       string `json:"tel"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Time      string `json:"time"`
}

func BuildFan(member *model.Member) Fans {
	tel := ""
	if member.Tel != "" {
		tel = waymon.WaymonTel(member.Tel)
	}
	return Fans{
		Id:        int64(member.ID),
		Tel:       tel,
		NickName:  member.NickName,
		AvatarUrl: member.AvatarUrl,
		Time:      waymon.OnTimeMonthDay(member.Time),
	}
}

func BuildFans(items []*model.Member) (fans []Fans) {
	for _, item := range items {
		fan := BuildFan(item)
		fans = append(fans, fan)
	}
	return
}

type FansOrder struct {
	Id         int64   `json:"id"`
	FilmName   string  `json:"film_name"`
	Cover      string  `json:"cover"`
	CinemaName string  `json:"cinema_name"`
	NickName   string  `json:"nick_name"`
	AvatarUrl  string  `json:"avatar_url"`
	TradeNo    string  `json:"trade_no"`
	Seats      string  `json:"seats"`
	Money      int     `json:"money"`
	Income     float64 `json:"income"`
	Num        int     `json:"num"`
	Time       string  `json:"time"`
}

func BuildFansOrder(order dao.FansOrder) FansOrder {
	configDao := dao.NewConfigDao()
	config, _ := configDao.ConfigInfo()
	price, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(order.Money/100)*config.AgentPre), 64)
	return FansOrder{
		FilmName:   order.FilmName,
		CinemaName: order.CinemaName,
		NickName:   order.NickName,
		AvatarUrl:  order.AvatarUrl,
		TradeNo:    order.TradeNo,
		Money:      order.Money,
		Income:     price,
		Num:        order.Num,
		Time:       waymon.OnTimeDay(order.Time),
	}
}

func BuildFansOrders(items []dao.FansOrder) (fansOrders []FansOrder) {
	for _, item := range items {
		fansOrder := BuildFansOrder(item)
		fansOrders = append(fansOrders, fansOrder)
	}
	return
}
