package models

// Charge to be paid
type Charge struct {
	Description string   `json:"description"`
	Amount      float64  `json:"amount"`
	Methods     []string `json:"methods"`
}

// ChargeResponse
type ChargeResponse struct {
	Embedded struct {
		ID          string  `json:"id"`
		Code        string  `json:"code"`
		Reference   string  `json:"reference"`
		DueDate     string  `json:"dueDate"`
		CheckoutURL string  `json:"checkoutUrl"`
		Amount      float64 `json:"amount"`
		Status      string  `json:"status"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"_embedded"`
}
