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

type collectionApi struct {
	collectionSvc services.ICollectionService
}

func newCollectionApi(collectionSvc services.ICollectionService) *collectionApi {
	return &collectionApi{
		collectionSvc: collectionSvc,
	}
}

func (api *collectionApi) setupRoute(rg *gin.RouterGroup) {
	rg.GET("", api.getCollections)
	rg.POST("", api.createCollection)
}

func (api *collectionApi) getCollections(ctx *gin.Context) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("get-collections"))

	var reqParams dto.GetCollectionsRequest
	if err := ctx.ShouldBindQuery(&reqParams); err != nil {
		restErr := errors.NewRestErrorInvalidFormat([]string{}, err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
		return
	}
	fmt.Println(reqParams)

	collections, err := api.collectionSvc.GetCollections(ctx,
		transformers.ToCRUDGetCollections(reqParams))
	if err != nil {
		restErr := transformers.RestErrTransformerInstance().Transform(err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
		return
	}

	transformers.ResponseOK(ctx, transformers.ToCollectionResponses(collections))
}

func (api *collectionApi) createCollection(ctx *gin.Context) {
	ctx.Set(constants.Prefix, uuid.MsgWithUUID("create-collection"))

	var reqBody dto.CreateCollectionRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		restErr := errors.NewRestErrorInvalidFormat([]string{}, err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
		return
	}

	if err := api.collectionSvc.CreateCollection(ctx,
		transformers.ToCRUDCreateCollection(reqBody),
	); err != nil {
		restErr := transformers.RestErrTransformerInstance().Transform(err)
		log.Errorln(ctx, restErr)
		transformers.ResponseErr(ctx, restErr)
	}

	transformers.ResponseOK(ctx, nil)
}
