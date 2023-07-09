package utils

import (
	"easy-deploy/utils/types"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfiguration() (types.Configuration, error) {
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		return types.Configuration{}, fmt.Errorf("no configuration file found. generate the configuration file using the CLI")
	}
	configuration := types.Configuration{}
	if err := yaml.Unmarshal(b, &configuration); err != nil {
		return types.Configuration{}, fmt.Errorf("failed to decode configuration")
	}
	return configuration, nil
}
