package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(zerolog.ConsoleWriter{
	Out:        os.Stdout,
	NoColor:    true,
	TimeFormat: time.Kitchen,
}).With().Timestamp().Logger().Level(zerolog.GlobalLevel())

func Log() *zerolog.Logger { return &logger }
