package handlers

import (
	"context"
	"fmt"

	"github.com/facilittei/checkout-listener/services"
)

// Version of the application handler
var Version = "dev"

// Checkout events
type Checkout struct {
	PaymentSrv services.PaymentContract
}

// NewCheckout creates a new instance of a checkout handler
func NewCheckout(paymentSrv services.PaymentContract) *Checkout {
	return &Checkout{
		PaymentSrv: paymentSrv,
	}
}

// Handle events requests
func (checkoutHandler *Checkout) Handle(ctx context.Context, params interface{}) (string, error) {
	return fmt.Sprintf("Version: %v", Version), nil
}
