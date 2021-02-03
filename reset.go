package main

import (
	"log"
	"os"

	"github.com/lyderic/tools"
)

func reset() {
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		tools.PrintGreenln("cache not found: nothing to reset")
		return
	}
	err := os.Remove(cache)
	if err != nil {
		log.Fatal(err)
	}
	tools.PrintYellowln("cache reset")
}
