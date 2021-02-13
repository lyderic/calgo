package internal

import (
	"log"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

func Debug(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		tools.PrintYellowf(format, args...)
	}
}
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func extractIdFromPath(path string) (id int) {
	bits := strings.Fields(path)
	last := bits[len(bits)-1]
	number := last[1 : len(last)-1]
	id, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return
}
