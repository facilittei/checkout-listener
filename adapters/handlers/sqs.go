package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/facilittei/checkout-listener/core/entities"
	"github.com/facilittei/checkout-listener/core/features/contracts"
)

// SQS queue
type SQS struct {
	PaymentService contracts.Payment
}

// NewSQS creates a new listener instance
func NewSQS(paymentService contracts.Payment) *SQS {
	return &SQS{
		PaymentService: paymentService,
	}
}

// Handler SQS messages
func (sqs *SQS) Handler(ctx context.Context, params interface{}) error {
	if evt, ok := params.(*events.SQSEvent); ok {
		for _, message := range evt.Records {
			sqs.PaymentService.Process(entities.Message{
				ID:   message.MessageId,
				Body: message.Body,
			})
		}
	}

	return nil
}
