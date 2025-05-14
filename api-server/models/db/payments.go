package db

import (
	"time"

	"github.com/birsanion/netopia/api-server/helpers/rand"
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusNew      PaymentStatus = "new"
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusError    PaymentStatus = "error"
	PaymentStatusApproved PaymentStatus = "approved"
)

type Payment struct {
	ID                int64 `gorm:"primaryKey"`
	TransactionID     string
	Amount            float64
	Currency          string
	Iban              string
	Description       string
	InternalReference string
	Status            PaymentStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.TransactionID == "" {
		// TODO: fixed scheme
		p.TransactionID = rand.RandStringBytes(10, rand.Digits+rand.UpperLetters)
	}
	return
}
