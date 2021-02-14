package cmd

import (
	"calgo/checks"
	. "calgo/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check consistency of database and filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		check()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func check() {
	calibreBooks := LoadFromCalibre()
	fmt.Printf("[Checking %d books]\n", len(calibreBooks))
	checks.Author()
	checks.Language()
	checks.Title()
}
