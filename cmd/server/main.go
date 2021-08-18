package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/gateways"
	"github.com/facilittei/checkout-listener/handlers"
	"github.com/facilittei/checkout-listener/services"
)

func main() {
	httpHandler := adapters.NewHTTPRequest()
	messageGtw := gateways.NewSQS()
	paymentSvc := services.NewPayment(httpHandler)
	checkoutHandler := handlers.NewCheckout(messageGtw, paymentSvc)
	lambda.Start(checkoutHandler.Handle)
}
