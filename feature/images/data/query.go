package data

import (
	"fmt"
	"log"
	"productapi/domain"

	"gorm.io/gorm"
)

type imagesData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ImagesData {
	return &imagesData{
		db: db,
	}
}

func (nd *imagesData) Insert(newData domain.Images) domain.Images {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Images{}
	}
	return cnv.ToDomain()
}

func (bd *imagesData) Update(imagesID int, updatedData domain.Images) domain.Images {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", imagesID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Images{}
	}
	cnv.ID = uint(imagesID)
	return cnv.ToDomain()
}

func (nd *imagesData) Delete(imagesID int) bool {
	err := nd.db.Where("ID = ?", imagesID).Delete(&Images{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

func (nd *imagesData) GetAll() []domain.Images {
	var data []Images

	err := nd.db.Where("enable != ?", "false").Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *imagesData) GetImagesID(imagesID int) []domain.Images {
	var data []Images
	err := nd.db.Where("ID = ?", imagesID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
