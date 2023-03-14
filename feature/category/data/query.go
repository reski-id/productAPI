package data

import (
	"fmt"
	"log"
	"productapi/domain"

	"gorm.io/gorm"
)

type categoryData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.CategoryData {
	return &categoryData{
		db: db,
	}
}

func (nd *categoryData) Insert(newData domain.Category) domain.Category {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Category{}
	}
	return cnv.ToDomain()
}

func (bd *categoryData) Update(categoryID int, updatedData domain.Category) domain.Category {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", categoryID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Category{}
	}
	cnv.ID = uint(categoryID)
	return cnv.ToDomain()
}

func (nd *categoryData) Delete(categoryID int) bool {
	err := nd.db.Where("ID = ?", categoryID).Delete(&Category{})
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

func (nd *categoryData) GetAll() []domain.Category {
	var data []Category

	err := nd.db.Where("enable != ?", "false").Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *categoryData) GetCategoryID(categoryID int) []domain.Category {
	var data []Category
	err := nd.db.Where("ID = ?", categoryID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
