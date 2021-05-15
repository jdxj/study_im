package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handlers() []func() (string, string, func(ctx *gin.Context)) {
	return []func() (string, string, func(ctx *gin.Context)){
		ping,
	}
}

func ping() (string, string, func(*gin.Context)) {
	method := http.MethodGet
	path := "/ping"

	handler := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
	return method, path, handler
}
