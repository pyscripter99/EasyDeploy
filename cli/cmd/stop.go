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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop [process name]",
	Short: "Runs stop action",
	Long: `Runs the stop action on the agent using the specified server.
	This will start the specified process, or all if none specified`,
	Run: func(cmd *cobra.Command, args []string) {
		server := GetServer(cmd)
		auth := GetAuth(cmd)
		process := ""
		if len(args) > 0 {
			process = args[0]
		}

		req, err := http.NewRequest("GET", server+"/stop/"+process, nil)
		if err != nil {
			panic("Could not create http request. " + err.Error())
		}
		req.Header.Add("Authorization", auth)
		client := &http.Client{}

		httpResp, err := client.Do(req)
		if err != nil {
			panic("Error running stop. " + err.Error())
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

			for _, proc := range resp.Processes {
				fmt.Println("Stopped: " + proc.Name)
			}
		} else {
			var resp types.WebProcessOrError
			if err := json.Unmarshal(textResp, &resp); err != nil {
				panic("Error unpacking values. " + err.Error())
			}
			fmt.Println("Stopped: " + resp.Process.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
