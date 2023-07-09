/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// reloadCmd represents the reload command
var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reloads the agent's configuration file",
	Long:  `Reloads the agent's configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		server := GetServer(cmd)
		auth := GetAuth(cmd)

		req, err := http.NewRequest("GET", server+"/reload/config", nil)
		if err != nil {
			panic("Could not create http request. " + err.Error())
		}
		req.Header.Add("Authorization", auth)
		client := &http.Client{}

		httpResp, err := client.Do(req)
		if err != nil {
			panic("Error reloading config. " + err.Error())
		}
		textResp, err := io.ReadAll(httpResp.Body)
		if err != nil {
			panic("Error reading response body. " + err.Error())
		}

		CheckError(textResp)
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}
