package models

// Payment to be processed
type Payment struct {
	Charge  Charge  `json:"charge"`
	Billing Billing `json:"billing"`
}

// PaymentRequest from client
type PaymentRequest struct {
	Payment    Payment    `json:"payment"`
	CreditCard CreditCard `json:"creditCard"`
}
