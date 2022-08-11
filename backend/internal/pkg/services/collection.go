package services

import (
	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ICollectionService interface {
	CreateCollection(ctx *gin.Context, crudCreateCollection entity.CRUDCreateCollection) *errors.DomainError
	GetCollections(ctx *gin.Context, crudGetCollections entity.CRUDGetCollections) ([]*entity.Collection, *errors.DomainError)
	GetCollectionInfo(
		ctx *gin.Context, 
		crudGetCollection entity.CRUDGetCollection,
	)  (*entity.Collection, *errors.DomainError)
}
