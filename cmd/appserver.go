package cmd

import (
	"github.com/hax0rr/transaction-service/app"
	"github.com/hax0rr/transaction-service/config"
	"github.com/hax0rr/transaction-service/httpserver"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func newStartAppServerCommand(config *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "app",
		Short: "Start the transport-service app server",
		RunE: func(command *cobra.Command, args []string) error {
			return runAppServer(config)
		},
	}
}

func runAppServer(config *config.Config) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	deps, err := app.NewDependencies(config)
	if err != nil {
		return err
	}

	router := app.NewRouter(deps)

	server := httpserver.New(config.Server, router)

	err = server.Start()
	if err != nil {
		return err
	}

	<-c
	return server.Shutdown()
}
