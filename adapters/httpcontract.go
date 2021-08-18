package adapters

// HTTP contract
type HTTP interface {
	Post(url string, body interface{}, options map[string]string) ([]byte, error)
}
