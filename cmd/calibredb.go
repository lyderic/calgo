package cmd

import (
	. "calgo/internal"
	"os"

	"github.com/spf13/cobra"
)

var calibredbCmd = &cobra.Command{
	Use:                "calibredb",
	Short:              "calibredb pass-thru",
	DisableFlagParsing: true,
	Run: func(cobraCmd *cobra.Command, args []string) {
		cmd := Calibredb(args...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		cmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(calibredbCmd)
}
