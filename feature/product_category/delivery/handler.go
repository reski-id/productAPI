package delivery

import (
	"log"
	"net/http"
	"productapi/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productCatHandler struct {
	productCatUsecase domain.ProductCatUseCase
}

func New(nu domain.ProductCatUseCase) domain.ProductCatHandler {
	return &productCatHandler{
		productCatUsecase: nu,
	}
}

func (nh *productCatHandler) InsertProductCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := nh.productCatUsecase.AddProductCategory(tmp.ToDomain())
		if err != nil {
			log.Println("Cannot proces data,ID sudah Ada", err)

			c.JSON(http.StatusInternalServerError, "Cannot proces data,ID sudah Ada")
			return (err)

		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})
	}
}

func (nh *productCatHandler) UpdateProductCategory() echo.HandlerFunc {
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

		data, err := nh.productCatUsecase.UpProductCategory(cnv, tmp.ToDomain())

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

func (nh *productCatHandler) DeleteProductCategory() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := nh.productCatUsecase.DelProductCategory(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete data")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data",
		})
	}
}

func (nh *productCatHandler) GetAllProductCategory() echo.HandlerFunc {
	return func(c echo.Context) error {

		data, err := nh.productCatUsecase.GetAllD()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "Cannot get data no data found")

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

func (nh *productCatHandler) GetProductCategoryID() echo.HandlerFunc {
	return func(c echo.Context) error {

		idProductCategory := c.Param("id")
		id, _ := strconv.Atoi(idProductCategory)
		data, err := nh.productCatUsecase.GetSpecificProductCategory(id)

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
