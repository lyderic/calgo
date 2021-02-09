package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"

	"github.com/lyderic/tools"
)

func title(books []Book) (result bool) {
	fmt.Println("Checking titles...")
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
			count++
		}
		if strings.Contains(book.Title, "\"") {
			report(book, "contains a double quote")
			count++
		}
		/*
			if secondWordIsCapitalized(words) {
				report(book, "second word is capitalized!")
				count++
			}
		*/
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

func isAccepted(s string) bool {
	for _, accepted := range c.Accepted {
		if s == accepted {
			return true
		}
	}
	return false
}
