package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var once sync.Once
var logger zerolog.Logger

func Log() *zerolog.Logger {
	once.Do(func() {
		logger = zerolog.New(os.Stdout)
	})
	return &logger
}
