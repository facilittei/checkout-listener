package runners

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facilittei/checkout-listener/adapters/handlers"
)

// Lambda AWS
type Lambda struct {
	Handler handlers.Handler
}

// NewLambda creates a new lambda instance
func NewLambda(handler handlers.Handler) *Lambda {
	return &Lambda{Handler: handler}
}

// Run lambda
func (app *Lambda) Run() {
	lambda.Start(app.Handler)
}
