package controller

import (
	"fmt"
	"kasir/cafe/config"
	"kasir/cafe/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Endpoint 15 : Add_payment
func Add_payment(c echo.Context) error {
	type nota_request struct {
		Buyer_name string
		Phone      string
	}

	var nota nota_request
	c.Bind(&nota)

	var cart []model.Cart
	err := config.DB.Where("buyer_name=? AND phone=?", nota.Buyer_name, nota.Phone).Find(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed add payment, buyer_name or phone not found",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	if len(cart) == 0 {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed add payment, buyer_name or phone not found",
			Data:    nil,
			Error:   err,
		})
	}

	var totalPrice int

	for _, item := range cart {
		var product model.Product
		config.DB.First(&product, item.Id_product)

		totalPrice += item.Price * item.Quantity
	}

	var payment model.Payment
	payment.Buyer_name = nota.Buyer_name
	payment.Phone = nota.Phone
	payment.Total_price = totalPrice
	payment.Paid = false

	err_cek2 := config.DB.Save(&payment).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Status:  500,
			Message: "failed add payment",
			Data:    nil,
			Error:   err_cek2.Error(),
		})
	}

	err_cek3 := config.DB.Where("buyer_name=? AND phone=?", nota.Buyer_name, nota.Phone).Delete(&model.Cart{}).Error
	if err_cek3 == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed delete cart, buyer_name or phone not found",
			Data:    nil,
			Error:   err_cek3.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success add payment",
		Data:    payment,
		Error:   nil,
	})
}

// Endpoint 16 : Paid
func Paid(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	type Paid struct {
		Pay int
	}
	var paid Paid
	c.Bind(&paid)

	var payment model.Payment
	err := config.DB.First(&payment, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed update, no payment with matches id",
			Data:    nil,
			Error:   err.Error(),
		})
	}
	fmt.Println(paid.Pay, "-", payment.Total_price)
	if paid.Pay-payment.Total_price < 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to paid, need more money or insufficient balance",
			"need":    (paid.Pay - payment.Total_price) * -1,
		})
	}

	err_cek4 := config.DB.Model(&model.Payment{}).Where("id=?", id).Updates(map[string]interface{}{"paid": true}).Error
	if err_cek4 != nil {
		return c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Status:  500,
			Message: "failed to paid",
			Data:    nil,
			Error:   err_cek4.Error(),
		})
	}

	err_cek3 := config.DB.Delete(&model.Payment{}, id).Error
	if err_cek3 == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, model.HttpResponse{
			Status:  404,
			Message: "failed delete cart, buyer_name or phone not found",
			Data:    nil,
			Error:   err_cek3.Error(),
		})
	}

	config.DB.First(&payment, id)

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success to paid",
		Data: map[string]interface{}{
			"pay":             paid.Pay,
			"change":          paid.Pay - payment.Total_price,
			"updated payment": payment,
		},
		Error: nil,
	})
}

// Endpoint 17 : total_income
func Total_income(c echo.Context) error {
	var income int
	err := config.DB.Table("payments").Select("SUM(total_price)").Scan(&income).Error
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
		Data: map[string]interface{}{
			"income": income,
		},
		Error: nil,
	})
}
