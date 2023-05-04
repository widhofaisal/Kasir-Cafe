package controller

import (
	_ "fmt"
	"net/http"
	"strconv"

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

	// checking 2 : is exist product in cart
	var cart2 model.Cart
	isExist_ProductinCart := config.DB.Where("id_product=?", cart.Id_product).First(&cart2)
	if isExist_ProductinCart.Error == nil {
		// id_product is exist in table cart

		cart2.Quantity += cart.Quantity
		config.DB.Save(&cart2)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update quantity",
			"user":    cart2,
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
	})
}

// Endpoint 11 : GetCart
func GetCarts(c echo.Context) error {
	var cart []model.Cart

	err := config.DB.Find(&cart).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get all carts",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all carts",
		"admins":  cart,
	})
}

// Endpoint 12 : GetCartById
func GetCartById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cart model.Cart

	err := config.DB.First(&cart, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get cart by id",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get cart by id",
		"admins":  cart,
	})
}
