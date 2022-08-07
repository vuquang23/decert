package transformers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"math/big"
	"time"

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
	issuerBytes, _ := json.Marshal(reqBody.CertData.Issuer)
	receiverBytes, _ := json.Marshal(reqBody.CertData.Receiver)

	var issuer entity.CertDataTypeIssuer
	var receiver entity.CertDataTypeReceiver

	json.Unmarshal(issuerBytes, &issuer)
	json.Unmarshal(receiverBytes, &receiver)

	// stringIssuer, stringReceiver := string(issuer), string(receiver)

	fmt.Println("issued at", reqBody.CertData.IssuedAt)

	return entity.CRUDCreateCertificate{
		CollectionId:	reqBody.CollectionId,
		CertData:		entity.CertDataType{
							CertTitle: reqBody.CertData.CertTitle,
							Issuer: issuer,
							Receiver: receiver,
							Description: reqBody.CertData.Description,
							IssuedAt: reqBody.CertData.IssuedAt,
							ExpiredAt: reqBody.CertData.ExpiredAt,
							CertImage: reqBody.CertData.CertImage,
							Platform: reqBody.CertData.Platform,
						},
		TxHash:   		reqBody.TxHash,
		Platform: 		reqBody.Platform,
	}
}

func TimeToBigInt(t time.Time) big.Int {
	var unixT int64
	unixT = t.Unix()
	fmt.Println("UNIXT", unixT)
	
	var unixTBigInt big.Int
	unixTBigInt.Mul(big.NewInt(unixT), big.NewInt(1000))
	
	return unixTBigInt
}

func ToCertificateResponses(certificates []*entity.Cert) []*dto.CertificateResponse {
	ret := []*dto.CertificateResponse{}
	for _, c := range certificates {
		var receiver dto.CertDataTypeReceiver;
		json.Unmarshal([]byte(c.Receiver), &receiver)

		ret = append(ret, &dto.CertificateResponse{
			ID:				c.ID,
			Description:	c.Description,
			IssuedAt:      	TimeToBigInt(c.IssuedAt),
			ExpiredAt:      TimeToBigInt(c.ExpiredAt),
			CollectionId:   c.CollectionId,
			CertNftId:      c.CertNftId,
			Data:  			c.Data,
			RevokedAt: 		TimeToBigInt(c.RevokedAt),
			RevokedReason:	c.RevokedReason,
			Receiver:		receiver,
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
