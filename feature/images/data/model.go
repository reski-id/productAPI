package data

import (
	"productapi/domain"

	"gorm.io/gorm"
)

type Images struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	File   string `json:"file" form:"file"`
	Enable bool   `json:"enable" form:"enable"`
}

func (b *Images) ToDomain() domain.Images {
	return domain.Images{
		ID:        int(b.ID),
		Name:      b.Name,
		File:      b.File,
		Enable:    b.Enable,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func ParseToArr(arr []Images) []domain.Images {
	var res []domain.Images

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Images) Images {
	var res Images
	res.ID = uint(data.ID)
	res.Name = data.Name
	res.File = data.File
	res.Enable = data.Enable
	return res
}
