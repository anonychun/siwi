package main

import (
	"github.com/anonychun/siwi/internal/directory"
	"github.com/anonychun/siwi/internal/logger"
	"github.com/anonychun/siwi/internal/server"
)

func main() {
	err := directory.Setup()
	if err != nil {
		logger.Log().Err(err).Msg("failed to setup directory")
	}

	err = server.Start()
	if err != nil {
		logger.Log().Err(err).Msg("failed to run server")
	}
}
