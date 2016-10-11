package cmd

import (
	"log"
	"strconv"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn device off",
	Long: `Turn device off.
	The first argument is the device number.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You must specify the device.")
		}

		dnum, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error reading int device number: %s\n", err)
		}

		domusLogin()
		dev, err := domus.GetDeviceState(godomus.NewDeviceKey(dnum))
		if err != nil {
			log.Fatal(err)
		}

		err = dev.Off()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(offCmd)
}
