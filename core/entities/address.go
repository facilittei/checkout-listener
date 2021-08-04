package entities

// Address information
type Address struct {
	Street   string `json:"street"`
	Number   string `json:"number"`
	City     string `json:"city"`
	State    string `json:"state"`
	PostCode string `json:"postCode"`
}
