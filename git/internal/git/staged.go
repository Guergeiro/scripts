package git

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func GetCleanStagedFiles() ([]string, error) {
	worktree, err := CurrentWorktree()
	if err != nil {
		return nil, err
	}
	status, err := worktree.Status()
	if err != nil {
		return nil, err
	}

	selectedFiles := []string{}
	for k, v := range status {
		if v.Worktree != git.Unmodified {
			// Only accept "clean" state files
			continue
		}
		if v.Staging == git.Deleted {
			// Deleted files do not require format
			continue
		}
		absoluteFile := filepath.Join(worktree.Filesystem.Root(), k)
		selectedFiles = append(selectedFiles, absoluteFile)
	}
	return selectedFiles, nil
}
