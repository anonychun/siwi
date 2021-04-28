package main

import (
	"github.com/anonychun/siwi/directory"
	"github.com/anonychun/siwi/logger"
	"github.com/anonychun/siwi/server"
)

func main() {
	err := directory.Setup()
	if err != nil {
		logger.Log().Err(err).Msg("failed to setup directory")
	}

	err = server.Start()
	if err != nil {
		logger.Log().Err(err).Msg("failed to start server")
	}
}
