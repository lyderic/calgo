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

func Calibredb(args ...string) (cmd *exec.Cmd) {
	var cli []string
	cli = append(cli, "--with-library="+libraryUrl())
	cli = append(cli, args...)
	cmd = exec.Command("calibredb", cli...)
	Debug(fmt.Sprintf("[XeQ]:%v\n", cmd.Args))
	return
}

func CalibreOutput(args ...string) (output []byte) {
	cmd := Calibredb(args...)
	cmd.Stderr = os.Stderr
	output, _ = cmd.Output()
	return
}

func CalibreOutputErr(args ...string) ([]byte, error) {
	cmd := Calibredb(args...)
	return cmd.Output()
}

func StartCalibre() {
	Debug("Looking for calibre... ")
	if portInUse(viper.GetString("port")) {
		Debug("calibre is running. All good!\n")
		return
	}
	Yellow("Calibre is not running! Starting, please wait..")
	cmd := exec.Command("calibre", "--detach", "--no-update-check", "--start-in-tray")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()
	if err != nil {
		Red("calibre couldn't be started!!! Aborting")
		log.Fatal(err)
	}
	for {
		if !portInUse(viper.GetString("port")) {
			Yellow(".")
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	Green(" done.\n")
}

func libraryUrl() string {
	return fmt.Sprintf("http://localhost:%s#%s",
		viper.GetString("port"),
		viper.GetString("library"))
}

func portInUse(port string) (inUse bool) {
	inUse = false
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Second)
	if conn != nil {
		conn.Close()
		inUse = true
		Debug("tcp/%s is in use... ", port)
	}
	return
}
