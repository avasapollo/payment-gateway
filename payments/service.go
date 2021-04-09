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
		lgr:   logrus.WithField("pkg","payments"),
		payer: storage,
	}
}

func (svc *Service) Authorize(ctx context.Context, req *AuthorizeReq) (*Transaction, error) {
	if err := req.Validate();err != nil {
		return nil, err
	}
	return svc.payer.Authorize(ctx,req)
}

func (svc *Service) Void(ctx context.Context, req *VoidReq) (*Transaction, error) {
	if err := req.Validate();err!= nil {
		return nil,err
	}
	return svc.payer.Void(ctx,req)
}

func (svc *Service) Capture(ctx context.Context, req *CaptureReq) (*Transaction, error) {
	if err := req.Validate();err!= nil {
		return nil,err
	}
	return svc.payer.Capture(ctx,req)
}

func (svc *Service) Refund(ctx context.Context, req *RefundReq) (*Transaction, error) {
	if err := req.Validate();err!= nil {
		return nil,err
	}
	return svc.payer.Refund(ctx,req)
}

