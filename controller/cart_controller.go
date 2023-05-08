package controller

import (
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"kasir/cafe/config"
	"kasir/cafe/model"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	_ "github.com/google/uuid"
)

// Endpoint 9 : add_cart
func Add_cart(c echo.Context) error {
	var cart model.Cart
	c.Bind(&cart)

	// checking 1 : is exist id_product
	var product model.Product
	isExist_IdProduct := config.DB.First(&product, cart.Id_product)
	if isExist_IdProduct.Error == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "id product not found",
			Data:    nil,
			Error:   isExist_IdProduct.Error.Error(),
		})
	}

	// checking 2 : is exist product in cart
	var cart2 model.Cart
	isExist_ProductinCart := config.DB.Where("id_product=?", cart.Id_product).First(&cart2)
	if isExist_ProductinCart.Error == nil {
		// id_product is exist in table cart

		cart2.Quantity += cart.Quantity
		err := config.DB.Save(&cart2).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.HttpResponse{
				Status:  400,
				Message: "failed add cart, invalid data format",
				Data:    nil,
				Error:   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, model.HttpResponse{
			Status:  200,
			Message: "success update quantity",
			Data:    cart2,
			Error:   nil,
		})
	}

	// fill price
	cart.Price = product.Price

	err := config.DB.Save(&cart).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.HttpResponse{
			Status:  400,
			Message: "failed add cart, invalid data format",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success create cart",
		Data:    cart,
		Error:   nil,
	})
}

// Endpoint 10 : update_cart
func Update_cart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cart model.Cart

	err := config.DB.First(&cart, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed update, no cart with matches id",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	err2 := c.Bind(&cart)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, model.HttpResponse{
			Status:  400,
			Message: "failed binding, invalid data format",
			Data:    nil,
			Error:   err2.Error(),
		})
	}

	err3 := config.DB.Save(&cart).Error
	if err3 != nil {
		return c.JSON(http.StatusBadRequest, model.HttpResponse{
			Status:  400,
			Message: "failed update cart, invalid data format",
			Data:    nil,
			Error:   err3.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success update cart by id",
		Data:    cart,
		Error:   nil,
	})
}

// Endpoint 11 : get_carts
func Get_carts(c echo.Context) error {
	var cart []model.Cart

	err := config.DB.Find(&cart).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Status:  500,
			Message: "failed get all cart",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success get all cart",
		Data:    cart,
		Error:   nil,
	})
}

// Endpoint 12 : get_carts_by_id
func Get_carts_by_id(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cart model.Cart

	err := config.DB.First(&cart, id).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed get cart, id not found",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success get cart by id",
		Data:    cart,
		Error:   nil,
	})
}

// Endpoint 13 : delete_cart
func Delete_cart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var cart model.Cart

	err := config.DB.First(&cart, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed delete, id not found",
			Data:    nil,
			Error:   err.Error(),
		})
	}
	config.DB.Delete(&cart)

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success delete cart",
		Data:    cart,
		Error:   nil,
	})
}

// Endpoint 14 : Get_nota
func Get_nota(c echo.Context) error {
	type nota_request struct {
		Buyer_name string
		Phone      string
	}

	var nota nota_request
	c.Bind(&nota)
	fmt.Println("nota", nota)
	var cart []model.Cart
	err := config.DB.Where("buyer_name=? AND phone=?", nota.Buyer_name, nota.Phone).Find(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed get nota, buyer_name or phone not found",
			Data:    nil,
			Error:   err.Error(),
		})
	}
	
	if len(cart)==0{
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed get nota, buyer_name or phone not found",
			Data:    nil,
			Error:   err,
		})
	}

	type menu struct {
		Id          int
		Name        string
		Quantity    int
		Price       int
		Total_price int
	}

	type nota_respone struct {
		Customer nota_request
		Menu     []menu
	}

	var response nota_respone
	response.Customer.Buyer_name = nota.Buyer_name
	response.Customer.Phone = nota.Phone
	for _, item := range cart {
		var product model.Product
		config.DB.First(&product, item.Id_product)

		fmt.Println(product)
		var temp menu
		temp.Id = int(product.ID)
		temp.Name = product.Name
		temp.Quantity = item.Quantity
		temp.Price = item.Price
		temp.Total_price = item.Price * item.Quantity
		response.Menu = append(response.Menu, temp)
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success get nota",
		Data:    response,
		Error:   nil,
	})
}
