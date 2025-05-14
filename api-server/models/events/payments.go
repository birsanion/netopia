package events

type CreatePayment struct {
	TransactionID     string  `json:"transaction_id"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	Iban              string  `json:"iban"`
	Description       string  `json:"description"`
	InternalReference string  `json:"internal_reference"`
}
