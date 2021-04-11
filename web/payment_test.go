package web

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/avasapollo/payment-gateway/payments"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/currency"
)

func TestPaymentApi_Authorize(t *testing.T) {
	t.Run("error 400 wrong json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: NewMockPayment(ctrl),
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/v1/authorize", bytes.NewBufferString("wrong json payload"))
		p.Authorize(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 400 currency not valid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: NewMockPayment(ctrl),
		}
		w := httptest.NewRecorder()
		req := &AuthorizeReq{
			Card: Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2020",
				CVV:         "211",
			},
			Amount: Amount{
				Value:    100,
				Currency: "BBB",
			},
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/authorize", bytes.NewBuffer(b))
		p.Authorize(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 400 validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, "4242424242424242", req.Card.CardNumber)
			// TODO check all request details
			return nil, payments.ErrValidation
		})
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &AuthorizeReq{
			Card: Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2020",
				CVV:         "211",
			},
			Amount: Amount{
				Value:    100,
				Currency: "EUR",
			},
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/authorize", bytes.NewBuffer(b))
		p.Authorize(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 500 auth failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, "4242424242424242", req.Card.CardNumber)
			// TODO check all request details
			return nil, payments.ErrAuthFailed
		})
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &AuthorizeReq{
			Card: Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2020",
				CVV:         "211",
			},
			Amount: Amount{
				Value:    100,
				Currency: "EUR",
			},
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/authorize", bytes.NewBuffer(b))
		p.Authorize(w, r)
		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("error 500 not catch error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, "4242424242424242", req.Card.CardNumber)
			// TODO check all request details
			return nil, errors.New("something wrong")
		})
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &AuthorizeReq{
			Card: Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2020",
				CVV:         "211",
			},
			Amount: Amount{
				Value:    100,
				Currency: "EUR",
			},
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/authorize", bytes.NewBuffer(b))
		p.Authorize(w, r)
		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("ok 200", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Authorize(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, req *payments.AuthorizeReq) (*payments.Transaction, error) {
			require.Equal(t, "4242424242424242", req.Card.CardNumber)
			// TODO check all request details
			tr := &payments.Transaction{
				AuthorizationID: uuid.NewString(),
				Status:          payments.Authorize,
				Amount: &payments.Amount{
					Value:    req.Amount,
					Currency: req.Currency,
				},
				CurrentAmount: &payments.Amount{
					Value:    0,
					Currency: req.Currency,
				},
				CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
				CardNumber: req.Card.CardNumber,
			}
			return tr, nil
		})
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &AuthorizeReq{
			Card: Card{
				Name:        "Andrea Vasapollo",
				CardNumber:  "4242424242424242",
				ExpireMonth: "12",
				ExpireYear:  "2020",
				CVV:         "211",
			},
			Amount: Amount{
				Value:    100,
				Currency: "EUR",
			},
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/authorize", bytes.NewBuffer(b))
		p.Authorize(w, r)
		require.Equal(t, http.StatusOK, w.Code)
		got := new(AuthorizeResp)
		err = json.NewDecoder(w.Body).Decode(got)
		require.NoError(t, err)
		require.NotEmpty(t, got.AuthorizationID)
		require.Equal(t, float64(100), got.Amount.Value)
		require.Equal(t, "EUR", got.Amount.Currency)
	})
}

