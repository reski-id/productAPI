package delivery

import (
	"productapi/domain"
)

type InsertRequest struct {
	Name   string `json:"name" form:"name"`
	Enable bool   `json:"enable" form:"enable"`
}

func (ni *InsertRequest) ToDomain() domain.Category {
	return domain.Category{
		Name:   ni.Name,
		Enable: ni.Enable,
	}
}
