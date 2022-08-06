package transformers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"decert/internal/pkg/constants"
	"decert/internal/pkg/dto"
	"decert/internal/pkg/entity"
	"decert/internal/pkg/errors"
)

func ResponseOK(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    constants.RestCodeOK,
		Message: constants.RestMsgOK,
		Data:    data,
	})
}

func ResponseErr(ctx *gin.Context, err *errors.RestError) {
	fmt.Println(ctx)
	ctx.AbortWithStatusJSON(int(err.HttpCode), err)
}

func ToCRUDGetCertificate(reqParams dto.GetCertificateRequest) entity.CRUDGetCertificate {
	return entity.CRUDGetCertificate{
		ID: 	reqParams.CertId,
	}
}


func ToCRUDGetCertificates(reqParams dto.GetCertificatesRequest) entity.CRUDGetCertificates {
	return entity.CRUDGetCertificates{
		CollectionId: 	reqParams.CollectionId,
		Limit:			reqParams.Limit,
		Offset:			reqParams.Offset,
	}
}

func ToCRUDCreateCertificate(reqBody dto.CreateCertificateRequest) entity.CRUDCreateCertificate {
	issuer, _ := json.Marshal(reqBody.CertData.Issuer)
	receiver, _ := json.Marshal(reqBody.CertData.Receiver)

	stringIssuer, stringReceiver := string(issuer), string(receiver)

	return entity.CRUDCreateCertificate{
		CollectionId:	reqBody.CollectionId,
		CertData:		entity.CertDataType{
							CertTitle: reqBody.CertData.CertTitle,
							Issuer: stringIssuer,
							Receiver: stringReceiver,
							Description: reqBody.CertData.Description,
							IssuedAt: reqBody.CertData.IssuedAt.Time,
							ExpiredAt: reqBody.CertData.ExpiredAt.Time,
							CertImage: reqBody.CertData.CertImage,
							Platform: reqBody.CertData.Platform,
						},
		TxHash:   		reqBody.TxHash,
		Platform: 		reqBody.Platform,
	}
}

func ToCertificateResponses(certificates []*entity.Cert) []*dto.CertificateResponse {
	ret := []*dto.CertificateResponse{}
	for _, c := range certificates {
		ret = append(ret, &dto.CertificateResponse{
			ID:				c.ID,
			Description:	c.Description,
			IssuedAt:      	c.IssuedAt,
			ExpiredAt:      c.ExpiredAt,
			CollectionId:   c.CollectionId,
			CertNftId:      c.CertNftId,
			Data:  			c.Data,
			RevokedAt: 		c.RevokedAt,
			RevokedReason:	c.RevokedReason,
			Receiver:		c.Receiver,
		})
	}
	return ret
}


func ToCRUDCreateCollection(reqBody dto.CreateCollectionRequest) entity.CRUDCreateCollection {
	return entity.CRUDCreateCollection{
		TxHash:   reqBody.TxHash,
		Platform: reqBody.Platform,
	}
}

func ToCRUDGetCollections(reqParams dto.GetCollectionsRequest) entity.CRUDGetCollections {
	return entity.CRUDGetCollections{
		Issuer: reqParams.Issuer,
		Limit:  reqParams.Limit,
		Offset: reqParams.Offset,
		Name:   reqParams.Name,
	}
}

func ToCollectionResponses(collections []*entity.Collection) []*dto.CollectionResponse {
	ret := []*dto.CollectionResponse{}
	for _, c := range collections {
		ret = append(ret, &dto.CollectionResponse{
			ID:                c.ID,
			CollectionName:    c.Title,
			CollectionSymbol:  c.Symbol,
			CollectionAddress: c.Address,
			Issuer:            c.Issuer,
			TotalIssued:       uint64(c.TotalIssued),
			TotalRevoked:      uint64(c.TotalRevoked),
		})
	}
	return ret
}
