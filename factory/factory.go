package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	pd "productapi/feature/product/data"
	productDelivery "productapi/feature/product/delivery"
	pu "productapi/feature/product/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {

	productData := pd.New(db)
	productCase := pu.New(productData)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteContent(e, productHandler)

}
