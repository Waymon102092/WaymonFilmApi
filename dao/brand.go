package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type BrandDao struct {
}

func NewBrandDao() BrandDao {
	return BrandDao{}
}

func (dao *BrandDao) BrandInfo(condition map[string]interface{}) (brand model.Brand, err error) {
	err = internal.DB.Model(&model.Brand{}).Where(condition).First(&brand).Error
	return
}

func (dao *BrandDao) BrandList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (brands []model.Brand, count int64, err error) {
	err = internal.DB.
		Model(&model.Brand{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("sort asc").
		Find(&brands).
		Error
	err = internal.DB.
		Model(&model.Brand{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *BrandDao) BrandAdd(brand *model.Brand) error {
	return internal.DB.Save(&brand).Error
}

func (dao *BrandDao) BrandEdit(brandId int64, brand *model.Brand) error {
	return internal.DB.Model(&model.Brand{}).Where("id=?", brandId).Updates(&brand).Error
}

func (dao *BrandDao) BrandStatus(brandId int64, status int) error {
	return internal.DB.Model(&model.Brand{}).Where("id=?", brandId).Update("status", status).Error
}

func (dao *BrandDao) BrandCount(aType int) (count int64, err error) {
	err = internal.DB.Model(&model.Brand{}).Where("type=?", aType).Count(&count).Error
	return
}

func (dao *BrandDao) BrandDelete(brandId int64) error {
	return internal.DB.Delete(&model.Brand{}, brandId).Error
}
