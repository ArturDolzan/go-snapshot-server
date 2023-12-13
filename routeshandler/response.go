package routerhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": msg,
	})
}

func sendSuccess(ctx *gin.Context, msg string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
	})
}
