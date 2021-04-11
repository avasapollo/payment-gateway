// +build integration

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

func TestMongoPayer_Authorize(t *testing.T) {
	getReq := func() *payments.AuthorizeReq {
		return &payments.AuthorizeReq{
			Card: &payments.Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
	}
	ctx := context.Background()
	payer, err := New()
	require.NoError(t, err)
	t.Run("error card 4000000000000119", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)
		req := getReq()
		req.Card.CardNumber = "4000000000000119"
		_, err := payer.Authorize(ctx, req)
		require.True(t, errors.Is(err, payments.ErrAuthFailed))
	})

	t.Run("ok", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)
		req := getReq()
		tr, err := payer.Authorize(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Authorize, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, req.Amount, tr.Amount.Value)
		require.Equal(t, req.Amount, tr.CaptureAmount.Value)
		require.Equal(t, float64(0), tr.RefundAmount.Value)
	})
}

func TestMongoPayer_Void(t *testing.T) {
	getAuthReq := func() *payments.AuthorizeReq {
		return &payments.AuthorizeReq{
			Card: &payments.Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
	}
	ctx := context.Background()
	payer, err := New()
	require.NoError(t, err)

	t.Run("error authorization_id not found", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)
		req := &payments.VoidReq{
			AuthorizationID: uuid.NewString(),
		}
		_, err := payer.Void(ctx, req)
		require.True(t, errors.Is(err, payments.ErrVoidFailed))
	})

	t.Run("ok", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		req := &payments.VoidReq{
			AuthorizationID: tr.AuthorizationID,
		}
		tr, err = payer.Void(ctx, req)
		require.NoError(t, err)
		require.Equal(t, req.AuthorizationID, tr.AuthorizationID)
		require.Equal(t, payments.Void, tr.Status)
	})
}

func TestMongoPayer_Capture(t *testing.T) {
	getAuthReq := func() *payments.AuthorizeReq {
		return &payments.AuthorizeReq{
			Card: &payments.Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
	}
	ctx := context.Background()
	payer, err := New()
	require.NoError(t, err)

	t.Run("error capture failed not found authorization_id", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)
		req := &payments.CaptureReq{
			AuthorizationID: uuid.NewString(),
			Amount:          10,
		}
		_, err := payer.Capture(ctx, req)
		require.True(t, errors.Is(err, payments.ErrCaptureFailed))
	})

	t.Run("error capture failed limit exceed", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		req := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          120,
		}
		_, err = payer.Capture(ctx, req)
		require.True(t, errors.Is(err, payments.ErrCaptureFailed))
	})

	t.Run("error capture failed card 4000000000000259", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		reqAuth.Card.CardNumber = "4000000000000259"
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		req := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}
		_, err = payer.Capture(ctx, req)
		require.True(t, errors.Is(err, payments.ErrCaptureFailed))
	})

	t.Run("ok capture 10", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		req := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}
		tr, err = payer.Capture(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(10), tr.RefundAmount.Value)
	})

	t.Run("ok 2 capture 10", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		req := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}
		tr, err = payer.Capture(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(10), tr.RefundAmount.Value)

		tr, err = payer.Capture(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(80), tr.CaptureAmount.Value)
		require.Equal(t, float64(20), tr.RefundAmount.Value)
	})

	t.Run("ok 1 capture 100 and 1 error exceed", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		req := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          100,
		}
		tr, err = payer.Capture(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(0), tr.CaptureAmount.Value)
		require.Equal(t, float64(100), tr.RefundAmount.Value)

		_, err = payer.Capture(ctx, req)
		require.True(t, errors.Is(err, payments.ErrCaptureFailed))
	})
}

func TestMongoPayer_Refund(t *testing.T) {
	getAuthReq := func() *payments.AuthorizeReq {
		return &payments.AuthorizeReq{
			Card: &payments.Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4000000000000000",
				ExpireMonth: "12",
				ExpireYear:  "2024",
				CVV:         "1234",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
	}
	ctx := context.Background()
	payer, err := New()
	require.NoError(t, err)

	t.Run("error refund failed not found authorization_id", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)
		req := &payments.RefundReq{
			AuthorizationID: uuid.NewString(),
			Amount:          10,
		}
		_, err := payer.Refund(ctx, req)
		require.True(t, errors.Is(err, payments.ErrRefundFailed))
	})

	t.Run("error refund failed limit exceed", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		reqCapture := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}
		tr, err = payer.Capture(ctx, reqCapture)
		require.NoError(t, err)

		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(10), tr.RefundAmount.Value)

		req := &payments.RefundReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          20,
		}
		tr, err = payer.Refund(ctx, req)
		require.True(t, errors.Is(err, payments.ErrRefundFailed))
	})

	t.Run("error capture failed card 4000000000003238", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		reqAuth.Card.CardNumber = "4000000000003238"
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		reqCapture := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}

		tr, err = payer.Capture(ctx, reqCapture)
		require.NoError(t, err)

		req := &payments.RefundReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          20,
		}
		_, err = payer.Refund(ctx, req)
		require.True(t, errors.Is(payments.ErrRefundFailed, err))
	})

	t.Run("ok refund 5", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		reqCapture := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}
		tr, err = payer.Capture(ctx, reqCapture)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(10), tr.RefundAmount.Value)

		req := &payments.RefundReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          5,
		}
		tr, err = payer.Refund(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Refund, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(5), tr.RefundAmount.Value)
	})
	t.Run("ok 2 refund 5", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		reqCapture := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}
		tr, err = payer.Capture(ctx, reqCapture)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(10), tr.RefundAmount.Value)

		req := &payments.RefundReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          5,
		}
		tr, err = payer.Refund(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Refund, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(5), tr.RefundAmount.Value)

		tr, err = payer.Refund(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Refund, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(0), tr.RefundAmount.Value)
	})

	t.Run("ok 2 refund 5 and error refund filed on third", func(t *testing.T) {
		defer payer.transactions.Drop(ctx)
		require.NoError(t, err)

		reqAuth := getAuthReq()
		tr, err := payer.Authorize(ctx, reqAuth)
		require.NoError(t, err)

		reqCapture := &payments.CaptureReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          10,
		}
		tr, err = payer.Capture(ctx, reqCapture)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Capture, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(10), tr.RefundAmount.Value)

		req := &payments.RefundReq{
			AuthorizationID: tr.AuthorizationID,
			Amount:          5,
		}
		tr, err = payer.Refund(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Refund, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(5), tr.RefundAmount.Value)

		tr, err = payer.Refund(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, tr.AuthorizationID)
		require.Equal(t, payments.Refund, tr.Status)
		require.Equal(t, currency.GBP, tr.Amount.Currency)
		require.Equal(t, float64(100), tr.Amount.Value)
		require.Equal(t, float64(90), tr.CaptureAmount.Value)
		require.Equal(t, float64(0), tr.RefundAmount.Value)

		tr, err = payer.Refund(ctx, req)
		require.True(t, errors.Is(payments.ErrRefundFailed, err))
	})
}
