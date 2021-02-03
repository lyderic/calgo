package main

import (
	"fmt"
	"strconv"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func report(book Book, message string) {
	fmt.Printf("[%04d] %s (%s): %s\n",
		book.Id, book.Title, book.Author,
		message)
}
