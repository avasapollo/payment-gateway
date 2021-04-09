package web

import (
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestPaymentApi_Authorize(t *testing.T) {
	t.Fatal("TODO")
	type fields struct {
		lgr     *logrus.Entry
		payment Payment
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &PaymentApi{
				lgr:     tt.fields.lgr,
				payment: tt.fields.payment,
			}
		})
	}
}

func TestPaymentApi_Capture(t *testing.T) {
	t.Fatal("TODO")

	type fields struct {
		lgr     *logrus.Entry
		payment Payment
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &PaymentApi{
				lgr:     tt.fields.lgr,
				payment: tt.fields.payment,
			}
		})
	}
}

func TestPaymentApi_Refund(t *testing.T) {
	t.Fatal("TODO")

	type fields struct {
		lgr     *logrus.Entry
		payment Payment
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &PaymentApi{
				lgr:     tt.fields.lgr,
				payment: tt.fields.payment,
			}
		})
	}
}

func TestPaymentApi_Void(t *testing.T) {
	t.Fatal("TODO")

	type fields struct {
		lgr     *logrus.Entry
		payment Payment
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &PaymentApi{
				lgr:     tt.fields.lgr,
				payment: tt.fields.payment,
			}
		})
	}
}
