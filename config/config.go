package config

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func getStringDefault(v *viper.Viper, key string, value string) string {
	v.SetDefault(key, value)
	return v.GetString(key)
}

func getIntDefault(v *viper.Viper, key string, value int) int {
	v.SetDefault(key, value)
	return v.GetInt(key)
}

type Config struct {
	AppLevel string
	AppPort  int

	DataUpload string
	DataPublic string
}

func load() Config {
	fang := viper.New()
	fang.AutomaticEnv()

	fang.SetConfigName("config")
	fang.SetConfigType("yaml")

	fang.AddConfigPath(".")
	value, available := os.LookupEnv("CONFIG_LOCATION")
	if available {
		fang.AddConfigPath(value)
	}

	_ = fang.ReadInConfig()
	homeDir, _ := os.UserHomeDir()

	return Config{
		AppLevel:   getStringDefault(fang, "app.level", gin.ReleaseMode),
		AppPort:    getIntDefault(fang, "app.port", 1427),
		DataUpload: path.Join(homeDir, "Documents/siwi/upload"),
		DataPublic: path.Join(homeDir, "Documents/siwi/public"),
	}
}

var config = load()

func Cfg() *Config { return &config }
