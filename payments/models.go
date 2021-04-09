package payments

import (
	"time"

	"golang.org/x/text/currency"
)

type TransactionStatus int

const (
	Authorized TransactionStatus = iota
	Voided
	Captured
	Refunded
)

func (d TransactionStatus) String() string {
	return [...]string{"Authorized", "Voided", "Captured", "Refunded"}[d]
}

type Transaction struct {
	AuthID           string
	Status           TransactionStatus
	Amount           float64
	CheckedAmount 	 float64
	Currency         currency.Unit
	CreatedAt        time.Time
	CardNumber string
}

type Card struct {
	Name 	   string
	CardNumber string
	ExpireMonth string
	ExpireYear string
	CVV string
}

type AuthorizeReq struct {
	Card *Card
	Amount float64
	Currency currency.Unit
}

type VoidReq struct {
	AuthID string
}

func (r *VoidReq) Validate() error  {
	if r.AuthID == "" {
		return ErrAuthIDEmpty
	}
	return nil
}

type RefundReq struct {
	AuthID string
	Amount float64
}

func (r *RefundReq) Validate() error  {
	switch  {
	case r.AuthID == "":
		return ErrAuthIDEmpty
	case r.Amount <= 0:
		return ErrAmountNotValid
	}
	return nil
}

type CaptureReq struct {
	AuthID string
	Amount float64
}

func (r *CaptureReq) Validate() error  {
	switch  {
	case r.AuthID == "":
		return ErrAuthIDEmpty
	case r.Amount <= 0:
		return ErrAmountNotValid
	}
	return nil
}