package data

import (
	"fmt"
	"log"
	"productapi/domain"

	"gorm.io/gorm"
)

type productCatData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductCatData {
	return &productCatData{
		db: db,
	}
}

func (nd *productCatData) Insert(newData domain.ProductCategory) domain.ProductCategory {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.ProductCategory{}
	}
	return cnv.ToDomain()
}

func (bd *productCatData) Update(dataID int, updatedData domain.ProductCategory) domain.ProductCategory {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", dataID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.ProductCategory{}
	}
	cnv.ID = uint(dataID)
	return cnv.ToDomain()
}

func (nd *productCatData) Delete(dataID int) bool {
	err := nd.db.Where("ID = ?", dataID).Delete(&ProductCategory{})
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

func (nd *productCatData) GetAll() []domain.ProductCategory {
	var data []ProductCategory
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *productCatData) GetbyID(dataID int) []domain.ProductCategory {
	var data []ProductCategory
	err := nd.db.Where("ID = ?", dataID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
