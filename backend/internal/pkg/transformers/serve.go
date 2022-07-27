package transformers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"decert/internal/pkg/constants"
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
	ctx.AbortWithStatusJSON(int(err.HttpCode), err)
}
