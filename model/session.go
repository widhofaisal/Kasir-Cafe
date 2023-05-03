package model

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}