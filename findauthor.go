package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/lyderic/tools"
)

func findauthor(books []Book, fix bool) {
	count := 0
	for _, book := range books {
		if book.Author != book.Sort {
			if fix {
				cmd := exec.Command("calibredb", "set_metadata", "-f",
					"author_sort:"+book.Author, strconv.Itoa(book.Id))
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
	tools.PrintRedln(count, "book(s) have author and author_sort mismatched!")
}
