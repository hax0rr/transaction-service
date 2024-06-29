package cmd

import (
	"github.com/spf13/cobra"
	"transaction-service/config"
	"transaction-service/db"
)

func newMigrateCommand(config *config.Config) *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run database migration",
		RunE: func(command *cobra.Command, args []string) error {
			return db.RunDatabaseMigrations(&config.Database)
		},
	}

	return migrateCmd
}

func newRollbackCommand(config *config.Config) *cobra.Command {
	var rollbackCmd = &cobra.Command{
		Use:   "rollback",
		Short: "Rollback database migration",
		RunE: func(command *cobra.Command, args []string) error {
			return db.RollbackLatestMigration(&config.Database)
		},
	}

	return rollbackCmd
}
