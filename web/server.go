package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	lgr *logrus.Entry
	router *mux.Router
	paymentApi *PaymentApi
}

func NewServer(payment Payment) *Server  {
	return &Server{
		lgr:        logrus.WithField("pkg","web"),
		router:     mux.NewRouter(),
		paymentApi: newPaymentApi(payment),
	}
}

func (srv *Server) Health(w http.ResponseWriter, r *http.Request) {
	resp := &HealthResp{Status: "UP"}
	WriteResponse(w,http.StatusOK,resp)
}

func (srv *Server) ListenServer(port int)  {
	srv.router.HandleFunc("/health",srv.Health).Methods(http.MethodGet)
	srv.router.HandleFunc("/v1/authorize",srv.paymentApi.Authorize).Methods(http.MethodPost)
	srv.router.HandleFunc("/v1/void",srv.paymentApi.Void).Methods(http.MethodPost)
	srv.router.HandleFunc("/v1/capture",srv.paymentApi.Capture).Methods(http.MethodPost)
	srv.router.HandleFunc("/v1/refund",srv.paymentApi.Refund).Methods(http.MethodPost)
	srv.lgr.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), srv.router))
}
