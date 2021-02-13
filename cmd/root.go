package cmd

import (
	"fmt"
	"os"

	. "calgo/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "calgo",
	Short: "Helper app to manage calibre database",
	//Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	debug := false
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calgo.yaml)")
	rootCmd.PersistentFlags().BoolP("debug", "", debug, "Show debugging information")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(os.Getenv("HOME"))
		viper.SetConfigName(".calgo")
	}
	//viper.AutomaticEnv() // other language in config file is ignored! (as
	// $LANGUAGE envvar takes precedence
	if err := viper.ReadInConfig(); err == nil {
		Debug("Using config file: %s\n", viper.ConfigFileUsed())
	}
	Debug("viper: %#v\n", viper.AllSettings())
}
