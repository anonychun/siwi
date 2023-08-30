package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"sync"

	"github.com/anonychun/siwi/config"
	"github.com/anonychun/siwi/logger"
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
				c.SaveUploadedFile(file, path.Join(config.Cfg().DataUpload, file.Filename))
				logger.Log().Info().Msgf("%s uploaded: %s (%s)", clientIP, file.Filename, ByteCountSI(file.Size))
			}(file)
		}

		wg.Wait()
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
