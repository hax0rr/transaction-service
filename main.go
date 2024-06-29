package main

import (
	"log"
	"transaction-service/cmd"
	"transaction-service/config"
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