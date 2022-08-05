package repositories

import (
	"encoding/json"
	"strconv"
	"fmt"
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

func (r *certificateRepository) SaveNewCertificate(
	ctx *gin.Context,
	crudCreateCertificate entity.CRUDCreateCertificate,
	nftId uint,
	// collectionId uint,
	// issuer string, 
	// recipient string,
	// certHash [32]byte, 
	// link string, 
	// issuedAt *big.Int,
) *errors.InfraError {
	// TODO
	fmt.Println("Saving cert")
	issuedAtInt, _ := strconv.ParseInt(crudCreateCertificate.CertData.IssuedAt.String(), 10, 32)
	expiredAtInt, _ := strconv.ParseInt(crudCreateCertificate.CertData.ExpiredAt.String(), 10, 32)

	certJsonData, _ := json.Marshal(crudCreateCertificate.CertData)
	certJsonString := string(certJsonData)

	certificate := &entity.Cert{
		Description:    crudCreateCertificate.CertData.Description,
		IssuedAt:       time.Unix(issuedAtInt, 7),
		ExpiredAt:      time.Unix(expiredAtInt, 7),
		CollectionId:   crudCreateCertificate.CollectionId,
		CertNftId:      nftId,
		Data:  			certJsonString,
		RevokedAt: 		time.Unix(expiredAtInt, 7),
		RevokedReason:	"",
		Receiver:		crudCreateCertificate.CertData.Receiver,
	}

	log.Debugf(ctx, "%+v", certificate)

	result := r.db.Create(certificate)
	if result.Error != nil {
		return errors.NewInfraErrorDBInsert([]string{}, result.Error)
	}
	return nil
}

func (r *certificateRepository) GetCertificatesByCollectionId(
	ctx *gin.Context, 
	collectionId uint,
	limit, offset uint,
) ([]*entity.Cert, *errors.InfraError) {
	log.Debugf(ctx, "Get certificate of collectionId: %s", collectionId)
	ret := []*entity.Cert{}
	result := r.db.
		Where("collection_id = ?", collectionId).
		Order("issued_at DESC").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&ret)
	if result.Error != nil {
		return nil, errors.NewInfraErrorDBSelect([]string{}, result.Error)
	}
	return ret, nil
}

func (r *certificateRepository) GetCertificateById(
	ctx *gin.Context, 
	certId uint,
) (*entity.Cert, *errors.InfraError) {
	log.Debugf(ctx, "Get certificate with id: %s", certId)
	ret := []*entity.Cert{}
	result := r.db.
		Where("id = ?", certId).
		Find(&ret)
	if result.Error != nil {
		return nil, errors.NewInfraErrorDBSelect([]string{}, result.Error)
	}

	if (len(ret) == 0) {
		return nil, nil
	}
	return ret[0], nil
}