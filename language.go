package main

import (
	"fmt"

	"github.com/lyderic/tools"
)

func language(books []Book) (result bool) {
	fmt.Println("Checking language... ")
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
		result = false
		tools.PrintRedln("> Failed!")
	} else {
		result = true
		tools.PrintGreenln("> Ok")
	}
	return
}
