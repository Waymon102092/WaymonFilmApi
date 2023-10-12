package serializer

import "Waymon_api/model"

type Help struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func BuildHelp(help model.Help) Help {
	return Help{
		Id:      int64(help.ID),
		Title:   help.Title,
		Content: help.Content,
	}
}

func BuildHelps(items []model.Help) (helps []Help) {
	for _, item := range items {
		help := BuildHelp(item)
		helps = append(helps, help)
	}
	return
}
