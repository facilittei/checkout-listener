package repositories

// AuthContract to manage authentication repositories
type AuthContract interface {
	GetToken() (string, error)
	StoreToken(token string) error
}
