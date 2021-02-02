package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
)

func findcomma(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if strings.Contains(book.Author, ",") {
			if fix {
				bits := strings.Split(book.Author, ",")
				correct := bits[1] + " " + bits[0]
				cmd := exec.Command("calibredb", "set_metadata",
					"-f", "authors:"+correct,
					"-f", "author_sort:"+correct,
					strconv.Itoa(book.Id))
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
			}
			fmt.Printf("[%04d] %s (%s - %s)\n", book.Id,
				book.Title, book.Author, book.Sort)
			count++
		}
	}
	tools.PrintRedln(count, "book(s) have a comma in the author field!")
}
