package config

import (
	"github.com/hax0rr/transaction-service/db"
	"github.com/hax0rr/transaction-service/httpserver"
	"github.com/spf13/viper"
	"strings"
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
