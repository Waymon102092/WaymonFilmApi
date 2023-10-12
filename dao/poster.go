package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type PosterDao struct {
}

func NewPosterDao() PosterDao {
	return PosterDao{}
}

func (dao *PosterDao) PosterList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (posters []model.Poster, count int64, err error) {
	err = internal.DB.
		Model(&model.Poster{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("sort asc").
		Find(&posters).
		Error
	err = internal.DB.
		Model(&model.Poster{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}
