package controllers

import (
	"net/http"
	"strconv"

	"kasir/cafe/config"
	"kasir/cafe/model"

	"github.com/labstack/echo/v4"
)

// Endpoint 4 : AddProduct
func AddProduct(c echo.Context) error {
	product := model.Product{}
	c.Bind(&product)

	err := config.DB.Save(&product).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed add product",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create product",
		"user":    product,
	})
}

// Endpoint 5 : GetProducts
func GetProducts(c echo.Context) error {
	var products []model.Product

	err := config.DB.Find(&products).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get all products",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all products",
		"admins":  products,
	})
}

// Endpoint 6 : GetProductsById
func GetProductsById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product model.Product

	err := config.DB.Find(&product, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get product by id",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product by id",
		"admins":  product,
	})
}

// Endpoint 7 : UpdateProduct
func UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product model.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed update, no product with matches id",
			"id":      id,
		})
	}

	err2 := c.Bind(&product)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	config.DB.Save(&product)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update product by id",
		"data":    product,
	})
}

// Endpoint 8 : DeleteProduct
func DeleteProduct(c echo.Context) error {
	var product model.Product
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.First(&product, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete, no product with matches id",
			"id":      id,
		})
	}
	config.DB.Delete(&product)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete product by id",
		"data":    product,
	})
}
