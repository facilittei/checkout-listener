package handlers

import (
	"context"
	"fmt"
	"log"

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
func NewCheckout(messageGtw gateways.MessageContract, paymentSrv services.PaymentContract) CheckoutContract {
	return &Checkout{
		MessageGtw: messageGtw,
		PaymentSrv: paymentSrv,
	}
}

// Handle events requests
func (checkoutHandler *Checkout) Handle(ctx context.Context, params interface{}) (string, error) {
	log.Print("Handle.params:")
	log.Println(params)
	payments := checkoutHandler.MessageGtw.GetPayments(params)
	for _, payment := range payments {
		log.Print("Handle.payment:")
		log.Printf("%+v", &payment)
		checkoutHandler.PaymentSrv.Process(payment)
	}
	return fmt.Sprintf("Version: %v", Version), nil
}
