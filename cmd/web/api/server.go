package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/jdxj/study_im/cmd/web/api/v1"
)

func NewServer() *gin.Engine {
	engine := gin.Default()
	api := engine.Group("/api")

	registerHandler(api, v1.Handlers())
	return engine
}

func registerHandler(rg *gin.RouterGroup, handlers []func() (string, string, func(ctx *gin.Context))) {
	for _, build := range handlers {
		method, path, handler := build()

		switch method {
		case http.MethodGet:
			rg.GET(path, handler)
		case http.MethodPost:
			rg.POST(path, handler)
		case http.MethodConnect:
		case http.MethodDelete:
			rg.DELETE(path, handler)
		case http.MethodHead:
			rg.HEAD(path, handler)
		case http.MethodOptions:
			rg.OPTIONS(path, handler)
		case http.MethodPatch:
			rg.PATCH(path, handler)
		case http.MethodPut:
			rg.PUT(path, handler)
		case http.MethodTrace:
		}
	}
}

