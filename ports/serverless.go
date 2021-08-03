package ports

import (
	"context"
)

// Serverless application to run
type Serverless interface {
	Handler(ctx context.Context, params interface{}) error
}
