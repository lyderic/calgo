package checks

import (
	"github.com/spf13/viper"
)

func Language() {
	set := SearchSet{
		Name: "languages",
		Searches: []Search{
			{
				Name:    "language not set",
				Pattern: "language:false",
			},
			{
				Name:    "language not " + viper.GetString("language"),
				Pattern: "not languages:" + viper.GetString("language"),
			},
		},
	}
	set.Display()
}
