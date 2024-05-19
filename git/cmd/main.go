package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/guergeiro/scripts/git/cmd/prune"
	"github.com/guergeiro/scripts/git/cmd/root"
	"github.com/guergeiro/scripts/git/cmd/staged"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Args: cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute(args)
		},
		Hidden:                true,
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagParsing:    true,
		DisableAutoGenTag:     true,
		DisableSuggestions:    true,
		DisableFlagsInUseLine: true,
	}

	rootCmd.AddCommand(root.Command())
	rootCmd.AddCommand(prune.Command())
	rootCmd.AddCommand(staged.Command())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func execute(args []string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}
