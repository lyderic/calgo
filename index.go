package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/lyderic/tools"
	"gopkg.in/yaml.v2"
)

type Index struct {
	Timestamp time.Time `yaml:"timestamp"`
	FSBooks   []FSBook  `yaml:"ebooks"`
}

func (i *Index) load() {
	if _, err := os.Stat(c.Index); os.IsNotExist(err) {
		i.create()
	} else {
		tools.PrintGreenf("Index found: %s\n", c.Index)
	}
	content, err := ioutil.ReadFile(c.Index)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &i)
	if err == nil {
		dbg("index loaded from %q\n", c.Index)
	} else {
		log.Fatal(err)
	}
	fmt.Printf("%d books indexed\n", len(i.FSBooks))
}

func (i *Index) create() {
	tools.PrintYellowf("Index not found at %q. Creating...\n", c.Index)
	i.Timestamp = time.Now()
	i.FSBooks = loadFromFilesystem()
	err := i.save()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Index) save() (err error) {
	data, err := yaml.Marshal(&i)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(c.Index, data, 0644)
	tools.PrintGreenf("index written to %q\n", c.Index)
	return
}
