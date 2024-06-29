package main

import (
	"github.com/hax0rr/transaction-service/cmd"
	"github.com/hax0rr/transaction-service/config"
	"log"
)

const configFile = "config/config.yml"

func main() {

	conf, err := config.Load()

	if err != nil {
		log.Fatal("error loading config file: ", err)
	}

	cli := cmd.NewCLI(conf)

	err = cli.Execute()
	if err != nil {
		log.Fatalf("failed to start the app due to %v", err)
	}
}
