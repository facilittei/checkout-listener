package services

import (
	"github.com/facilittei/checkout-listener/gateways"
	"github.com/facilittei/checkout-listener/models"
)

// Payment gateway
type Payment struct {
	MessageGtw gateways.MessageContract
}

// NewPayment creates a new instance of a payment intent
func NewPayment(messageGtw gateways.MessageContract) *Payment {
	return &Payment{
		MessageGtw: messageGtw,
	}
}

// Process payment
func (paymentSvc Payment) Process(payment models.Payment) error {
	return nil
}
