package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type CustomDao struct {
}

func NewCustomDao() CustomDao {
	return CustomDao{}
}

func (dao *CustomDao) CustomInfo(condition map[string]interface{}) (custom model.Custom, err error) {
	err = internal.DB.
		Model(&model.Custom{}).
		Where(condition).
		First(&custom).
		Error
	return
}
