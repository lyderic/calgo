package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/lyderic/tools"
	"github.com/spf13/viper"
)

func Calibre(args []string) {
	StartCalibre()
	var cli []string
	cli = append(cli, "--with-library="+libraryUrl())
	cli = append(cli, args...)
	cmd := exec.Command("calibredb", cli...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	Debug(fmt.Sprintf("[XeQ]:%v", cmd.Args))
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
	Debug(fmt.Sprintf("[XeQ]:%v", cmd.Args))
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
	Debug(fmt.Sprintf("[XeQ]:%v", cmd.Args))
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
	Debug("Looking for calibre...\n")
	if localhostPortIsInUse(viper.GetString("port")) {
		Debug("calibre is running")
		return
	}
	tools.PrintRed("\nCalibre is not running! Starting, please wait..")
	cmd := exec.Command("calibre", "--detach", "--no-update-check", "--start-in-tray")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()
	if err != nil {
		tools.PrintRedln("calibre couldn't be started!!! Aborting")
		log.Fatal(err)
	}
	for {
		if !localhostPortIsInUse(viper.GetString("port")) {
			tools.PrintRed(".")
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	tools.PrintRedln(" done.")

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
		Debug("Port %s is in use\n", port)
	}
	return
}
