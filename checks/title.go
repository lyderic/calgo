package checks

func Title() {
	set := SearchSet{
		Name: "titles",
		Searches: []Search{
			{
				Name:    "title contains a dandling hyphen",
				Pattern: `title:" - "`,
			},
			{
				Name:    "title contains a double quote",
				Pattern: `title:"\""`,
			},
			{
				Name:    "title starts with lower case",
				Pattern: `title:"~^[a-z]" and not title:"meurtre.com"`,
			},
			{
				Name:    "title with at least 3 upper case letters",
				Pattern: `title:"~[A-Z][A-Z][A-Z]" and not title:"~(ABC|CQFD|K-PAX|SAS|SSN|URSS)" and not title:"~(III|VII|VIII|XII|XIV|XVI|XIX|XXI)"`,
			},
		},
	}
	set.Display()
}
