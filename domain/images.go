package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Images struct {
	ID        int
	Name      string
	File      string
	Enable    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ImagesHandler interface {
	InsertImages() echo.HandlerFunc
	UpdateImages() echo.HandlerFunc
	DeleteImages() echo.HandlerFunc
	GetAllImages() echo.HandlerFunc
}

type ImagesUseCase interface {
	AddImages(IDUser int, useImages Images) (Images, error)
	UpImages(IDImages int, updateData Images) (Images, error)
	DelImages(IDImages int) (bool, error)
	GetSpecificImages(ImagesID int) ([]Images, error)
	GetAllP() ([]Images, error)
}

type ImagesData interface {
	Insert(inserImages Images) Images
	Update(IDImages int, updatedImages Images) Images
	Delete(IDImages int) bool
	GetImagesID(ImagesID int) []Images
	GetAll() []Images
}
