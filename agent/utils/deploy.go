package utils

import (
	"easy-deploy/utils/types"

	"github.com/go-git/go-git/v5/plumbing/transport/http"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func Deploy(process types.ConfigProcess) (commit string, err error) {
	// Pull changes
	repo, err := git.PlainOpen(process.WorkingDirectory)
	if err != nil {
		return "", err
	}

	work, err := repo.Worktree()
	if err != nil {
		return "", err
	}

	// Check for authentication
	username := process.GitUsername
	password := process.GitToken

	if username == "" {
		if err := work.Pull(&git.PullOptions{ReferenceName: plumbing.ReferenceName("refs/heads/" + process.GitBranch)}); err != nil {
			return "", err
		}
	} else {
		if err := work.Pull(&git.PullOptions{ReferenceName: plumbing.ReferenceName("refs/heads/" + process.GitBranch), Auth: &http.BasicAuth{Username: username, Password: password}}); err != nil {
			return "", err
		}
	}

	// Run deploy script
	for _, command := range process.Commands.Deploy {
		cmd := makeCommand(command, process.WorkingDirectory)
		cmd.Start()
		cmd.Wait()
	}

	// Get commit
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	com, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return "", err
	}
	return com.ID().String(), nil
}
