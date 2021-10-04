package services

import "github.com/facilittei/checkout-listener/models"

// PaymentContract to process requests
type PaymentContract interface {
	Process(payment models.PaymentRequest) error
}
