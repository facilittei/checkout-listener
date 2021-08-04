package main

import (
	"github.com/facilittei/checkout-listener/adapters/handlers"
	"github.com/facilittei/checkout-listener/adapters/runners"
	"github.com/facilittei/checkout-listener/core/services"
)

func main() {
	payment := services.NewPayment()
	handler := handlers.NewSQS(payment)
	runner := runners.NewLambda(handler)
	app := NewApp(runner)
	app.Run()
}
