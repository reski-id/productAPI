package data

import (
	"fmt"
	"log"
	"productapi/domain"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductData {
	return &productData{
		db: db,
	}
}

func (nd *productData) Insert(newData domain.Product) domain.Product {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Product{}
	}
	return cnv.ToDomain()
}

func (bd *productData) Update(productID int, updatedData domain.Product) domain.Product {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", productID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Product{}
	}
	cnv.ID = uint(productID)
	return cnv.ToDomain()
}

func (nd *productData) Delete(productID int) bool {
	err := nd.db.Where("ID = ?", productID).Delete(&Product{})
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

func (nd *productData) GetAll() []domain.Product {
	var data []Product

	err := nd.db.Where("enable != ?", "false").Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *productData) GetProductID(productID int) []domain.Product {
	var data []Product
	err := nd.db.Where("ID = ?", productID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
