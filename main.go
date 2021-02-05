package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/lyderic/tools"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	err := loadConfiguration()
	if err != nil {
		panic(err)
	}
	doprobeauthor := flag.Bool("a", false, "probe author")
	doprobetitle := flag.Bool("t", false, "probe title")
	doprobelanguage := flag.Bool("l", false, "probe language")
	doprobemetadata := flag.Bool("m", false, "probe metadata")
	dodebug := flag.Bool("d", false, "debug")
	flag.Usage = usage
	flag.Parse()
	fix := false
	if len(flag.Args()) != 0 {
		switch flag.Args()[0] {
		case "fix":
			fix = true
		case "reset":
			reset()
		default:
			tools.PrintYellowln("no action specified, showing only")
		}
	}
	books := calibreToJson()
	fmt.Println("Loaded", len(books), "books")
	if *dodebug {
		dumpConfiguration()
	}
	if *doprobeauthor {
		author(books, fix)
	}
	if *doprobetitle {
		title(books, fix)
	}
	if *doprobelanguage {
		language(books, fix)
	}
	if *doprobemetadata {
		metadata(books)
	}
	if fix {
		reset()
	}
	err = saveConfiguration()
	if err != nil {
		panic(err)
	}
}

func calibreToJson() (books []Book) {
	json.Unmarshal(load(), &books)
	return
}

func usage() {
	fmt.Println("calgo <option> [fix|reset]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
