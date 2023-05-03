package controller

import (
	"net/http"
	"strconv"
	_ "fmt"
	"net/http"

	"kasir/cafe/config"
	"kasir/cafe/model"

	"github.com/labstack/echo/v4"
  
  "gorm.io/gorm"

	_ "github.com/google/uuid"
)

// Endpoint 9 : AddProduct
func AddCart(c echo.Context) error {
	var cart model.Cart
	c.Bind(&cart)

	// checking 1 : is exist id_product
	var product model.Product
	isExist_IdProduct := config.DB.First(&product, cart.Id_product)
	if isExist_IdProduct.Error == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id product not found",
			"error":   isExist_IdProduct.Error.Error(),
		})
	}

	// fill price
	cart.Price = product.Price

	err := config.DB.Save(&cart).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed add cart",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create cart",
		"user":    cart,
    })
}

// Endpoint 10 : UpdateCart
func UpdateCart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cart model.Cart

	err := config.DB.First(&cart, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed update, no cart with matches id",
			"id":      id,
		})
	}

	err2 := c.Bind(&cart)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	config.DB.Save(&cart)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update cart by id",
		"data":    cart,
	
)

	
