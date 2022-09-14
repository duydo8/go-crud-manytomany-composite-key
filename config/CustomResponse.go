package config

import (
	"github.com/gin-gonic/gin"
)

func CustomResponse(context *gin.Context, code int, message string, C any) {
	context.JSON(
		code,
		gin.H{
			"data":    C,
			"message": message,
			"code":    code,
		},
	)
}
