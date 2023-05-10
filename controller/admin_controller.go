package controller

import (
	"net/http"

	"kasir/cafe/config"
	"kasir/cafe/middleware"
	"kasir/cafe/model"

	"github.com/labstack/echo/v4"
)

// Endpoint 0 : hello_world
func Hello_world(c echo.Context) error {
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
		Message: "Welcome to the Kasir Cafe application",
		Data: map[string]interface{}{
			"owner":     "Widho Faisal Hakim",
			"github":    "https://github.com/widhofaisal/Kasir-Cafe",
			"instagram": "https://www.instagram.com/whydhoo/",
			"income":    income,
			"CI/CD":     true,
		},
		Error: nil,
	})
}

// Endpoint 2 : login
func Login(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)
	c.Get("user")
	err := config.DB.Where("username=? AND password=?", admin.Username, admin.Password).First(&admin).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.HttpResponse{
			Status:  401,
			Message: "failed login, username or password false",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	// generate jwt token
	token, err := middleware.CreateToken(admin.Username, admin.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Status:  500,
			Message: "failed generate token",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	adminResponse := model.LoginResponse{
		ID:       int(admin.ID),
		Username: admin.Username,
		Phone:    admin.Phone,
		Token:    token,
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  200,
		Message: "success login",
		Data:    adminResponse,
		Error:   nil,
	})
}

// Endpoint 3 : get_admins
func Get_admins(c echo.Context) error {
	var admins []model.Admin
	var adminResponse []model.AdminResponse

	if err := config.DB.Find(&admins).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.HttpResponse{
			Status:  500,
			Message: "failed get all admin",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	for _, admin := range admins {
		person := model.AdminResponse{
			ID:       int(admin.ID),
			Username: admin.Username,
			Phone:    admin.Phone,
		}
		adminResponse = append(adminResponse, person)
	}

	return c.JSON(http.StatusOK, model.HttpResponse{
		Status:  500,
		Message: "success get all admin",
		Data:    adminResponse,
		Error:   nil,
	})
}
