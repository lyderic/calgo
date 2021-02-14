package cmd

import (
	"calgo/checks"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	var searches []checks.Search
	err := viper.UnmarshalKey("checks", &searches)
	if err != nil {
		panic(err)
	}
	for _, search := range searches {
		result := search.Process()
		result.Display()
	}
}
