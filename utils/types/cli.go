package types

type CliConfig struct {
	Name      string `yaml:"name"`
	Server    string `yaml:"server"`
	AuthToken string `yaml:"auth_token"`
}
