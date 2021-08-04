package handlers

import "context"

// Handler request
type Handler interface {
	Handler(ctx context.Context, params interface{}) error
}
