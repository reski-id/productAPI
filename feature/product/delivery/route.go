package delivery

import (
	"productapi/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteProduct(e *echo.Echo, ph domain.ProductHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/products", ph.InsertProduct())
	e.PUT("/products/:id", ph.UpdateProduct())
	e.DELETE("/products/:id", ph.DeleteProduct())
	e.GET("/products", ph.GetAllProduct())
}
