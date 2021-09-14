package handlers

import "context"

// CheckoutContract handler to process payments
type CheckoutContract interface {
	Handle(ctx context.Context, params interface{}) (string, error)
}
