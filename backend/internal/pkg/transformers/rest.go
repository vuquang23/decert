package transformers

import (
	"decert/internal/pkg/constants"
	"decert/internal/pkg/errors"
)

type restErrConstructor func(errorEntities []string, rootCause error) *errors.RestError

type IRestErrTransformer interface {
	Transform(domainErr *errors.DomainError) *errors.RestError
	RegisterTransformFunc(domainErrCode string, f restErrConstructor)
}

type restErrTransformer struct {
	mapping map[string]restErrConstructor
}

var restErrTrans *restErrTransformer

func InitRestErrTransformer() {
	if restErrTrans == nil {
		restErrTrans = &restErrTransformer{
			mapping: map[string]restErrConstructor{},
		}
		restErrTrans.RegisterTransformFunc(constants.DomainErrCodeRequired, errors.NewRestErrorRequired)
		restErrTrans.RegisterTransformFunc(constants.DomainErrCodeNotAcceptedValue, errors.NewRestErrorNotAcceptedValue)
		restErrTrans.RegisterTransformFunc(constants.DomainErrCodeInvalidFormat, errors.NewRestErrorInvalidFormat)
		restErrTrans.RegisterTransformFunc(constants.DomainErrCodeUnauthenticated, errors.NewRestErrorUnauthenticated)
		restErrTrans.RegisterTransformFunc(constants.DomainErrCodeNotFound, errors.NewRestErrorNotFound)
		restErrTrans.RegisterTransformFunc(constants.DomainErrCodeUnknown, errors.NewRestErrorInternal)
	}
}

func RestErrTransformerInstance() IRestErrTransformer {
	return restErrTrans
}

func (t *restErrTransformer) Transform(domainErr *errors.DomainError) *errors.RestError {
	f := t.mapping[domainErr.Code]
	if f == nil {
		return errors.NewRestErrorInternal(domainErr.ErrorEntities, domainErr)
	}
	return f(domainErr.ErrorEntities, domainErr)
}

func (t *restErrTransformer) RegisterTransformFunc(domainErrCode string, f restErrConstructor) {
	t.mapping[domainErrCode] = f
}
