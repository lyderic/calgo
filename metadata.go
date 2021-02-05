package main

import "github.com/lyderic/tools"

func metadata(books []Book) {
	count := 0
	for _, book := range books {
		if len(book.Formats) < 2 {
			report(book, "book doesn't at least two formats!")
			count++
		}
	}
	if count > 0 {
		tools.PrintRedln(count, "book(s) don't have at least two formats!")
	} else {
		tools.PrintGreenln("All books have at least two formats. All good!")
	}
}
