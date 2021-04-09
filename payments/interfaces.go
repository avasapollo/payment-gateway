package payments

import (
	"context"

)

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=payments
type Payer interface {
	Authorize(ctx context.Context,req *AuthorizeReq) (*Transaction,error)
	Void(ctx context.Context,req *VoidReq) (*Transaction,error)
	Capture(ctx context.Context,req *CaptureReq) (*Transaction,error)
	Refund(ctx context.Context,req *RefundReq) (*Transaction,error)
}


