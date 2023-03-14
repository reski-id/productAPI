package mysql

import (
	"fmt"
	"log"
	"productapi/config"

	categoryData "productapi/feature/category/data"
	imagesData "productapi/feature/images/data"
	productData "productapi/feature/product/data"
	pcData "productapi/feature/product_category/data"
	piData "productapi/feature/product_images/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(productData.Product{}, imagesData.Images{}, categoryData.Category{}, pcData.ProductCategory{}, piData.ProductImages{})

}
