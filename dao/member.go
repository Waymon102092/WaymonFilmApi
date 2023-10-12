package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
)

type MemberDao struct {
}

func NewMemberDao() MemberDao {
	return MemberDao{}
}

func (dao *MemberDao) MemberInfo(condition map[string]interface{}) (member model.Member, err error) {
	err = internal.DB.
		Model(&model.Member{}).
		Where(condition).
		First(&member).
		Error
	return
}

func (dao *MemberDao) MemberList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (members []*model.Member, count int64, err error) {
	err = internal.DB.
		Model(&model.Member{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("id asc").
		Find(&members).
		Error
	err = internal.DB.
		Model(&model.Member{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *MemberDao) MemberAdd(member *model.Member) error {
	return internal.DB.Save(&member).Error
}

func (dao *MemberDao) MemberEdit(memberId int64, member *model.Member) error {
	return internal.DB.Model(&model.Member{}).Where("id= ?", memberId).Updates(&member).Error
}

func (dao *MemberDao) MemberStatus(memberId int64, status int) error {
	return internal.DB.Model(&model.Member{}).Where("id= ?", memberId).Update("status", status).Error
}

func (dao *MemberDao) MemberStaff(staffId int64) (member *model.Member, err error) {
	err = internal.DB.
		Model(&model.Member{}).
		Where("id = ?", staffId).
		First(&member).
		Error
	return
}
