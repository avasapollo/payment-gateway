package payments

import (
	"context"

	creditcard "github.com/durango/go-credit-card"
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

func (svc *Service) validateAuthorizeReq(req *AuthorizeReq)  error {
	switch  {
	case req.Card == nil :
		return ErrCardNil
	case req.Card.Name == "":
		return ErrCardName
	}

	card :=  creditcard.Card{
		 Number:  req.Card.CardNumber,
		 Cvv:     req.Card.CVV,
		 Month:   req.Card.ExpireMonth,
		 Year:    req.Card.ExpireYear,
	 }

	 if err:=  card.Validate(true);err != nil {
	 	return err
	 }

	if  req.Amount <= 0 {
		return ErrAmountNotValid
	}
	return nil
}

func (svc *Service) Authorize(ctx context.Context, req *AuthorizeReq) (*Transaction, error) {
	if err := svc.validateAuthorizeReq(req);err != nil {
		return nil, err
	}
	return svc.payer.Authorize(ctx,req)
}

func (svc *Service) validateVoidReq(req *VoidReq) error  {
	return req.Validate()
}

func (svc *Service) Void(ctx context.Context, req *VoidReq) (*Transaction, error) {
	if err := svc.validateVoidReq(req);err!= nil {
		return nil,err
	}
	return svc.payer.Void(ctx,req)
}

func (svc *Service) validateCaptureReq(req *CaptureReq) error  {
	return req.Validate()
}

func (svc *Service) Capture(ctx context.Context, req *CaptureReq) (*Transaction, error) {
	if err := svc.validateCaptureReq(req);err!= nil {
		return nil,err
	}
	return svc.payer.Capture(ctx,req)
}

func (svc *Service) validateRefundReq(req *RefundReq) error  {
	return req.Validate()
}

func (svc *Service) Refund(ctx context.Context, req *RefundReq) (*Transaction, error) {
	if err := svc.validateRefundReq(req);err!= nil {
		return nil,err
	}
	return svc.payer.Refund(ctx,req)
}

