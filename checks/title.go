package checks

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	. "calgo/internal"

	"github.com/lyderic/tools"
)

func Title(calibreBooks []CalibreBook) {
	fmt.Println("Checking titles...")
	var reports []Report
	reports = append(reports, checkSavedSearches(calibreBooks))
	/*
		for _, calibreBook := range calibreBooks {
				words := strings.Fields(calibreBook.Title)
				if isFirstLetterLowerCase(calibreBook.Title) {
					printReport(calibreBook, "first letter is not capitalised!")
					count++
				}
				if containsUpperCaseWord(words) {
					printReport(calibreBook, "contains at least one word that is all upper case!")
					count++
				}
				if strings.Contains(calibreBook.Title, " - ") {
					printReport(calibreBook, "contains a dandling hyphen!")
					count++
				}
				if strings.Contains(calibreBook.Title, "\"") {
					printReport(calibreBook, "contains a double quote")
					count++
				}
		}
	*/
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

func checkSavedSearches(calibreBooks []CalibreBook) (report Report) {
	fmt.Println("Running saved searches...")
	output := CalibreOutput("saved_searches", "list")
	var savedSearches []SavedSearch
	if len(output) > 0 {
		scanner := bufio.NewScanner(bytes.NewReader(output))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "Name: ") {
				var savedSearch SavedSearch
				savedSearch.Name = strings.Replace(line, "Name: ", "", 1)
				scanner.Scan()
				nextLine := scanner.Text()
				savedSearch.SearchString = strings.Replace(nextLine,
					"Search string: ", "", 1)
				savedSearches = append(savedSearches, savedSearch)
			}
		}
	}
	var results []SavedSearchResult
	for _, savedSearch := range savedSearches {
		var result SavedSearchResult
		result.Search = savedSearch
		output, err := CalibreOutputErr("search", savedSearch.SearchString)
		if err != nil {
			result.Books = []CalibreBook{}
		} else {
			result.Books = getCalibreBooksFromIdList(calibreBooks, strings.Split(string(output), ","))
		}
		results = append(results, result)
	}
	for _, result := range results {
		if len(result.Books) == 0 {
			fmt.Printf("No books found for search %q\n", result.Search.Name)
		} else {
			fmt.Printf("%d book%s found for search %q\n[%s]\n",
				len(result.Books),
				tools.Ternary(len(result.Books) > 1, "s", ""),
				result.Search.Name, result.Search.SearchString)
			for _, calibreBook := range result.Books {
				fmt.Println(calibreBook.Id, calibreBook.Title, calibreBook.Author)
			}
		}
	}
	return
}

func getCalibreBooksFromIdList(calibreBooks []CalibreBook, ids []string) (books []CalibreBook) {
	for _, book := range calibreBooks {
		for _, id := range ids {
			idint, err := strconv.Atoi(id)
			if err != nil {
				tools.PrintRedf("[%s,%d]", id, idint)
			}
			if book.Id == idint {
				books = append(books, book)
				break
			}
		}
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
