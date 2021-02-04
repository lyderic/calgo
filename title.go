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
		words := strings.Fields(book.Title)
		if isFirstLetterLowerCase(book.Title) {
			report(book, "first letter is not capitalised!")
			count++
		}
		if containsUpperCaseWord(words) {
			report(book, "contains at least one word that is all upper case!")
			count++
		}
		if strings.Contains(book.Title, " - ") {
			report(book, "contains a dandling hyphen!")
			if fix {
				fixDandlingHyphenInTitle(book)
			}
			count++
		}
		if strings.Contains(book.Title, "\"") {
			report(book, "contains a double quote")
			if fix {
				fixDoubleQuoteInTitle(book)
			}
			count++
		}
		if secondWordIsCapitalized(words) {
			report(book, "second word is capitalized!")
			if fix {
				proposeFixTitle(book)
			}
			count++
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

func containsUpperCaseWord(words []string) bool {
	for _, word := range words {
		for _, r := range word {
			if !unicode.IsUpper(r) && unicode.IsLetter(r) {
				return false
			}
		}
	}
	return true
}

func secondWordIsCapitalized(words []string) bool {
	if len(words) < 2 {
		return false
	}
	if strings.HasPrefix(words[0], "[") { // for series
		return false
	}
	secondWord := words[1]
	runes := []rune(secondWord)
	return unicode.IsUpper(runes[0])
}

func isAccepted(s string) bool {
	for _, accepted := range configuration.Accepted {
		if s == accepted {
			return true
		}
	}
	return false
}
