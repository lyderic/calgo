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

func exmain() {
	debug = flag.Bool("debug", false, "debug")
	flag.Usage = usage
	flag.Parse()
	c.load()
	if len(flag.Args()) != 0 {
		switch flag.Args()[0] {
		case "check":
			calibreBooks := loadFromCalibre()
			fsentries := loadFromFilesystem()
			check(calibreBooks, fsentries)
		case "index":
			var i Index
			i.load()
		case "backup":
			tools.PrintBlueln("In preparation.....")
		default:
			dbg("Executing calibredb command...")
			calibre(flag.Args())
		}
	} else {
		usage()
	}
	err := c.save()
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Println("calgo <option> [check|backup|index]")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
