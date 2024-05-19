package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func CurrentWorkingDirectoryRepository() (*git.Repository, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return git.PlainOpenWithOptions(
		cwd,
		&git.PlainOpenOptions{
			DetectDotGit:          true,
			EnableDotGitCommonDir: true,
		},
	)
}

func CurrentWorktree() (*git.Worktree, error) {
	repository, err := CurrentWorkingDirectoryRepository()
	if err != nil {
		return nil, err
	}
	return repository.Worktree()
}

func CurrentWorktreeStatus() (git.Status, error) {
	worktree, err := CurrentWorktree()
	if err != nil {
		return nil, err
	}
	return worktree.Status()
}
