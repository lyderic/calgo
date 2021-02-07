package main

import (
	"fmt"
	"strings"

	"github.com/lyderic/tools"
)

func author(books []Book) (result bool) {
	fmt.Println("Checking authors... ")
	count := 0
	for _, book := range books {
		if book.Author != book.Sort {
			report(book, "["+book.Sort+"] authors and author_sort mismatch!")
			count++
		}
		if strings.Contains(book.Author, ",") {
			report(book, "comma in authors field!")
			count++
		}
	}
	if count > 0 {
		result = false
		tools.PrintRedln("> Failed!")
	} else {
		result = true
		tools.PrintGreenln("> Ok")
	}
	return
}
