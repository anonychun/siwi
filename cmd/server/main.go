package main

import (
	"fmt"

	"github.com/anonychun/siwi/internal/service/infra/directory"
	"github.com/anonychun/siwi/internal/service/infra/logger"
	"github.com/anonychun/siwi/internal/service/server"
)

func main() {
	logger.Setup()
	directory.Setup()

	err := server.Start()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to run server: %s", err.Error()))
	}
}
