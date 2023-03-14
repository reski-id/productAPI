package delivery

import (
	"productapi/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteCategory(e *echo.Echo, ph domain.CategoryHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/category", ph.InsertCategory())
	e.PUT("/category/:id", ph.UpdateCategory())
	e.DELETE("/category/:id", ph.DeleteCategory())
	e.GET("/category", ph.GetAllCategory())
}
