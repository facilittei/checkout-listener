package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// AuthRedis to manage authentication tokens on Redis
type AuthRedis struct {
	Ctx context.Context
	DB  *redis.Client
}

// NewAuthRedis creates a new instance of a auth redis repository
func NewAuthRedis(ctx context.Context, db *redis.Client) AuthContract {
	return &AuthRedis{
		Ctx: ctx,
		DB:  db,
	}
}

// GetToken authentication stored in redis
func (authRedis *AuthRedis) GetToken() (string, error) {
	result, err := authRedis.DB.Get(authRedis.Ctx, "token").Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// StoreToken authentication in redis
func (authRedis *AuthRedis) StoreToken(token string) error {
	return authRedis.DB.Set(authRedis.Ctx, "token", token, time.Hour).Err()
}
