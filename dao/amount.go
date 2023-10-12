package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type AmountDao struct {
}

func NewAmountDao() AmountDao {
	return AmountDao{}
}

func (dao *AmountDao) AmountList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (amounts []model.Amount, count int64, err error) {
	err = internal.DB.
		Model(&model.Amount{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Find(&amounts).
		Error
	err = internal.DB.
		Model(&model.Amount{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *AmountDao) AmountMoney(condition map[string]interface{}, likeCondition string) (amount int, err error) {
	err = internal.DB.
		Model(&model.Amount{}).
		Where(condition).
		Where(likeCondition).
		Pluck("COALESCE(SUM(money), 0) as amount", &amount).
		Error
	return
}

func (dao *AmountDao) AmountSettle(memberId int64) (amount int, err error) {
	err = internal.DB.
		Model(&model.Amount{}).
		Where("member_id = ? and status = 0", memberId).
		Pluck("COALESCE(SUM(money), 0) as amount", &amount).
		Error
	return
}

func (dao *AmountDao) AmountAccumulate(memberId int64) (amount int, err error) {
	err = internal.DB.
		Model(&model.Amount{}).
		Where("member_id = ? and status = 1", memberId).
		Pluck("COALESCE(SUM(money), 0) as amount", &amount).
		Error
	return
}
