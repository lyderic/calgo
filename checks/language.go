package checks

import (
	. "calgo/internal"

	"github.com/spf13/viper"
)

func Language(calibreBooks []CalibreBook) {
	set := SearchSet{
		Name: "languages",
		Searches: []Search{
			{
				Name:    "language not set",
				Pattern: "language:false",
				BookSet: calibreBooks,
			},
			{
				Name:    "language not " + viper.GetString("language"),
				Pattern: "not languages:" + viper.GetString("language"),
				BookSet: calibreBooks,
			},
		},
	}
	set.Display()
}
