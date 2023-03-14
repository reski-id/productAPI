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

func SeedCategoryData(db *gorm.DB) {
	categories := []categoryData.Category{
		{Name: "Electronics", Enable: true},
		{Name: "Books", Enable: true},
		{Name: "Clothing", Enable: true},
		{Name: "Food", Enable: true},
	}
	for i := range categories {
		err := db.Create(&categories[i]).Error
		if err != nil {
			log.Println("Error seeding categories:", err.Error())
		}
	}
}

func SeedProductData(db *gorm.DB) {
	products := []productData.Product{
		{Name: "Iphone", Description: "Iphone Best Phone", Enable: true},
		{Name: "Samsung", Description: "Samsung Best Phone", Enable: true},
		{Name: "Harry Potter 1", Description: "Book of Harry Potter 1", Enable: true},
		{Name: "Flash Universe", Description: "Book of Flash", Enable: true},
		{Name: "Hat", Description: "Description for hat", Enable: true},
		{Name: "SmartWatch", Description: "Best Smartwach on Market", Enable: true},
		{Name: "Laptop HP", Description: "This is product HP Laptop", Enable: true},
		{Name: "Somay", Description: "Food in Indonesia", Enable: true},
		{Name: "Fizza", Description: "This is product no 9", Enable: true},
		{Name: "Golang for Dummy", Description: "This is product no 10", Enable: true},
	}
	for i := range products {
		err := db.Create(&products[i]).Error
		if err != nil {
			log.Println("Error seeding products:", err.Error())
		}
	}
}
