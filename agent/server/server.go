package server

import (
	"easy-deploy/agent/utils"
	"easy-deploy/utils/types"
	"errors"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Configuration types.Configuration

func FindProcess(name string) (types.ConfigProcess, error) {
	for _, process := range Configuration.Processes {
		if process.Name == name {
			return process, nil
		}
	}
	return types.ConfigProcess{}, errors.New("no process with that name found")
}

func StartServer(logger *zap.Logger) {
	// Setup server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// Routes
	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"online": true, "services": [...]string{"API", "Agent"}, "processes": []types.WebProcess{}})
	})

	r.GET("/start", func(ctx *gin.Context) {
		processes := []types.WebProcess{}
		for _, process := range Configuration.Processes {
			processes = append(processes, utils.StartProcess(process))

		}
		ctx.JSON(http.StatusOK, processes)
	})

	r.GET("/start/:process", func(ctx *gin.Context) {
		process, err := FindProcess(ctx.Param("process"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, types.WebProcessOrError{WebError: types.WebError{Error: true, Message: err.Error()}})
		}
		ctx.JSON(http.StatusOK, types.WebProcessOrError{Process: utils.StartProcess(process), WebError: types.WebError{}})
	})

	r.GET("/stop", func(ctx *gin.Context) {
		processes := []types.WebProcess{}
		for _, process := range Configuration.Processes {
			processes = append(processes, utils.StopProcess(process))
		}
		ctx.JSON(http.StatusOK, processes)
	})

	r.GET("/stop/:process", func(ctx *gin.Context) {
		process, err := FindProcess(ctx.Param("process"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, types.WebProcessOrError{WebError: types.WebError{Error: true, Message: err.Error()}})
		}
		ctx.JSON(http.StatusOK, types.WebProcessOrError{Process: utils.StopProcess(process)})
	})

	r.GET("/deploy", func(ctx *gin.Context) {
		// Deploy update to every process
		for _, process := range Configuration.Processes {
			// Stop process
			utils.StopProcess(process)

			// Update
			if _, err := utils.Deploy(process); err != nil {
				utils.StartProcess(process)
				ctx.JSON(http.StatusInternalServerError, types.WebError{Error: true, Message: err.Error()})
				return
			}

			// Start process
			utils.StartProcess(process)
		}
		ctx.JSON(http.StatusOK, types.WebError{})
	})

	// Run the server
	r.Run(":8900")
}