func TestPaymentApi_Void(t *testing.T) {
	t.Run("error 400 wrong json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: NewMockPayment(ctrl),
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/v1/void", bytes.NewBufferString("wrong json payload"))
		p.Void(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 400 validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Void(gomock.Any(), &payments.VoidReq{AuthorizationID: "id_1"}).Return(nil, payments.ErrValidation)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &VoidReq{
			AuthorizationID: "id_1",
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/void", bytes.NewBuffer(b))
		p.Void(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 404 authorization_id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Void(gomock.Any(), &payments.VoidReq{AuthorizationID: "id_1"}).Return(nil, payments.ErrAutIDNotFound)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &VoidReq{
			AuthorizationID: "id_1",
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/void", bytes.NewBuffer(b))
		p.Void(w, r)
		require.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("error 500 generic error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Void(gomock.Any(), &payments.VoidReq{AuthorizationID: "id_1"}).Return(nil, errors.New("something wrong"))
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &VoidReq{
			AuthorizationID: "id_1",
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/void", bytes.NewBuffer(b))
		p.Void(w, r)
		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		tr := &payments.Transaction{
			AuthorizationID: uuid.NewString(),
			Status:          payments.Authorize,
			Amount: &payments.Amount{
				Value:    100,
				Currency: currency.EUR,
			},
			CurrentAmount: &payments.Amount{
				Value:    0,
				Currency: currency.EUR,
			},
			CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
			CardNumber: "xxxxxxxxxxxxxx",
		}
		mockPayment.EXPECT().Void(gomock.Any(), &payments.VoidReq{AuthorizationID: "id_1"}).Return(tr, nil)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &VoidReq{
			AuthorizationID: "id_1",
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/void", bytes.NewBuffer(b))
		p.Void(w, r)
		require.Equal(t, http.StatusOK, w.Code)
		got := new(VoidResp)
		err = json.NewDecoder(w.Body).Decode(got)
		require.NoError(t, err)
		require.Equal(t, float64(100), got.Amount.Value)
		require.Equal(t, "EUR", got.Amount.Currency)
	})
}

func TestPaymentApi_Capture(t *testing.T) {
	t.Run("error 400 wrong json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: NewMockPayment(ctrl),
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/v1/capture", bytes.NewBufferString("wrong json payload"))
		p.Capture(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 400 validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Capture(gomock.Any(), &payments.CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, payments.ErrValidation)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/capture", bytes.NewBuffer(b))
		p.Capture(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 404 authorization_id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Capture(gomock.Any(), &payments.CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, payments.ErrAutIDNotFound)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/capture", bytes.NewBuffer(b))
		p.Capture(w, r)
		require.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("error 500 capture failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Capture(gomock.Any(), &payments.CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, payments.ErrCaptureFailed)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/capture", bytes.NewBuffer(b))
		p.Capture(w, r)
		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("error 500 generic error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Capture(gomock.Any(), &payments.CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, errors.New("something wrong"))
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/capture", bytes.NewBuffer(b))
		p.Capture(w, r)
		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tr := &payments.Transaction{
			AuthorizationID: uuid.NewString(),
			Status:          payments.Capture,
			Amount: &payments.Amount{
				Value:    100,
				Currency: currency.EUR,
			},
			CurrentAmount: &payments.Amount{
				Value:    10,
				Currency: currency.EUR,
			},
			CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
			CardNumber: "xxxxxxxxxxxxxx",
		}
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Capture(gomock.Any(), &payments.CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(tr, nil)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &CaptureReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/capture", bytes.NewBuffer(b))
		p.Capture(w, r)
		require.Equal(t, http.StatusOK, w.Code)
		got := new(CaptureResp)
		err = json.NewDecoder(w.Body).Decode(got)
		require.NoError(t, err)
		require.Equal(t, float64(90), got.Amount.Value)
		require.Equal(t, "EUR", got.Amount.Currency)
	})
}

func TestPaymentApi_Refund(t *testing.T) {
	t.Run("error 400 wrong json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: NewMockPayment(ctrl),
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/v1/refund", bytes.NewBufferString("wrong json payload"))
		p.Refund(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 400 validation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Refund(gomock.Any(), &payments.RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, payments.ErrValidation)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/refund", bytes.NewBuffer(b))
		p.Refund(w, r)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("error 404 authorization_id not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Refund(gomock.Any(), &payments.RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, payments.ErrAutIDNotFound)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/refund", bytes.NewBuffer(b))
		p.Refund(w, r)
		require.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("error 500 refund failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Refund(gomock.Any(), &payments.RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, payments.ErrRefundFailed)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/refund", bytes.NewBuffer(b))
		p.Refund(w, r)
		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("error 500 generic error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Refund(gomock.Any(), &payments.RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(nil, errors.New("something wrong"))
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/capture", bytes.NewBuffer(b))
		p.Refund(w, r)
		require.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		tr := &payments.Transaction{
			AuthorizationID: uuid.NewString(),
			Status:          payments.Refund,
			Amount: &payments.Amount{
				Value:    100,
				Currency: currency.EUR,
			},
			CurrentAmount: &payments.Amount{
				Value:    10,
				Currency: currency.EUR,
			},
			CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
			CardNumber: "xxxxxxxxxxxxxx",
		}
		mockPayment := NewMockPayment(ctrl)
		mockPayment.EXPECT().Refund(gomock.Any(), &payments.RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}).Return(tr, nil)
		p := &PaymentApi{
			lgr:     logrus.WithField("pkg", "web"),
			payment: mockPayment,
		}
		w := httptest.NewRecorder()
		req := &RefundReq{
			AuthorizationID: "id_1",
			Amount:          10,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		r := httptest.NewRequest(http.MethodPost, "/v1/refund", bytes.NewBuffer(b))
		p.Refund(w, r)
		require.Equal(t, http.StatusOK, w.Code)
		got := new(RefundResp)
		err = json.NewDecoder(w.Body).Decode(got)
		require.NoError(t, err)
		require.Equal(t, float64(10), got.Amount.Value)
		require.Equal(t, "EUR", got.Amount.Currency)
	})
}
