package cmd

import (
	"log"
	"strconv"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch device on or off",
	Long: `Switch device on or off.
	The first argument should be 'up', 'off', or 'toggle'. The second is the device number.`,
	Run: func(cmd *cobra.Command, args []string) {
		property := godomus.PropClassId("CLSID-DEVC-PROP-TOR-SW")
		if len(args) < 2 {
			log.Fatal("You must specify the state and the device.")
		}
		d, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Error reading device number (int): %s\n", err)
		}
		var action godomus.ActionClassId
		switch args[0] {
		case "on":
			action = godomus.ActionClassId("CLSID-ACTION-ON")
		case "off":
			action = godomus.ActionClassId("CLSID-ACTION-OFF")
		case "toggle":
			action = godomus.ActionClassId("CLSID-ACTION-TOGGLE")
		default:
			log.Fatal("Unknown action. Use 'on', 'off' or 'toggle'.\n")
		}
		domusLogin()
		err = domus.ExecuteAction(action, property, godomus.NewDeviceKey(d))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(switchCmd)
}
