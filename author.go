package main

import (
	"fmt"
	"strings"

	"github.com/lyderic/tools"
)

func author(calibreBooks []CalibreBook) (result bool) {
	fmt.Println("Checking authors... ")
	count := 0
	for _, calibreBook := range calibreBooks {
		if calibreBook.Author != calibreBook.Sort {
			report(calibreBook, "["+calibreBook.Sort+"] authors and author_sort mismatch!")
			count++
		}
		if strings.Contains(calibreBook.Author, ",") {
			report(calibreBook, "comma in authors field!")
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
