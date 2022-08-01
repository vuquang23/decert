package services

import (
	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ICertificateService interface {
	CreateCertificate(ctx *gin.Context, crudCreateCertificate entity.CRUDCreateCertificate) *errors.DomainError
	GetCertificates(ctx *gin.Context, crudGetCertificates entity.CRUDGetCertificates) ([]*entity.Certificate, *errors.DomainError)
	GetCertificateInfo(ctx *gin.Context, crudGetCertificateInfo entity.CRUDGetCertificateInfo) (*entity.Certificate, *errors.DomainError)
	RevokeCertificate(ctx *gin.Context, crudRevokeCertificate entity.CRUDRevokeCertificate) *errors.DomainError
}
