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
	if len(flag.Args()) != 0 {
		switch flag.Args()[0] {
		case "check":
			calibreBooks := loadFromCalibre()
			fsentries := loadFromFilesystem()
			check(calibreBooks, fsentries)
		case "backup":
			tools.PrintBlueln("In preparation.....")
		default:
			dbg("Executing calibredb command...")
			calibre(flag.Args())
		}
	} else {
		usage()
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
