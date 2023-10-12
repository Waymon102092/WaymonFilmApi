package dao

import (
	"Waymon_api/internal"
	"Waymon_api/model"
	"sync"
	"time"
)

type ReportDao struct {
}

func NewReportDao() ReportDao {
	return ReportDao{}
}

func (dao *ReportDao) ReportInfo(condition map[string]interface{}) (report model.Report, err error) {
	err = internal.DB.Model(&model.Report{}).Where(condition).First(&report).Error
	return
}

func (dao *ReportDao) ReportList(condition map[string]interface{}, likeCondition string, limit model.BaseLimit) (reports []model.Report, count int64, err error) {
	err = internal.DB.
		Model(&model.Report{}).
		Where(condition).
		Where(likeCondition).
		Offset((limit.Page - 1) * limit.Size).
		Limit(limit.Size).
		Order("sort asc").
		Find(&reports).
		Error
	err = internal.DB.
		Model(&model.Report{}).
		Where(condition).
		Where(likeCondition).
		Count(&count).
		Error
	return
}

func (dao *ReportDao) ReportAdd(report *model.Report, imgs []string) error {
	tx := internal.DB.Begin()
	err := tx.Save(&report).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(imgs) > 0 {
		wg := new(sync.WaitGroup)
		wg.Add(len(imgs))
		for _, img := range imgs {
			reportImg := &model.ReportImg{
				ReportId: int64(report.ID),
				Img:      img,
				Time:     time.Now().Unix(),
				Status:   1,
			}
			err = tx.Save(&reportImg).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			wg.Done()
		}
		wg.Wait()
	}
	tx.Commit()
	return nil
}

func (dao *ReportDao) ReportEdit(reportId int64, report *model.Report) error {
	return internal.DB.Model(&model.Report{}).Where("id=?", reportId).Updates(&report).Error
}

func (dao *ReportDao) ReportStatus(reportId int64, status int) error {
	return internal.DB.Model(&model.Report{}).Where("id=?", reportId).Update("status", status).Error
}

func (dao *ReportDao) ReportCount() (count int64, err error) {
	err = internal.DB.Model(&model.Report{}).Count(&count).Error
	return
}

func (dao *ReportDao) ReportDelete(reportId int64) error {
	return internal.DB.Delete(&model.Report{}, reportId).Error
}
