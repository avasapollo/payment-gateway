package payments

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/currency"
)

type testSuite struct {
	payer *MockPayer
}

func newTestSuite(ctrl *gomock.Controller) *testSuite {
	return &testSuite{
		payer: NewMockPayer(ctrl),
	}
}

func TestService_Authorize(t *testing.T) {
	tr := &Transaction{
		AuthID: "transaction_id_1",
		Status: Authorized,
		Amount: &Amount{
			Value:    100,
			Currency: currency.EUR,
		},
		CurrentAmount: &Amount{
			Value:    0,
			Currency: currency.EUR,
		},
		CreatedAt:  time.Now().UTC(),
		CardNumber: "4242424242424242",
	}
	type args struct {
		ctx context.Context
		req *AuthorizeReq
	}
	tests := []struct {
		name string
		init func(req *AuthorizeReq, suite *testSuite)
		args args
		want func(got *Transaction, err error)
	}{
		{
			name: "error validation",
			init: func(req *AuthorizeReq, suite *testSuite) {
			},
			args: args{
				ctx: context.Background(),
				req: &AuthorizeReq{
					Card: &Card{
						Name:        "",
						CardNumber:  "4242424242424242",
						ExpireMonth: "12",
						ExpireYear:  "2052",
						CVV:         "211",
					},
					Amount:   100,
					Currency: currency.EUR,
				},
			},
			want: func(got *Transaction, err error) {
				require.True(t, errors.Is(err, ErrCardName))
			},
		},
		{
			name: "ok",
			init: func(req *AuthorizeReq, suite *testSuite) {
				suite.payer.EXPECT().Authorize(gomock.Any(), req).Return(tr, nil)
			},
			args: args{
				ctx: context.Background(),
				req: &AuthorizeReq{
					Card: &Card{
						Name:        "Andrea Vasapollo",
						CardNumber:  "4242424242424242",
						ExpireMonth: "12",
						ExpireYear:  "2052",
						CVV:         "211",
					},
					Amount:   100,
					Currency: currency.EUR,
				},
			},
			want: func(got *Transaction, err error) {
				require.NoError(t, err)
				require.Equal(t, tr, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			suite := newTestSuite(ctrl)
			tt.init(tt.args.req, suite)

			svc := &Service{
				lgr:   logrus.WithField("pkg", "payments"),
				payer: suite.payer,
			}
			got, err := svc.Authorize(tt.args.ctx, tt.args.req)
			tt.want(got, err)
		})
	}
}

func TestService_Capture(t *testing.T) {
	tr := &Transaction{
		AuthID: "transaction_id_1",
		Status: Captured,
		Amount: &Amount{
			Value:    100,
			Currency: currency.EUR,
		},
		CurrentAmount: &Amount{
			Value:    10,
			Currency: currency.EUR,
		},
		CreatedAt:  time.Now().UTC(),
		CardNumber: "4242424242424242",
	}

	type args struct {
		ctx context.Context
		req *CaptureReq
	}
	tests := []struct {
		name string
		init func(req *CaptureReq, suite *testSuite)
		args args
		want func(got *Transaction, err error)
	}{
		{
			name: "error validation",
			init: func(req *CaptureReq, suite *testSuite) {
			},
			args: args{
				ctx: context.Background(),
				req: &CaptureReq{
					AuthID: "",
					Amount: 10,
				},
			},
			want: func(got *Transaction, err error) {
				require.True(t, errors.Is(err, ErrAuthIDEmpty))
			},
		},
		{
			name: "ok",
			init: func(req *CaptureReq, suite *testSuite) {
				suite.payer.EXPECT().Capture(gomock.Any(), req).Return(tr, nil)
			},
			args: args{
				ctx: context.Background(),
				req: &CaptureReq{
					AuthID: "transaction_id_1",
					Amount: 10,
				},
			},
			want: func(got *Transaction, err error) {
				require.NoError(t, err)
				require.Equal(t, tr, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			suite := newTestSuite(ctrl)
			tt.init(tt.args.req, suite)

			svc := &Service{
				lgr:   logrus.WithField("pkg", "payments"),
				payer: suite.payer,
			}
			got, err := svc.Capture(tt.args.ctx, tt.args.req)
			tt.want(got, err)
		})
	}
}

func TestService_Refund(t *testing.T) {
	tr := &Transaction{
		AuthID: "transaction_id_1",
		Status: Refunded,
		Amount: &Amount{
			Value:    100,
			Currency: currency.EUR,
		},
		CurrentAmount: &Amount{
			Value:    0,
			Currency: currency.EUR,
		},
		CreatedAt:  time.Now().UTC(),
		CardNumber: "4242424242424242",
	}
	type args struct {
		ctx context.Context
		req *RefundReq
	}
	tests := []struct {
		name string
		init func(req *RefundReq, suite *testSuite)
		args args
		want func(got *Transaction, err error)
	}{
		{
			name: "error validation",
			init: func(req *RefundReq, suite *testSuite) {

			},
			args: args{
				ctx: context.Background(),
				req: &RefundReq{
					AuthID: "",
					Amount: 10,
				},
			},
			want: func(got *Transaction, err error) {
				require.True(t, errors.Is(err, ErrAuthIDEmpty))
			},
		},
		{
			name: "ok",
			init: func(req *RefundReq, suite *testSuite) {
				suite.payer.EXPECT().Refund(gomock.Any(), req).Return(tr, nil)
			},
			args: args{
				ctx: context.Background(),
				req: &RefundReq{
					AuthID: "transaction_id_1",
					Amount: 10,
				},
			},
			want: func(got *Transaction, err error) {
				require.NoError(t, err)
				require.Equal(t, tr, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			suite := newTestSuite(ctrl)
			tt.init(tt.args.req, suite)

			svc := &Service{
				lgr:   logrus.WithField("pkg", "payments"),
				payer: suite.payer,
			}
			got, err := svc.Refund(tt.args.ctx, tt.args.req)
			tt.want(got, err)
		})
	}
}

func TestService_Void(t *testing.T) {
	tr := &Transaction{
		AuthID: "transaction_id_1",
		Status: Voided,
		Amount: &Amount{
			Value:    100,
			Currency: currency.EUR,
		},
		CurrentAmount: &Amount{
			Value:    100,
			Currency: currency.EUR,
		},
		CreatedAt:  time.Now().UTC(),
		CardNumber: "4242424242424242",
	}

	type args struct {
		ctx context.Context
		req *VoidReq
	}
	tests := []struct {
		name string
		init func(req *VoidReq, suite *testSuite)
		args args
		want func(got *Transaction, err error)
	}{
		{
			name: "error validation",
			init: func(req *VoidReq, suite *testSuite) {

			},
			args: args{
				ctx: context.Background(),
				req: &VoidReq{
					AuthID: "",
				},
			},
			want: func(got *Transaction, err error) {
				require.True(t, errors.Is(err, ErrAuthIDEmpty))
			},
		},
		{
			name: "ok",
			init: func(req *VoidReq, suite *testSuite) {
				suite.payer.EXPECT().Void(gomock.Any(), req).Return(tr, nil)
			},
			args: args{
				ctx: context.Background(),
				req: &VoidReq{
					AuthID: "transaction_id_1",
				},
			},
			want: func(got *Transaction, err error) {
				require.NoError(t, err)
				require.Equal(t, tr, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			suite := newTestSuite(ctrl)
			tt.init(tt.args.req, suite)

			svc := &Service{
				lgr:   logrus.WithField("pkg", "payments"),
				payer: suite.payer,
			}
			got, err := svc.Void(tt.args.ctx, tt.args.req)
			tt.want(got, err)
		})
	}
}
