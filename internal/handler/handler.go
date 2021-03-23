package handler

import (
	"mime/multipart"
	"net/http"
	"path"
	"sync"

	"github.com/anonychun/siwi/internal/config"
	"github.com/anonychun/siwi/internal/logger"
	"github.com/gin-gonic/gin"
)

type AppHandler interface {
	Index() gin.HandlerFunc
	Upload() gin.HandlerFunc
}

type appHandler struct{}

func NewAppHandler() AppHandler {
	return &appHandler{}
}

func (appHandler) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func (appHandler) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			logger.Log().Err(err).Msg("failed to recieve form data")
			return
		}

		var wg sync.WaitGroup
		clientIP := c.ClientIP()

		files := form.File["files"]
		wg.Add(len(files))

		for _, file := range files {
			go func(file *multipart.FileHeader) {
				defer wg.Done()
				c.SaveUploadedFile(file, path.Join(config.Config().DataUpload, file.Filename))
				logger.Log().Info().Msgf("%s uploaded: %s", clientIP, file.Filename)
			}(file)
		}

		wg.Wait()
		c.Redirect(http.StatusSeeOther, "/")
	}
}
