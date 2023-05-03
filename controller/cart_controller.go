package controller

import (
	"net/http"
	"strconv"

	"kasir/cafe/config"
	"kasir/cafe/model"

	"github.com/labstack/echo/v4"
)

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
