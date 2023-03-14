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
)

func Initfactory(e *echo.Echo, db *gorm.DB) {

	productData := pd.New(db)
	productCase := pu.New(productData)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteContent(e, productHandler)

	imagesData := id.New(db)
	imagesCase := iu.New(imagesData)
	imagesHandler := imagesDelivery.New(imagesCase)
	imagesDelivery.RouteContent(e, imagesHandler)

	categoryData := ci.New(db)
	categoryCase := cu.New(categoryData)
	categoryHandler := categoryDelivery.New(categoryCase)
	categoryDelivery.RouteContent(e, categoryHandler)

}
