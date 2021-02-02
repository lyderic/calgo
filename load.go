package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/lyderic/tools"
)

func load() (formachine []byte) {
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		fmt.Print("loading...")
		cmd := exec.Command("calibredb", "list", "-f", "all", "--for-machine")
		cmd.Stderr = os.Stderr
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Print("\r          \r")
		ioutil.WriteFile(cache, output, 0644)
	} else {
		tools.PrintGreenln("cache found:", cache)
	}
	formachine, err := ioutil.ReadFile(cache)
	if err != nil {
		panic(err)
	}
	return
}
