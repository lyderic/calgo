package checks

import (
	. "calgo/internal"
)

func Title(calibreBooks []CalibreBook) {
	set := SearchSet{
		Name: "titles",
		Searches: []Search{
			{
				Name:    "title contains a dandling hyphen",
				Pattern: `title:" - "`,
				BookSet: calibreBooks,
			},
			{
				Name:    "title contains a double quote",
				Pattern: `title:"\""`,
				BookSet: calibreBooks,
			},
			{
				Name:    "title starts with lower case",
				Pattern: `title:"~^[a-z]" and not title:"meurtre.com"`,
				BookSet: calibreBooks,
			},
			{
				Name:    "title with at least 3 upper case letters",
				Pattern: `title:"~[A-Z][A-Z][A-Z]" and not title:"~(ABC|CQFD|K-PAX|SAS|SSN|URSS)" and not title:"~(III|VII|VIII|XII|XIV|XVI|XIX|XXI)"`,
				BookSet: calibreBooks,
			},
		},
	}
	set.Display()
}
