package delivery

import (
	"productapi/domain"
)

type InsertRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Enable      bool   `json:"enable" form:"enable"`
}

func (ni *InsertRequest) ToDomain() domain.Product {
	return domain.Product{
		Name:        ni.Name,
		Description: ni.Description,
		Enable:      ni.Enable,
	}
}
