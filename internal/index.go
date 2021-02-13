package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func (i *Index) Load() {
	if _, err := os.Stat(viper.GetString("index")); os.IsNotExist(err) {
		i.create()
	} else {
		tools.PrintGreenf("Index found: %s\n", viper.GetString("index"))
	}
	content, err := ioutil.ReadFile(viper.GetString("index"))
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(content, &i)
	if err == nil {
		Debug("index loaded from %q\n", viper.GetString("index"))
	} else {
		log.Fatal(err)
	}
	fmt.Printf("%d books indexed\n", len(i.FSBooks))
}

func (i *Index) Save() (err error) {
	data, err := yaml.Marshal(&i)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(viper.GetString("index"), data, 0644)
	tools.PrintGreenf("index written to %q\n", viper.GetString("index"))
	return
}

func (i *Index) create() {
	tools.PrintYellowf("Index not found at %q. Creating...\n", viper.GetString("index"))
	i.Timestamp = time.Now()
	i.FSBooks = LoadFromFilesystem()
	err := i.Save()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Index) reindex() (err error) {
	err = os.Remove(viper.GetString("index"))
	if err != nil {
		return
	}
	tools.PrintYellowf("Index %q removed\n", viper.GetString("index"))
	i.Load()
	return
}
