package cmd

import (
	"log"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// roomsCmd represents the rooms command
var roomsCmd = &cobra.Command{
	Use:   "rooms",
	Short: "Get list of rooms",
	Long:  `Get list of rooms for the given site and user`,
	Run: func(cmd *cobra.Command, args []string) {
		rooms := godomus.RoomsFromInfos(domusInfos())
		if len(rooms) < 1 {
			log.Fatal("No rooms found, are site and userid correct?")
		}
		output(outputFormat, rooms)
	},
}

func init() {
	getCmd.AddCommand(roomsCmd)
}
