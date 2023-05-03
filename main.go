package main

import (
	"kasir/cafe/route"
)

func main() {
	e := route.New()
	e.Start(":8000")
}
