package config

import (
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

func GetStringDefault(viper *viper.Viper, key string, defaultValue string) string {
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
	DocsPath   string
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

	err := fang.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	appConfig := AppConfig{
		viper:      fang,
		AppLevel:   GetStringDefault(fang, "app.level", "debug"),
		LogLevel:   GetStringDefault(fang, "log.level", "debug"),
		AppPort:    GetStringDefault(fang, "app.port", "4444"),
		Template:   fang.GetString("template"),
		Static:     fang.GetString("static"),
		DataUpload: fang.GetString("data.upload"),
		DataPublic: fang.GetString("data.public"),
		DocsPath:   fang.GetString("docs.path"),
	}

	return appConfig
}

func Config() AppConfig {
	once.Do(func() {
		config = load()
	})

	return config
}
