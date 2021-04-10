package payments

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/text/currency"
)

func TestAuthorizeReq_Validate(t *testing.T) {
	getReq := func() *AuthorizeReq {
		return &AuthorizeReq{
			Card: &Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2052",
				CVV:         "123",
			},
			Amount:   100,
			Currency: currency.GBP,
		}
	}
	tests := []struct {
		name string
		init func(req *AuthorizeReq)
		want func(err error)
	}{
		{
			name: "error card nil",
			init: func(req *AuthorizeReq) {
				req.Card = nil
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrCardNil))
			},
		},
		{
			name: "error card nil",
			init: func(req *AuthorizeReq) {
				req.Card.Name = ""
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrCardName))
			},
		},
		{
			name: "error card invalid",
			init: func(req *AuthorizeReq) {
				req.Card.CardNumber = ""
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrCardNotValid))
			},
		},
		{
			name: "error amount not valid",
			init: func(req *AuthorizeReq) {
				req.Amount = -10
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrAmountNotValid))
			},
		},
		{
			name: "ok",
			init: func(req *AuthorizeReq) {
			},
			want: func(err error) {
				require.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := getReq()
			tt.init(req)
			err := req.Validate()
			tt.want(err)
		})
	}
}

func TestVoidReq_Validate(t *testing.T) {
	getReq := func() *VoidReq {
		return &VoidReq{
			AuthorizationID: "transaction_id_1",
		}
	}
	tests := []struct {
		name string
		init func(req *VoidReq)
		want func(err error)
	}{
		{
			name: "error auth_id empty",
			init: func(req *VoidReq) {
				req.AuthorizationID = ""
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrAuthIDEmpty))
			},
		},
		{
			name: "ok",
			init: func(req *VoidReq) {
			},
			want: func(err error) {
				require.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := getReq()
			tt.init(r)
			err := r.Validate()
			tt.want(err)
		})
	}
}

func TestRefundReq_Validate(t *testing.T) {
	getReq := func() *RefundReq {
		return &RefundReq{
			AuthorizationID: "transaction_id_1",
			Amount:          10,
		}
	}
	tests := []struct {
		name string
		init func(req *RefundReq)
		want func(err error)
	}{
		{
			name: "error auth_id empty",
			init: func(req *RefundReq) {
				req.AuthorizationID = ""
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrAuthIDEmpty))
			},
		},
		{
			name: "error amount < 0",
			init: func(req *RefundReq) {
				req.Amount = -10
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrAmountNotValid))
			},
		},
		{
			name: "ok",
			init: func(req *RefundReq) {
			},
			want: func(err error) {
				require.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := getReq()
			tt.init(r)
			err := r.Validate()
			tt.want(err)
		})
	}
}

func TestCaptureReq_Validate(t *testing.T) {
	getReq := func() *CaptureReq {
		return &CaptureReq{
			AuthorizationID: "transaction_id_1",
			Amount:          10,
		}
	}
	tests := []struct {
		name string
		init func(req *CaptureReq)
		want func(err error)
	}{
		{
			name: "error auth_id empty",
			init: func(req *CaptureReq) {
				req.AuthorizationID = ""
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrAuthIDEmpty))
			},
		},
		{
			name: "error amount < 0",
			init: func(req *CaptureReq) {
				req.Amount = -10
			},
			want: func(err error) {
				require.True(t, errors.Is(err, ErrAmountNotValid))
			},
		},
		{
			name: "ok",
			init: func(req *CaptureReq) {
			},
			want: func(err error) {
				require.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := getReq()
			tt.init(r)
			err := r.Validate()
			tt.want(err)
		})
	}
}
