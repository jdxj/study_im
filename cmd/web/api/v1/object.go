package v1

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	Temp string
)

func objectsPOST() (string, string, func(*gin.Context)) {
	method := http.MethodPost
	path := "/objects"

	handler := func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		dst := filepath.Join(Temp, file.Filename)
		err = ctx.SaveUploadedFile(file, dst)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
	return method, path, handler
}

func objectsGET() (string, string, func(*gin.Context)) {
	method := http.MethodGet
	path := "/objects/:name"

	handler := func(ctx *gin.Context) {
		name := ctx.Param("name")

		fp := filepath.Join(Temp, name)
		ctx.File(fp)
	}
	return method, path, handler
}
