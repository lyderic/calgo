package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"

	"github.com/lyderic/tools"
)

func title(calibreBooks []CalibreBook) (result bool) {
	// use saved searches by parsing 'calibredb saved_searches list' and applying!
	fmt.Println("Checking titles...")
	count := 0
	for _, calibreBook := range calibreBooks {
		words := strings.Fields(calibreBook.Title)
		if isFirstLetterLowerCase(calibreBook.Title) {
			report(calibreBook, "first letter is not capitalised!")
			count++
		}
		if containsUpperCaseWord(words) {
			report(calibreBook, "contains at least one word that is all upper case!")
			count++
		}
		if strings.Contains(calibreBook.Title, " - ") {
			report(calibreBook, "contains a dandling hyphen!")
			count++
		}
		if strings.Contains(calibreBook.Title, "\"") {
			report(calibreBook, "contains a double quote")
			count++
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
