package cmd

import (
	"fmt"
	"log"

	"github.com/sgrimee/godomus"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Continuously listen for device update events",
	Run: func(cmd *cobra.Command, args []string) {
		domusLogin()

		events := make(chan godomus.EventMsg, 1)
		errs := make(chan error, 1)
		go domus.ListenForEvents(events, errs)

		fmt.Println("Waiting for events, Ctrl+C to exit")
		for {
			select {
			case ev := <-events:
				//fmt.Printf("Event received: %+v\n", ev)
				d, err := domus.GetDeviceState(ev.DeviceKey)
				if err != nil {
					log.Fatal(err)
				}
				val := d.States[0].Values[0]
				fmt.Printf("%20s: %5d - %25s - %s %s %s\n",
					d.RoomLabel, d.Key.Num(), d.Label, val.Label, val.Value, val.Unit)
			case err := <-errs:
				log.Fatal(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listenCmd)
}
