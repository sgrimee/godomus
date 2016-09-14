package cmd

import (
	"fmt"
	"os"

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
			fmt.Println(err)
			os.Exit(-1)
		}
		output(outputFormat, sites)
	},
}

func init() {
	getCmd.AddCommand(sitesCmd)
}
