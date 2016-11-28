package cmd

import (
	"log"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

var devicesCmdRoom int

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Get devices",
	Long:  `Get list of devices in the given room, or all rooms (slow))`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			roomKeys []godomus.RoomKey
			devices  []godomus.Device
		)
		// only one room if specified, otherwise all rooms
		if devicesCmdRoom != 0 {
			roomKeys = append(roomKeys, godomus.NewRoomKey(devicesCmdRoom))
		} else {
			infos, err := domus.LoginInfos()
			if err != nil {
				log.Fatal(err)
			}
			rooms := godomus.RoomsFromInfos(infos)
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
			for i := range devs {
				devs[i].RoomKey = &rk
			}
			devices = append(devices, devs...)
		}
		output(outputFormat, devices)
	},
}

func init() {
	getCmd.AddCommand(devicesCmd)
	devicesCmd.Flags().IntVarP(&devicesCmdRoom, "room", "r", 0, "room number")
}
