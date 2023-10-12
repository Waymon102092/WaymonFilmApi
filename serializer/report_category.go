package serializer

import (
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
)

type ReportCategory struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Sort   int    `json:"sort"`
	Time   string `json:"time"`
	Status int    `json:"status"`
}

func BuildReportCategory(reportCategory model.ReportCategory) ReportCategory {
	return ReportCategory{
		Id:     int64(reportCategory.ID),
		Title:  reportCategory.Title,
		Sort:   reportCategory.Sort,
		Time:   waymon.OnTime(reportCategory.Time),
		Status: reportCategory.Status,
	}
}

func BuildReportCategories(items []model.ReportCategory) (reportCategories []ReportCategory) {
	for _, item := range items {
		reportCategory := BuildReportCategory(item)
		reportCategories = append(reportCategories, reportCategory)
	}
	return
}
