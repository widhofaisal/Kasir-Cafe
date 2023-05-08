package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Buyer_name string `json:"buyer_name" form:"buyer_name"`
	Phone      string `json:"phone" form:"phone"`
	Id_product string `json:"id_product" form:"id_product"`
	Quantity   int    `json:"quantity" form:"quantity"`
	Price      int    `json:"price" form:"price"`
}
