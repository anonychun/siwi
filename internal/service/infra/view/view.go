package view

import (
	"html/template"
	"os"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

type View interface {
	LoadTemplate() (*template.Template, error)
}

type view struct {
	router *gin.Engine
}

func NewView(router *gin.Engine) View {
	return &view{router}
}

func (view *view) LoadTemplate() (*template.Template, error) {
	tpl := template.New("")

	box, err := rice.FindBox("../../../../template")
	if err != nil {
		return nil, err
	}

	files, err := findTemplates(box)
	if err != nil {
		return nil, err
	}

	for _, name := range files {
		tplString, err := box.String(name)
		if err != nil {
			return nil, err
		}
		tpl.New(name).Parse(tplString)
	}

	view.router.StaticFS("/static", rice.MustFindBox("../../../../static").HTTPBox())

	return tpl, nil
}

func findTemplates(box *rice.Box) ([]string, error) {
	var files []string

	err := box.Walk("", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return err
	})

	return files, err
}
