package checks

import (
	"fmt"

	. "calgo/internal"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

func language(calibreBooks []CalibreBook) (result bool) {
	fmt.Println("Checking language... ")
	count := 0
	for _, calibreBook := range calibreBooks {
		if len(calibreBook.Languages) == 0 {
			Report(calibreBook, "no language set!")
			count++
			continue
		}
		if calibreBook.Languages[0] != viper.GetString("language") {
			Report(calibreBook, fmt.Sprintf("language not %q!", viper.GetString("language")))
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
