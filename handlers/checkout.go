package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/facilittei/checkout-listener/gateways"
	"github.com/facilittei/checkout-listener/services"
)

// Version of the application handler
var Version = "dev"

// Checkout events
type Checkout struct {
	MessageGtw gateways.MessageContract
	PaymentSrv services.PaymentContract
}

// NewCheckout creates a new instance of a checkout handler
func NewCheckout(messageGtw gateways.MessageContract, paymentSrv services.PaymentContract) *Checkout {
	return &Checkout{
		MessageGtw: messageGtw,
		PaymentSrv: paymentSrv,
	}
}

// Handle events requests
func (checkoutHandler *Checkout) Handle(ctx context.Context, params interface{}) (string, error) {
	log.Println(params)

	if evt, ok := params.(*events.SQSEvent); ok {
		log.Println(evt)
		for _, message := range evt.Records {
			log.Println(message)
		}
	}

	payments := checkoutHandler.MessageGtw.GetPayments(params)
	log.Println(payments)
	return fmt.Sprintf("Version: %v", Version), nil
}
