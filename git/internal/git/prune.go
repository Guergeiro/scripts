package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func PruneCurrentWorktreeRemotes(
	origins []string,
	privateKeyFile string,
	passphrase string,
) error {
	repository, err := CurrentWorkingDirectoryRepository()
	if err != nil {
		return err
	}
	publicKeys, err := ssh.NewPublicKeysFromFile(
		"git",
		privateKeyFile,
		passphrase,
	)
	if err != nil {
		return err
	}
	for _, value := range origins {
		err := repository.Fetch(&git.FetchOptions{
			RemoteName: value,
			Prune:      true,
			Auth:       publicKeys,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
