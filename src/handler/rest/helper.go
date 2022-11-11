package rest

import (
	"agrokan-backend/src/business/entity"

	"github.com/gin-gonic/gin"
)

func (r *rest) httpRespSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: message,
			Code:    code,
		},
		Data: data,
	}

	ctx.JSON(code, resp)
}

func (r *rest) httpRespError(ctx *gin.Context, code int, err error) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: err.Error(),
			Code:    code,
		},
		Data: nil,
	}
	ctx.JSON(code, resp)
}
