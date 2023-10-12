package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type FansOrder struct {
	FilmId     string `json:"film_id"`
	FilmName   string `json:"film_name"`
	CinemaCode int64  `json:"cinema_code"`
	CinemaName string `json:"cinema_name"`
	NickName   string `json:"nick_name"`
	AvatarUrl  string `json:"avatar_url"`
	TradeNo    string `json:"trade_no"`
	Money      int    `json:"money"`
	Num        int    `json:"num"`
	Time       int64  `json:"time"`
}

type FansDao struct {
}

func NewFansDao() FansDao {
	return FansDao{}
}

func (dao *FansDao) FansList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (fans []*model.Member, count int64, err error) {
	err = internal.DB.
		Model(&model.Member{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("time desc").
		Find(&fans).
		Error
	err = internal.DB.
		Model(&model.Member{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *FansDao) FansCount(condition map[string]interface{}, starTime, endTime int64) (count int64, err error) {
	err = internal.DB.
		Model(&model.Member{}).
		Where(condition).
		Where("time > ? and time < ?", starTime, endTime).
		Count(&count).
		Error
	return
}

func (dao *FansDao) FansOrder(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (order []FansOrder, count int64, err error) {
	err = internal.DB.
		Select("order.money,order.trade_no,order.num,order.film_id,order.film_name,order.cinema_code,order.cinema_name,order.time,member.nick_name,member.avatar_url").
		Model(&model.Order{}).
		Joins("left join member on order.member_id = member.id").
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("order.time desc").
		Find(&order).
		Error
	err = internal.DB.
		Select("order.money,order.trade_no,order.num,order.film_id,order.film_name,order.cinema_code,order.cinema_name,order.time,member.nick_name,member.avatar_url").
		Model(&model.Order{}).
		Joins("left join member on order.member_id = member.id").
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *FansDao) FansOrderCount(condition map[string]interface{}, likeCondition string) (count int64, err error) {
	err = internal.DB.
		Model(&model.Order{}).
		Joins("left join member on order.member_id = member.id").
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}
