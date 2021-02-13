package checks

import (
	. "calgo/internal"

	"github.com/lyderic/tools"
)

func metadata(calibreBooks []CalibreBook) {
	count := 0
	for _, calibreBook := range calibreBooks {
		if len(calibreBook.Formats) < 2 {
			Report(calibreBook, "book doesn't at least two formats!")
			count++
		}
	}
	if count > 0 {
		tools.PrintRedln(count, "book(s) don't have at least two formats!")
	} else {
		tools.PrintGreenln("All books have at least two formats. All good!")
	}
}
