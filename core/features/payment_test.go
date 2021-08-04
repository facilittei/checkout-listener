package features

import (
	"os"
	"testing"

	"github.com/facilittei/checkout-listener/adapters/handlers"
	"github.com/facilittei/checkout-listener/core/entities"
)

var customer entities.Customer

func init() {
	customer = entities.Customer{
		Name:     "",
		Email:    "",
		Document: "",
	}
}

func TestPaymentCreateChargeWithoutGatewayURL(t *testing.T) {
	os.Unsetenv("PAYMENT_GATEWAY_URL")
	httpHandler := handlers.NewHTTP()
	payment := NewPayment(httpHandler)

	if payment.CreateCharge(customer) == nil {
		t.Errorf("it should return an error from missing payment gateway URL")
	}
}
