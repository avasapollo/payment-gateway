package web

import (
	"reflect"
	"testing"

	"github.com/avasapollo/payment-gateway/payments"
)

func Test_toPaymentAuthorizeReq(t *testing.T) {
	type args struct {
		req *AuthorizeReq
	}
	tests := []struct {
		name    string
		args    args
		want    *payments.AuthorizeReq
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toPaymentAuthorizeReq(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("toPaymentAuthorizeReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPaymentAuthorizeReq() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toPaymentCaptureReq(t *testing.T) {
	t.Fatal("TODO")
	type args struct {
		req *CaptureReq
	}
	tests := []struct {
		name string
		args args
		want *payments.CaptureReq
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toPaymentCaptureReq(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPaymentCaptureReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toPaymentRefundReq(t *testing.T) {
	t.Fatal("TODO")

	type args struct {
		req *RefundReq
	}
	tests := []struct {
		name string
		args args
		want *payments.RefundReq
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toPaymentRefundReq(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPaymentRefundReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toPaymentVodReq(t *testing.T) {
	t.Fatal("TODO")

	type args struct {
		req *VoidReq
	}
	tests := []struct {
		name string
		args args
		want *payments.VoidReq
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toPaymentVodReq(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPaymentVodReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toCaptureResp(t *testing.T) {
	t.Fatal("TODO")
	type args struct {
		tr *payments.Transaction
	}
	tests := []struct {
		name string
		args args
		want *CaptureResp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCaptureResp(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toCaptureResp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toRefundResp(t *testing.T) {
	t.Fatal("TODO")

	type args struct {
		tr *payments.Transaction
	}
	tests := []struct {
		name string
		args args
		want *RefundResp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toRefundResp(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toRefundResp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toVoidResp(t *testing.T) {
	t.Fatal("TODO")

	type args struct {
		tr *payments.Transaction
	}
	tests := []struct {
		name string
		args args
		want *VoidResp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toVoidResp(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toVoidResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
