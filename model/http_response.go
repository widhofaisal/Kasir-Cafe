package model

type HttpResponse struct {
	Status  int
	Message string
	Data    interface{}
	Error   interface{}
}
