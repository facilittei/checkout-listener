package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/config"
	"github.com/facilittei/checkout-listener/gateways"
	"github.com/facilittei/checkout-listener/handlers"
	"github.com/facilittei/checkout-listener/repositories"
	"github.com/facilittei/checkout-listener/services"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load(".")
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	httpHandler := adapters.NewHTTPResty()
	messageGtw := gateways.NewSQS()
	authRepository := repositories.NewAuthRedis(ctx, rdb)
	authToken := handlers.NewAuth(cfg, httpHandler, authRepository)

	headers, err := authToken.GetCredentials()
	if err != nil {
		log.Fatal(err)
	}

	paymentSvc := services.NewPayment(cfg, httpHandler, headers)
	checkoutHandler := handlers.NewCheckout(messageGtw, paymentSvc)
	lambda.Start(checkoutHandler.Handle)
}
