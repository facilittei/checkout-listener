package models

// Payment to be processed
type Payment struct {
	Charge  Charge  `json:"charge"`
	Billing Billing `json:"billing"`
}
