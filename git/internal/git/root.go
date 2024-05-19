package git

func WorktreeRoot() (string, error) {
	repository, err := CurrentWorkingDirectoryRepository()
	if err != nil {
		return "", err
	}
	// Workaround: https://github.com/go-git/go-git/issues/405
	worktree, err := repository.Worktree()
	if err != nil {
		return "", err
	}
	return worktree.Filesystem.Root(), nil
}
