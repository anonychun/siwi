package server

import (
	"net/http"
	"path"

	"github.com/anonychun/siwi/internal/service/infra/config"
	uploadHandler "github.com/anonychun/siwi/internal/service/upload/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.Config().AppLevel)
	router := gin.New()

	router.Use(gin.Recovery())
	router.LoadHTMLGlob(path.Join(config.Config().Template, "*"))
	router.MaxMultipartMemory = 10000 << 20

	uploadHandler := uploadHandler.NewUploadHTTPHandler()

	router.GET("/", uploadHandler.Get())
	router.POST("/upload", uploadHandler.Post())

	router.StaticFS("/public", http.Dir(path.Join(config.Config().DataPublic)))
	router.StaticFS("/static", http.Dir(path.Join(config.Config().Static)))

	return router
}
