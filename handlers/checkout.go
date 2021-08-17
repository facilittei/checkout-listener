package handlers

import (
	"context"
	"fmt"

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
func (checkoutHandler *Checkout) Handle(ctx context.Context, params interface{}) error {
	fmt.Println("It worked!")
	return nil
}
