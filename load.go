package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func loadFromCalibre() (calibreBooks []CalibreBook) {
	var raw []byte
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		fmt.Print("loading calibre data...")
		raw = calibreOutput("list", "-f", "all", "--for-machine")
		fmt.Print("\r                              \r")
	}
	json.Unmarshal(raw, &calibreBooks)
	dbg("Loaded from calibre: %d books\n", len(calibreBooks))
	return
}

func loadFromFilesystem() (fsBooks []FSBook) {
	fmt.Print("loading filesystem data...")
	err := filepath.Walk(c.CalibreDir,
		func(path string, finfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".epub" {
				var fsentry FSBook
				fsentry.DirPath = filepath.Dir(path)
				fsentry.Epub = finfo.Name()
				fsentry.OriginalEpub = strings.Replace(finfo.Name(),
					".epub", ".original_epub", 1)
				fsentry.Id = extractIdFromPath(fsentry.DirPath)
				fsBooks = append(fsBooks, fsentry)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\r                              \r")
	dbg("Loaded from filesystem: %d books\n", len(fsBooks))
	return
}
