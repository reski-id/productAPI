package delivery

import (
	"productapi/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePC(e *echo.Echo, pc domain.ProductCatHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/productcategory", pc.InsertProductCategory())
	e.PUT("/productcategory/:id", pc.UpdateProductCategory())
	e.DELETE("/productcategory/:id", pc.DeleteProductCategory())
	e.GET("/productcategory", pc.GetAllProductCategory())
	e.GET("/productcategory/:id", pc.GetProductCategoryID())
}
