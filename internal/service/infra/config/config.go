package config

import (
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

func getStringDefault(viper *viper.Viper, key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

var once sync.Once
var config AppConfig

type AppConfig struct {
	viper      *viper.Viper
	AppLevel   string
	LogLevel   string
	AppPort    string
	Template   string
	Static     string
	DataUpload string
	DataPublic string
}

func load() AppConfig {
	fang := viper.New()

	fang.SetConfigName("config")
	fang.SetConfigType("yaml")
	fang.SetEnvPrefix("APP")
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()

	fang.AddConfigPath(".")
	value, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(value)
	}

	_ = fang.ReadInConfig()

	appConfig := AppConfig{
		viper:      fang,
		AppLevel:   getStringDefault(fang, "app.level", "release"),
		LogLevel:   getStringDefault(fang, "log.level", "debug"),
		AppPort:    getStringDefault(fang, "app.port", "4444"),
		Template:   getStringDefault(fang, "template", "template"),
		Static:     getStringDefault(fang, "static", "static"),
		DataUpload: getStringDefault(fang, "data.upload", "data/upload"),
		DataPublic: getStringDefault(fang, "data.public", "data/public"),
	}

	return appConfig
}

func Config() AppConfig {
	once.Do(func() {
		config = load()
	})

	return config
}
