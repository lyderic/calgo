package cmd

import (
	"os"

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
	backupCmd.PersistentFlags().BoolP("embed", "e", false, "embed metadata before syncing")
	viper.BindPFlag("embed", backupCmd.PersistentFlags().Lookup("embed"))
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
	if viper.GetBool("embed") {
		Embed()
	}
	Rsync(ebooksdir, backupdir+"/ebooks")
	Rsync(configdir, backupdir+"/dotconfig_calibre")
}
