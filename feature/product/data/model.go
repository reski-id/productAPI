package data

import (
	"productapi/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name"  validate:"required"`
	Description string `json:"description" form:"description"  validate:"required"`
	Enable      bool   `json:"enable" form:"enable"  validate:"required"`
}

func (b *Product) ToDomain() domain.Product {
	return domain.Product{
		ID:          int(b.ID),
		Name:        b.Name,
		Description: b.Description,
		Enable:      b.Enable,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

func ParseToArr(arr []Product) []domain.Product {
	var res []domain.Product

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Product) Product {
	var res Product
	res.ID = uint(data.ID)
	res.Name = data.Name
	res.Description = data.Description
	res.Enable = data.Enable
	return res
}
