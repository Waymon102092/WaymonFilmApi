package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type DistrictDao struct {
}

func NewDistrictDao() DistrictDao {
	return DistrictDao{}
}

func (dao *DistrictDao) DistrictInfo(condition map[string]interface{}) (district *model.District, err error) {
	err = internal.DB.Model(&model.District{}).Where(condition).First(&district).Error
	return
}

func (dao *DistrictDao) DistrictAdd(district *model.District) error {
	return internal.DB.Save(&district).Error
}
