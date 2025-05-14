package responses

import (
	models_db "github.com/birsanion/netopia/api-server/models/db"
)

type InitPaymentRespose struct {
	TransactionID string                  `json:"transaction_id"`
	Status        models_db.PaymentStatus `json:"status"`
}
