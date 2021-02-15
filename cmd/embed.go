package cmd

import (
	. "calgo/internal"

	"github.com/spf13/cobra"
)

var embedCmd = &cobra.Command{
	Use:   "embed",
	Short: "embed metadata into epubs",
	Run: func(cmd *cobra.Command, args []string) {
		Embed()
	},
}

func init() {
	rootCmd.AddCommand(embedCmd)
}
