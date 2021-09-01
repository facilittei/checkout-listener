package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

// Config settings
type Config struct {
	PaymentGatewayURL string `mapstructure:"PAYMENT_GATEWAY_URL" required:"true"`
}

// Load configuration from environment
func Load(path string) (Config, error) {
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigName("application")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		loadFromEnvironment(&cfg)
		return cfg, nil
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	if err := validateEnvironment(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// loadFromEnvironment when there is no .env file provided
func loadFromEnvironment(cfg *Config) {
	cfg.PaymentGatewayURL = viper.GetString("PAYMENT_GATEWAY_URL")
}

// validateEnvironment checks whether all required variables are presented
func validateEnvironment(cfg Config) error {
	properties := reflect.TypeOf(cfg)

	for i := 0; i < properties.NumField(); i++ {
		property := properties.Field(i)
		tag := property.Tag
		environment := tag.Get("mapstructure")
		required := tag.Get("required")
		value := os.Getenv(environment)

		if value == "" && required == "true" {
			return fmt.Errorf("%s is required", environment)
		}
	}

	return nil
}
