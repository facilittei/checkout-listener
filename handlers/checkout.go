package handlers

import (
	"context"
	"fmt"

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
	payments := checkoutHandler.MessageGtw.GetPayments(params)
	checkoutHandler.PaymentSrv.Process(payments[0])
	return fmt.Sprintf("Version: %v", Version), nil
}
