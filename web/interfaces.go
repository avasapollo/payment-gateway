package web

import (
	"context"

	"github.com/avasapollo/payment-gateway/payments"
)

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=web
type Payment interface {
	Authorize(ctx context.Context, req *payments.AuthorizeReq) (*payments.Transaction, error)
	Void(ctx context.Context, req *payments.VoidReq) (*payments.Transaction, error)
	Capture(ctx context.Context, req *payments.CaptureReq) (*payments.Transaction, error)
	Refund(ctx context.Context, req *payments.RefundReq) (*payments.Transaction, error)
}
