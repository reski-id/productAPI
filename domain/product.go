package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Enable      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type ProductHandler interface {
	InsertProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	GetAllProduct() echo.HandlerFunc
}

type ProductUseCase interface {
	AddProduct(IDUser int, useProduct Product) (Product, error)
	UpProduct(IDProduct int, updateData Product) (Product, error)
	DelProduct(IDProduct int) (bool, error)
	GetSpecificProduct(ProductID int) ([]Product, error)
	GetAllP() ([]Product, error)
}

type ProductData interface {
	Insert(inserProduct Product) Product
	Update(IDProduct int, updatedProduct Product) Product
	Delete(IDProduct int) bool
	GetProductID(ProductID int) []Product
	GetAll() []Product
}
