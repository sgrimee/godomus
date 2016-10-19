package cmd

import (
	"log"
	"strconv"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Get a single device",
	Long: `Get infos on a single device
	ldclient get device 110`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("You must specify a device.")
		}
		d, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error reading device number (int): %s\n", err)
		}
		domusLogin()
		dev, err := domus.GetDeviceState(godomus.NewDeviceKey(d))
		if err != nil {
			log.Fatal(err)
		}
		output(outputFormat, dev)
	},
}

func init() {
	getCmd.AddCommand(deviceCmd)
}
