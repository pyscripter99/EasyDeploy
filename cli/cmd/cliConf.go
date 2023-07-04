/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"easy-deploy/utils/types"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// cliConfCmd represents the cliConf command
var cliConfCmd = &cobra.Command{
	Use:   "cliConf",
	Short: "Generates a CLI configuration file",
	Long:  `The cliConf command will generate a configuration file for the client to save repeatedly entering server values`,
	Run: func(cmd *cobra.Command, args []string) {
		questions := []*survey.Question{
			{
				Name:     "Name",
				Prompt:   &survey.Input{Message: "Enter name for profile:"},
				Validate: survey.Required,
			},
			{
				Name:     "Server",
				Prompt:   &survey.Input{Message: "Enter server ip:"},
				Validate: survey.Required,
			},
			{
				Name:     "AuthToken",
				Prompt:   &survey.Password{Message: "Enter authentication token:"},
				Validate: survey.Required,
			},
		}
		profile := types.CliConfig{}
		survey.Ask(questions, &profile)

		profileStr, err := yaml.Marshal(profile)
		if err != nil {
			panic("Error marshaling profile. " + err.Error())
		}

		os.WriteFile(".deploy", profileStr, 0600)
	},
}

func init() {
	rootCmd.AddCommand(cliConfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliConfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliConfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
