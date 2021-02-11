package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func report(calibreBook CalibreBook, message string) {
	fmt.Printf("[%04d] %s (%s): %s\n",
		calibreBook.Id, calibreBook.Title, calibreBook.Author,
		message)
}

func extractIdFromPath(path string) (id int) {
	bits := strings.Fields(path)
	last := bits[len(bits)-1]
	number := last[1 : len(last)-1]
	id, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func dbg(format string, args ...interface{}) {
	if *debug {
		tools.PrintYellowf(format, args...)
	}
}
