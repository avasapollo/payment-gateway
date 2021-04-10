package web

import (
	"reflect"
	"testing"

	"github.com/avasapollo/payment-gateway/payments"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/currency"
)

func Test_toPaymentAuthorizeReq(t *testing.T) {
	type args struct {
		req *AuthorizeReq
	}
	tests := []struct {
		name string
		args args
		want func(got *payments.AuthorizeReq, err error)
	}{
		{
			name: "error not found currency",
			args: args{
				req: &AuthorizeReq{
					Card: Card{
						Name:        "Andrea Vasapollo",
						CardNumber:  "card_id_1",
						ExpireMonth: "12",
						ExpireYear:  "2020",
						CVV:         "111",
					},
					Amount: Amount{
						Value:    100,
						Currency: "BBB",
					},
				},
			},
			want: func(got *payments.AuthorizeReq, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "ok",
			args: args{
				req: &AuthorizeReq{
					Card: Card{
						Name:        "Andrea Vasapollo",
						CardNumber:  "card_id_1",
						ExpireMonth: "12",
						ExpireYear:  "2020",
						CVV:         "111",
					},
					Amount: Amount{
						Value:    100,
						Currency: "EUR",
					},
				},
			},
			want: func(got *payments.AuthorizeReq, err error) {
				require.NoError(t, err)
				want := &payments.AuthorizeReq{
					Card: &payments.Card{
						Name:        "Andrea Vasapollo",
						CardNumber:  "card_id_1",
						ExpireMonth: "12",
						ExpireYear:  "2020",
						CVV:         "111",
					},
					Amount:   100,
					Currency: currency.EUR,
				}
				require.Equal(t, want, got)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toPaymentAuthorizeReq(tt.args.req)
			tt.want(got, err)
		})
	}
}

func Test_toPaymentCaptureReq(t *testing.T) {
	type args struct {
		req *CaptureReq
	}
	tests := []struct {
		name string
		args args
		want *payments.CaptureReq
	}{
		{
			name: "ok",
			args: args{
				req: &CaptureReq{
					AuthorizationID: "id_1",
					Amount:          100,
				},
			},
			want: &payments.CaptureReq{
				AuthorizationID: "id_1",
				Amount:          100,
			},
		},
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
	type args struct {
		req *RefundReq
	}
	tests := []struct {
		name string
		args args
		want *payments.RefundReq
	}{
		{
			name: "ok",
			args: args{
				req: &RefundReq{
					AuthorizationID: "id_1",
					Amount:          100,
				},
			},
			want: &payments.RefundReq{
				AuthorizationID: "id_1",
				Amount:          100,
			},
		},
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
	type args struct {
		req *VoidReq
	}
	tests := []struct {
		name string
		args args
		want *payments.VoidReq
	}{
		{
			name: "ok",
			args: args{
				req: &VoidReq{
					AuthorizationID: "id_1",
				},
			},
			want: &payments.VoidReq{
				AuthorizationID: "id_1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toPaymentVoidReq(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPaymentVodReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toCaptureResp(t *testing.T) {
	type args struct {
		tr *payments.Transaction
	}
	tests := []struct {
		name string
		args args
		want *CaptureResp
	}{
		{
			name: "ok",
			args: args{
				tr: &payments.Transaction{
					AuthorizationID: "id_1",
					Amount: &payments.Amount{
						Value:    10,
						Currency: currency.EUR,
					},
					CurrentAmount: &payments.Amount{
						Value:    1,
						Currency: currency.EUR,
					},
				},
			},
			want: &CaptureResp{
				Amount: &Amount{
					Value:    9,
					Currency: currency.EUR.String(),
				},
			},
		},
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
	type args struct {
		tr *payments.Transaction
	}
	tests := []struct {
		name string
		args args
		want *RefundResp
	}{
		{
			name: "ok",
			args: args{
				tr: &payments.Transaction{
					AuthorizationID: "id_1",
					Amount: &payments.Amount{
						Value:    10,
						Currency: currency.EUR,
					},
					CurrentAmount: &payments.Amount{
						Value:    1,
						Currency: currency.EUR,
					},
				},
			},
			want: &RefundResp{
				Amount: &Amount{
					Value:    9,
					Currency: currency.EUR.String(),
				},
			},
		},
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
	type args struct {
		tr *payments.Transaction
	}
	tests := []struct {
		name string
		args args
		want *VoidResp
	}{
		{
			name: "ok",
			args: args{
				tr: &payments.Transaction{
					AuthorizationID: "id_1",
					Amount: &payments.Amount{
						Value:    10,
						Currency: currency.EUR,
					},
					CurrentAmount: &payments.Amount{
						Value:    0,
						Currency: currency.EUR,
					},
				},
			},
			want: &VoidResp{
				Amount: &Amount{
					Value:    10,
					Currency: currency.EUR.String(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toVoidResp(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toVoidResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
