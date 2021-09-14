package handlers

import "context"

// HandlerContract interface
type HandlerContract interface {
	Handle(ctx context.Context, params interface{}) (string, error)
}
