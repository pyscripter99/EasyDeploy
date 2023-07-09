package types

type ConfigProcessCommands struct {
	Start  []string `yaml:"start"`
	Deploy []string `yaml:"deploy"`
	Stop   []string `yaml:"stop"`
}

type ConfigProcess struct {
	Name             string                `yaml:"name"`
	WorkingDirectory string                `yaml:"working_directory"`
	GitUrl           string                `yaml:"git_url"`
	GitBranch        string                `yaml:"git_branch"`
	GitUsername      string                `yaml:"git_username"`
	GitToken         string                `yaml:"git_token"`
	Commands         ConfigProcessCommands `yaml:"commands"`
}

type Configuration struct {
	Name      string          `yaml:"name"`
	Processes []ConfigProcess `yaml:"processes"`
	AuthToken string          `yaml:"auth_token"`
}
