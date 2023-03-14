package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ProductCategory struct {
	ID         int
	ProductID  int
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type ProductCatHandler interface {
	InsertProductCategory() echo.HandlerFunc
	UpdateProductCategory() echo.HandlerFunc
	DeleteProductCategory() echo.HandlerFunc
	GetAllProductCategory() echo.HandlerFunc
	GetProductCategoryID() echo.HandlerFunc
}

type ProductCatUseCase interface {
	AddProductCategory(newData ProductCategory) (ProductCategory, error)
	UpProductCategory(IDProductCategory int, updateData ProductCategory) (ProductCategory, error)
	DelProductCategory(IDProductCategory int) (bool, error)
	GetAllD() ([]ProductCategory, error)
	GetSpecificProductCategory(ProductCategoryID int) ([]ProductCategory, error)
}

type ProductCatData interface {
	Insert(insertProductCategory ProductCategory) ProductCategory
	Update(IDProductCategory int, updatedProductCategory ProductCategory) ProductCategory
	Delete(IDProductCategory int) bool
	GetAll() []ProductCategory
	GetbyID(ProductCategoryID int) []ProductCategory
}
