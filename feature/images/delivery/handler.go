package delivery

import (
	"log"
	"net/http"
	"productapi/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type imagesHandler struct {
	imagesUsecase domain.ImagesUseCase
}

func New(nu domain.ImagesUseCase) domain.ImagesHandler {
	return &imagesHandler{
		imagesUsecase: nu,
	}
}

func (nh *imagesHandler) InsertImages() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := nh.imagesUsecase.AddImages(1, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}

func (nh *imagesHandler) UpdateImages() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp InsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}
		data, err := nh.imagesUsecase.UpImages(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"data":    data,
		})
	}
}

func (nh *imagesHandler) DeleteImages() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := nh.imagesUsecase.DelImages(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data",
		})
	}
}

func (nh *imagesHandler) GetImagesID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idImages := c.Param("id")
		id, _ := strconv.Atoi(idImages)
		data, err := nh.imagesUsecase.GetSpecificImages(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my data",
			"users":   data,
		})
	}
}

func (nh *imagesHandler) GetAllImages() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := nh.imagesUsecase.GetAllP()
		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    data,
		})
	}
}
