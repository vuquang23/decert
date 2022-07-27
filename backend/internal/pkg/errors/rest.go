package errors

import (
	"decert/internal/pkg/constants"
	"fmt"
)

type RestError struct {
	HttpCode      int64    `json:"-"`
	Code          int64    `json:"code"`
	Message       string   `json:"message"`
	ErrorEntities []string `json:"errorEntities"`
	RootCause     error    `json:"-"`
}

func (e *RestError) Error() string {
	return fmt.Sprintf(
		"API Error: {Code: %d, Message: %s, ErrorEntities: %v, RootCause: %v}",
		e.Code,
		e.Message,
		e.ErrorEntities,
		e.RootCause,
	)
}

func NewRestError(httpCode int64, code int64, message string, errorEntities []string, rootCause error) *RestError {
	return &RestError{
		HttpCode:      httpCode,
		Code:          code,
		Message:       message,
		ErrorEntities: errorEntities,
		RootCause:     rootCause,
	}
}

func NewRestErrorRequired(errorEntities []string, rootCause error) *RestError {
	return NewRestError(400, constants.RestErrCodeRequired, constants.RestErrMsgRequired, errorEntities, rootCause)
}

func NewRestErrorNotAcceptedValue(errorEntities []string, rootCause error) *RestError {
	return NewRestError(400, constants.RestErrCodeNotAcceptedValue, constants.RestErrMsgNotAcceptedValue, errorEntities, rootCause)
}

func NewRestErrorInvalidFormat(errorEntities []string, rootCause error) *RestError {
	return NewRestError(400, constants.RestErrCodeInvalidFormat, constants.RestErrMsgInvalidFormat, errorEntities, rootCause)
}

func NewRestErrorUnauthenticated(errorEntities []string, rootCause error) *RestError {
	return NewRestError(401, constants.RestErrCodeUnauthenticated, constants.RestErrMsgUnauthenticated, errorEntities, rootCause)
}

func NewRestErrorNotFound(errorEntities []string, rootCause error) *RestError {
	return NewRestError(404, constants.RestErrCodeNotFound, constants.RestErrMsgNotFound, errorEntities, rootCause)
}

func NewRestErrorInternal(errorEntities []string, rootCause error) *RestError {
	return NewRestError(500, constants.RestErrCodeInternal, constants.RestErrMsgInternal, errorEntities, rootCause)
}
