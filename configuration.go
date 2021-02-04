package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/lyderic/tools"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Accepted []string `yaml:"accepted"`
}

var configuration Configuration

func loadConfiguration() (err error) {
	if _, err = os.Stat(conf); os.IsNotExist(err) {
		tools.PrintYellowln(conf, ": not found")
		return nil
	}
	content, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &configuration)
	if err == nil {
		tools.PrintGreenln("configuration loaded from", conf)
	}
	return
}

func saveConfiguration() (err error) {
	data, err := yaml.Marshal(&configuration)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(conf, data, 0644)
	if err == nil {
		tools.PrintGreenln("configuration saved to", conf)
	}
	return
}

func dumpConfiguration() {
	tools.PrintYellowf("Configuration: %#v\n", configuration)
}
