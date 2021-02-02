package main

import (
	"fmt"
	"os"

	"github.com/lyderic/tools"
)

func reset() {
	err := os.Remove(cache)
	if err != nil {
		panic(err)
	}
	tools.PrintYellowln("cache reset")
}

func findlanguage(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if len(book.Languages) == 0 {
			fmt.Printf("[%04d] %s %s\n", book.Id,
				book.Title, book.Author)
			count++
		}
	}
	tools.PrintRedln(count, "book(s) don't have a language set up!")
}

func findfrench(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if len(book.Languages) == 0 {
			continue
		}
		if book.Languages[0] != "fra" {
			fmt.Printf("[%04d] %s %s '%s'\n", book.Id,
				book.Title, book.Author, book.Languages[0])
			count++
		}
	}
	tools.PrintRedln(count, "book(s) don't have French set up as language")
}

func findupper(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if IsUpper(book.Title) {
			fmt.Printf("[%04d] %s (%s - %s)\n", book.Id,
				book.Title, book.Author, book.Sort)
			count++
		}
	}
	tools.PrintRedln(count, "book(s) have their title all uppercase!")
}
