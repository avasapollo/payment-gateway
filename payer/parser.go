package payer

import (
	"time"

	"github.com/avasapollo/payment-gateway/payments"
	"github.com/google/uuid"
	"golang.org/x/text/currency"
)

func toTransactionDtoFromAuthorizeReq(req *payments.AuthorizeReq) *TransactionDto {
	now := time.Now().UTC().Truncate(time.Millisecond)
	return &TransactionDto{
		AuthorizationID: uuid.NewString(),
		CardNumber:      req.Card.CardNumber,
		Status:          payments.Authorize.String(),
		Amount: &AmountDto{
			Value:    req.Amount,
			Currency: req.Currency.String(),
		},
		CaptureAmount: &AmountDto{
			Value:    req.Amount,
			Currency: req.Currency.String(),
		},
		RefundAmount: &AmountDto{
			Value:    0,
			Currency: req.Currency.String(),
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

}

func toTransaction(dto *TransactionDto) *payments.Transaction {
	c, _ := currency.ParseISO(dto.Amount.Currency)

	return &payments.Transaction{
		AuthorizationID: dto.AuthorizationID,
		CardNumber:      dto.CardNumber,
		Status:          payments.ToTransactionStatus(dto.Status),
		Amount: &payments.Amount{
			Value:    dto.Amount.Value,
			Currency: c,
		},
		CaptureAmount: &payments.Amount{
			Value:    dto.CaptureAmount.Value,
			Currency: c,
		},
		RefundAmount: &payments.Amount{
			Value:    dto.RefundAmount.Value,
			Currency: c,
		},
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
