package cmd

import (
	"log"
	"strconv"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// upCmd represents the on command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up device",
	Long: `Bring device up.
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

		err = dev.Up()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
