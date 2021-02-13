package cmd

import (
	. "calgo/internal"

	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "index files on filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		var i Index
		i.Load()
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
