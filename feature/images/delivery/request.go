package delivery

import (
	"productapi/domain"
)

type InsertRequest struct {
	Name   string `json:"name" form:"name"`
	File   string `json:"file" form:"file"`
	Enable bool   `json:"enable" form:"enable"`
}

func (ni *InsertRequest) ToDomain() domain.Images {
	return domain.Images{
		Name:   ni.Name,
		File:   ni.File,
		Enable: ni.Enable,
	}
}
