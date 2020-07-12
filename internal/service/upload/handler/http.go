package handler

import (
	"mime/multipart"
	"net/http"
	"path"
	"sync"

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

		var wg sync.WaitGroup
		clientIP := c.ClientIP()

		files := form.File["files"]
		for _, file := range files {
			wg.Add(1)
			go func(file *multipart.FileHeader) {
				c.SaveUploadedFile(file, path.Join(config.Config().DataUpload, file.Filename))
				logger.Info(clientIP, " uploaded: ", file.Filename)
				wg.Done()
			}(file)
		}

		wg.Wait()
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
