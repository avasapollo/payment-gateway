package web

type HealthResp struct {
	Status string `json:"status"`
}

type Card struct {
	Name 	   string `json:"name"`
	CardNumber string `json:"card_number"`
	ExpireMonth string `json:"expire_month"`
	ExpireYear string `json:"expire_year"`
	CVV 		string `json:"cvv"`
}

type AuthorizeReq struct {
	Card *Card `json:"card"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
}

type VoidReq struct {
	AuthorizationID string `json:"authorization_id"`
}

type CaptureReq struct {
	AuthorizationID string `json:"authorization_id"`
	Amount float64
}

type RefundReq struct {
	AuthorizationID string `json:"authorization_id"`
	Amount float64
}

type ErrorResp struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

func newErrResp(code,message string) *ErrorResp  {
	return &ErrorResp{
		Code:    code,
		Message: message,
	}
}

type AuthorizeResp struct {
	AuthorizationID string `json:"authorization_id"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
}

type VoidResp struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
}

type CaptureResp struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
}

type RefundResp struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
}