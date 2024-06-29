package db

import "fmt"

type Config struct {
	Name          string `mapstructure:"NAME"`
	Host          string `mapstructure:"HOST"`
	User          string `mapstructure:"USER"`
	Password      string `mapstructure:"PASSWORD"`
	Port          int    `mapstructure:"PORT"`
	Driver        string `mapstructure:"DRIVER"`
	TimeoutInMs   int    `mapstructure:"TIMEOUT_IN_MS"`
	MigrationPath string `mapstructure:"MIGRATION_PATH"`
}

func (db Config) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
