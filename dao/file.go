package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type FileDao struct {
}

func NewFileDao() FileDao {
	return FileDao{}
}

func (dao *FileDao) FileList(condition map[string]interface{}, limit model.BaseLimit) (files []model.File, count int64, err error) {
	err = internal.DB.
		Model(&model.File{}).
		Where(condition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Find(&files).
		Error
	err = internal.DB.
		Model(&model.File{}).
		Where(condition).
		Count(&count).
		Error
	return
}

func (dao *FileDao) FileAdd(file *model.File) error {
	return internal.DB.Save(&file).Error
}

func (dao *FileDao) FileStatus(fileId int64, status int) error {
	return internal.DB.Model(&model.File{}).Where("id=?", fileId).Update("status", status).Error
}

func (dao *FileDao) FileDelete(fileId int64) error {
	return internal.DB.Delete(&model.File{}, fileId).Error
}
