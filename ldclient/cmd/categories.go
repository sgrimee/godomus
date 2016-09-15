package cmd

import (
	"log"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Get categories of devices in a room",
	Long:  `Get categories of devices in a room`,
	Run: func(cmd *cobra.Command, args []string) {
		if room == 0 {
			log.Fatal("You must specify the room number")
		}
		domusLogin()
		categories, err := domus.GetCategories(godomus.NewRoomKey(room))
		if err != nil {
			log.Fatalf("Could not get categories for room %d: %s\n", room, err)
		}
		output(outputFormat, categories)
	},
}

func init() {
	getCmd.AddCommand(categoriesCmd)
	categoriesCmd.Flags().IntVarP(&room, "room", "r", 0, "Room number")
	categoriesCmd.MarkFlagRequired("room")
}
