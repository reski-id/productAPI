package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ProductImages struct {
	ID        int
	ProductID int
	ImagesID  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ProductImgHandler interface {
	InsertProductImages() echo.HandlerFunc
	UpdateProductImages() echo.HandlerFunc
	DeleteProductImages() echo.HandlerFunc
	GetAllProductImages() echo.HandlerFunc
	GetProductImagesID() echo.HandlerFunc
}

type ProductImgUseCase interface {
	AddProductImages(newData ProductImages) (ProductImages, error)
	UpProductImages(IDProductImages int, updateData ProductImages) (ProductImages, error)
	DelProductImages(IDProductImages int) (bool, error)
	GetAllD() ([]ProductImages, error)
	GetSpecificProductImages(ProductImagesID int) ([]ProductImages, error)
}

type ProductImgData interface {
	Insert(insertProductImages ProductImages) ProductImages
	Update(IDProductImages int, updatedProductImages ProductImages) ProductImages
	Delete(IDProductImages int) bool
	GetAll() []ProductImages
	GetbyID(ProductImagesID int) []ProductImages
}
