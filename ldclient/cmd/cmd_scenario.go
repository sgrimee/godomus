package cmd

import (
	"log"
	"strconv"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// scenarioCmd represents the scenario command
var scenarioCmd = &cobra.Command{
	Use:   "scenario",
	Short: "Launch a scenario",
	Long:  `Launch a scenario.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You must specify the scenario.")
		}

		snum, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error reading scenario number: %s\n", err)
		}
		sk := godomus.NewScenarioKey(snum)

		domusLogin()
		if err = domus.RunScenario(sk); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(scenarioCmd)
}
