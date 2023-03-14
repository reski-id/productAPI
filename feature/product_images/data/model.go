package data

import (
	"productapi/domain"

	"gorm.io/gorm"
)

type ProductImages struct {
	gorm.Model
	ProductID int     `gorm:"unique" json:"ProductID" form:"ProductID" validate:"required"`
	ImagesID  int     `gorm:"unique" json:"ImagesID" form:"ImagesID" validate:"required"`
	Product   Product `gorm:"foreignKey:ProductID; references:ID; constraint:OnDelete:CASCADE"`
	Images    Images  `gorm:"foreignKey:ImagesID; references:ID; constraint:OnDelete:CASCADE"`
}
type Product struct {
	gorm.Model
}

type Images struct {
	gorm.Model
}

func (b *ProductImages) ToDomain() domain.ProductImages {
	return domain.ProductImages{
		ID:        int(b.ID),
		ProductID: b.ProductID,
		ImagesID:  b.ImagesID,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func ParseToArr(arr []ProductImages) []domain.ProductImages {
	var res []domain.ProductImages

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.ProductImages) ProductImages {
	var res ProductImages
	res.ID = uint(data.ID)
	res.ProductID = data.ProductID
	res.ImagesID = data.ImagesID
	return res
}
