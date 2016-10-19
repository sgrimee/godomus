package cmd

import (
	"log"

	"github.com/spf13/cobra"
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
		log.Fatal(err)
	}
}
