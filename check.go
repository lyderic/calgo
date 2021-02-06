package main

import (
	"fmt"

	"github.com/lyderic/tools"
)

func check(calibreBooks []Book, fsentries []FSEntry) {
	fmt.Println("Number of ebooks recorded in calibre:", len(calibreBooks))
	fmt.Println("Number of ebooks found on filesystem:", len(fsentries))
	if checkIds(calibreBooks, fsentries) &&
		checkContentOpfInEpub(fsentries) {
		tools.PrintGreenln("Data OK")
	} else {
		tools.PrintRedln("Inconsistent data!")
	}
}

func checkIds(books []Book, entries []FSEntry) bool {
	for _, entry := range entries {
		found := false
		for _, book := range books {
			if entry.Id == book.Id {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("ID#%04d on filesystem not found in calibre DB!\n", entry.Id)
			return false
		}
	}
	return true
}

func checkContentOpfInEpub(entries []FSEntry) bool {
	return false
}
