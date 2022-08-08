package transformers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"strconv"

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

func stringToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return n
	} else {
		fmt.Println(s, "is not an integer.")
		return 0
	}
}

func ToCRUDCreateCertificate(reqBody dto.CreateCertificateRequest) entity.CRUDCreateCertificate {
	issuerBytes, _ := json.Marshal(reqBody.CertData.Issuer)
	receiverBytes, _ := json.Marshal(reqBody.CertData.Receiver)

	var issuer entity.CertDataTypeIssuer
	var receiver entity.CertDataTypeReceiver

	json.Unmarshal(issuerBytes, &issuer)
	json.Unmarshal(receiverBytes, &receiver)

	fmt.Println("issued at", reqBody.CertData.IssuedAt)
	fmt.Println("ExpiredAt", reqBody.CertData.ExpiredAt)
	if (reqBody.CertData.ExpiredAt == "") {
		reqBody.CertData.ExpiredAt = "253402127999000" // biggest value of year 9999
	}

	return entity.CRUDCreateCertificate{
		CollectionId:	reqBody.CollectionId,
		CertData:		entity.CertDataType{
							CertTitle: reqBody.CertData.CertTitle,
							Issuer: issuer,
							Receiver: receiver,
							Description: reqBody.CertData.Description,
							IssuedAt: stringToInt64(reqBody.CertData.IssuedAt),
							ExpiredAt: stringToInt64(reqBody.CertData.ExpiredAt),
							CertImage: reqBody.CertData.CertImage,
							Platform: reqBody.CertData.Platform,
						},
		TxHash:   		reqBody.TxHash,
		Platform: 		reqBody.Platform,
	}
}

func ToCRUDRevokeCertificate(reqBody dto.RevokeCertificateRequest) entity.CRUDRevokeCertificate {
	return entity.CRUDRevokeCertificate{
		ID:			reqBody.CertId,
		TxHash:   		reqBody.TxHash,
		Platform: 		reqBody.Platform,
	}
}

func TimeToMsString(t time.Time) string {
	var unixT int64
	unixT = t.Unix() * 1000
	
	if unixT < 24*3600*1000 {
		return "null"
	} else {
		return strconv.FormatInt(unixT, 10)
	}
}

func ToCertificateResponses(certificates []*entity.Cert) []*dto.CertificateResponse {
	ret := []*dto.CertificateResponse{}
	for _, c := range certificates {
		var receiver dto.CertDataReceiver;
		json.Unmarshal([]byte(c.Receiver), &receiver)

		var certDataFromDB dto.CertDataTypeFromDB;
		json.Unmarshal([]byte(c.Data), &certDataFromDB)

		ret = append(ret, &dto.CertificateResponse{
			ID:				c.ID,
			Description:	c.Description,
			IssuedAt:      	TimeToMsString(c.IssuedAt),
			ExpiredAt:      TimeToMsString(c.ExpiredAt),
			CollectionId:   c.CollectionId,
			CertNftId:      c.CertNftId,
			Data:  			dto.CertDataType{
								CertTitle: certDataFromDB.CertTitle,
								Issuer: certDataFromDB.Issuer,
								Receiver: certDataFromDB.Receiver,
								Description: certDataFromDB.Description,
								IssuedAt: strconv.FormatInt(certDataFromDB.IssuedAt, 10),
								ExpiredAt: strconv.FormatInt(certDataFromDB.ExpiredAt, 10),
								CertImage: certDataFromDB.CertImage,
								Platform: certDataFromDB.Platform,
							},
			RevokedAt: 		TimeToMsString(c.RevokedAt),
			RevokedReason:	c.RevokedReason,
			Receiver:		receiver,
		})
	}
	return ret
}

func ToCertificateResponse(certificate *entity.Cert) *dto.CertificateResponse {
	if certificate == nil {
		return nil
	}

	var receiver dto.CertDataReceiver;
	json.Unmarshal([]byte(certificate.Receiver), &receiver)

	var certDataFromDB dto.CertDataTypeFromDB;
	json.Unmarshal([]byte(certificate.Data), &certDataFromDB)

	ret := dto.CertificateResponse{
		ID:				certificate.ID,
		Description:	certificate.Description,
		IssuedAt:      	TimeToMsString(certificate.IssuedAt),
		ExpiredAt:      TimeToMsString(certificate.ExpiredAt),
		CollectionId:   certificate.CollectionId,
		CertNftId:      certificate.CertNftId,
		Data:  			dto.CertDataType{
							CertTitle: certDataFromDB.CertTitle,
							Issuer: certDataFromDB.Issuer,
							Receiver: certDataFromDB.Receiver,
							Description: certDataFromDB.Description,
							IssuedAt: strconv.FormatInt(certDataFromDB.IssuedAt, 10),
							ExpiredAt: strconv.FormatInt(certDataFromDB.ExpiredAt, 10),
							CertImage: certDataFromDB.CertImage,
							Platform: certDataFromDB.Platform,
						},
		RevokedAt: 		TimeToMsString(certificate.RevokedAt),
		RevokedReason:	certificate.RevokedReason,
		Receiver:		receiver,
	}
	return &ret
}


func ToCRUDCreateCollection(reqBody dto.CreateCollectionRequest) entity.CRUDCreateCollection {
	return entity.CRUDCreateCollection{
		TxHash:   reqBody.TxHash,
		Platform: reqBody.Platform,
	}
}

func ToCRUDGetCollection(reqParams dto.GetCollectionRequest) entity.CRUDGetCollection {
	return entity.CRUDGetCollection{
		ID: reqParams.ID,
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

func ToCollectionResponse(collection *entity.Collection) *dto.CollectionResponse {
	ret := &dto.CollectionResponse{
		ID:                collection.ID,
		CollectionName:    collection.Title,
		CollectionSymbol:  collection.Symbol,
		CollectionAddress: collection.Address,
		Issuer:            collection.Issuer,
		TotalIssued:       uint64(collection.TotalIssued),
		TotalRevoked:      uint64(collection.TotalRevoked),
	}
	return ret
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
