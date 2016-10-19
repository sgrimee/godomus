package cmd

import "github.com/spf13/cobra"

var outputFormat string // set by command-line parameter

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve data from the LD server",
	Long:  `Retrieve objects from the LD server`,
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "text", "Output format: text, json or yaml")
}
