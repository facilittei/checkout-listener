package main

import (
	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/core/services"
	"github.com/facilittei/checkout-listener/ports"
)

func main() {
	payment := services.NewPayment()
	handler := ports.NewSQS(payment)
	app := adapters.NewLambda(handler)
	app.Run()
}
