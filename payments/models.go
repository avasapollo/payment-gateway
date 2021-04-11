package payments

import (
	"fmt"
	"time"

	creditcard "github.com/durango/go-credit-card"
	"golang.org/x/text/currency"
)

type TransactionStatus int

const (
	Unknown TransactionStatus = iota
	Authorize
	Void
	Capture
	Refund
)

func ToTransactionStatus(s string) TransactionStatus {
	switch s {
	case Authorize.String():
		return Authorize
	case Void.String():
		return Void
	case Capture.String():
		return Capture
	case Refund.String():
		return Refund
	default:
		return Unknown
	}
}

func (d TransactionStatus) String() string {
	return [...]string{"Unknown", "Authorize", "Void", "Capture", "Refund"}[d]
}

type Transaction struct {
	AuthorizationID string
	CardNumber      string
	Status          TransactionStatus
	Amount          *Amount
	CaptureAmount   *Amount
	RefundAmount    *Amount
	CurrentAmount   *Amount
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Amount struct {
	Value    float64
	Currency currency.Unit
}

type Card struct {
	Name        string
	CardNumber  string
	ExpireMonth string
	ExpireYear  string
	CVV         string
}

type AuthorizeReq struct {
	Card     *Card
	Amount   float64
	Currency currency.Unit
}

func (req *AuthorizeReq) Validate() error {
	switch {
	case req.Card == nil:
		return ErrCardNil
	case req.Card.Name == "":
		return ErrCardName
	}

	card := creditcard.Card{
		Number: req.Card.CardNumber,
		Cvv:    req.Card.CVV,
		Month:  req.Card.ExpireMonth,
		Year:   req.Card.ExpireYear,
	}

	if err := card.Validate(true); err != nil {
		return fmt.Errorf("%w: %s", ErrCardNotValid, err.Error())
	}

	if req.Amount <= 0 {
		return ErrAmountNotValid
	}
	return nil
}

type VoidReq struct {
	AuthorizationID string
}

func (r *VoidReq) Validate() error {
	if r.AuthorizationID == "" {
		return ErrAuthIDEmpty
	}
	return nil
}

type RefundReq struct {
	AuthorizationID string
	Amount          float64
}

func (r *RefundReq) Validate() error {
	switch {
	case r.AuthorizationID == "":
		return ErrAuthIDEmpty
	case r.Amount <= 0:
		return ErrAmountNotValid
	}
	return nil
}

type CaptureReq struct {
	AuthorizationID string
	Amount          float64
}

func (r *CaptureReq) Validate() error {
	switch {
	case r.AuthorizationID == "":
		return ErrAuthIDEmpty
	case r.Amount <= 0:
		return ErrAmountNotValid
	}
	return nil
}
