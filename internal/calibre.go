package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/viper"
)

func Calibre(args []string) {
	StartCalibre()
	var cli []string
	cli = append(cli, "--with-library="+libraryUrl())
	cli = append(cli, args...)
	cmd := exec.Command("calibredb", cli...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	Debug(fmt.Sprintf("[XeQ]:%v\n", cmd.Args))
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func CalibreOutput(args ...string) (output []byte) {
	StartCalibre()
	var cli []string
	cli = append(cli, "--with-library="+libraryUrl())
	cli = append(cli, args...)
	cmd := exec.Command("calibredb", cli...)
	cmd.Stderr = os.Stderr
	Debug(fmt.Sprintf("[XeQ]:%v\n", cmd.Args))
	output, err := cmd.Output()
	if err != nil {
		Debug("%#v\n", err)
	}
	return
}

func CalibreOutputErr(args ...string) ([]byte, error) {
	StartCalibre()
	var cli []string
	cli = append(cli, "--with-library="+libraryUrl())
	cli = append(cli, args...)
	cmd := exec.Command("calibredb", cli...)
	Debug(fmt.Sprintf("[XeQ]:%v\n", cmd.Args))
	return cmd.Output()
}

func Calibredb(args []string) {
	StartCalibre()
	var cli []string
	cli = append(cli, "--with-library="+libraryUrl())
	cli = append(cli, args...)
	cmd := exec.Command("calibredb", cli...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	Debug("[XeQ]:%v\n", cmd.Args)
	cmd.Run()
}

func StartCalibre() {
	Debug("Looking for calibre... ")
	if localhostPortIsInUse(viper.GetString("port")) {
		Debug("calibre is running. All good!\n")
		return
	}
	Yellow("\nCalibre is not running! Starting, please wait..")
	cmd := exec.Command("calibre", "--detach", "--no-update-check", "--start-in-tray")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()
	if err != nil {
		Red("calibre couldn't be started!!! Aborting")
		log.Fatal(err)
	}
	for {
		if !localhostPortIsInUse(viper.GetString("port")) {
			Yellow(".")
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	Yellow(" done.\n")

}

func libraryUrl() string {
	return fmt.Sprintf("http://localhost:%s#%s",
		viper.GetString("port"),
		viper.GetString("library"))
}

func localhostPortIsInUse(port string) (inUse bool) {
	inUse = false
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Second)
	if conn != nil {
		conn.Close()
		inUse = true
		Debug("tcp/%s is in use... ", port)
	}
	return
}
