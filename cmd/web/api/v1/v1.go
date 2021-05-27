package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers() []func() (string, string, func(ctx *gin.Context)) {
	return []func() (string, string, func(ctx *gin.Context)){
		ping,
		objectsPOST,
		objectsGET,
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
