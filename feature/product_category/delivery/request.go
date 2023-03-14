package delivery

import (
	"productapi/domain"
	"time"
)

type InsertRequest struct {
	ProductID  int `json:"ProductID" form:"ProductID" validate:"required"`
	CategoryID int `json:"CategoryID" form:"CategoryID" validate:"required"`
}

func (ni *InsertRequest) ToDomain() domain.ProductCategory {
	return domain.ProductCategory{
		ProductID:  ni.ProductID,
		CategoryID: ni.CategoryID,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  time.Time{},
	}
}
