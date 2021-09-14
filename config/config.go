package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

// Config settings
type Config struct {
	PaymentGatewayURL   string `mapstructure:"PAYMENT_GATEWAY_URL" required:"true"`
	PaymentGatewayAuth  string `mapstructure:"PAYMENT_GATEWAY_AUTH" required:"true"`
	PaymentGatewayToken string `mapstructure:"PAYMENT_GATEWAY_TOKEN" required:"true"`
	RedisHost           string `mapstructure:"REDIS_HOST" required:"true"`
	RedisPort           string `mapstructure:"REDIS_PORT" required:"true"`
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
	cfg.PaymentGatewayAuth = viper.GetString("PAYMENT_GATEWAY_AUTH")
	cfg.PaymentGatewayToken = viper.GetString("PAYMENT_GATEWAY_TOKEN")
	cfg.RedisHost = viper.GetString("REDIS_HOST")
	cfg.RedisPort = viper.GetString("REDIS_PORT")
}

// validateEnvironment checks whether all required variables are presented
func validateEnvironment(cfg Config) error {
	properties := reflect.TypeOf(cfg)

	for i := 0; i < properties.NumField(); i++ {
		property := properties.Field(i)
		tag := property.Tag
		environment := tag.Get("mapstructure")
		required := tag.Get("required")
		value := reflect.ValueOf(cfg).Field(i)

		if value.String() == "" && required == "true" {
			return fmt.Errorf("%s is required", environment)
		}
	}

	return nil
}
