package data

import (
	"fmt"
	"log"
	"productapi/domain"

	"gorm.io/gorm"
)

type productImgData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductImgData {
	return &productImgData{
		db: db,
	}
}

func (nd *productImgData) Insert(newData domain.ProductImages) domain.ProductImages {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.ProductImages{}
	}
	return cnv.ToDomain()
}

func (bd *productImgData) Update(dataID int, updatedData domain.ProductImages) domain.ProductImages {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", dataID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.ProductImages{}
	}
	cnv.ID = uint(dataID)
	return cnv.ToDomain()
}

func (nd *productImgData) Delete(dataID int) bool {
	err := nd.db.Where("ID = ?", dataID).Delete(&ProductImages{})
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

func (nd *productImgData) GetAll() []domain.ProductImages {
	var data []ProductImages
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *productImgData) GetbyID(dataID int) []domain.ProductImages {
	var data []ProductImages
	err := nd.db.Where("ID = ?", dataID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
