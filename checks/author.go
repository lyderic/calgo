package checks

import (
	. "calgo/internal"
)

func Author() {
	set := SearchSet{
		Name: "authors",
		Searches: []Search{
			{
				Name:    "comma in author",
				Pattern: `authors:","`,
			},
		},
	}
	set.Display()
}

func checkAuthorMatchSort(calibreBook CalibreBook) (report Report) {
	report.Book = calibreBook
	if calibreBook.Author != calibreBook.Sort {
		report.Message = "[" + calibreBook.Sort + "] authors and author_sort mismatch!"
	}
	return
}
