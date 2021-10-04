package gateways

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/facilittei/checkout-listener/models"
	"github.com/mitchellh/mapstructure"
)

// SQS queue
type SQS struct{}

// NewSQS creates a new listener instance
func NewSQS() MessageContract {
	return &SQS{}
}

// GetPayments SQS messages
func (sqs *SQS) GetPayments(params interface{}) []models.PaymentRequest {
	var paymentsRequest []models.PaymentRequest
	var sqsEvent events.SQSEvent

	err := mapstructure.Decode(params, &sqsEvent)
	if err != nil {
		log.Printf("could not decode message interface: %v", err)
		return paymentsRequest
	}

	for _, message := range sqsEvent.Records {
		var paymentRequest models.PaymentRequest
		json.Unmarshal([]byte(message.Body), &paymentRequest)
		paymentsRequest = append(paymentsRequest, paymentRequest)
	}

	return paymentsRequest
}
