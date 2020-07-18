package server

import (
	"net/http"

	"github.com/anonychun/siwi/internal/service/infra/config"
	"github.com/anonychun/siwi/internal/service/infra/view"
	uploadHandler "github.com/anonychun/siwi/internal/service/upload/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	gin.SetMode(config.Config().AppLevel)
	router := gin.New()

	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 100000 << 20 // 100GB

	view := view.NewView(router)
	template, err := view.LoadTemplate()
	if err != nil {
		return nil, err
	}
	router.SetHTMLTemplate(template)

	uploadHandler := uploadHandler.NewUploadHTTPHandler()

	router.GET("/", uploadHandler.Get())
	router.POST("/upload", uploadHandler.Post())
	router.StaticFS("/public", http.Dir(config.Config().DataPublic))

	return router, nil
}
