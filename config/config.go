package config

import (
	"github.com/spf13/viper"
	"strings"
	"transaction-service/db"
	"transaction-service/httpserver"
)

const configFile = "config/config.yml"

type Config struct {
	Server   httpserver.Config `mapstructure:"SERVER"`
	Database db.Config         `mapstructure:"DATABASE"`
}

func Load() (*Config, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetConfigName("/config/")
	v.SetConfigFile(configFile)
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	unmarshalErr := v.Unmarshal(&config)

	return &config, unmarshalErr
}
