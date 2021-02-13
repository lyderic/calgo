package checks

import (
	. "calgo/internal"
	"fmt"
)

type SavedSearch struct {
	Name         string
	SearchString string
}

type SavedSearchResult struct {
	Search SavedSearch
	Books  []CalibreBook
}

func (s SavedSearch) String() string {
	return fmt.Sprintf("> Name  : %s\n> Search: %s", s.Name, s.SearchString)
}
