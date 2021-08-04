package features

import (
	"errors"
	"fmt"

	"github.com/facilittei/checkout-listener/adapters/handlers"
	"github.com/facilittei/checkout-listener/config"
	"github.com/facilittei/checkout-listener/core/entities"
	"github.com/facilittei/checkout-listener/core/features/contracts"
)

// payment to the payment gateway
type payment struct {
	Message     entities.Message
	HTTPRequest handlers.HTTPHandler
}

// NewPayment creates a new instance of a payment intent
func NewPayment(httpRequest handlers.HTTPHandler) contracts.Payment {
	return &payment{
		HTTPRequest: httpRequest,
	}
}

// Process the payment intent
func (paymentService *payment) Process(message entities.Message) error {
	fmt.Println(message)
	return nil
}

// CreateCharge bill
func (paymentService *payment) CreateCharge(customer entities.Customer) error {
	if config.Load().PaymentGatewayURL == "" {
		return errors.New("payment gateway URL is required")
	}

	endpoint := config.Load().PaymentGatewayURL + "/api-integration/charges"
	resp, err := paymentService.HTTPRequest.Post(endpoint, customer, nil)
	if err != nil {
		return err
	}

	data := string(resp)
	fmt.Println(data)

	return nil
}
