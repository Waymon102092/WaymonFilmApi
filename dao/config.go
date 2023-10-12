package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type ConfigDao struct {
}

func NewConfigDao() ConfigDao {
	return ConfigDao{}
}

func (dao *ConfigDao) ConfigInfo() (config model.Config, err error) {
	err = internal.DB.Model(&model.Config{}).First(&config).Error
	return
}

func (dao *ConfigDao) ConfigEdit(configId int64, config *model.Config) error {
	return internal.DB.Model(&model.Config{}).Where("id=?", configId).Updates(&config).Error
}
