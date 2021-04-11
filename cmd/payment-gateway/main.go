package main

import (
	"github.com/avasapollo/payment-gateway/payer"
	"github.com/avasapollo/payment-gateway/payments"
	"github.com/avasapollo/payment-gateway/web"
	"github.com/sirupsen/logrus"
)

func main() {
	lgr := logrus.WithField("pkg", "main")
	storage, err := payer.New()
	if err != nil {
		lgr.Fatal(err)
	}
	paymentSvc := payments.New(storage)
	srv := web.NewServer(paymentSvc)
	srv.ListenServer(8080)
}
