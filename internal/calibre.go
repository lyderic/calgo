package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

func calibre(args []string) {
	var cli []string
	cli = append(cli, "--with-library="+viper.GetString("url"))
	cli = append(cli, args...)
	cmd := exec.Command("calibredb", cli...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	Debug(fmt.Sprintf("[XeQ]:%v", cmd.Args))
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func calibreOutput(args ...string) (output []byte) {
	var cli []string
	cli = append(cli, "--with-library="+viper.GetString("url"))
	cli = append(cli, args...)
	cmd := exec.Command("calibredb", cli...)
	cmd.Stderr = os.Stderr
	Debug(fmt.Sprintf("[XeQ]:%v", cmd.Args))
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return
}
