package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facilittei/checkout-listener/gateways"
	"github.com/facilittei/checkout-listener/handlers"
	"github.com/facilittei/checkout-listener/services"
)

func main() {
	messageGtw := gateways.NewSQS()
	paymentSvc := services.NewPayment(messageGtw)
	checkoutHandler := handlers.NewCheckout(paymentSvc)
	lambda.Start(checkoutHandler.Handle)
}
