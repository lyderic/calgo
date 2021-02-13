package main

import (
	"calgo/cmd"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	cmd.Execute()
}
