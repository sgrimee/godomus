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
		var (
			roomKeys []godomus.RoomKey
			devices  godomus.Devices
		)
		// only one room if specified, otherwise all rooms
		if room != 0 {
			roomKeys = append(roomKeys, godomus.NewRoomKey(room))
			domusLogin()
		} else {
			rooms := godomus.RoomsFromInfos(domusInfos())
			if len(rooms) < 1 {
				log.Fatal("No rooms found, are site and userid correct?")
			}
			for _, r := range rooms {
				roomKeys = append(roomKeys, r.Key)
			}
		}

		for _, rk := range roomKeys {
			devs, err := domus.DevicesInRoom(rk, godomus.CategoryClassId(""))
			if err != nil {
				log.Fatalf("Could not get devices for room %d: %s\n", rk.Num(), err)
			}
			for i, _ := range devs {
				devs[i].RoomKey = rk
			}
			devices = append(devices, devs...)
		}
		output(outputFormat, devices)
	},
}

func init() {
	getCmd.AddCommand(devicesCmd)
	devicesCmd.Flags().IntVarP(&room, "room", "r", 0, "Only devices in room number")
}
