package staged

import (
	"strings"

	"github.com/guergeiro/scripts/git/internal/git"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := cobra.Command{
		Use:   "staged",
		Short: "Selects the current full staged files",
		Long: `This command takes the current full staged (that are not modified) and prints
it to stdout.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			separator, err := cmd.Flags().GetString("separator")
			if err != nil {
				return err
			}
			result, err := execute(separator)
			_, err = cmd.OutOrStdout().Write([]byte(result))
			return err
		},
	}
	command.Flags().StringP("separator", "S", "\n", "Separator of files")

	return &command
}

func execute(separator string) (string, error) {
	selectedFiles, err := git.GetCleanStagedFiles()
	if err != nil {
		return "", err
	}
	return strings.Join(selectedFiles, separator), nil
}
