package main

import (
	"github.com/avasapollo/payment-gateway/payer"
	"github.com/avasapollo/payment-gateway/payments"
	"github.com/avasapollo/payment-gateway/web"
)

func main() {
	storage := payer.New()
	paymentSvc := payments.New(storage)
	srv := web.NewServer(paymentSvc)
	srv.ListenServer(8080)
}
