package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// scenariosCmd represents the scenarios command
var scenariosCmd = &cobra.Command{
	Use:   "scenarios",
	Short: "Get available scenarios",
	Long:  `Get a list of scenarios that can be launched.`,
	Run: func(cmd *cobra.Command, args []string) {
		infos, err := domus.LoginInfos()
		if err != nil {
			log.Fatal(err)
		}
		scenarios := infos.Scenarios
		if len(scenarios) < 1 {
			log.Fatal("No scenarios found, are site and userid correct?")
		}
		output(outputFormat, scenarios)
	},
}

func init() {
	getCmd.AddCommand(scenariosCmd)
}
