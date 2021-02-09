package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/lyderic/tools"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Url      string   `yaml:"url"`
	Language string   `yaml:"language"`
	Accepted []string `yaml:"accepted"`
}

var c Configuration

func loadConfiguration() {
	if _, err := os.Stat(conf); os.IsNotExist(err) {
		tools.PrintYellowln(conf, ": not found. Using defaults:")
		c.Url = "http://localhost:8080"
		c.Language = "eng"
		tools.PrintYellowf("%v\n", c)
		return
	}
	content, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &c)
	if err == nil {
		dbg("configuration loaded from " + conf)
	} else {
		log.Fatal(err)
	}
	return
}

func saveConfiguration() (err error) {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(conf, data, 0644)
	if err == nil {
		dbg("configuration saved to " + conf)
	}

	return
}

func dumpConfiguration() {
	tools.PrintYellowf("Configuration: %#v\n", c)
}
