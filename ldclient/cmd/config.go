package cmd

import (
	"fmt"
	"log"

	"github.com/sgrimee/godomus"
	"github.com/spf13/viper"
)

var cfgFile string

type Config struct {
	DomusConfig godomus.Config
	Debug       bool
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ldclient.yaml)")

	RootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debugging")

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

// getConfig reads in config file and ENV variables if set and returns a config
func getConfig() Config {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".ldclient") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	gdCfg := godomus.Config{
		SiteKey:    cfgSiteKey(),
		UserKey:    cfgUserKey(),
		Password:   cfgPassword(),
		Url:        cfgUrl(),
		SocketPort: cfgSocketPort(),
		Debug:      viper.GetBool("debug"),
	}

	return Config{
		DomusConfig: gdCfg,
		Debug:       viper.GetBool("debug"),
	}
}

func cfgSiteKey() godomus.SiteKey {
	if !viper.IsSet("site") || (viper.GetInt("site") < 1) {
		log.Fatal("You must give a site (int), or set it in config file")
	}
	return godomus.NewSiteKey(viper.GetInt("site"))
}

func cfgUserKey() godomus.UserKey {
	if !viper.IsSet("user") || (viper.GetInt("user") < 1) {
		log.Fatal("You must give a user (int), or set it in config file")
	}
	return godomus.NewUserKey(viper.GetInt("user"))
}

func cfgPassword() string {
	if !viper.IsSet("password") || (viper.GetString("password") == "") {
		log.Fatal("You must give a password, or set it in config file")
	}
	return viper.GetString("password")
}

func cfgUrl() string {
	if !viper.IsSet("url") || (viper.GetString("url") == "") {
		log.Fatal("You must give a url, or set it in config file")
	}
	return viper.GetString("url")
}

func cfgSocketPort() int {
	if !viper.IsSet("socket_port") || (viper.GetInt("socket_port") < 1) {
		log.Fatal("You must give a socket_port (int), or set it in config file")
	}
	return viper.GetInt("socket_port")
}
