package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

func main() {
	doreset := flag.Bool("reset", false, "reset cache")
	dofindauthor := flag.Bool("a", false, "find author")
	dofindlanguage := flag.Bool("l", false, "find language")
	dofindfrench := flag.Bool("f", false, "find French")
	dofindupper := flag.Bool("u", false, "find uppercase titles")
	dofindcomma := flag.Bool("c", false, "find authors with commas in their name")
	dofindsuspect := flag.Bool("s", false, "find books with suspect title")
	flag.Usage = usage
	flag.Parse()
	fix := false
	if len(flag.Args()) != 0 {
		if flag.Args()[0] == "fix" {
			fix = true
		}
	}
	if *doreset {
		reset()
	}
	books := calibreToJson()
	fmt.Println("Loaded", len(books), "books")
	if *dofindauthor {
		findauthor(books, fix)
	}
	if *dofindlanguage {
		findlanguage(books, fix)
	}
	if *dofindfrench {
		findfrench(books, fix)
	}
	if *dofindupper {
		findupper(books, fix)
	}
	if *dofindcomma {
		findcomma(books, fix)
	}
	if *dofindsuspect {
		findsuspect(books, fix)
	}
	if fix {
		reset()
	}
}

func calibreToJson() (books []Book) {
	json.Unmarshal(load(), &books)
	return
}

func usage() {
	flag.PrintDefaults()
}
