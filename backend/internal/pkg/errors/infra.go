package errors

import (
	"fmt"
	"decert/internal/pkg/constants"
)

type InfraError struct {
	Code          string
	Message       string
	ErrorEntities []string
	RootCause     error
}

func (e *InfraError) Error() string {
	return fmt.Sprintf(
		"INFRA ERROR: {Code: %s, Message: %s, entities: %v, rootCause: %v}",
		e.Code,
		e.Message,
		e.ErrorEntities,
		e.RootCause,
	)
}

func NewInfraError(code string, message string, entities []string, rootCause error) *InfraError {
	return &InfraError{
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func NewInfraErrorDBConnect(entities []string, rootCause error) *InfraError {
	return NewInfraError(constants.InfraErrCodeDBConnect, constants.InfraErrMsgDBConnect, entities, rootCause)
}

func NewInfraErrorDBNotFound(entities []string, rootCause error) *InfraError {
	return NewInfraError(constants.InfraErrCodeDBNotFound, constants.InfraErrMsgDBNotFound, entities, rootCause)
}

func NewInfraErrorDBSelect(entities []string, rootCause error) *InfraError {
	return NewInfraError(constants.InfraErrCodeDBSelect, constants.InfraErrMsgDBSelect, entities, rootCause)
}

func NewInfraErrorDBInsert(entities []string, rootCause error) *InfraError {
	return NewInfraError(constants.InfraErrCodeDBInsert, constants.InfraErrMsgDBInsert, entities, rootCause)
}

func NewInfraErrorDBUpdate(entities []string, rootCause error) *InfraError {
	return NewInfraError(constants.InfraErrCodeDBUpdate, constants.InfraErrMsgDBUpdate, entities, rootCause)
}

func NewInfraErrorDBDelete(entities []string, rootCause error) *InfraError {
	return NewInfraError(constants.InfraErrCodeDBDelete, constants.InfraErrMsgDBDelete, entities, rootCause)
}

func NewInfraErrorDBUnknown(entities []string, rootCause error) *InfraError {
	return NewInfraError(constants.InfraErrCodeDBUnknown, constants.InfraErrMsgDBUnknown, entities, rootCause)
}
