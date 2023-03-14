package data

import (
	"productapi/domain"

	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	ProductID  int      `gorm:"unique" json:"ProductID" form:"ProductID" validate:"required"`
	CategoryID int      `gorm:"unique" json:"CategoryID" form:"CategoryID" validate:"required"`
	Product    Product  `gorm:"foreignKey:ProductID; references:ID; constraint:OnDelete:CASCADE"`
	Category   Category `gorm:"foreignKey:CategoryID; references:ID; constraint:OnDelete:CASCADE"`
}
type Product struct {
	gorm.Model
}

type Category struct {
	gorm.Model
}

func (b *ProductCategory) ToDomain() domain.ProductCategory {
	return domain.ProductCategory{
		ID:         int(b.ID),
		ProductID:  b.ProductID,
		CategoryID: b.CategoryID,
		CreatedAt:  b.CreatedAt,
		UpdatedAt:  b.UpdatedAt,
	}
}

func ParseToArr(arr []ProductCategory) []domain.ProductCategory {
	var res []domain.ProductCategory

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.ProductCategory) ProductCategory {
	var res ProductCategory
	res.ID = uint(data.ID)
	res.ProductID = data.ProductID
	res.CategoryID = data.CategoryID
	return res
}
