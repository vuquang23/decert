package api

import (
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"

	"decert/internal/pkg/constants"
	"decert/internal/pkg/dto"
	"decert/internal/pkg/errors"
	"decert/internal/pkg/services"
	"decert/internal/pkg/transformers"
	"decert/internal/pkg/utils/log"
	"decert/internal/pkg/utils/uuid"
)

type certificateApi struct {
	certificateSvc services.ICertificateService
}

func newCertificateApi(certificateSvc services.ICertificateService) *certificateApi {
	return &certificateApi{
		certificateSvc: certificateSvc,
	}
}

func (api *certificateApi) setupRoute(rg *gin.RouterGroup) {
	rg.GET("", api.getCertificates)
	rg.GET(":certId", api.getCertificateInfo)
	rg.PUT(":certId", api.revokeCertificate)
	rg.POST("", api.createCertificate)
}

func (api * certificateApi) getCertificateInfo(ctx *gin.Context) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("get-certificate-info"))

	dtoGetCertificate := dto.GetCertificateRequest{}

	certIdStr, _ := strconv.ParseUint(ctx.Param("certId"), 10, 32)
	dtoGetCertificate.CertId = uint(certIdStr)

	if certificate, err := api.certificateSvc.GetCertificateInfo(ctx,
		transformers.ToCRUDGetCertificate(dtoGetCertificate),
	); err != nil {
		restErr := transformers.RestErrTransformerInstance().Transform(err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
	} else {
		transformers.ResponseOK(ctx, certificate)
	}
}

func (api *certificateApi) getCertificates(ctx *gin.Context) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("get-certificates"))

	query := ctx.Request.URL.Query()

	var reqParams dto.GetCertificatesRequest
	collectionIdUint64, _ := strconv.ParseUint(query["collectionId"][0], 10, 32)
	limitUint64, _ := strconv.ParseUint(query["limit"][0], 10, 32)
	offsetUint64, _ := strconv.ParseUint(query["offset"][0], 10, 32)
	reqParams.CollectionId = uint(collectionIdUint64)
	reqParams.Limit = uint(limitUint64)
	reqParams.Offset = uint(offsetUint64)
	
	certificates, err := api.certificateSvc.GetCertificates(ctx,
		transformers.ToCRUDGetCertificates(reqParams))
	if err != nil {
		restErr := transformers.RestErrTransformerInstance().Transform(err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
		return
	}
	fmt.Println("Certs", certificates)

	transformers.ResponseOK(ctx, transformers.ToCertificateResponses(certificates))
}

func (api *certificateApi) createCertificate(ctx *gin.Context) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("create-certificate"))
	
	var reqBody dto.CreateCertificateRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		restErr := errors.NewRestErrorInvalidFormat([]string{}, err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
		return
	}

	fmt.Println("BODY", reqBody)
	fmt.Println("CRUD CREATE", transformers.ToCRUDCreateCertificate(reqBody))

	if err := api.certificateSvc.CreateCertificate(ctx,
		transformers.ToCRUDCreateCertificate(reqBody),
	); err != nil {
		restErr := transformers.RestErrTransformerInstance().Transform(err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
	}

	transformers.ResponseOK(ctx, nil)
}

func (api *certificateApi) revokeCertificate(ctx *gin.Context) {

}
