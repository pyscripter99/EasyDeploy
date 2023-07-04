package cmd

import (
	"easy-deploy/utils/types"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func GetProfile() (types.CliConfig, error) {
	if _, err := os.Stat(".deploy"); err == nil {
		profileString, err := os.ReadFile(".deploy")
		if err != nil {
			return types.CliConfig{}, err
		}
		var profile types.CliConfig
		if err := yaml.Unmarshal(profileString, &profile); err != nil {
			return types.CliConfig{}, err
		}
		return profile, nil
	}
	return types.CliConfig{}, nil
}

func GetAuth(cmd *cobra.Command) string {
	var auth string

	profile, err := GetProfile()
	if err != nil {
		panic("Error loading profile: " + err.Error())
	}

	if profile.AuthToken == "" {
		auth, err = cmd.Flags().GetString("auth-token")
		if err != nil {
			panic("Error parsing auth token. " + err.Error())
		}
	} else {
		auth = profile.AuthToken
	}

	return auth
}

func GetServer(cmd *cobra.Command) string {
	var server string

	profile, err := GetProfile()
	if err != nil {
		panic("Error loading profile: " + err.Error())
	}

	if profile.Server == "" {
		server, err = cmd.Flags().GetString("server")
		if err != nil {
			panic("Error parsing server. " + err.Error())
		}
	} else {
		server = profile.Server
	}

	return server
}

func CheckError(resp []byte) {
	var webErr types.WebError
	if err := json.Unmarshal(resp, &webErr); err != nil {
		panic("Error decoding response. " + err.Error())
	}

	if webErr.Error {
		fmt.Println("Server error: " + webErr.Message)
		os.Exit(1)
	}
}
