package web

import (
	"github.com/avasapollo/payment-gateway/payments"
	"golang.org/x/text/currency"
)

func toPaymentAuthorizeReq(req *AuthorizeReq) (*payments.AuthorizeReq,error)  {
	cu,err := currency.ParseISO(req.Currency)
	if err!= nil {
		return nil,err
	}
	return &payments.AuthorizeReq{
		Card:     &payments.Card{
			Name:        req.Card.Name,
			CardNumber:  req.Card.CardNumber,
			ExpireMonth: req.Card.ExpireMonth,
			ExpireYear:  req.Card.ExpireYear,
			CVV:        req.Card.CVV,
		},
		Amount:   req.Amount,
		Currency: cu,
	},nil
}

func toPaymentVodReq(req *VoidReq) *payments.VoidReq  {
	return &payments.VoidReq{
		AuthID: req.AuthorizationID,
	}
}

func toPaymentCaptureReq(req *CaptureReq) *payments.CaptureReq  {
	return &payments.CaptureReq{
		AuthID: req.AuthorizationID,
		Amount: req.Amount,
	}
}

func toPaymentRefundReq(req *RefundReq) *payments.RefundReq  {
	return &payments.RefundReq{
		AuthID: req.AuthorizationID,
		Amount: req.Amount,
	}
}

func toAuthorizeResp(tr *payments.Transaction) *AuthorizeResp  {
	return &AuthorizeResp{
		AuthorizationID: tr.AuthID,
		Amount:          tr.Amount,
		Currency:        tr.Currency.String(),
	}
}

func toVoidResp(tr *payments.Transaction) *VoidResp  {
	return &VoidResp{
		Amount:          tr.Amount,
		Currency:        tr.Currency.String(),
	}
}

func toCaptureResp(tr *payments.Transaction) *CaptureResp  {
	return &CaptureResp{
		Amount:          tr.Amount - tr.CheckedAmount,
		Currency:        tr.Currency.String(),
	}
}

func toRefundResp(tr *payments.Transaction) *RefundResp  {
	return &RefundResp{
		Amount:          tr.Amount - tr.CheckedAmount,
		Currency:        tr.Currency.String(),
	}
}