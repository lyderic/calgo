package main

import (
	"log"
	"regexp"
	"strings"
	"unicode"

	"github.com/lyderic/tools"
)

func title(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if isAccepted(book.Title) { // list of valid title in data.go
			continue
		}
		if isFirstLetterLowerCase(book.Title) {
			report(book, "first letter is not capitalised!")
			count++
			continue
		}
		if containsUpperCaseWord(book.Title) {
			report(book, "contains at least one word that is all upper case!")
			count++
			continue
		}
		if strings.Contains(book.Title, " - ") {
			report(book, "contains a dandling hyphen!")
			if fix {
				fixDandlingHyphenInTitle(book)
			}
			count++
			continue
		}
		if strings.Contains(book.Title, "\"") {
			report(book, "contains a double quote")
			if fix {
				fixDoubleQuoteInTitle(book)
			}
			count++
			continue
		}
	}
	if count > 0 {
		tools.PrintRedln(count, "book(s) have incorrect title!")
	} else {
		tools.PrintGreenln("All books have correct title. All good!")
	}
}

func isFirstLetterLowerCase(s string) (match bool) {
	match, err := regexp.MatchString("^[a-z]", s)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func containsUpperCaseWord(s string) bool {
	bits := strings.Fields(s)
	for _, bit := range bits {
		for _, r := range bit {
			if !unicode.IsUpper(r) && unicode.IsLetter(r) {
				return false
			}
		}
	}
	return true
}

func isAccepted(s string) bool {
	for _, accepted := range valids {
		if s == accepted {
			return true
		}
	}
	return false
}
