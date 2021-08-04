package config

import (
	"os"
	"sync"
)

var once sync.Once
var cfg Config

// Config settings
type Config struct {
	PaymentGatewayURL string
}

// Load configuration from environment
func Load() Config {
	once.Do(func() {
		cfg.PaymentGatewayURL = os.Getenv("PAYMENT_GATEWAY_URL")
	})

	return cfg
}
