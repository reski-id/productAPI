package delivery

import (
	"productapi/domain"
	"time"
)

type InsertRequest struct {
	ProductID int `json:"ProductID" form:"ProductID" validate:"required"`
	ImagesID  int `json:"ImagesID" form:"ImagesID" validate:"required"`
}

func (ni *InsertRequest) ToDomain() domain.ProductImages {
	return domain.ProductImages{
		ProductID: ni.ProductID,
		ImagesID:  ni.ImagesID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
}
