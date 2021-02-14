package cmd

import (
	"calgo/checks"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchesCmd = &cobra.Command{
	Use:   "searches",
	Short: "run searches defined in configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		searches()
	},
}

func init() {
	rootCmd.AddCommand(searchesCmd)
}

func searches() {
	var searches []checks.Search
	err := viper.UnmarshalKey("searches", &searches)
	if err != nil {
		panic(err)
	}
	for _, search := range searches {
		result := search.Process()
		result.Display()
	}
}
