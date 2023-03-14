package data

import (
	"productapi/domain"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Enable bool   `json:"enable" form:"enable"`
}

func (b *Category) ToDomain() domain.Category {
	return domain.Category{
		ID:        int(b.ID),
		Name:      b.Name,
		Enable:    b.Enable,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func ParseToArr(arr []Category) []domain.Category {
	var res []domain.Category

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Category) Category {
	var res Category
	res.ID = uint(data.ID)
	res.Name = data.Name
	res.Enable = data.Enable
	return res
}
