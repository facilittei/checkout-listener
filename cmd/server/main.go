package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/config"
	"github.com/facilittei/checkout-listener/gateways"
	"github.com/facilittei/checkout-listener/handlers"
	"github.com/facilittei/checkout-listener/services"
)

func main() {
	cfg, err := config.Load(".")
	if err != nil {
		log.Fatal(err)
	}

	httpHandler := adapters.NewHTTPResty()
	messageGtw := gateways.NewSQS()
	paymentSvc := services.NewPayment(cfg, httpHandler)
	checkoutHandler := handlers.NewCheckout(messageGtw, paymentSvc)
	lambda.Start(checkoutHandler.Handle)
}
