package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Buyer_name  string `json:"buyer_name" form:"buyer_name"`
	Phone       string `json:"phone" form:"phone"`
	Total_price int    `json:"total_price" form:"total_price"`
	Paid        bool `json:"paid" form:"paid"`
}
