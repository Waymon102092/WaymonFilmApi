package serializer

import (
	"Waymon_api/dao"
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
)

type Report struct {
	Id             int    `json:"id"`
	AvatarUrl      string `json:"avatar_url"`
	NickName       string `json:"nick_name"`
	Tel            string `json:"tel"`
	ReportCategory string `json:"report_category"`
	Content        string `json:"content"`
	Sort           int    `json:"sort"`
	Time           string `json:"time"`
	Status         int    `json:"status"`
}

func BuildReport(report model.Report, member *model.Member, reportCategory model.ReportCategory) Report {
	return Report{
		Id:             int(report.ID),
		AvatarUrl:      member.AvatarUrl,
		NickName:       member.NickName,
		Tel:            member.Tel,
		ReportCategory: reportCategory.Title,
		Content:        report.Content,
		Sort:           report.Sort,
		Time:           waymon.OnTime(report.Time),
		Status:         report.Status,
	}
}

func BuildReports(items []model.Report) (reports []Report) {
	for _, item := range items {
		condition := make(map[string]interface{})
		condition["id"] = item.MemberId
		memberDao := dao.NewMemberDao()
		member, _ := memberDao.MemberInfo(condition)

		likeCondition := make(map[string]interface{})
		likeCondition["id"] = item.ReportCategoryId
		reportCategoryDao := dao.NewReportCategoryDao()
		reportCategory, _ := reportCategoryDao.ReportCategoryInfo(likeCondition)

		report := BuildReport(item, &member, reportCategory)
		reports = append(reports, report)
	}
	return
}
