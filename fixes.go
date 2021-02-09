package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
)

func fixauthormismatch(book Book) {
	cmd := exec.Command("calibredb", "set_metadata", "-f",
		"author_sort:"+book.Author, strconv.Itoa(book.Id))
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func fixcomma(book Book) {
	bits := strings.Split(book.Author, ",")
	correct := bits[1] + " " + bits[0]
	cmd := exec.Command("calibredb", "set_metadata",
		"-f", "authors:"+correct,
		"-f", "author_sort:"+correct,
		strconv.Itoa(book.Id))
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func fixalluppertitle(title string, id int) {
	lowered := strings.ToLower(title)
	first := string(title)
	capitalized := first + lowered
	cmd := exec.Command("calibredb", "set_metadata",
		"-f", "title:"+capitalized, strconv.Itoa(id))
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	tools.PrintYellowln("-->", capitalized)
}

func fixDandlingHyphenInTitle(book Book) {
	tools.PrintYellowf("Fixing %q... ", book.Title)
	cleanTitle := strings.ReplaceAll(book.Title, " - ", " ")
	cmd := exec.Command("calibredb", "set_metadata",
		"-f", "title:"+cleanTitle, strconv.Itoa(book.Id))
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	tools.PrintYellowln("done.")
}

func fixDoubleQuoteInTitle(book Book) {
	tools.PrintYellowf("Removing double quote from %q... ", book.Title)
	cleanTitle := strings.ReplaceAll(book.Title, "\"", "")
	cmd := exec.Command("calibredb", "set_metadata",
		"-f", "title:"+cleanTitle, strconv.Itoa(book.Id))
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	tools.PrintYellowln("done.")
}

func proposeFixTitle(book Book) {
	cleanTitle, err := tools.VimString(book.Title)
	if err != nil {
		log.Fatal(err)
	}
	if cleanTitle == book.Title {
		c.Accepted = append(c.Accepted, book.Title)
		return
	}
	cmd := exec.Command("calibredb", "set_metadata",
		"-f", "title:"+cleanTitle, strconv.Itoa(book.Id))
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	tools.PrintYellowln(book.Title, "->", cleanTitle)
}
