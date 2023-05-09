package route

import (
	"kasir/cafe/constant"
	"kasir/cafe/controller"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/hello", controller.Hello_world)
	e.POST("/admins/login", controller.Login)

	// Authentication JWT
	authJWT := e.Group("")
	authJWT.Use(mid.JWT([]byte(constant.SECRET_JWT)))

	authJWT.GET("/admins", controller.Get_admins)

	authJWT.GET("/products", controller.Get_products)
	authJWT.GET("/products/:id", controller.Get_product_by_id)
	authJWT.POST("/products", controller.Add_product)
	authJWT.PUT("/products/:id", controller.Update_product)
	authJWT.DELETE("/products/:id", controller.Delete_product)

	authJWT.POST("/carts", controller.Add_cart)
	authJWT.PUT("/carts/:id", controller.Update_cart)
	authJWT.GET("/carts", controller.Get_carts)
	authJWT.GET("/carts/:id", controller.Get_carts_by_id)
	authJWT.GET("/carts/nota", controller.Get_nota)
	authJWT.DELETE("/carts/:id", controller.Delete_cart)
	
	authJWT.POST("/payments", controller.Add_payment)
	authJWT.PUT("/payments/:id", controller.Paid)
	
	return e
}
