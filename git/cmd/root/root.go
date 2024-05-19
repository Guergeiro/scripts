package root

import (
	"github.com/guergeiro/scripts/git/internal/git"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := cobra.Command{
		Use:   "root",
		Args:  cobra.NoArgs,
		Short: "Outputs the current worktree root directory to stdin",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := git.WorktreeRoot()
			_, err = cmd.OutOrStdout().Write([]byte(result))
			return err
		},
	}

	return &command
}
