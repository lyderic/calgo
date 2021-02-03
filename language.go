package main

import (
	"github.com/lyderic/tools"
)

func language(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if len(book.Languages) == 0 {
			report(book, "no language set!")
			count++
			continue
		}
		if book.Languages[0] != "fra" {
			report(book, "language not French!")
			count++
			continue
		}
	}
	if count > 0 {
		tools.PrintRedln(count, "book(s) have incorrect language")
	} else {
		tools.PrintGreenln("All books have correct language. All good!")
	}
}
