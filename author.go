package main

import (
	"strings"

	"github.com/lyderic/tools"
)

func author(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if book.Author != book.Sort {
			report(book, "["+book.Sort+"] authors and author_sort mismatch!")
			count++
			continue
		}
		if strings.Contains(book.Author, ",") {
			report(book, "comma in authors field!")
			count++
			continue
		}
	}
	if count > 0 {
		tools.PrintRedln(count, "book(s) have incorrect author!")
	} else {
		tools.PrintGreenln("All books have correct author. All good!")
	}
}
