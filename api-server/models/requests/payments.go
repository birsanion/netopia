package requests

type InitPaymentPayload struct {
	Amount            float64 `json:"amount" binding:"required"`
	Currency          string  `json:"currency" binding:"required"`
	Iban              string  `json:"iban" binding:"required"`
	Description       string  `json:"description" binding:"required"`
	InternalReference string  `json:"internal_reference" binding:"required"`
}
