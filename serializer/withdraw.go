package serializer

import (
	"Waymon_api/model"
	"Waymon_api/pkg/waymon"
)

type Withdraw struct {
	Id     int64  `json:"id"`
	Money  int    `json:"money"`
	Price  int    `json:"price"`
	Time   string `json:"time"`
	Status int    `json:"status"`
}

func BuildWithdraw(amount model.Withdraw) Withdraw {
	return Withdraw{
		Id:     int64(amount.ID),
		Money:  amount.Money,
		Price:  amount.Price,
		Time:   waymon.OnTime(amount.Time),
		Status: amount.Status,
	}
}

func BuildWithdraws(items []model.Withdraw) (withdraws []Withdraw) {
	for _, item := range items {
		withdraw := BuildWithdraw(item)
		withdraws = append(withdraws, withdraw)
	}
	return
}
