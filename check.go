package main

import (
	"archive/zip"
	"fmt"
	"os"
	"os/exec"

	"github.com/lyderic/tools"
)

func check(calibreBooks []Book, fsentries []FSEntry) {
	fmt.Println("Number of ebooks recorded in calibre:", len(calibreBooks))
	fmt.Println("Number of ebooks found on filesystem:", len(fsentries))
	if performCalibreBuiltinCheck() &&
		checkContentOpfInEpub(fsentries) &&
		title(calibreBooks) && author(calibreBooks) && language(calibreBooks) {
		tools.PrintGreenln("Data OK")
	} else {
		tools.PrintRedln("Inconsistent data!")
	}
}

func performCalibreBuiltinCheck() (result bool) {
	fmt.Println("Performing calibre built-in library check... ")
	cmd := exec.Command("calibredb", "check_library")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()
	if err == nil {
		result = true
		tools.PrintGreenln("> Ok")
	} else {
		result = false
		tools.PrintRedln("> Failed!")
	}
	return
}

func checkIds(books []Book, entries []FSEntry) bool {
	fmt.Println("Checking IDs on file system exist in the calibre DB...")
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
	fmt.Println("Checking epub on filesystem have a content.opf file...")
	count := 0
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
			count++
			result = false
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
