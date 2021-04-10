package web

import (
	"errors"
	"net/http"

	"github.com/avasapollo/payment-gateway/payments"
	"github.com/sirupsen/logrus"
)

type PaymentApi struct {
	lgr *logrus.Entry
	payment Payment
}

func newPaymentApi(payment Payment) *PaymentApi  {
	return &PaymentApi{
		lgr:     logrus.WithField("pkg","web"),
		payment: payment,
	}
}

func(api *PaymentApi) Authorize(w http.ResponseWriter, r *http.Request) {
	req := new(AuthorizeReq)
	if err := UnmarshalRequest(r,req);err!= nil {
		WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1","wrong body"))
		return
	}
	paymentReq,err := toPaymentAuthorizeReq(req)
	if err!= nil {
		WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1",err.Error()))
		return
	}
	tr,err := api.payment.Authorize(r.Context(),paymentReq)
	if err!= nil {
		switch  {
		case  errors.Is(err,payments.ErrValidation):
			WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1",err.Error()))
			return
		case errors.Is(err,payments.ErrAuthFailed):
			WriteResponse(w,http.StatusInternalServerError,newErrResp("err_ie2",err.Error()))
			return
		default:
			WriteResponse(w,http.StatusInternalServerError,newErrResp("err_ie1",err.Error()))
			return
		}
	}
	WriteResponse(w,http.StatusOK, toAuthorizeResp(tr))
}

func(api *PaymentApi) Void(w http.ResponseWriter, r *http.Request) {
	req := new(VoidReq)
	if err := UnmarshalRequest(r,req);err!= nil {
		WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1","wrong body"))
		return
	}
	paymentReq := toPaymentVodReq(req)

	tr,err := api.payment.Void(r.Context(),paymentReq)
	if err!= nil {
		switch  {
		case  errors.Is(err,payments.ErrValidation):
			WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1",err.Error()))
			return
		case errors.Is(err,payments.ErrAutIDNotFound):
			WriteResponse(w,http.StatusNotFound,newErrResp("err_nf1",err.Error()))
			return
		default:
			WriteResponse(w,http.StatusInternalServerError,newErrResp("err_ie1",err.Error()))
			return
		}
	}
	WriteResponse(w,http.StatusOK, toVoidResp(tr))
}

func(api *PaymentApi) Capture(w http.ResponseWriter, r *http.Request) {
	req := new(CaptureReq)
	if err := UnmarshalRequest(r,req);err!= nil {
		WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1","wrong body"))
		return
	}
	paymentReq := toPaymentCaptureReq(req)

	tr,err := api.payment.Capture(r.Context(),paymentReq)
	if err!= nil {
		switch  {
		case  errors.Is(err,payments.ErrValidation):
			WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1",err.Error()))
			return
		case errors.Is(err,payments.ErrAutIDNotFound):
			WriteResponse(w,http.StatusNotFound,newErrResp("err_nf1",err.Error()))
			return
		case errors.Is(err,payments.ErrCaptureFailed):
			WriteResponse(w,http.StatusInternalServerError,newErrResp("err_ie2",err.Error()))
			return
		default:
			WriteResponse(w,http.StatusInternalServerError,newErrResp("err_ie1",err.Error()))
			return
		}
	}
	WriteResponse(w,http.StatusOK, toCaptureResp(tr))
}

func(api *PaymentApi) Refund(w http.ResponseWriter, r *http.Request) {
	req := new(RefundReq)
	if err := UnmarshalRequest(r,req);err!= nil {
		WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1","wrong body"))
		return
	}
	paymentReq := toPaymentRefundReq(req)

	tr,err := api.payment.Refund(r.Context(),paymentReq)
	if err!= nil {
		switch  {
		case  errors.Is(err,payments.ErrValidation):
			WriteResponse(w,http.StatusBadRequest,newErrResp("err_br1",err.Error()))
			return
		case errors.Is(err,payments.ErrAutIDNotFound):
			WriteResponse(w,http.StatusNotFound,newErrResp("err_nf1",err.Error()))
			return
		case errors.Is(err,payments.ErrRefundFailed):
			WriteResponse(w,http.StatusInternalServerError,newErrResp("err_ie2",err.Error()))
			return
		default:
			WriteResponse(w,http.StatusInternalServerError,newErrResp("err_ie1",err.Error()))
			return
		}
	}
	WriteResponse(w,http.StatusOK, toRefundResp(tr))
}