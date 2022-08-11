package errors

import (
	"decert/internal/pkg/constants"
	"fmt"
)

type DomainError struct {
	Code          string
	Message       string
	ErrorEntities []string
	RootCause     error
}

func (e *DomainError) Error() string {
	return fmt.Sprintf(
		"DOMAIN ERROR: {Code: %s, Message: %s, entities: %v, rootCause: %v}",
		e.Code,
		e.Message,
		e.ErrorEntities,
		e.RootCause,
	)
}

func NewDomainError(code string, message string, errorEntities []string, rootCause error) *DomainError {
	return &DomainError{
		Code:          code,
		Message:       message,
		ErrorEntities: errorEntities,
		RootCause:     rootCause,
	}
}

func NewDomainErrorRequired(errorEntities []string, rootCause error) *DomainError {
	return NewDomainError(constants.DomainErrCodeRequired, constants.DomainErrMsgRequired, errorEntities, rootCause)
}

func NewDomainErrorNotAcceptedValue(errorEntities []string, rootCause error) *DomainError {
	return NewDomainError(constants.DomainErrCodeNotAcceptedValue, constants.DomainErrMsgNotAcceptedValue, errorEntities, rootCause)
}

func NewDomainErrorInvalidFormat(errorEntities []string, rootCause error) *DomainError {
	return NewDomainError(constants.DomainErrCodeInvalidFormat, constants.DomainErrMsgInvalidFormat, errorEntities, rootCause)
}

func NewDomainErrorUnauthenticated(errorEntities []string, rootCause error) *DomainError {
	return NewDomainError(constants.DomainErrCodeUnauthenticated, constants.DomainErrMsgUnauthenticated, errorEntities, rootCause)
}

func NewDomainErrorNotFound(errorEntities []string, rootCause error) *DomainError {
	return NewDomainError(constants.DomainErrCodeNotFound, constants.DomainErrMsgNotFound, errorEntities, rootCause)
}

func NewDomainErrorUnknown(errorEntities []string, rootCause error) *DomainError {
	return NewDomainError(constants.DomainErrCodeUnknown, constants.DomainErrMsgUnknown, errorEntities, rootCause)
}
