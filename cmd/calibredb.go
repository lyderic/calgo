package cmd

import (
	. "calgo/internal"

	"github.com/spf13/cobra"
)

var calibredbCmd = &cobra.Command{
	Use:                "calibredb",
	Short:              "calibredb pass-thru",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		Calibredb(args)
	},
}

func init() {
	rootCmd.AddCommand(calibredbCmd)
}
