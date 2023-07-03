/*
Copyright © 2023 Ryder Retzlaff <ryder@retzlaff.family>
*/
package main

import (
	"easy-deploy/agent/server"
	"easy-deploy/utils/types"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func main() {
	// Initialize logger
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	suger := logger.Sugar()

	// Load configuration file
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		suger.Panic("No configuration file found. Generate the configuration file using the CLI")
	}
	server.Configuration = types.Configuration{}
	if err := yaml.Unmarshal(b, &server.Configuration); err != nil {
		suger.Panic("Error reading configuration file")
	}
	suger.Info("Loaded configuration file", server.Configuration)
	suger.Info("Loaded project: ", server.Configuration.Name)

	// Connect database
	server.ConnectDatabase()
	suger.Info("Connected database")

	// Start web server
	suger.Info("Starting web server")
	server.StartServer(logger)
	defer suger.Info("Stopped web server")
}
