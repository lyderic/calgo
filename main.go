package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lyderic/tools"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	debug = flag.Bool("d", false, "debug")
	flag.Usage = usage
	flag.Parse()
	loadConfiguration()
	calibreBooks := loadFromCalibre()
	fsentries := loadFromFilesystem()
	if len(flag.Args()) != 0 {
		switch flag.Args()[0] {
		case "check":
			check(calibreBooks, fsentries)
		default:
			tools.PrintRedln("no action specified")
		}
	}
	err := saveConfiguration()
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Println("calgo <option> [check]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
