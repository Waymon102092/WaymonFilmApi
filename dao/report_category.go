package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type ReportCategoryDao struct {
}

func NewReportCategoryDao() ReportCategoryDao {
	return ReportCategoryDao{}
}

func (dao *ReportCategoryDao) ReportCategoryInfo(condition map[string]interface{}) (reportCategory model.ReportCategory, err error) {
	err = internal.DB.Model(&model.ReportCategory{}).Where(condition).First(&reportCategory).Error
	return
}

func (dao *ReportCategoryDao) ReportCategoryList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (reportCategories []model.ReportCategory, count int64, err error) {
	err = internal.DB.
		Model(&model.ReportCategory{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("sort asc").
		Find(&reportCategories).
		Error
	err = internal.DB.
		Model(&model.ReportCategory{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *ReportCategoryDao) ReportCategoryAdd(reportCategory *model.ReportCategory) error {
	return internal.DB.Save(&reportCategory).Error
}

func (dao *ReportCategoryDao) ReportCategoryEdit(reportCategoryId int64, reportCategory *model.ReportCategory) error {
	return internal.DB.Model(&model.ReportCategory{}).Where("id=?", reportCategoryId).Updates(&reportCategory).Error
}

func (dao *ReportCategoryDao) ReportCategoryStatus(reportCategoryId int64, status int) error {
	return internal.DB.Model(&model.ReportCategory{}).Where("id=?", reportCategoryId).Update("status", status).Error
}

func (dao *ReportCategoryDao) ReportCategoryCount() (count int64, err error) {
	err = internal.DB.Model(&model.ReportCategory{}).Count(&count).Error
	return
}

func (dao *ReportCategoryDao) ReportCategoryDelete(reportCategoryId int64) error {
	return internal.DB.Delete(&model.ReportCategory{}, reportCategoryId).Error
}
