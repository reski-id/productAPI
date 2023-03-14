package delivery

import (
	"productapi/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePI(e *echo.Echo, ph domain.ProductImgHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/productimages", ph.InsertProductImages())
	e.PUT("/productimages/:id", ph.UpdateProductImages())
	e.DELETE("/productimages/:id", ph.DeleteProductImages())
	e.GET("/productimages", ph.GetAllProductImages())
	e.GET("/productimages/:id", ph.GetProductImagesID())
}
