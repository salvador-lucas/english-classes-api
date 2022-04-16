package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/salvador-lucas/english-classes-api/views"
)

func SetResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, views.BaseResponse{
		StatusCode: statusCode,
		Data:       data,
	})
}
