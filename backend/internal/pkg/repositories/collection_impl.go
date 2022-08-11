package repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"
	"decert/internal/pkg/utils/log"
)

type collectionRepository struct {
	db *gorm.DB
}

var collectionRepo *collectionRepository

func InitCollectionRepository(db *gorm.DB) {
	if collectionRepo == nil {
		collectionRepo = &collectionRepository{
			db: db,
		}
	}
}

func CollectionRepositoryInstance() ICollectionRepository {
	return collectionRepo
}

func (r *collectionRepository) SaveNewCollection(ctx *gin.Context, ethAddress, title, symbol, issuer string) *errors.InfraError {
	collection := &entity.Collection{
		Title:        title,
		Symbol:       symbol,
		Address:      ethAddress,
		Issuer:       issuer,
		TotalIssued:  0,
		TotalRevoked: 0,
	}

	log.Debugf(ctx, "%+v", collection)

	result := r.db.Create(collection)
	if result.Error != nil {
		return errors.NewInfraErrorDBInsert([]string{}, result.Error)
	}
	return nil
}

func (r *collectionRepository) GetCollectionsByIssuerAddress(
	ctx *gin.Context, issuer string, limit, offset uint64, name string,
) ([]*entity.Collection, *errors.InfraError) {
	log.Debugf(ctx, "Get collection of issuer: %s", issuer)
	ret := []*entity.Collection{}
	tx := r.db.Where("issuer = ?", issuer)
	if name != "" {
		tx = tx.Where("title LIKE ?", "%"+name+"%")
	}
	result := tx.
		Order("created_at DESC").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&ret)
	if result.Error != nil {
		return nil, errors.NewInfraErrorDBSelect([]string{}, result.Error)
	}
	return ret, nil
}

func (r *collectionRepository) GetCollectionById(
	ctx *gin.Context, collectionId uint,
) (*entity.Collection, *errors.InfraError) {
	log.Debugf(ctx, "Get collection of id: %s", collectionId)
	ret := []*entity.Collection{}
	tx := r.db.Where("id = ?", collectionId)
	result := tx.Find(&ret)

	if result.Error != nil {
		return nil, errors.NewInfraErrorDBSelect([]string{}, result.Error)
	}

	if (len(ret) == 0) {
		return nil, nil
	}
	return ret[0], nil
}