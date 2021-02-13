package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func LoadFromCalibre() (calibreBooks []CalibreBook) {
	var raw []byte
	fmt.Print("loading calibre data...")
	raw = calibreOutput("list", "-f", "all", "--for-machine")
	fmt.Print("\r                              \r")
	json.Unmarshal(raw, &calibreBooks)
	Debug("Loaded from calibre: %d books\n", len(calibreBooks))
	return
}

func LoadFromFilesystem() (fsBooks []FSBook) {
	fmt.Print("loading filesystem data...")
	err := filepath.Walk(viper.GetString("calibredir"),
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
	Debug("Loaded from filesystem: %d books\n", len(fsBooks))
	return
}
