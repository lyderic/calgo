package checks

import (
	"fmt"
	"strings"

	. "calgo/internal"

	"github.com/lyderic/tools"
)

func Author(calibreBooks []CalibreBook) {
	fmt.Println("Checking authors... ")
	var reports []Report
	for _, calibreBook := range calibreBooks {
		reports = append(reports, checkAuthorMatchSort(calibreBook))
		reports = append(reports, checkCommaInAuthorsField(calibreBook))
	}
	count := 0
	for _, report := range reports {
		if report.Message != "" {
			fmt.Println(report)
			count++
		}
	}
	if count == 0 {
		tools.PrintGreenln("> Ok")
	}
}

func checkAuthorMatchSort(calibreBook CalibreBook) (report Report) {
	report.Book = calibreBook
	if calibreBook.Author != calibreBook.Sort {
		report.Message = "[" + calibreBook.Sort + "] authors and author_sort mismatch!"
	}
	return
}

func checkCommaInAuthorsField(calibreBook CalibreBook) (report Report) {
	report.Book = calibreBook
	if strings.Contains(calibreBook.Author, ",") {
		report.Message = "comma in authors field!"
	}
	return
}
