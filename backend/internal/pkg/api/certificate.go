package api

import (
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
}

func (api *certificateApi) getCertificates(ctx *gin.Context) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("get-certificates"))

	// var callOpts = bind.CallOpts{}
	// log.Debugln(ctx, "callOpts")
	// var scCaller sc_certificate.ScCaller = sc_certificate.ScCaller{}
	// var bs = scCaller.Test()
	// fmt.Println(callOpts)
	// fmt.Println(bs)

	// response := make(map[string]string)
	// response["msg"] = "Hello a Quang"
	// response["batchSize"] = string(bs)
	// transformers.ResponseOK(ctx, bs)

	// var reqParams dto.GetCertificatesRequest
	// if err := ctx.ShouldBindQuery(&reqParams); err != nil {
	// 	restErr := errors.NewRestErrorInvalidFormat([]string{}, err)
	// 	log.Errorln(ctx, restErr)
	// 	transformers.ResponseErr(ctx, restErr)
	// 	return
	// }
	// fmt.Println(reqParams)

	// certificates, err := api.certificateSvc.GetCertificates(ctx,
	// 	transformers.ToCRUDGetCertificates(reqParams))
	// if err != nil {
	// 	restErr := transformers.RestErrTransformerInstance().Transform(err)
	// 	log.Errorln(ctx, restErr)
	// 	transformers.ResponseErr(ctx, restErr)
	// 	return
	// }

	// transformers.ResponseOK(ctx, transformers.ToCertificateResponses(certificates))
}

func (api *certificateApi) createCertificate(ctx *gin.Context) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("create-certificate"))
	
	var reqBody dto.CreateCertificateRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		fmt.Println("REQBODY", reqBody)
		restErr := errors.NewRestErrorInvalidFormat([]string{}, err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
		return
	}

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
