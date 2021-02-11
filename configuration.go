package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/lyderic/tools"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Url        string `yaml:"url"`
	Language   string `yaml:"language"`
	CalibreDir string `yaml:"calibredir"`
	Index      string `yaml:"index"`
}

var c Configuration

func (c *Configuration) load() {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		tools.PrintYellowln(configFile, ": not found. Using defaults:")
		c.Url = "http://localhost:8080"
		c.Language = "fra"
		c.CalibreDir = os.Getenv("HOME") + "/Calibre Library"
		tools.PrintYellowf("%v\n", c)
		return
	}
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &c)
	if err == nil {
		dbg("configuration loaded from %s\n", configFile)
		dbg("%#v\n", c)
	} else {
		log.Fatal(err)
	}
}

func (c *Configuration) save() (err error) {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(configFile, data, 0644)
	if err == nil {
		dbg("configuration saved to %s\n", configFile)
	}

	return
}
