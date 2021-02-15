package cmd

import (
	"calgo/checks"
	"sync"

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
	var wg sync.WaitGroup
	for _, search := range searches {
		wg.Add(1)
		go routine(search, &wg)
	}
	wg.Wait()
}

func routine(search checks.Search, wg *sync.WaitGroup) {
	defer wg.Done()
	result := search.Process()
	result.Display()
}
