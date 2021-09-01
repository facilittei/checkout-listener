package services

import (
	"log"

	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/config"
	"github.com/facilittei/checkout-listener/models"
	"github.com/facilittei/checkout-listener/repositories"
)

// Payment gateway
type Payment struct {
	Config         config.Config
	HTTPHandler    adapters.HTTP
	AuthRepository repositories.AuthContract
}

// NewPayment creates a new instance of a payment intent
func NewPayment(
	config config.Config,
	httpHandler adapters.HTTP,
	authRepository repositories.AuthContract,
) *Payment {
	return &Payment{
		Config:         config,
		HTTPHandler:    httpHandler,
		AuthRepository: authRepository,
	}
}

// Process payment
func (paymentSvc Payment) Process(payment models.Payment) error {
	log.Println(payment)
	return nil
}
