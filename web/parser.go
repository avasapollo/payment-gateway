package web

import (
	"github.com/avasapollo/payment-gateway/payments"
	"golang.org/x/text/currency"
)

func toPaymentAuthorizeReq(req *AuthorizeReq) (*payments.AuthorizeReq, error) {
	cu, err := currency.ParseISO(req.Amount.Currency)
	if err != nil {
		return nil, err
	}
	r := &payments.AuthorizeReq{
		Card: &payments.Card{
			Name:        req.Card.Name,
			CardNumber:  req.Card.CardNumber,
			ExpireMonth: req.Card.ExpireMonth,
			ExpireYear:  req.Card.ExpireYear,
			CVV:         req.Card.CVV,
		},
		Amount:   req.Amount.Value,
		Currency: cu,
	}
	return r, nil
}

func toPaymentVoidReq(req *VoidReq) *payments.VoidReq {
	return &payments.VoidReq{
		AuthorizationID: req.AuthorizationID,
	}
}

func toPaymentCaptureReq(req *CaptureReq) *payments.CaptureReq {
	return &payments.CaptureReq{
		AuthorizationID: req.AuthorizationID,
		Amount:          req.Amount,
	}
}

func toPaymentRefundReq(req *RefundReq) *payments.RefundReq {
	return &payments.RefundReq{
		AuthorizationID: req.AuthorizationID,
		Amount:          req.Amount,
	}
}

func toAuthorizeResp(tr *payments.Transaction) *AuthorizeResp {
	return &AuthorizeResp{
		AuthorizationID: tr.AuthorizationID,
		Amount: &Amount{
			Value:    tr.Amount.Value,
			Currency: tr.Amount.Currency.String(),
		},
	}
}

func toVoidResp(tr *payments.Transaction) *VoidResp {
	return &VoidResp{
		Amount: &Amount{
			Value:    tr.Amount.Value,
			Currency: tr.Amount.Currency.String(),
		},
	}
}

func toCaptureResp(tr *payments.Transaction) *CaptureResp {
	return &CaptureResp{
		Amount: &Amount{
			Value:    tr.CaptureAmount.Value,
			Currency: tr.Amount.Currency.String(),
		},
	}
}

func toRefundResp(tr *payments.Transaction) *RefundResp {
	return &RefundResp{
		Amount: &Amount{
			Value:    tr.RefundAmount.Value,
			Currency: tr.Amount.Currency.String(),
		},
	}
}
