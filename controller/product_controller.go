package controller

import (
	"net/http"
	"strconv"

	"kasir/cafe/config"
	"kasir/cafe/model"

	"github.com/labstack/echo/v4"
)

// Endpoint 4 : add_product
func Add_product(c echo.Context) error {
	product := model.Product{}
	c.Bind(&product)

	err := config.DB.Save(&product).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Status:  500,
			Message: "failed add product",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success create product",
		Data:    product,
		Error:   nil,
	})
}

// Endpoint 5 : get_products
func Get_products(c echo.Context) error {
	var products []model.Product

	err := config.DB.Find(&products).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Status:  500,
			Message: "failed get all product",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success get all products",
		Data:    products,
		Error:   nil,
	})
}

// Endpoint 6 : get_product_by_id
func Get_product_by_id(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product model.Product

	err := config.DB.Find(&product, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed get product by id, id not found",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success get product by id",
		Data:    product,
		Error:   nil,
	})
}

// Endpoint 7 : update_product
func Update_product(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product model.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed update, no product with matches id",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	err2 := c.Bind(&product)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, model.HttpResponse{
			Status:  400,
			Message: "failed binding, invalid data format",
			Data:    nil,
			Error:   err.Error(),
		})
	}
	config.DB.Save(&product)

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success update product by id",
		Data:    product,
		Error:   nil,
	})
}

// Endpoint 8 : delete_product
func Delete_product(c echo.Context) error {
	var product model.Product
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.First(&product, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed delete, no product with matches id",
			Data:    nil,
			Error:   err.Error(),
		})
	}
	config.DB.Delete(&product)

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success delete product by id",
		Data:    product,
		Error:   nil,
	})
}
