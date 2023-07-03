package utils

import (
	"easy-deploy/utils/types"
	"os/exec"
	"strings"
	"time"
)

func makeCommand(command string, workingDirectory string) *exec.Cmd {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = workingDirectory
	return cmd
}

func StartProcess(process types.ConfigProcess) types.WebProcess {
	// Run start commands
	for _, command := range process.Commands.Start {
		cmd := makeCommand(command, process.WorkingDirectory)
		cmd.Start()
		cmd.Wait()
	}
	return types.WebProcess{Name: process.Name, Active: true, StartedAt: time.Now()}
}

func StopProcess(process types.ConfigProcess) types.WebProcess {
	// Run stop commands
	for _, command := range process.Commands.Stop {
		cmd := makeCommand(command, process.WorkingDirectory)
		cmd.Start()
		cmd.Wait()
	}
	return types.WebProcess{Name: process.Name, Active: false, StartedAt: time.Now()}
}
