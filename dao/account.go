package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type AccountDao struct {
}

func NewAccountDao() AccountDao {
	return AccountDao{}
}

func (dao *AccountDao) AccountInfo(condition map[string]interface{}) (account *model.Account, err error) {
	err = internal.DB.Model(&model.Account{}).Where(condition).First(&account).Error
	return
}

func (dao *AccountDao) AccountAdd(account *model.Account) error {
	return internal.DB.Save(&account).Error
}
