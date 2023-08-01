/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"easy-deploy/utils/types"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Example configuration
var exampleConfiguration = types.Configuration{Name: "example-project", Processes: []types.ConfigProcess{
	{Name: "hello-console", WorkingDirectory: ".", GitUrl: "https://github.com/example/repo.git", GitBranch: "production", GitUsername: "github username (only if repo is private)", GitToken: "github PAC token (only if repo is private)", Commands: types.ConfigProcessCommands{
		Start:  []string{"echo \"Hello World!\""},
		Deploy: []string{"echo \"Apply any changes you need after updating\""},
		Stop:   []string{"echo \"Stopping\""},
	}},
}}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize EasyDeploy in the current directory",
	Long: `The init command sets up the required files for Easy Deploy to run.
Optionally will run a wizard to help configure the configuration files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating configuration file")
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic("Error reading flags")
		}
		if example {
			// Write example configuration
			if err := writeConfiguration(exampleConfiguration); err != nil {
				panic("Error saving configuration")
			}
			fmt.Println("Saved configuration to: 'config.yaml'")
			return
		}

		// Wizard
		configuration := types.Configuration{}
		survey.AskOne(&survey.Input{Message: "Project name:"}, &configuration.Name)
		for {
			processName := ""
			workingDirectory := ""
			gitUrl := ""
			gitBranch := ""
			gitUsername := ""
			gitPassword := ""
			commands := types.ConfigProcessCommands{}
			another := false
			privacy := ""
			survey.AskOne(&survey.Input{Message: "Process name:"}, &processName)
			survey.AskOne(&survey.Input{Message: "Process working directory:"}, &workingDirectory)
			survey.AskOne(&survey.Input{Message: "Process git url:"}, &gitUrl)
			survey.AskOne(&survey.Select{Message: "Choose your repository privacy", Options: []string{"private", "public"}}, &privacy)
			if privacy == "private" {
				survey.AskOne(&survey.Input{Message: "Git username: "}, &gitUsername)
				survey.AskOne(&survey.Input{Message: "Git password/token: "}, &gitPassword)
			}
			survey.AskOne(&survey.Input{Message: "Process git branch:"}, &gitBranch)
			fmt.Println("Enter start commands:")
			for {
				command := ""
				survey.AskOne(&survey.Input{Message: ""}, &command)
				if !(command == "") {
					commands.Start = append(commands.Start, command)
				} else {
					break
				}
			}
			fmt.Println("Enter deploy commands:")
			for {
				command := ""
				survey.AskOne(&survey.Input{Message: ""}, &command)
				if !(command == "") {
					commands.Deploy = append(commands.Deploy, command)
				} else {
					break
				}
			}
			fmt.Println("Enter stop commands:")
			for {
				command := ""
				survey.AskOne(&survey.Input{Message: ""}, &command)
				if !(command == "") {
					commands.Stop = append(commands.Stop, command)
				} else {
					break
				}
			}
			survey.AskOne(&survey.Confirm{Message: "Add another process?", Default: false}, &another)
			process := types.ConfigProcess{Name: processName, WorkingDirectory: workingDirectory, GitUrl: gitUrl, GitBranch: gitBranch, Commands: commands}
			if privacy == "private" {
				process.GitUsername = gitUsername
				process.GitToken = gitPassword
			}
			configuration.Processes = append(configuration.Processes, process)
			if !another {
				break
			}
		}

		// Generate and add token
		token := uuid.New().String()
		if token == "" {
			fmt.Println("Error generating authentication token")
			os.Exit(1)
		}
		configuration.AuthToken = token
		fmt.Printf("Your authentication token is: '%s'\n IMPORTANT: Keep this private and in a safe location, without it you will not be able to send commands to EasyDeploy\n", token)

		// Save configuration
		if err := writeConfiguration(configuration); err != nil {
			panic("Error saving configuration")
		}
		fmt.Println("Saved configuration to: 'config.yaml'")
	},
}

func writeConfiguration(configuration types.Configuration) error {
	configString, err := yaml.Marshal(configuration)
	if err != nil {
		return err
	}
	os.WriteFile("config.yaml", configString, 0644)
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("example", "e", false, "Skips the wizard and generates an example configuration file")
}
