package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
	Id_admin string `json:"id_admin" form:"id_admin"`
}