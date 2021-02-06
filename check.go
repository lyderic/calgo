package main

import (
	"archive/zip"
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
	fmt.Print("Checking IDs on file system exist in the calibre DB... ")
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
	fmt.Println("done.")
	return true
}

func checkContentOpfInEpub(entries []FSEntry) (result bool) {
	fmt.Print("Checking epub on filesystem have a content.opf file... ")
	result = true
	for _, entry := range entries {
		reader, err := zip.OpenReader(entry.Fullpath)
		if err != nil {
			tools.PrintRedln("Problem with this file:", entry.Fullpath)
			panic(err)
		}
		found := false
		for _, file := range reader.File {
			if file.FileInfo().Name() == "content.opf" {
				found = true
				break
			}
		}
		if !found {
			tools.PrintRedln("File with no content.opf:", entry.Fullpath)
			result = false
		}
	}
	fmt.Println("done.")
	return
}
