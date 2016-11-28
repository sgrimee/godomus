package cmd

import (
	"log"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

var categoriesCmdRoom int

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Get categories",
	Long:  `Get categories of devices in a room, or all rooms (slow)`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			roomKeys   []godomus.RoomKey
			categories []godomus.Category
		)
		// only one room if specified, otherwise all rooms
		if categoriesCmdRoom != 0 {
			roomKeys = append(roomKeys, godomus.NewRoomKey(categoriesCmdRoom))
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

		// use a map for de-duplication and device count addition
		cmap := make(map[godomus.CategoryClassId]godomus.Category)
		for _, rk := range roomKeys {
			cir, err := domus.CategoriesInRoom(rk)
			if err != nil {
				log.Fatalf("Could not get categories for room %d: %s\n", rk.Num(), err)
			}
			for _, v := range cir {
				if _, ok := cmap[v.CatClsId]; ok {
					v.DevicesCount += cmap[v.CatClsId].DevicesCount
					cmap[v.CatClsId] = v
					continue
				}
				cmap[v.CatClsId] = v
			}
		}
		for _, v := range cmap {
			categories = append(categories, v)
		}
		output(outputFormat, categories)
	},
}

func init() {
	getCmd.AddCommand(categoriesCmd)
	categoriesCmd.Flags().IntVarP(&categoriesCmdRoom, "room", "r", 0, "room number")
}
