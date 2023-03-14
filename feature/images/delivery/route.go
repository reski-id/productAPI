package delivery

import (
	"productapi/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteContent(e *echo.Echo, ph domain.ImagesHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/file", ph.InsertImages())
	e.PUT("/file/:id", ph.UpdateImages())
	e.DELETE("/file/:id", ph.DeleteImages())
	e.GET("/file", ph.GetAllImages())
}
