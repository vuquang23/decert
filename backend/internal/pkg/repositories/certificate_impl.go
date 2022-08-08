package repositories

import (
	"encoding/json"
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
) *errors.InfraError {
	fmt.Println("Saving cert")

	certJsonData, _ := json.Marshal(crudCreateCertificate.CertData)
	certJsonString := string(certJsonData)

	receiverJsonData, _ := json.Marshal(crudCreateCertificate.CertData.Receiver)
	receiverJsonString := string(receiverJsonData)

	fmt.Println("certJsonString", certJsonString)

	certificate := &entity.Cert{
		Description:    crudCreateCertificate.CertData.Description,
		IssuedAt:       time.Unix(crudCreateCertificate.CertData.IssuedAt/1000, 7),
		ExpiredAt:      time.Unix(crudCreateCertificate.CertData.ExpiredAt/1000, 7),
		CollectionId:   crudCreateCertificate.CollectionId,
		CertNftId:      nftId,
		Data:  			certJsonString,
		RevokedAt: 		time.Unix(0, 7),
		RevokedReason:	"",
		Receiver:		receiverJsonString,
	}

	log.Debugf(ctx, "%+v", certificate)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		resultCreate := r.db.Create(certificate)
		if resultCreate.Error != nil {
			return errors.NewInfraErrorDBInsert([]string{}, resultCreate.Error)
		}
		
		// add 1 to collection cert count
		var currentCollection entity.Collection
		if err := tx.Where("id = ?", crudCreateCertificate.CollectionId).First(&currentCollection).Error; err != nil {
            return err
        }
		resultInc := tx.
				Model(&currentCollection).
				Where("id = ?", crudCreateCertificate.CollectionId).
				Updates(entity.Collection{TotalIssued: currentCollection.TotalIssued+1})
		if resultInc.Error != nil {
			return errors.NewInfraErrorDBInsert([]string{}, resultInc.Error)
		}
		return nil
    })

	if err != nil {
		return errors.NewInfraErrorDBInsert([]string{}, err)
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

func (r *certificateRepository) RevokeCertificate(
	ctx *gin.Context,
	crudRevokeCertificate entity.CRUDRevokeCertificate,
) *errors.InfraError {
	log.Debugf(ctx, "Revoking certificate")

	err := r.db.Transaction(func(tx *gorm.DB) error {
		var currentEntity *entity.Cert
		currentEntity, _ = r.GetCertificateById(ctx, crudRevokeCertificate.ID)

		resultRevoke := r.db.
			Model(&currentEntity).
			Where("id = ?", crudRevokeCertificate.ID).
			Updates(entity.Cert{RevokedAt: time.Unix(crudRevokeCertificate.RevokedAt/1000, 7), RevokedReason: crudRevokeCertificate.RevokedReason})
		if resultRevoke.Error != nil {
			return errors.NewInfraErrorDBSelect([]string{}, resultRevoke.Error)
		}
		
		// add 1 to collection cert revoke count
		var currentCollection entity.Collection
		if err := tx.Where("id = ?", currentEntity.CollectionId).First(&currentCollection).Error; err != nil {
            return err
        }
		resultDec := tx.
				Model(&currentCollection).
				Where("id = ?", currentEntity.CollectionId).
				Updates(entity.Collection{TotalRevoked: currentCollection.TotalRevoked+1})
		if resultDec.Error != nil {
			return errors.NewInfraErrorDBInsert([]string{}, resultDec.Error)
		}
		return nil
    })

	if err != nil {
		return errors.NewInfraErrorDBInsert([]string{}, err)
	}


	return nil
}