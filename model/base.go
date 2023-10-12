package model

type BaseLimit struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}
