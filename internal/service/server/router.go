package server

import (
	"path"

	rice "github.com/GeertJohan/go.rice"
	"github.com/anonychun/siwi/internal/pkg/tpl"
	"github.com/anonychun/siwi/internal/service/infra/config"
	uploadHandler "github.com/anonychun/siwi/internal/service/upload/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	gin.SetMode(config.Config().AppLevel)
	router := gin.New()

	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 100000 << 20 // 100GB

	tpl, err := tpl.LoadTemplate()
	if err != nil {
		return nil, err
	}
	router.SetHTMLTemplate(tpl)

	uploadHandler := uploadHandler.NewUploadHTTPHandler()

	router.GET("/", uploadHandler.Get())
	router.POST("/upload", uploadHandler.Post())

	router.StaticFS("/public", rice.MustFindBox(path.Join("../../../", config.Config().DataPublic)).HTTPBox())
	router.StaticFS("/static", rice.MustFindBox(path.Join("../../../", config.Config().Static)).HTTPBox())

	return router, nil
}
