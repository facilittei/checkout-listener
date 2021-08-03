package adapters

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facilittei/checkout-listener/ports"
)

// Lambda AWS
type Lambda struct {
	Handler ports.Serverless
}

// NewLambda creates a new lambda instance
func NewLambda(handler ports.Serverless) *Lambda {
	return &Lambda{Handler: handler}
}

// Run lambda
func (app *Lambda) Run() {
	lambda.Start(app.Handler)
}
