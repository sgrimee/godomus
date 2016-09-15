package cmd

import (
	"log"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

var room int

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Get devices in a room",
	Long:  `Get list of devices in the given room`,
	Run: func(cmd *cobra.Command, args []string) {
		if room == 0 {
			log.Fatal("You must specify the room number")
		}
		domusLogin()
		devices, err := domus.DevicesInRoom(godomus.NewRoomKey(room), godomus.CategoryClassId(""))
		if err != nil {
			log.Fatalf("Could not get devices for room %d: %s\n", room, err)
		}
		output(outputFormat, devices)
	},
}

func init() {
	getCmd.AddCommand(devicesCmd)
	devicesCmd.Flags().IntVarP(&room, "room", "r", 0, "Room number")
	devicesCmd.MarkFlagRequired("room")
}
