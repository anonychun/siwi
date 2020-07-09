package handler

import (
	"net/http"
	"path"

	"github.com/anonychun/siwi/internal/service/infra/config"
	"github.com/anonychun/siwi/internal/service/infra/logger"
	"github.com/gin-gonic/gin"
)

type UploadHTTPHandler interface {
	Get() gin.HandlerFunc
	Post() gin.HandlerFunc
}

type uploadHTTPHandler struct {
}

func NewUploadHTTPHandler() UploadHTTPHandler {
	return &uploadHTTPHandler{}
}

func (httpHandler *uploadHTTPHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	}
}

func (httpHandler *uploadHTTPHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			logger.Error(err.Error())
			return
		}

		files := form.File["files"]
		for _, file := range files {
			c.SaveUploadedFile(file, path.Join(config.Config().DataUpload, file.Filename))
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
