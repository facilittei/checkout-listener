package handlers

import (
	"context"
)

// AppHandler request
type AppHandler interface {
	Handler(ctx context.Context, params interface{}) error
}

// HTTPHandler to handle request
type HTTPHandler interface {
	Post(url string, body interface{}, options map[string]string) ([]byte, error)
}
