package repositories

import (
	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ICertificateRepository interface {
	SaveNewCertificate(
		ctx *gin.Context, 
		crudCreateCertificate entity.CRUDCreateCertificate,
		nftId uint,
	) *errors.InfraError
	GetCertificatesByCollectionId(ctx *gin.Context, crudGetCertificates entity.CRUDGetCertificates) ([]*entity.Cert, *errors.InfraError)
	GetCertificateById(ctx *gin.Context, certId uint) (*entity.Cert, *errors.InfraError)
	RevokeCertificate(
		ctx *gin.Context,
		crudRevokeCertificate entity.CRUDRevokeCertificate,
	) *errors.InfraError
}
