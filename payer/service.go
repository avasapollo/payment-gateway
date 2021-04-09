package payer

import (
	"context"
	"time"

	"github.com/avasapollo/payment-gateway/payments"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var _ payments.Payer = (*Service)(nil)

type Service struct {
	lgr *logrus.Entry
	transactions map[string]*payments.Transaction
}

func New() *Service  {
	return &Service{
		lgr: logrus.WithField("pkg","payer"),
	}
}

func (svc *Service) Authorize(ctx context.Context, req *payments.AuthorizeReq) (*payments.Transaction, error) {
	if req.Card.CardNumber == "4000000000000119" {
		return nil,payments.ErrAuthFailed
	}
	tr := &payments.Transaction{
		AuthID:        uuid.NewString(),
		Status:        payments.Authorized,
		Amount:        req.Amount,
		CheckedAmount: 0,
		Currency:      req.Currency,
		CreatedAt:     time.Now().UTC(),
		CardNumber: req.Card.CardNumber,
	}

	svc.transactions[tr.AuthID] = tr
	return tr,nil
}

func (svc *Service) Void(ctx context.Context, req *payments.VoidReq) (*payments.Transaction, error) {
	tr := svc.transactions[req.AuthID]
	switch  {
	case tr == nil:
		return nil,payments.ErrAutIDNotFound
	case tr.Status != payments.Authorized:
		return nil,payments.ErrVoidFailed
	}
	tr.Status = payments.Voided
	svc.transactions[req.AuthID] = tr
	return tr,nil
}

func (svc *Service) Capture(ctx context.Context, req *payments.CaptureReq) (*payments.Transaction, error) {
	tr := svc.transactions[req.AuthID]
	switch  {
	case tr == nil:
		 return nil,payments.ErrAutIDNotFound
	case tr.CardNumber == "4000000000000259":
		return nil,payments.ErrCaptureFailed
	case tr.Status != payments.Authorized :
		return nil,payments.ErrCaptureFailed
	}

	tr.Status = payments.Captured
	amount := tr.CheckedAmount + req.Amount
	if amount > tr.Amount {
		return nil,payments.ErrCaptureLimitExceeded
	}
	tr.CheckedAmount = amount
	svc.transactions[req.AuthID] = tr
	return tr,nil
}

func (svc *Service) Refund(ctx context.Context, req *payments.RefundReq) (*payments.Transaction, error) {
	panic("implement me")
}

