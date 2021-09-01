package models

// Charge to be paid
type Charge struct {
	Description string   `json:"description"`
	Amount      float64  `json:"amount"`
	Methods     []string `json:"methods"`
}
