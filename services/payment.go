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
	HTTPHandler    adapters.HTTPContract
	AuthRepository repositories.AuthContract
}

// NewPayment creates a new instance of a payment intent
func NewPayment(
	config config.Config,
	httpHandler adapters.HTTPContract,
	authRepository repositories.AuthContract,
) PaymentContract {
	return &Payment{
		Config:         config,
		HTTPHandler:    httpHandler,
		AuthRepository: authRepository,
	}
}

// Process payment
func (paymentSvc Payment) Process(payment models.Payment) error {
	paymentSvc.AuthRepository.StoreToken("abc1234")
	log.Println(paymentSvc.AuthRepository.GetToken())
	return nil
}
