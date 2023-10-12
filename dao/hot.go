package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type HotDao struct {
}

func NewHotDao() HotDao {
	return HotDao{}
}

func (dao *HotDao) HotInfo(condition map[string]interface{}) (hot *model.Hot, err error) {
	err = internal.DB.Model(&model.Hot{}).Where(condition).First(&hot).Error
	return
}

func (dao *HotDao) HotAdd(hot *model.Hot) error {
	return internal.DB.Save(&hot).Error
}
