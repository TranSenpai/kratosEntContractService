package errors

import kerrors "github.com/go-kratos/kratos/v2/errors"

const (
	ReasonContractNotFound     = "CONTRACT_NOT_FOUND"
	ReasonInvalidContract      = "INVALID_CONTRACT_DATA"
	ReasonMissingRequiredField = "MISSING_REQUIRED_FIELD"
	ReasonInternalError        = "INTERNAL_ERROR"
)

func BadRequest(reason, msg string) error {
	return kerrors.New(400, reason, msg)
}

func NotFound(reason, msg string) error {
	return kerrors.New(404, reason, msg)
}

func Internal(reason, msg string) error {
	return kerrors.New(500, reason, msg)
}
