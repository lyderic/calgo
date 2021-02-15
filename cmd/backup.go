package cmd

import (
	"os"
	"os/exec"
	"time"

	. "calgo/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup ebooks and calibre configuration",
	Run: func(cmd *cobra.Command, args []string) {
		backup()
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}

// write a report stating the date + possibly other useful information
func backup() {
	backupdir := viper.GetString("backup-dir")
	ebooksdir := viper.GetString("ebooks-dir")
	configdir := viper.GetString("calibre-config-dir")
	if _, err := os.Stat(backupdir); os.IsNotExist(err) {
		Red("Backup directory not found! %q\n", backupdir)
		return
	}
	rsync(ebooksdir, backupdir+"/ebooks")
	rsync(configdir, backupdir+"/calibre-config")
}

func rsync(src, dst string) {
	start := time.Now()
	cmd := exec.Command("rsync", "-Pav", "--delete", src+"/", dst+"/")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	Debug("[XeQ]-%v\n", cmd.Args)
	cmd.Start()
	cmd.Wait()
	Green("Synced %q to %q in %v\n", src, dst, time.Since(start))
}
