package cmd

import (
	"os"
	"path/filepath"

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

// rename this command as 'backup'
// write a report stating the date + possibly other useful information
func backup() {
	basedir := viper.GetString("backup-dir")
	datadir := filepath.Join(basedir, viper.GetString("library"))
	//configdir := filepath.Join(basedir, viper.GetString("calibre-config-dir"))
	if _, err := os.Stat(basedir); os.IsNotExist(err) {
		Red("Backup directory not found! %q\n", basedir)
		return
	}
	cmd := Calibredb("export", "--all", "--progress", "--to-dir="+datadir)
	cmd.Stdout = os.Stdout
	cmd.Start()
	cmd.Wait()
}
