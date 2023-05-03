package route

import (
	"kasir/cafe/constant"
	"kasir/cafe/controller"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/hello", controller.GetHelloWorld)
	e.POST("/admins/login", controller.Login)

	// Authentication JWT
	authJWT := e.Group("")
	authJWT.Use(mid.JWT([]byte(constant.SECRET_JWT)))
	authJWT.GET("/admins", controller.GetAdmins)
	authJWT.GET("/products", controller.GetProducts)
	authJWT.GET("/products/:id", controller.GetProductsById)
	authJWT.POST("/products", controller.AddProduct)
	authJWT.PUT("/products/:id", controller.UpdateProduct)
	authJWT.DELETE("/products/:id", controller.DeleteProduct)

	return e
}
