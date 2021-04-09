package payments

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation")
	ErrAuthFailed = errors.New("authorisation failure")

	ErrCardNil = fmt.Errorf("%w: card must to be set",ErrValidation)
	ErrCardName = fmt.Errorf("%w: card name is not valid",ErrValidation)
	ErrAmountNotValid = fmt.Errorf("%w: amount is not valid",ErrValidation)
	ErrAuthIDEmpty = fmt.Errorf("%w: authorization id must to be set",ErrValidation)
)
