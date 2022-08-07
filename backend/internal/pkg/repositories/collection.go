package repositories

import (
	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ICollectionRepository interface {
	SaveNewCollection(ctx *gin.Context, ethAddress, title, symbol, issuer string) *errors.InfraError
	GetCollectionsByIssuerAddress(ctx *gin.Context, issuer string, limit, offset uint64, name string) ([]*entity.Collection, *errors.InfraError)
	GetCollectionById(ctx *gin.Context, collectionId uint) (*entity.Collection, *errors.InfraError)
}
