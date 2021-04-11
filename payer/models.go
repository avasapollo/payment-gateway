package payer

import (
	"time"
)

type TransactionDto struct {
	AuthorizationID string     `bson:"_id"`
	CardNumber      string     `bson:"card_number"`
	Status          string     `bson:"status"`
	Amount          *AmountDto `bson:"amount"`
	CaptureAmount   *AmountDto `bson:"capture_amount"`
	RefundAmount    *AmountDto `bson:"refund_amount"`
	CurrentAmount   *AmountDto `bson:"current_amount"`
	CreatedAt       time.Time  `bson:"created_at"`
	UpdatedAt       time.Time  `bson:"updated_at"`
}

type AmountDto struct {
	Value    float64 `bson:"value"`
	Currency string  `bson:"currency"`
}
