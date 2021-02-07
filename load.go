package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lyderic/tools"
)

func loadFromCalibre() (books []Book) {
	var raw []byte
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		fmt.Print("loading calibre data...")
		cmd := exec.Command("calibredb", "list", "-f", "all", "--for-machine")
		cmd.Stderr = os.Stderr
		raw, err = cmd.Output()
		if err != nil {
			tools.PrintRedln("An instance of calibre is most probably already running!")
			log.Fatal(err)
		}
		fmt.Print("\r                              \r")
	}
	json.Unmarshal(raw, &books)
	dbg(fmt.Sprintf("Loaded from calibre: %d books", len(books)))
	return
}

func loadFromFilesystem() (fsentries []FSEntry) {
	fmt.Print("loading filesystem data...")
	err := filepath.Walk(basedir,
		func(path string, finfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".epub" {
				var entry FSEntry
				entry.Fullpath = path
				entry.Filename = finfo.Name()
				entry.Parentdir = filepath.Dir(path)
				entry.Id = extractIdFromPath(entry.Parentdir)
				fsentries = append(fsentries, entry)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\r                              \r")
	dbg(fmt.Sprintf("Loaded from filesystem: %d entries", len(fsentries)))
	return
}
