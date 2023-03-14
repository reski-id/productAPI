package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Category struct {
	ID        int
	Name      string
	Enable    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CategoryHandler interface {
	InsertCategory() echo.HandlerFunc
	UpdateCategory() echo.HandlerFunc
	DeleteCategory() echo.HandlerFunc
	GetAllCategory() echo.HandlerFunc
}

type CategoryUseCase interface {
	AddCategory(IDUser int, useCategory Category) (Category, error)
	UpCategory(IDCategory int, updateData Category) (Category, error)
	DelCategory(IDCategory int) (bool, error)
	GetSpecificCategory(CategoryID int) ([]Category, error)
	GetAllP() ([]Category, error)
}

type CategoryData interface {
	Insert(inserCategory Category) Category
	Update(IDCategory int, updatedCategory Category) Category
	Delete(IDCategory int) bool
	GetCategoryID(CategoryID int) []Category
	GetAll() []Category
}
