package checks

import (
	. "calgo/internal"
	"encoding/json"
	"fmt"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

type Search struct {
	Name    string `yaml:"name"`
	Pattern string `yaml:"pattern"`
}

type SearchResult struct {
	Search Search
	Books  []CalibreBook
}

func (s Search) Process() (result SearchResult) {
	Blue("Running search: %q...\n", s.Name)
	if viper.GetBool("verbose") {
		Blue("[%s]\n", s.Pattern)
	}
	result.Search = s
	output := CalibreOutput("list",
		"-s", s.Pattern, "-f", "all", "--for-machine")
	json.Unmarshal(output, &result.Books)
	return
}

func (r SearchResult) Display() {
	if len(r.Books) == 0 {
		fmt.Printf("No books found for search %q%s\n",
			r.Search.Name, showPattern(r.Search.Pattern))
		return
	}
	fmt.Printf("Found %d book%s:\n", len(r.Books),
		tools.Ternary(len(r.Books) > 1, "s", ""))
	for _, book := range r.Books {
		fmt.Println(book)
	}
}

func showPattern(s string) (showing string) {
	if viper.GetBool("verbose") {
		return fmt.Sprintf(" [%s]", s)
	}
	return
}
