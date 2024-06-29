package cmd

import (
	"github.com/hax0rr/transaction-service/config"
	"github.com/spf13/cobra"
)

func NewCLI(config *config.Config) *cobra.Command {
	cli := &cobra.Command{
		Use:   "transaction-service",
		Short: "transaction-service",
		Long:  "transaction service stores the accounts and transaction routines",
	}

	cli.AddCommand(newStartAppServerCommand(config))
	cli.AddCommand(newMigrateCommand(config))
	cli.AddCommand(newRollbackCommand(config))

	return cli
}
