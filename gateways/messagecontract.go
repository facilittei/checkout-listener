package gateways

import (
	"github.com/facilittei/checkout-listener/models"
)

// MessageContract that has the payment checkout request
type MessageContract interface {
	GetPayments(params interface{}) []models.PaymentRequest
}
