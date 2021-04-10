package web

type HealthResp struct {
	Status string `json:"status"`
}

type Card struct {
	Name        string `json:"name"`
	CardNumber  string `json:"card_number"`
	ExpireMonth string `json:"expire_month"`
	ExpireYear  string `json:"expire_year"`
	CVV         string `json:"cvv"`
}

type AuthorizeReq struct {
	Card   Card   `json:"card"`
	Amount Amount `json:"amount"`
}

type VoidReq struct {
	AuthorizationID string `json:"authorization_id"`
}

type CaptureReq struct {
	AuthorizationID string `json:"authorization_id"`
	Amount          float64
}

type RefundReq struct {
	AuthorizationID string `json:"authorization_id"`
	Amount          float64
}

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func newErrResp(code, message string) *ErrorResp {
	return &ErrorResp{
		Code:    code,
		Message: message,
	}
}

type Amount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type AuthorizeResp struct {
	AuthorizationID string  `json:"authorization_id"`
	Amount          *Amount `json:"amount"`
}

type VoidResp struct {
	Amount *Amount `json:"amount"`
}

type CaptureResp struct {
	Amount *Amount `json:"amount"`
}

type RefundResp struct {
	Amount *Amount `json:"amount"`
}
