package directory

import (
	"os"

	"github.com/anonychun/siwi/internal/service/infra/config"
	"github.com/anonychun/siwi/internal/service/infra/logger"
)

func Setup() {
	err := os.MkdirAll(config.Config().DataUpload, os.ModePerm)
	if err != nil {
		logger.Panic(err.Error())
	}

	err = os.MkdirAll(config.Config().DataPublic, os.ModePerm)
	if err != nil {
		logger.Panic(err.Error())
	}
}
