package cmd

import (
	"fmt"
	"os"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// scenariosCmd represents the scenarios command
var scenariosCmd = &cobra.Command{
	Use:   "scenarios",
	Short: "Get available scenarios",
	Long:  `Get a list of scenarios that can be launched.`,
	Run: func(cmd *cobra.Command, args []string) {
		scenarios := godomus.ScenariosFromInfos(domusInfos())
		if len(scenarios) < 1 {
			fmt.Println("No scenarios found, are site and userid correct?")
			os.Exit(-1)
		}
		output(outputFormat, scenarios)
	},
}

func init() {
	getCmd.AddCommand(scenariosCmd)
}
