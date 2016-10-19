package cmd

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig, initDomus)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ldclient.yaml)")

	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debugging")

	RootCmd.PersistentFlags().StringP("password", "p", "", "password")
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))

	RootCmd.PersistentFlags().StringP("ws", "w", "", "Web services url")
	viper.BindPFlag("ws", RootCmd.PersistentFlags().Lookup("ws"))

	RootCmd.PersistentFlags().IntP("socketPort", "k", 51023, "Port for the events socket")
	viper.BindPFlag("socket_port", RootCmd.PersistentFlags().Lookup("socketPort"))

	RootCmd.PersistentFlags().IntP("user", "u", 0, "User id")
	viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user"))

	RootCmd.PersistentFlags().IntP("site", "s", 0, "Site id")
	viper.BindPFlag("site", RootCmd.PersistentFlags().Lookup("site"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".ldclient") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && debug {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	if debug {
		fmt.Println("Config:")
		spew.Dump(viper.AllSettings())
	}
}
