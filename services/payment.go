package services

import (
	"log"

	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/config"
	"github.com/facilittei/checkout-listener/models"
)

// Payment gateway
type Payment struct {
	Config      config.Config
	HTTPHandler adapters.HTTP
}

// NewPayment creates a new instance of a payment intent
func NewPayment(config config.Config, httpHandler adapters.HTTP) *Payment {
	return &Payment{
		Config:      config,
		HTTPHandler: httpHandler,
	}
}

// Process payment
func (paymentSvc Payment) Process(payment models.Payment) error {
	log.Println(payment)
	return nil
}
