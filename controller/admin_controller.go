package controller

import (
	"net/http"

	"kasir/cafe/config"
	"kasir/cafe/middleware"
	"kasir/cafe/model"

	"github.com/labstack/echo/v4"
)

// Endpoint 0 : GetHelloWorld
func GetHelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello World My Name Widho Faisal Hakim",
	})
}

// Endpoint 2 : Login
func Login(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)

	err := config.DB.Where("username=? AND password=?", admin.Username, admin.Password).First(&admin).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	// generate jwt token
	token, err := middleware.CreateToken(admin.Username, admin.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed generate token",
			"error":   err.Error(),
		})
	}

	adminResponse := model.LoginResponse{
		ID:       int(admin.ID),
		Username: admin.Username,
		Phone:    admin.Phone,
		Token:    token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    adminResponse,
	})
}

// Endpoint 3 : GetAdmins
func GetAdmins(c echo.Context) error {
	var admins []model.Admin
	var adminResponse []model.AdminResponse

	if err := config.DB.Find(&admins).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, admin := range admins {
		person := model.AdminResponse{
			ID:       int(admin.ID),
			Username: admin.Username,
			Phone:    admin.Phone,
		}
		adminResponse = append(adminResponse, person)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get admins",
		"admins":  adminResponse,
	})
}
