package repositories

import (
	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ICertificateRepository interface {
	SaveNewCertificate(ctx *gin.Context, ethAddress, title, symbol, issuer string) *errors.InfraError
	GetCertificatesByCollectionId(ctx *gin.Context, collectionId uint) ([]*entity.Certificate, *errors.InfraError)
	GetCertificateById(ctx *gin.Context, certId uint) (*entity.Certificate, *errors.InfraError)
}
