package model

type Admin struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
	Token    string `json:"token" form:"token"`
}

type AdminResponse struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
}