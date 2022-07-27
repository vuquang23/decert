package api

import (
	"github.com/gin-gonic/gin"

	"decert/internal/pkg/constants"
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

	log.Debugln(ctx, "hello", "my", "fen")
	log.Debugf(ctx, "%d %d %s", 2, 3, "1200")
	// example
	transformers.ResponseOK(ctx, "Hello Hoang Tung")
}

func (api *collectionApi) createCollection(ctx *gin.Context) {
	transformers.ResponseErr(ctx, errors.NewRestErrorNotFound([]string{"wallet"}, nil))
}
