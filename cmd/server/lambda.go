package server

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facilittei/checkout-listener/gateways"
	"github.com/facilittei/checkout-listener/handlers"
	"github.com/facilittei/checkout-listener/services"
)

// Lambda AWS
type Lambda struct{}

// NewLambda instance
func NewLambda() *Lambda {
	return &Lambda{}
}

// Start application
func (app *Lambda) Start() {
	messageGtw := gateways.NewSQS()
	paymentSvc := services.NewPayment(messageGtw)
	checkoutHandler := handlers.NewCheckout(paymentSvc)
	lambda.Start(checkoutHandler.Handle)
}
