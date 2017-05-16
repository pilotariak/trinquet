package utils

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	completionShells = map[string]func(out io.Writer, cmd *cobra.Command) error{
		"bash": runCompletionBash,
		"zsh":  runCompletionZsh,
	}
)

func NewCompletionCommand(out io.Writer, example string) *cobra.Command {
	shells := []string{}
	for s := range completionShells {
		shells = append(shells, s)
	}

	cmd := &cobra.Command{
		Use:     "completion SHELL",
		Short:   "Output shell completion code for the specified shell (bash or zsh)",
		Example: example,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := RunCompletion(out, cmd, args); err != nil {
				fmt.Println(RedOut(err))
			}
			return nil
		},
		ValidArgs: shells,
	}

	return cmd
}

func RunCompletion(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Shell not specified.")
	}
	if len(args) > 1 {
		return fmt.Errorf("Too many arguments. Expected only the shell type.")
	}
	run, found := completionShells[args[0]]
	if !found {
		return fmt.Errorf("Unsupported shell type %q.", args[0])
	}

	return run(out, cmd.Parent())
}

func runCompletionBash(out io.Writer, diabloctl *cobra.Command) error {
	return diabloctl.GenBashCompletion(out)
}

func runCompletionZsh(out io.Writer, diabloctl *cobra.Command) error {
	return fmt.Errorf("Zsh is currently Unsupported")
}
