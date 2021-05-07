package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/anonychun/siwi/config"
	"github.com/anonychun/siwi/logger"
	"github.com/gin-gonic/gin"
)

func Start() error {
	router, err := NewRouter()
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Cfg().AppPort),
		Handler: router,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		defer close(idleConnsClosed)

		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			logger.Log().Err(err).Msg("failed to shutdown server")
		}
	}()

	var ipAddr string
	conn, err := net.Dial("udp", "8.8.8.8:80")
	switch {
	case err != nil && config.Cfg().AppLevel == gin.ReleaseMode:
		return err
	case err != nil && config.Cfg().AppLevel == gin.DebugMode:
		ipAddr = "127.0.0.1"
	case err == nil:
		ipAddr = conn.LocalAddr().(*net.UDPAddr).IP.String()
		defer conn.Close()
	}

	logger.Log().Info().Msgf(
		"SIWI starting on [%s%s]",
		ipAddr, httpServer.Addr,
	)

	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	<-idleConnsClosed

	logger.Log().Info().Msg("stopped server gracefully")
	return nil
}
