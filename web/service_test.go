package web

import (
	"context"
	"testing"

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
	getAuthReq := &v1.AuthorizeReq{
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

	t.Run("error card nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq
		req.Card = nil
		err := svc.validateAuthorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error amount nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq
		req.Amount = nil
		err := svc.validateAuthorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestPaymentGatewayService_Authorize(t *testing.T) {
	getAuthReq := &v1.AuthorizeReq{
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

	t.Run("error card nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq
		req.Card = nil
		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error parse currency", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		payment := NewMockPayment(ctrl)
		svc := NewPaymentGatewayService(payment)
		req := getAuthReq
		req.Amount.Currency = "BBB"
		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("error validation", func(t *testing.T) {
		// TODO: add stuff here

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		req := getAuthReq
		payment := NewMockPayment(ctrl)
		payment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, r *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, req.Card.CardNumber, r.Card.CardNumber)
			// TODO: add more check
			return &payments.Transaction{
				AuthorizationID: "authorization_id_1",
				Amount: &payments.Amount{
					Value:    100,
					Currency: currency.EUR,
				},
			}, nil
		})
		svc := NewPaymentGatewayService(payment)

		_, err := svc.Authorize(context.Background(), req)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}
