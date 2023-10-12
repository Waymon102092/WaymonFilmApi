package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
	"gorm.io/gorm"
)

type WithdrawDao struct {
}

func NewWithdrawDao() WithdrawDao {
	return WithdrawDao{}
}

func (dao *WithdrawDao) WithdrawInfo(condition map[string]interface{}) (withdraw model.Withdraw, err error) {
	err = internal.DB.Model(&model.Withdraw{}).Where(condition).First(&withdraw).Error
	return
}

func (dao *WithdrawDao) WithdrawList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (withdraws []model.Withdraw, count int64, err error) {
	err = internal.DB.
		Model(&model.Withdraw{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Find(&withdraws).
		Error
	err = internal.DB.
		Model(&model.Withdraw{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *WithdrawDao) WithdrawAdd(withdraw *model.Withdraw) error {
	tx := internal.DB.Begin()
	//更新余额
	err := tx.Model(&model.Member{}).Where("id=?", withdraw.MemberId).Update("money", gorm.Expr("money- ?", withdraw.Money)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//
	member := model.Member{}
	err = tx.Model(&model.Member{}).Where("id=?", withdraw.MemberId).First(&member).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	withdraw.Price = member.Money
	err = tx.Save(&withdraw).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (dao *WithdrawDao) WithdrawEdit(withdrawId int64, withdraw *model.Withdraw) error {
	return internal.DB.Model(&model.Withdraw{}).Where("id=?", withdrawId).Updates(&withdraw).Error
}

func (dao *WithdrawDao) WithdrawStatus(withdrawId int64, status int) error {
	return internal.DB.Model(&model.Withdraw{}).Where("id=?", withdrawId).Update("status", status).Error
}

func (dao *WithdrawDao) WithdrawDelete(withdrawId int64) error {
	return internal.DB.Model(&model.Withdraw{}).Delete(&model.Withdraw{}, withdrawId).Error
}

func (dao *WithdrawDao) WithdrawMoney(condition map[string]interface{}, likeCondition string) (amount int, err error) {
	err = internal.DB.
		Model(&model.Withdraw{}).
		Where(condition).
		Where(likeCondition).
		Pluck("COALESCE(SUM(money), 0) as amount", &amount).
		Error
	return
}
