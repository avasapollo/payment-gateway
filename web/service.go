package web

import (
	"context"
	"errors"

	"github.com/avasapollo/payment-gateway/payments"
	v1 "github.com/avasapollo/payment-gateway/web/proto/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ v1.PaymentGatewayServer = (*PaymentGatewayService)(nil)

type PaymentGatewayService struct {
	lgr     *logrus.Entry
	payment Payment
}

func NewPaymentGatewayService(payment Payment) *PaymentGatewayService {
	return &PaymentGatewayService{
		lgr:     logrus.WithField("pkg", "web"),
		payment: payment,
	}
}

func (p *PaymentGatewayService) Health(ctx context.Context, e *empty.Empty) (*v1.HealthResp, error) {
	resp := &v1.HealthResp{
		Status: "UP",
	}
	return resp, nil
}

func (p *PaymentGatewayService) validateAuthorize(ctx context.Context, req *v1.AuthorizeReq) error {
	switch {
	case req.Card == nil:
		return status.Error(codes.InvalidArgument, "card must be specified")
	case req.Amount == nil:
		return status.Error(codes.InvalidArgument, "amount must be specified")
	}
	return nil
}

func (p *PaymentGatewayService) Authorize(ctx context.Context, req *v1.AuthorizeReq) (*v1.AuthorizeResp, error) {
	if err := p.validateAuthorize(ctx, req); err != nil {
		return nil, err
	}

	c, err := currency.ParseISO(req.Amount.Currency)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "amount currency must be respect ISO 4217. example EUR,GBP")
	}
	authReq := &payments.AuthorizeReq{
		Card: &payments.Card{
			Name:        req.Card.Name,
			CardNumber:  req.Card.CardNumber,
			ExpireMonth: req.Card.ExpireMonth,
			ExpireYear:  req.Card.ExpireYear,
			CVV:         req.Card.Cvv,
		},
		Amount:   req.Amount.Value,
		Currency: c,
	}

	tr, err := p.payment.Authorize(ctx, authReq)
	if err != nil {
		switch {
		case errors.Is(err, payments.ErrValidation):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, payments.ErrAuthFailed):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	resp := &v1.AuthorizeResp{
		Result:          "ok",
		AuthorizationId: tr.AuthorizationID,
		Amount: &v1.Amount{
			Value:    tr.Amount.Value,
			Currency: tr.Amount.Currency.String(),
		},
	}
	return resp, nil
}

func (p *PaymentGatewayService) Void(ctx context.Context, req *v1.VoidReq) (*v1.AmountResp, error) {
	voidReq := &payments.VoidReq{
		AuthorizationID: req.AuthorizationId,
	}
	tr, err := p.payment.Void(ctx, voidReq)
	if err != nil {
		switch {
		case errors.Is(err, payments.ErrValidation):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, payments.ErrVoidFailed):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	resp := &v1.AmountResp{
		Result: "ok",
		Amount: &v1.Amount{
			Value:    tr.Amount.Value,
			Currency: tr.Amount.Currency.String(),
		},
	}
	return resp, nil
}

func (p *PaymentGatewayService) Capture(ctx context.Context, req *v1.CaptureReq) (*v1.AmountResp, error) {
	captureReq := &payments.CaptureReq{
		AuthorizationID: req.AuthorizationId,
		Amount:          req.Amount,
	}

	tr, err := p.payment.Capture(ctx, captureReq)
	if err != nil {
		switch {
		case errors.Is(err, payments.ErrValidation):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, payments.ErrCaptureFailed):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	resp := &v1.AmountResp{
		Result: "ok",
		Amount: &v1.Amount{
			Value:    tr.CaptureAmount.Value,
			Currency: tr.CaptureAmount.Currency.String(),
		},
	}
	return resp, nil
}

func (p *PaymentGatewayService) Refund(ctx context.Context, req *v1.RefundReq) (*v1.AmountResp, error) {
	refundReq := &payments.RefundReq{
		AuthorizationID: req.AuthorizationId,
		Amount:          req.Amount,
	}

	tr, err := p.payment.Refund(ctx, refundReq)
	if err != nil {
		switch {
		case errors.Is(err, payments.ErrValidation):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, payments.ErrRefundFailed):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	resp := &v1.AmountResp{
		Result: "ok",
		Amount: &v1.Amount{
			Value:    tr.RefundAmount.Value,
			Currency: tr.RefundAmount.Currency.String(),
		},
	}
	return resp, nil
}
