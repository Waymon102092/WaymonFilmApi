package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type CityDao struct {
}

func NewCityDao() CityDao {
	return CityDao{}
}

func (dao *CityDao) CityInfo(condition map[string]interface{}) (city *model.City, err error) {
	err = internal.DB.Model(&model.City{}).Where(condition).First(&city).Error
	return
}

func (dao *CityDao) CityAdd(city *model.City) error {
	return internal.DB.Save(&city).Error
}
