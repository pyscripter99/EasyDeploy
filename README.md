# EasyDeploy

Easy Deploy is a deployment utility designed to be simple, and easy. This project was built to make managing and hosting projects on a remote machine easy, and as simple as possible. It relies on GitHub for hosting the latest deployments/updates.

## Usage

### Help screen

```
A deployment utility and command-line interface for deploying
github projects with ease.

Usage:
  easy-deploy [command]

Available Commands:
  cliConf     Generates a CLI configuration file
  completion  Generate the autocompletion script for the specified shell
  deploy      Runs deploy action
  help        Help about any command
  init        Initialize EasyDeploy in the current directory
  start       Runs start action
  stop        Runs stop action

Flags:
  -t, --auth-token string   Authorization token for server
  -h, --help                help for easy-deploy
  -s, --server string       Agent root url (default "http://127.0.0.1:8900")

Use "easy-deploy [command] --help" for more information about a command.
```
*Help screen from deploy-cli*

### Setting up server configuration
``` bash
$ deploy-cli init
```

This will run a wizard for generating the configuration file `config.yaml`

### Creating a CLI profile
``` bash
$ deploy-cli cliConf
```

This will run a wizard that removes the repetitive process of supplying the authorization token and server ip address by saving a profile to the current directory. `.deploy` will be generated in the current directory

### Starting and stopping deployments and services
``` bash
$ deploy-cli start [process name]
$ deploy-cli stop [process name]
```
The commands respectively will either start or stop the specified process, or all if not set by running the command(s) found in the Easy Deploy configuration (`config.yaml`)

### Upgrading/deploying
``` bash
$ deploy-cli deploy [process name]
```

This command will upgrade/deploy all/the specified processes by stopping the specified process (if any), and pulling the latest change from github with the branch specified in the servers `config.yaml>[process]>git_branch`, typically `production`, then running any deploy commands. (like installing new packages)
