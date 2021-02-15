package cmd

import (
	. "calgo/internal"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:                   "completion",
	Short:                 "Generate bash completion script",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletionFile("calgo.completion")
		Green("written completion file: calgo.completion\n")
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
