package server

import (
	"net/http"

	"github.com/anonychun/siwi/config"
	"github.com/anonychun/siwi/handler"
	"github.com/anonychun/siwi/webui"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	gin.SetMode(config.Cfg().AppLevel)
	router := gin.New()
	router.Use(gin.Recovery())

	tpl, err := webui.NewView(router).LoadTemplate()
	if err != nil {
		return nil, err
	}
	router.SetHTMLTemplate(tpl)

	appHandler := handler.NewAppHandler()

	router.GET("/", appHandler.Index())
	router.POST("/upload", appHandler.Upload())
	router.StaticFS("/public", http.Dir(config.Cfg().DataPublic))

	return router, nil
}
