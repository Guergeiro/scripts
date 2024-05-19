package prune

import (
	"os"
	"path/filepath"

	"github.com/guergeiro/scripts/git/internal/git"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	origins := []string{"origin"}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	defaultPrivateKeyFile := filepath.Join(home, ".ssh", "id_rsa")
	passphrase := ""

	command := cobra.Command{
		Use:   "prune {...origins}",
		Short: "Prunes remotes passed as arguments, origin otherwise",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				args = append(args, origins...)
			}
			privateKeyFile, err := filepath.Abs(defaultPrivateKeyFile)
			if err != nil {
				return err
			}

			return git.PruneCurrentWorktreeRemotes(
				origins,
				privateKeyFile,
				passphrase,
			)
		},
	}

	command.PersistentFlags().StringVar(
		&defaultPrivateKeyFile,
		"private-key",
		defaultPrivateKeyFile,
		"Path to private key file",
	)

	command.PersistentFlags().StringVarP(
		&passphrase,
		"passphrase",
		"P",
		passphrase,
		"Passphrase",
	)

	return &command
}
