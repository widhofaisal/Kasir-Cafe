package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id_cart     string `json:"id_cart" form:"id_cart"`
	Id_payment  string `json:"id_payment" form:"id_payment"`
	Id_product  string `json:"id_product" form:"id_product"`
	Quantity    string `json:"quantity" form:"quantity"`
	Total_price string `json:"total_price" form:"total_price"`
}
