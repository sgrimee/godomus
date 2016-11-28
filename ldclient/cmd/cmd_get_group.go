package cmd

import (
	"log"
	"strconv"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Get a single group",
	Long:  `Get a single group given as first argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You must specify the group number.")
		}
		gnum, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error reading int group number: %s\n", err)
		}
		gk := godomus.NewGroupKey(gnum)

		group, err := domus.GetGroup(gk)
		if err != nil {
			log.Fatal(err)
		}

		output(outputFormat, group)
	},
}

func init() {
	getCmd.AddCommand(groupCmd)
}
