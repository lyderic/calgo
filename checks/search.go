package checks

import (
	. "calgo/internal"
	"fmt"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
)

type Search struct {
	Name    string
	Pattern string
	BookSet []CalibreBook
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
}

func (s Search) Process() (result SearchResult) {
	Debug("Running search: %#v\n", s)
	result.Search = s
	output, err := CalibreOutputErr("search", s.Pattern)
	if err != nil {
		result.BooksFound = []CalibreBook{}
	} else {
		result.BooksFound = parseSearchOutput(s.BookSet, output)
	}
	return
}

func (r SearchResult) Display() {
	if len(r.BooksFound) == 0 {
		fmt.Printf("No books found for search %q [%s]\n",
			r.Search.Name, r.Search.Pattern)
		return
	}
	fmt.Printf("Search %q [%s] found %d book%s:\n",
		r.Search.Name, r.Search.Pattern,
		len(r.BooksFound), tools.Ternary(len(r.BooksFound) > 1, "s", ""))
	for _, calibreBook := range r.BooksFound {
		fmt.Println(calibreBook)
	}
}

func parseSearchOutput(calibreBooks []CalibreBook,
	output []byte) (books []CalibreBook) {
	ids := strings.Split(string(output), ",")
	for _, book := range calibreBooks {
		for _, id := range ids {
			idint, err := strconv.Atoi(id)
			if err != nil {
				tools.PrintRedf("[%s,%d]", id, idint)
			}
			if book.Id == idint {
				books = append(books, book)
				break
			}
		}
	}
	return
}
