package payments

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation")
	ErrAuthFailed = errors.New("authorisation failure")
	ErrVoidFailed = errors.New("void failure")
	ErrCaptureFailed = errors.New("capture failure")
	ErrAutIDNotFound = errors.New("authorization id is not found")
	ErrCaptureLimitExceeded = errors.New("capture failure, limit exceeded")

	ErrCardNil = fmt.Errorf("%w: card must to be set",ErrValidation)
	ErrCardName = fmt.Errorf("%w: card name is not valid",ErrValidation)
	ErrAmountNotValid = fmt.Errorf("%w: amount is not valid",ErrValidation)
	ErrAuthIDEmpty = fmt.Errorf("%w: authorization id must to be set",ErrValidation)
)
