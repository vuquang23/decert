package transformers

import (
	"decert/internal/pkg/constants"
	"decert/internal/pkg/errors"
)

type domainErrConstructor func(errorEntities []string, rootCause error) *errors.DomainError

type IDomainErrTransformer interface {
	Transform(infraErr *errors.InfraError) *errors.DomainError
	RegisterTransformFunc(infraErrCode string, f domainErrConstructor)
}

type domainErrTransformer struct {
	mapping map[string]domainErrConstructor
}

var domainErrTrans *domainErrTransformer

func InitDomainErrorTransformer() {
	if domainErrTrans == nil {
		domainErrTrans = &domainErrTransformer{
			mapping: map[string]domainErrConstructor{},
		}
		domainErrTrans.RegisterTransformFunc(constants.InfraErrCodeDBConnect, errors.NewDomainErrorUnknown)
		domainErrTrans.RegisterTransformFunc(constants.InfraErrCodeDBNotFound, errors.NewDomainErrorNotFound)
		domainErrTrans.RegisterTransformFunc(constants.InfraErrCodeDBSelect, errors.NewDomainErrorUnknown)
		domainErrTrans.RegisterTransformFunc(constants.InfraErrCodeDBInsert, errors.NewDomainErrorUnknown)
		domainErrTrans.RegisterTransformFunc(constants.InfraErrCodeDBUpdate, errors.NewDomainErrorUnknown)
		domainErrTrans.RegisterTransformFunc(constants.InfraErrCodeDBDelete, errors.NewDomainErrorUnknown)
		domainErrTrans.RegisterTransformFunc(constants.InfraErrCodeDBUnknown, errors.NewDomainErrorUnknown)
	}
}

func DomainErrTransformerInstance() IDomainErrTransformer {
	return domainErrTrans
}

func (t *domainErrTransformer) Transform(infraErr *errors.InfraError) *errors.DomainError {
	f := t.mapping[infraErr.Code]
	if f == nil {
		return errors.NewDomainErrorUnknown(infraErr.ErrorEntities, infraErr)
	}
	return f(infraErr.ErrorEntities, infraErr)
}

func (t *domainErrTransformer) RegisterTransformFunc(infraErrCode string, f domainErrConstructor) {
	t.mapping[infraErrCode] = f
}
