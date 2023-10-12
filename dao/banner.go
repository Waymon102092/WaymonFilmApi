package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type BannerDao struct {
}

func NewBannerDao() BannerDao {
	return BannerDao{}
}

func (dao *BannerDao) BannerList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (banners []model.Banner, count int64, err error) {
	err = internal.DB.
		Model(&model.Banner{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("sort asc").
		Find(&banners).
		Error
	err = internal.DB.
		Model(&model.Banner{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}
