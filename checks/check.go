package checks

import (
	"archive/zip"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	. "calgo/internal"

	"github.com/lyderic/tools"
)

func performCalibreBuiltinCheck() (result bool) {
	// close running calibre (calibre -s)
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
	// restart calibre (with option to run in bg)
	return
}

func checkIds(calibreBooks []CalibreBook, fsBooks []FSBook) bool {
	// not necessary, will be done by performCalibreBuiltinCheck
	fmt.Println("Checking IDs on file system exist in the calibre DB...")
	for _, fsBook := range fsBooks {
		found := false
		for _, calibreBook := range calibreBooks {
			if fsBook.Id == calibreBook.Id {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("ID#%04d on filesystem not found in calibre DB!\n", fsBook.Id)
			return false
		}
	}
	fmt.Println("done.")
	return true
}

func checkContentOpfInEpub(fsBooks []FSBook) (result bool) {
	fmt.Println("Checking epub on filesystem have a content.opf file...")
	count := 0
	for _, fsBook := range fsBooks {
		fullpath := filepath.Join(fsBook.DirPath, fsBook.Epub)
		reader, err := zip.OpenReader(fullpath)
		if err != nil {
			tools.PrintRedln("Problem with this file:", fullpath)
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
			tools.PrintRedln("File with no content.opf:", fullpath)
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
