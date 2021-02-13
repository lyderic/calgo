package checks

import (
	"fmt"

	. "calgo/internal"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

func Language(calibreBooks []CalibreBook) {
	fmt.Println("Checking language... ")
	var reports []Report
	for _, calibreBook := range calibreBooks {
		reports = append(reports, checkLanguageIsSet(calibreBook))
		reports = append(reports, checkLanguageIsCorrect(calibreBook))
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

func checkLanguageIsSet(calibreBook CalibreBook) (report Report) {
	report.Book = calibreBook
	if len(calibreBook.Languages) == 0 {
		report.Message = "no language set!"
	}
	return
}

func checkLanguageIsCorrect(calibreBook CalibreBook) (report Report) {
	report.Book = calibreBook
	if calibreBook.Languages[0] != viper.GetString("language") {
		report.Message = fmt.Sprintf("language not %q!", viper.GetString("language"))
	}
	return
}
