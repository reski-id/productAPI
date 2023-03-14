package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	pd "productapi/feature/product/data"
	productDelivery "productapi/feature/product/delivery"
	pu "productapi/feature/product/usecase"

	id "productapi/feature/images/data"
	imagesDelivery "productapi/feature/images/delivery"
	iu "productapi/feature/images/usecase"

	ci "productapi/feature/category/data"
	categoryDelivery "productapi/feature/category/delivery"
	cu "productapi/feature/category/usecase"

	pci "productapi/feature/product_category/data"
	prodCategoryDelivery "productapi/feature/product_category/delivery"
	pcu "productapi/feature/product_category/usecase"

	pii "productapi/feature/product_images/data"
	imgImgDelivery "productapi/feature/product_images/delivery"
	piu "productapi/feature/product_images/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {

	productData := pd.New(db)
	productCase := pu.New(productData)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteProduct(e, productHandler)

	imagesData := id.New(db)
	imagesCase := iu.New(imagesData)
	imagesHandler := imagesDelivery.New(imagesCase)
	imagesDelivery.RouteImages(e, imagesHandler)

	categoryData := ci.New(db)
	categoryCase := cu.New(categoryData)
	categoryHandler := categoryDelivery.New(categoryCase)
	categoryDelivery.RouteCategory(e, categoryHandler)

	prodCategoryData := pci.New(db)
	prodCategoryCase := pcu.New(prodCategoryData)
	prodCategoryHandler := prodCategoryDelivery.New(prodCategoryCase)
	prodCategoryDelivery.RoutePC(e, prodCategoryHandler)

	imgCategoryData := pii.New(db)
	imgCategoryCase := piu.New(imgCategoryData)
	imgCategoryHandler := imgImgDelivery.New(imgCategoryCase)
	imgImgDelivery.RoutePI(e, imgCategoryHandler)

}
