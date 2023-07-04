/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"easy-deploy/utils/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy [process name]",
	Short: "Runs deploy action",
	Long: `Runs the deploy action on the agent specified server.
This will stop all services, pull the latest changes, and start them up again`,
	Run: func(cmd *cobra.Command, args []string) {
		server := GetServer(cmd)
		auth := GetAuth(cmd)
		process := ""
		if len(args) > 0 {
			process = args[0]
		}

		req, err := http.NewRequest("GET", server+"/deploy/"+process, nil)
		if err != nil {
			panic("Could not create http request. " + err.Error())
		}
		req.Header.Add("Authorization", auth)
		client := &http.Client{}

		httpResp, err := client.Do(req)
		if err != nil {
			panic("Error running deploy. " + err.Error())
		}
		textResp, err := io.ReadAll(httpResp.Body)
		if err != nil {
			panic("Error reading response body. " + err.Error())
		}

		CheckError(textResp)

		if process == "" {
			var resp types.WebProcessListOrError
			if err := json.Unmarshal(textResp, &resp); err != nil {
				panic("Error unpacking values. " + err.Error())
			}

			if resp.Error {
				fmt.Println("Error: " + resp.Message)
			} else {
				for _, proc := range resp.Processes {
					fmt.Println("Updated: " + proc.Name)
				}
			}
		} else {
			var resp types.WebProcessOrError
			if err := json.Unmarshal(textResp, &resp); err != nil {
				panic("Error unpacking values. " + err.Error())
			}

			if resp.Error {
				fmt.Println("Error: " + resp.Message)
			} else {
				fmt.Println("Updated: " + resp.Process.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
