package repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"
	"decert/internal/pkg/utils/log"

	"time"
)

type certificateRepository struct {
	db *gorm.DB
}

var certificateRepo *certificateRepository

func InitCertificateRepository(db *gorm.DB) {
	if certificateRepo == nil {
		certificateRepo = &certificateRepository{
			db: db,
		}
	}
}

func CertificateRepositoryInstance() ICertificateRepository {
	return certificateRepo
}

func (r *certificateRepository) SaveNewCertificate(ctx *gin.Context, ethAddress, title, symbol, issuer string) *errors.InfraError {
	// certificate := &entity.Certificate{
	// 	Title:        title,
	// 	Symbol:       symbol,
	// 	Address:      ethAddress,
	// 	Issuer:       issuer,
	// 	TotalIssued:  0,
	// 	TotalRevoked: 0,
	// }

	// log.Debugf(ctx, "%+v", certificate)

	// result := r.db.Create(certificate)
	// if result.Error != nil {
	// 	return errors.NewInfraErrorDBInsert([]string{}, result.Error)
	// }
	return nil
}

func (r *certificateRepository) GetCertificatesByCollectionId(
	ctx *gin.Context, collectionId uint,
) ([]*entity.Certificate, *errors.InfraError) {
	log.Debugf(ctx, "Get certificate of collectionId: %s", collectionId)
	ret := []*entity.Certificate{}
	// result := r.db.
	// 	Where("issuer = ?", issuer).
	// 	Order("created_at DESC").
	// 	Limit(int(limit)).
	// 	Offset(int(offset)).
	// 	Find(&ret)
	// if result.Error != nil {
	// 	return nil, errors.NewInfraErrorDBSelect([]string{}, result.Error)
	// }
	return ret, nil
}

func (r *certificateRepository) GetCertificateById(
	ctx *gin.Context, 
	certId uint,
) (*entity.Certificate, *errors.InfraError) {
	log.Debugf(ctx, "Get certificate with id: %s", certId)
	ret := &entity.Certificate{
		1, 
		"desc", 
		time.Now(),
		time.Now(),
		1,
		123,
		"{}",
		time.Now(),
		"reason",
		"receiver",
	}

	return ret, nil
}