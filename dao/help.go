package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type HelpDao struct {
}

func NewHelpDao() HelpDao {
	return HelpDao{}
}

func (dao *HelpDao) HelpInfo(helpId int64) (help model.Help, err error) {
	err = internal.DB.Model(&model.Help{}).Where("id = ?", helpId).First(&help).Error
	return
}

func (dao *HelpDao) HelpList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (helps []model.Help, count int64, err error) {
	err = internal.DB.
		Model(&model.Help{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Find(&helps).
		Error
	err = internal.DB.
		Model(&model.Help{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *HelpDao) HelpAdd(help *model.Help) error {
	return internal.DB.Save(&help).Error
}

func (dao *HelpDao) HelpEdit(helpId int64, help *model.Help) error {
	return internal.DB.Model(&model.Help{}).Where("id = ?", helpId).Updates(&help).Error
}

func (dao *HelpDao) HelpStatus(helpId int64, status int) error {
	return internal.DB.Model(&model.Help{}).Where("id = ?", helpId).Update("status", status).Error
}

func (dao *HelpDao) HelpCount() (count int64, err error) {
	err = internal.DB.Model(&model.Help{}).Count(&count).Error
	return
}

func (dao *HelpDao) HelpDelete(helpId int64) error {
	return internal.DB.Model(&model.Help{}).Delete(&model.Help{}, helpId).Error
}
