package db

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewDB(config *Config) (*sqlx.DB, error) {
	d, err := sqlx.Open(config.Driver, config.URL())
	if err != nil {
		return nil, err
	}

	if err = d.Ping(); err != nil {
		return nil, err
	}

	return d, err
}

func RollbackLatestMigration(config *Config) error {
	m, err := migrate.New(config.MigrationPath, config.URL())
	if err != nil {
		return err
	}

	err = m.Steps(-1)
	if errors.Is(err, migrate.ErrNoChange) || err == nil {
		log.Info().Msg("DB last migration rollback successful")
		return nil
	}

	return err
}

func RunDatabaseMigrations(config *Config) error {
	db, err := sql.Open(config.Driver, config.URL())
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(config.MigrationPath, config.Driver, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) || err == nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info().Msg("DB already upto date")
		} else {
			log.Info().Msg("DB Migrate successful")
		}
		return nil
	}

	return err
}
