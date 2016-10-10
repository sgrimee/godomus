package cmd

import (
	"fmt"
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
		if len(args) < 2 {
			log.Fatal("You must specify the state and the device.")
		}

		dnum, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Error reading int device number: %s\n", err)
		}

		domusLogin()
		dev, err := domus.GetDeviceState(godomus.NewDeviceKey(dnum))
		if err != nil {
			log.Fatal(err)
		}

		err = nil
		action := args[0]
		switch action {
		case "on":
			err = domus.On(dev)
		case "off":
			err = domus.Off(dev)
		case "toggle":
			err = domus.Toggle(dev)
		default:
			err = fmt.Errorf("Unsupported action: %s\n", action)
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(switchCmd)
}
