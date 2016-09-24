package cmd

import (
	"log"
	"strconv"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// motorCmd represents the motor command
var motorCmd = &cobra.Command{
	Use:   "motor",
	Short: "Action motor up, down, or stop it.",
	Long:  `The first argument should be 'up', 'down' or 'stop'. The second is the device number.`,
	Run: func(cmd *cobra.Command, args []string) {
		var action godomus.ActionClassId
		var property godomus.PropClassId
		if len(args) < 2 {
			log.Fatal("You must specify the state and the device.")
		}
		d, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Error reading device number (int): %s\n", err)
		}
		switch args[0] {
		case "up":
			property = godomus.PropClassId("CLSID-DEVC-PROP-MOTOR-UD")
			action = godomus.ActionClassId("CLSID-ACTION-UP")
		case "down":
			property = godomus.PropClassId("CLSID-DEVC-PROP-MOTOR-UD")
			action = godomus.ActionClassId("CLSID-ACTION-DOWN")
		case "stop":
			property = godomus.PropClassId("CLSID-DEVC-PROP-MOTOR-SW-STOP")
			action = godomus.ActionClassId("CLSID-ACTION-STOP")
		default:
			log.Fatal("Unknown action. Use 'up' or 'down'.\n")
		}
		domusLogin()
		err = domus.ExecuteAction(action, property, godomus.NewDeviceKey(d))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(motorCmd)
}
