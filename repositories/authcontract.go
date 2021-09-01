package repositories

// AuthContract to manage authentication repositories
type AuthContract interface {
	GetToken(token string) (string, error)
	StoreToken(token string) error
}
