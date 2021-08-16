package models

// Customer that is making a payment
type Customer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	Address
}
