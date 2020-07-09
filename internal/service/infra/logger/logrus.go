package logger

import (
	"os"

	"github.com/anonychun/siwi/internal/service/infra/config"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Setup() {
	if logger != nil {
		return
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(config.Config().LogLevel)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.SetLevel(logLevel)

	logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logLevel,
	}
}

func AddHook(hook logrus.Hook) {
	logger.Hooks.Add(hook)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func LogErrors(err error, action string, args ...interface{}) {
	if err != nil {
		Error("Failed to ", action, " with errors ", err, " and data ", args)
	} else {
		Debug("Success to ", action, " with data ", args)
	}
}
