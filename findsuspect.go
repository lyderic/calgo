package main

import (
	"fmt"
	"regexp"

	"github.com/lyderic/tools"
)

func findsuspect(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		matchFirstLetterNotCap, err := regexp.MatchString("^[a-z]", book.Title)
		if err != nil {
			panic(err)
		}
		if matchFirstLetterNotCap {
			fmt.Printf("[%04d] %s (%s): First letter is not capitalised!\n", book.Id, book.Title, book.Author)
			count++
		}
	}
	tools.PrintRedln(count, "book(s) affected!")
}
