package internal

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func Debug(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		Cyan(format, args...)
	}
}

func Embed() {
	start := time.Now()
	cmd := Calibredb("embed_metadata", "all")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	cmd.Run()
	Green("Embedded metadata in %v\n", time.Since(start))
}

func Rsync(src, dst string) {
	start := time.Now()
	cmd := exec.Command("rsync", "-Pav", "--delete", src+"/", dst+"/")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	Debug("[XeQ]-%v\n", cmd.Args)
	cmd.Start()
	cmd.Wait()
	Green("Synced %q to %q in %v\n", src, dst, time.Since(start))
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
