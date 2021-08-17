package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/facilittei/checkout-listener/services"
)

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
	return fmt.Sprintf("Version: %v", os.Getenv("VERSION")), nil
}
