package repositories

import "github.com/go-redis/redis/v8"

// AuthRedis to manage authentication tokens on Redis
type AuthRedis struct {
	DB *redis.Client
}

// NewAuthRedis creates a new instance of a auth redis repository
func NewAuthRedis(db *redis.Client) AuthContract {
	return &AuthRedis{
		DB: db,
	}
}

// GetToken authentication stored in redis
func (authRedis *AuthRedis) GetToken(token string) (string, error) {
	return "", nil
}

// StoreToken authentication in redis
func (authRedis *AuthRedis) StoreToken(token string) error {
	return nil
}
