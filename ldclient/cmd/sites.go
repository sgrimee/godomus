package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// sitesCmd represents the sites command
var sitesCmd = &cobra.Command{
	Use:   "sites",
	Short: "Get sites managed by the LD server",
	Long:  `Get sites managed by the LD server`,
	Run: func(cmd *cobra.Command, args []string) {
		sites, err := domus.GetSites()
		if err != nil {
			log.Fatal(err)
		}
		output(outputFormat, sites)
	},
}

func init() {
	getCmd.AddCommand(sitesCmd)
}
