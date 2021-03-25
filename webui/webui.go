package webui

import (
	"embed"
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

//go:embed templates/* static/*
var embedded embed.FS

type WebUI interface {
	LoadTemplate() (*template.Template, error)
}

type webui struct {
	router *gin.Engine
}

func NewView(router *gin.Engine) WebUI {
	return &webui{router}
}

func (webui *webui) LoadTemplate() (*template.Template, error) {
	filenames, err := findTemplates("templates", &embedded)
	if err != nil {
		return nil, err
	}

	tpl, err := template.ParseFS(embedded, filenames...)
	if err != nil {
		return nil, err
	}

	webui.router.StaticFS("/assets", http.FS(embedded))
	webui.router.GET("/favicon.ico", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/assets/static/img/favicon.ico")
	})

	return tpl, nil
}

func findTemplates(root string, embedFS *embed.FS) ([]string, error) {
	var filenames []string
	entries, err := embedFS.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		entryPath := path.Join(root, entry.Name())
		switch entry.IsDir() {
		case true:
			temp, err := findTemplates(entryPath, embedFS)
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, temp...)
		default:
			filenames = append(filenames, entryPath)
		}

	}

	return filenames, nil
}
