package checks

import (
	. "calgo/internal"
	"fmt"
)

type Report struct {
	Book    CalibreBook
	Message string
}

func (r Report) String() string {
	return fmt.Sprintf("[%04d] %s (%s): %s\n",
		r.Book.Id, r.Book.Title, r.Book.Author,
		r.Message)
}

func printReport(calibreBook CalibreBook, message string) {
	fmt.Printf("[%04d] %s (%s): %s\n",
		calibreBook.Id, calibreBook.Title, calibreBook.Author,
		message)
}
