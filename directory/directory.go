package directory

import (
	"os"

	"github.com/anonychun/siwi/config"
)

func Setup() error {
	for _, dir := range []string{
		config.Config().DataUpload,
		config.Config().DataPublic,
	} {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
