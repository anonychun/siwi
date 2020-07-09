package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/anonychun/siwi/internal/pkg/ip"
	"github.com/anonychun/siwi/internal/service/infra/config"
	"github.com/anonychun/siwi/internal/service/infra/logger"
)

func Start() error {
	appPort := ":" + config.Config().AppPort

	server := NewRouter()
	httpServer := &http.Server{
		Addr:    appPort,
		Handler: server,
	}

	logger.Info("SIWI starting on ", ip.GetLocalIP(), appPort)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			logger.Error("Received an Interrupt Signal", err)
		}
	}()

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Error("HTTP Server Failed", err)
		close(idleConnsClosed)
	}

	<-idleConnsClosed

	logger.Info("Stopped server gracefully")
	return nil
}
