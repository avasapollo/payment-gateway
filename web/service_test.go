package web

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/avasapollo/payment-gateway/payments"
	v1 "github.com/avasapollo/payment-gateway/web/proto/v1"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestPaymentGatewayService_Health(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	payment := NewMockPayment(ctrl)
	svc := NewPaymentGatewayService(payment)
	got, err := svc.Health(context.Background(), &empty.Empty{})
	require.NoError(t, err)
	require.Equal(t, "UP", got.Status)
}

func TestPaymentGatewayService_validateAuthorize(t *testing.T) {
	getAuthReq := func() *v1.AuthorizeReq {
		return &v1.AuthorizeReq{
			Card: &v1.Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2052",
				Cvv:         "123",
			},
			Amount: &v1.Amount{
				Value:    100,
				Currency: "EUR",
			},
		}
	}

	t.Run("error card nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq()
		req.Card = nil
		err := svc.validateAuthorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error amount nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq()
		req.Amount = nil
		err := svc.validateAuthorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestPaymentGatewayService_Authorize(t *testing.T) {
	getAuthReq := func() *v1.AuthorizeReq {
		return &v1.AuthorizeReq{
			Card: &v1.Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2052",
				Cvv:         "123",
			},
			Amount: &v1.Amount{
				Value:    100,
				Currency: "EUR",
			},
		}
	}

	t.Run("error card nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq()
		req.Card = nil
		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error parse currency", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq()
		req.Amount.Currency = "BBB"
		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getAuthReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, req.Card.CardNumber, r.Card.CardNumber)
			// TODO: add more check
			return nil, payments.ErrValidation
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getAuthReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, req.Card.CardNumber, r.Card.CardNumber)
			// TODO: add more check
			return nil, payments.ErrAuthFailed
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.PermissionDenied, status.Code(err))
	})

	t.Run("error internal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getAuthReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, req.Card.CardNumber, r.Card.CardNumber)
			// TODO: add more check
			return nil, errors.New("something wrong")
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getAuthReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, req.Card.CardNumber, r.Card.CardNumber)
			// TODO: add more check
			return &payments.Transaction{
				AuthorizationID: "authorization_id_1",
				CardNumber:      "4242424242424242",
				Status:          payments.Authorize,
				Amount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
				CaptureAmount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
				RefundAmount: &payments.Amount{
					Value:    0,
					Currency: currency.EUR,
				},
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}, nil
		})
		svc := NewPaymentGatewayService(payment)

		resp, err := svc.Authorize(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, "authorization_id_1", resp.AuthorizationId)
		require.Equal(t, req.Amount.Value, resp.Amount.Value)
		require.Equal(t, req.Amount.Currency, resp.Amount.Currency)
	})
}

func TestPaymentGatewayService_Void(t *testing.T) {
	getVoidReq := func() *v1.VoidReq {
		return &v1.VoidReq{
			AuthorizationId: "authorization_id_1",
		}
	}

	t.Run("error validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getVoidReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Void(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.VoidReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			return nil, payments.ErrValidation
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Void(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getVoidReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Void(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.VoidReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			return nil, payments.ErrVoidFailed
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Void(context.Background(), req)
		require.Equal(t, codes.PermissionDenied, status.Code(err))
	})

	t.Run("error internal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getVoidReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Void(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.VoidReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			return nil, errors.New("something wrong")
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Void(context.Background(), req)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getVoidReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Void(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.VoidReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			return &payments.Transaction{
				AuthorizationID: "authorization_id_1",
				CardNumber:      "4242424242424242",
				Status:          payments.Void,
				Amount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
				CaptureAmount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
				RefundAmount: &payments.Amount{
					Value:    0,
					Currency: currency.EUR,
				},
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}, nil
		})
		svc := NewPaymentGatewayService(payment)

		resp, err := svc.Void(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, float64(100), resp.Amount.Value)
		require.Equal(t, currency.EUR.String(), resp.Amount.Currency)
	})
}

func TestPaymentGatewayService_Capture(t *testing.T) {
	getCaptureReq := func() *v1.CaptureReq {
		return &v1.CaptureReq{
			AuthorizationId: "authorization_id_1",
			Amount:          100,
		}

	}
	t.Run("error validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getCaptureReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Capture(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.CaptureReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return nil, payments.ErrValidation
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Capture(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getCaptureReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Capture(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.CaptureReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return nil, payments.ErrCaptureFailed
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Capture(context.Background(), req)
		require.Equal(t, codes.PermissionDenied, status.Code(err))
	})

	t.Run("error internal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getCaptureReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Capture(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.CaptureReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return nil, errors.New("something wrong")
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Capture(context.Background(), req)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getCaptureReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Capture(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.CaptureReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return &payments.Transaction{
				AuthorizationID: "authorization_id_1",
				CardNumber:      "4242424242424242",
				Status:          payments.Capture,
				Amount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
				CaptureAmount: &payments.Amount{
					Value:    0,
					Currency: currency.EUR,
				},
				RefundAmount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}, nil
		})
		svc := NewPaymentGatewayService(payment)
		resp, err := svc.Capture(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, float64(0), resp.Amount.Value)
		require.Equal(t, currency.EUR.String(), resp.Amount.Currency)
	})
}

func TestPaymentGatewayService_Refund(t *testing.T) {
	getRefundReq := func() *v1.RefundReq {
		return &v1.RefundReq{
			AuthorizationId: "authorization_id_1",
			Amount:          100,
		}
	}

	t.Run("error validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getRefundReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Refund(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.RefundReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return nil, payments.ErrValidation
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Refund(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getRefundReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Refund(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.RefundReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return nil, payments.ErrRefundFailed
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Refund(context.Background(), req)
		require.Equal(t, codes.PermissionDenied, status.Code(err))
	})

	t.Run("error internal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getRefundReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Refund(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.RefundReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return nil, errors.New("something wrong")
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Refund(context.Background(), req)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getRefundReq()
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Refund(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.RefundReq) (*payments.Transaction, error) {
			require.Equal(t, req.AuthorizationId, r.AuthorizationID)
			require.Equal(t, req.Amount, r.Amount)
			return &payments.Transaction{
				AuthorizationID: "authorization_id_1",
				CardNumber:      "4242424242424242",
				Status:          payments.Capture,
				Amount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
				CaptureAmount: &payments.Amount{
					Value:    0,
					Currency: currency.EUR,
				},
				RefundAmount: &payments.Amount{
					Value:    0,
					Currency: currency.EUR,
				},
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}, nil
		})
		svc := NewPaymentGatewayService(payment)
		resp, err := svc.Refund(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, float64(0), resp.Amount.Value)
		require.Equal(t, currency.EUR.String(), resp.Amount.Currency)
	})
}
