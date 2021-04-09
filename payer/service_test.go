package payer

import (
	"context"
	"errors"
	"testing"

	"github.com/avasapollo/payment-gateway/payments"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/currency"
)

func TestService_Authorize(t *testing.T) {
	type args struct {
		ctx context.Context
		req *payments.AuthorizeReq
	}
	tests := []struct {
		name    string
		init func(svc *Service)
		args    args
		want func(got *payments.Transaction,err error)
	}{
		{
			name: "error authorization failure",
			init: func(svc *Service) {

			},
			args: args{
				ctx: context.Background(),
				req: &payments.AuthorizeReq{
					Card:     &payments.Card{
						CardNumber:  "4000000000000119",
					},
					Amount:   100,
					Currency: currency.GBP,
				},
			},
			want: func(got *payments.Transaction, err error) {
				require.True(t,errors.Is(err,payments.ErrAuthFailed))
			},
		},
		{
			name: "ok",
			init: func(svc *Service) {

			},
			args: args{
				ctx: context.Background(),
				req: &payments.AuthorizeReq{
					Card:     &payments.Card{
						Name:        "Andrea",
						CardNumber:  "4000000000000000",
						ExpireMonth: "12",
						ExpireYear:  "2024",
						CVV:         "1234",
					},
					Amount:   100,
					Currency: currency.GBP,
				},
			},
			want: func(got *payments.Transaction, err error) {
				require.NoError(t,err)
				require.NotEmpty(t,got.AuthID)
				require.Equal(t,payments.Authorized,got.Status)
				require.Equal(t,float64(100),got.Amount)
				require.Equal(t,"4000000000000000",got.CardNumber)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := New()
			tt.init(svc)
			got, err := svc.Authorize(tt.args.ctx, tt.args.req)
			tt.want(got,err)
		})
	}
}


func TestService_Void(t *testing.T) {
	ctx := context.Background()

	t.Run("error not found", func(t *testing.T) {
		svc := New()
		req := &payments.VoidReq{
			AuthID: uuid.NewString(),
		}
		_, err := svc.Void(ctx, req)
		require.True(t,errors.Is(err,payments.ErrAutIDNotFound))
	})

	t.Run("error status not authorized", func(t *testing.T) {
		svc := New()

		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)
		tr.Status = payments.Refunded
		svc.transactions.Store(tr.AuthID,*tr)

		voidReq := &payments.VoidReq{
			AuthID:tr.AuthID,
		}

		tr, err = svc.Void(ctx, voidReq)
		require.True(t,errors.Is(err,payments.ErrVoidFailed))
	})

	t.Run("ok", func(t *testing.T) {
		svc := New()

		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		voidReq := &payments.VoidReq{
			AuthID:tr.AuthID,
		}
		tr, err = svc.Void(ctx, voidReq)
		require.NoError(t,err)
		require.Equal(t, voidReq.AuthID,tr.AuthID)
		require.Equal(t, payments.Voided,tr.Status)
	})
}


func TestService_Capture(t *testing.T) {
	ctx := context.Background()

	t.Run("error not found", func(t *testing.T) {
		svc := New()
		req := &payments.CaptureReq{
			AuthID: uuid.NewString(),
			Amount: 10,
		}
		_, err := svc.Capture(ctx, req)
		require.True(t,errors.Is(err,payments.ErrAutIDNotFound))
	})

	t.Run("error capture failed", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000259",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}
		tr, err = svc.Capture(ctx, captureReq)
		require.True(t,errors.Is(err,payments.ErrCaptureFailed))
	})

	t.Run("error capture failed", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		tr.Status = payments.Voided
		svc.transactions.Store(tr.AuthID,*tr)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.True(t,errors.Is(err,payments.ErrCaptureFailed))
	})

	t.Run("error capture limit exceeded", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 110,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.True(t,errors.Is(err,payments.ErrCaptureLimitExceeded))
	})

	t.Run("ok", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.NoError(t,err)
		require.Equal(t, payments.Captured,tr.Status)
		require.Equal(t, float64(100),tr.Amount)
		require.Equal(t, float64(10),tr.CheckedAmount)
	})

	t.Run("ok 2 times", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.NoError(t,err)
		require.Equal(t, payments.Captured,tr.Status)
		require.Equal(t, float64(100),tr.Amount)
		require.Equal(t, float64(10),tr.CheckedAmount)

		captureReq = &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.NoError(t,err)
		require.Equal(t, payments.Captured,tr.Status)
		require.Equal(t, float64(100),tr.Amount)
		require.Equal(t, float64(20),tr.CheckedAmount)
	})
}

func TestService_Refund(t *testing.T) {
	ctx := context.Background()

	t.Run("error not found", func(t *testing.T) {
		svc := New()
		req := &payments.RefundReq{
			AuthID: uuid.NewString(),
			Amount: 10,
		}
		_, err := svc.Refund(ctx, req)
		require.True(t,errors.Is(err,payments.ErrAutIDNotFound))
	})

	t.Run("error refund failed", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000003238",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		refundReq := &payments.RefundReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}
		tr, err = svc.Refund(ctx, refundReq)
		require.True(t,errors.Is(err,payments.ErrRefundFailed))
	})

	t.Run("error refund failed", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		tr.Status = payments.Voided
		svc.transactions.Store(tr.AuthID,*tr)

		refundReq := &payments.RefundReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Refund(ctx, refundReq)
		require.True(t,errors.Is(err,payments.ErrRefundFailed))
	})

	t.Run("error refund limit exceeded", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.NoError(t,err)
		require.Equal(t, float64(10),tr.CheckedAmount)

		refundReq := &payments.RefundReq{
			AuthID: tr.AuthID,
			Amount: 20,
		}
		tr,err =svc.Refund(ctx,refundReq)
		require.True(t, errors.Is(err,payments.ErrRefundLimitExceeded))
	})

	t.Run("ok", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.NoError(t,err)
		require.Equal(t, float64(10),tr.CheckedAmount)

		refundReq := &payments.RefundReq{
			AuthID: tr.AuthID,
			Amount: 9,
		}
		tr,err =svc.Refund(ctx,refundReq)
		require.NoError(t, err)
		require.Equal(t, payments.Refunded,tr.Status)
		require.Equal(t, float64(100),tr.Amount)
		require.Equal(t, float64(1),tr.CheckedAmount)
	})

	t.Run("ok 2 times", func(t *testing.T) {
		svc := New()
		authReq := &payments.AuthorizeReq{
			Card:     &payments.Card{
				Name:        "Andrea",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
		tr,err := svc.Authorize(ctx,authReq)
		require.NoError(t,err)

		captureReq := &payments.CaptureReq{
			AuthID:tr.AuthID,
			Amount: 10,
		}

		tr, err = svc.Capture(ctx, captureReq)
		require.NoError(t,err)
		require.Equal(t, float64(10),tr.CheckedAmount)

		refundReq := &payments.RefundReq{
			AuthID: tr.AuthID,
			Amount: 9,
		}
		tr,err =svc.Refund(ctx,refundReq)
		require.NoError(t, err)
		require.Equal(t, payments.Refunded,tr.Status)
		require.Equal(t, float64(100),tr.Amount)
		require.Equal(t, float64(1),tr.CheckedAmount)

		refundReq = &payments.RefundReq{
			AuthID: tr.AuthID,
			Amount: 1,
		}
		tr,err =svc.Refund(ctx,refundReq)
		require.NoError(t, err)
		require.Equal(t, payments.Refunded,tr.Status)
		require.Equal(t, float64(100),tr.Amount)
		require.Equal(t, float64(0),tr.CheckedAmount)
	})
}
