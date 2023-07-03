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
	Use:   "deploy",
	Short: "Runs deploy action",
	Long: `Runs the deploy action on the agent specified server.
This will stop all services, pull the latest changes, and start them up again`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := cmd.Flags().GetString("server")
		if err != nil {
			panic("Error parsing server. " + err.Error())
		}
		httpResp, err := http.Get(server + "/deploy")
		if err != nil {
			panic("Error running deploy. " + err.Error())
		}
		textResp, err := io.ReadAll(httpResp.Body)
		if err != nil {
			panic("Error reading response body. " + err.Error())
		}
		var resp types.WebError
		if err := json.Unmarshal(textResp, &resp); err != nil {
			panic("Error unpacking values. " + err.Error())
		}

		if resp.Error {
			fmt.Println("Error: " + resp.Message)
		} else {
			fmt.Println("Done")
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
