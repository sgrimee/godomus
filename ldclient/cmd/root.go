package cmd

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	debug   bool
	domus   *godomus.Domus
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ldclient",
	Short: "Lifedomus command-line client",
	Long: `ldclient implements a partial client for the Lifedomus home 
automation server. You can get rooms, get and launch scenarions,
get and set devices state and also continuously listen for events.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

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

// initDomus sets-up the domus object
func initDomus() {
	var err error
	domus, err = godomus.New(
		viper.GetString("url"),
		viper.GetInt("socket_port"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if debug {
		domus.Debug = true
	}
}

// validateSiteSet ensures a site number was provided
func validateSiteSet() {
	if !viper.IsSet("site") || (viper.GetInt("site") < 1) {
		fmt.Println("You must give a site (int), or set it in config file")
		os.Exit(-1)
	}
}

// validateUserSet ensures a userid and password were provided
func validateUserSet() {
	if !viper.IsSet("user") || (viper.GetInt("user") < 1) {
		fmt.Println("You must give a user (int), or set it in config file")
		os.Exit(-1)
	}
	if !viper.IsSet("password") || (viper.GetString("password") == "") {
		fmt.Println("You must give a password, or set it in config file")
		os.Exit(-1)
	}
}

// domusInfos logs in an returns infos
func domusInfos() *godomus.LoginInfos {
	validateSiteSet()
	validateUserSet()
	sk := godomus.NewSiteKey(viper.GetInt("site"))
	uk := godomus.NewUserKey(viper.GetInt("user"))
	pass := viper.GetString("password")
	infos, err := domus.LoginInfos(sk, uk, pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return infos
}
