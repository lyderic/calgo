package main

import (
	"fmt"

	"github.com/lyderic/tools"
)

func language(calibreBooks []CalibreBook) (result bool) {
	fmt.Println("Checking language... ")
	count := 0
	for _, calibreBook := range calibreBooks {
		if len(calibreBook.Languages) == 0 {
			report(calibreBook, "no language set!")
			count++
			continue
		}
		if calibreBook.Languages[0] != c.Language {
			report(calibreBook, fmt.Sprintf("language not %q!", c.Language))
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
