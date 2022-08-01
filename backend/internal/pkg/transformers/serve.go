package transformers

import (
	"fmt"
	"net/http"

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

func ToCRUDCreateCertificate(reqBody dto.CreateCertificateRequest) entity.CRUDCreateCertificate {
	return entity.CRUDCreateCertificate{
		CollectionId:	reqBody.CollectionId,
		CertData:		entity.CertDataType{
							CertTitle: reqBody.CertData.CertTitle,
							Issuer: reqBody.CertData.Issuer,
							Receiver: reqBody.CertData.Receiver,
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
