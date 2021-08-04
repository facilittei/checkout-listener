package services

import (
	"fmt"

	"github.com/facilittei/checkout-listener/core/entities"
	"github.com/facilittei/checkout-listener/core/services/contracts"
)

// payment to the payment gateway
type payment struct {
	Message entities.Message
}

// NewPayment creates a new instance of a payment intent
func NewPayment() contracts.Payment {
	return &payment{}
}

// Process the payment intent
func (paymentService *payment) Process(message entities.Message) error {
	fmt.Println(message)
	return nil
}
