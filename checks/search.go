package checks

import (
	. "calgo/internal"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

type Search struct {
	Name    string
	Pattern string
}

type SearchSet struct {
	Name     string
	Searches []Search
}

func (s SearchSet) Display() {
	tools.PrintBluef("Checking %s...\n", s.Name)
	for _, search := range s.Searches {
		result := search.Process()
		result.Display()
	}
}

type SearchResult struct {
	Search     Search
	BooksFound []CalibreBook
	Ids        []string
}

func (s Search) Process() (result SearchResult) {
	Debug("Running search: %#v\n", s)
	result.Search = s
	output, err := CalibreOutputErr("search", s.Pattern)
	if err != nil {
		result.Ids = []string{}
	} else {
		result.Ids = strings.Split(string(output), ",")
	}
	return
}

func (r SearchResult) Display() {
	if len(r.Ids) == 0 {
		fmt.Printf("No books found for search %q%s\n",
			r.Search.Name, showPattern(r.Search.Pattern))
		return
	}
	for _, id := range r.Ids {
		fmt.Printf("Search %q%s found %d book%s:\n",
			r.Search.Name, showPattern(r.Search.Pattern),
			len(r.Ids), tools.Ternary(len(r.Ids) > 1, "s", ""))
		output := CalibreOutput("show_metadata", id, "--as-opf")
		var opf Opf
		xml.Unmarshal(output, &opf)
		Debug("%#v\n", opf)
		fmt.Printf("%s (%s)\n", opf.Metadata.Title, opf.Metadata.Creator)
	}
}

func showPattern(s string) (showing string) {
	if viper.GetBool("verbose") {
		return fmt.Sprintf(" [%s]", s)
	}
	return
}
