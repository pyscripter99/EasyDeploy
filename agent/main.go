/*
Copyright © 2023 Ryder Retzlaff <ryder@retzlaff.family>
*/
package main

import (
	"easy-deploy/agent/server"
	"easy-deploy/agent/utils"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	suger := logger.Sugar()
	zap.ReplaceGlobals(logger)

	// Load configuration file
	config, err := utils.LoadConfiguration()
	if err != nil {
		suger.Fatal("Failed to load config. " + err.Error())
	}

	server.Configuration = config

	suger.Info("Loaded configuration file", server.Configuration)
	suger.Info("Loaded project: ", server.Configuration.Name)

	// Start web server
	suger.Info("Starting web server")
	server.StartServer(logger)
	defer suger.Info("Stopped web server")
}
