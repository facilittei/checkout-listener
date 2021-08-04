package main

import (
	"github.com/facilittei/checkout-listener/adapters/handlers"
	"github.com/facilittei/checkout-listener/adapters/runners"
	"github.com/facilittei/checkout-listener/core/features"
)

func main() {
	httpHandler := handlers.NewHTTP()
	payment := features.NewPayment(httpHandler)
	SQShandler := handlers.NewSQS(payment)
	runner := runners.NewLambda(SQShandler)
	app := NewApp(runner)
	app.Run()
}
