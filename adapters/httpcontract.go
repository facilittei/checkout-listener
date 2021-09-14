package adapters

// HTTPContract interface
type HTTPContract interface {
	Post(url string, body interface{}, options map[string]string) ([]byte, error)
}
