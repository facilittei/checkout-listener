package entities

// Message order requested to be processed by a payment gateway
type Message struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}
