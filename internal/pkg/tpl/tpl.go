package tpl

import (
	"html/template"
	"os"
	"path"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/anonychun/siwi/internal/service/infra/config"
)

func LoadTemplate() (*template.Template, error) {
	tpl := template.New("")

	box, err := rice.FindBox(path.Join("../../../", config.Config().Template))
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

	return tpl, nil
}

func findTemplates(box *rice.Box) ([]string, error) {
	var files []string

	root := "/"
	err := box.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return err
	})

	return files, err
}
