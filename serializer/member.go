package serializer

import (
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
)

type User struct {
	Code      int64  `json:"code"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Tel       string `json:"tel"`
	Money     int    `json:"money"`
	Status    int    `json:"status"`
}

func BuildUser(member *model.Member) User {
	tel := ""
	if member.Tel != "" {
		tel = waymon.WaymonTel(member.Tel)
	}
	return User{
		Code:      int64(member.ID),
		Tel:       tel,
		NickName:  member.NickName,
		AvatarUrl: member.AvatarUrl,
		Money:     member.Money,
		Status:    member.Status,
	}
}
