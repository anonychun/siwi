package directory

import (
	"os"

	"github.com/anonychun/siwi/config"
)

func Setup() error {
	for _, dir := range []string{
		config.Cfg().DataUpload,
		config.Cfg().DataPublic,
	} {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
