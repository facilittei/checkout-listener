package gateways

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/facilittei/checkout-listener/models"
)

// SQS queue
type SQS struct{}

// NewSQS creates a new listener instance
func NewSQS() *SQS {
	return &SQS{}
}

// GetPayments SQS messages
func (sqs *SQS) GetPayments(params interface{}) []models.Payment {
	var payments []models.Payment

	log.Println(params.(*events.SQSEvent))

	if evt, ok := params.(*events.SQSEvent); ok {
		log.Println(evt)
		for _, message := range evt.Records {
			var payment models.Payment
			json.Unmarshal([]byte(message.Body), &payment)
			payments = append(payments, payment)
		}
	}

	return payments
}
