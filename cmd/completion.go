package cmd

import (
	"github.com/lyderic/tools"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:                   "completion",
	Short:                 "Generate bash completion script",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletionFile("calgo.completion")
		tools.PrintGreenln("written completion file: calgo.completion")
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
