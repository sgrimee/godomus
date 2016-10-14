package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Get available device groups",
	Long:  `Get a list of device groups.`,
	Run: func(cmd *cobra.Command, args []string) {
		domusLogin()
		groups, err := domus.GetGroups()
		if err != nil {
			log.Fatal(err)
		}
		if len(groups) < 1 {
			log.Fatal("No groups found, are site and userid correct?")
		}
		output(outputFormat, groups)
	},
}

func init() {
	getCmd.AddCommand(groupsCmd)
}
