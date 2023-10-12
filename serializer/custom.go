package serializer

import (
	"Waymon_api/model"
)

type Custom struct {
	Id  int64  `json:"id"`
	Img string `json:"img"`
}

func BuildCustom(custom model.Custom) Custom {
	return Custom{
		Id:  int64(custom.ID),
		Img: custom.Img,
	}
}
