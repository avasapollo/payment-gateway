package payments

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Service struct {
	lgr   *logrus.Entry
	payer Payer
}

func New(storage Payer) *Service {
	return &Service{
		lgr:   logrus.WithField("pkg", "payments"),
		payer: storage,
	}
}

// Authorize create the transaction
func (svc *Service) Authorize(ctx context.Context, req *AuthorizeReq) (*Transaction, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// truncate the float example 10.988 to 10.98
	req.Amount = truncateFloat(req.Amount)
	return svc.payer.Authorize(ctx, req)
}

// Void the transaction, after that it will not more accessible to capture and refund
func (svc *Service) Void(ctx context.Context, req *VoidReq) (*Transaction, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return svc.payer.Void(ctx, req)
}

// Capture collect the credit from the amount, it is possible to capture multiple times until the actual amount of the transaction
func (svc *Service) Capture(ctx context.Context, req *CaptureReq) (*Transaction, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	req.Amount = truncateFloat(req.Amount)
	return svc.payer.Capture(ctx, req)
}

// Refund the credit from the amount that is captured, it is possible to refund multiple times until the actual captured
func (svc *Service) Refund(ctx context.Context, req *RefundReq) (*Transaction, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	req.Amount = truncateFloat(req.Amount)
	return svc.payer.Refund(ctx, req)
}
