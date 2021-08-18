package services

import (
	"log"

	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/models"
)

// Payment gateway
type Payment struct {
	HTTPHandler adapters.HTTP
}

// NewPayment creates a new instance of a payment intent
func NewPayment(httpHandler adapters.HTTP) *Payment {
	return &Payment{
		HTTPHandler: httpHandler,
	}
}

// Process payment
func (paymentSvc Payment) Process(payment models.Payment) error {
	log.Println(payment)
	return nil
}
