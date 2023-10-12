package serializer

import (
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
)

type Amount struct {
	Id     int64  `json:"id"`
	Type   int    `json:"type"`
	Money  int    `json:"money"`
	Price  int    `json:"price"`
	Time   string `json:"time"`
	Status int    `json:"status"`
}

func BuildAmount(amount model.Amount) Amount {
	return Amount{
		Id:     int64(amount.ID),
		Type:   amount.Type,
		Money:  amount.Money,
		Price:  amount.Price,
		Time:   waymon.OnTime(amount.Time),
		Status: amount.Status,
	}
}

func BuildAmounts(items []model.Amount) (amounts []Amount) {
	for _, item := range items {
		amount := BuildAmount(item)
		amounts = append(amounts, amount)
	}
	return
}
