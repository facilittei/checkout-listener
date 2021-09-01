package config

import (
	"os"
	"testing"
)

func TestLoadFromEnvironment(t *testing.T) {
	paymentGatewayURL := "https://payments.facilittei.com"
	os.Setenv("PAYMENT_GATEWAY_URL", paymentGatewayURL)

	cfg, err := Load(".")
	if err != nil {
		t.Errorf("got an error %s", err)
	}

	os.Unsetenv("PAYMENT_GATEWAY_URL")
	if cfg.PaymentGatewayURL != paymentGatewayURL {
		t.Errorf("want %s but got %s", cfg.PaymentGatewayURL, paymentGatewayURL)
	}
}

func TestRequiredEnvironmentVariable(t *testing.T) {
	err := validateEnvironment(Config{})
	want := "PAYMENT_GATEWAY_URL is required"
	got := err.Error()
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}
