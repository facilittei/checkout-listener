package models

// Payment to be processed
type Payment struct {
	Description string   `json:"description"`
	Amount      float64  `json:"amount"`
	Methods     []string `json:"methods"`
	Customer
}
