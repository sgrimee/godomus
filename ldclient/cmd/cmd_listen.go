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

		devices := make(chan godomus.Device, 1)
		errs := make(chan error, 1)
		done := make(chan struct{})
		defer close(done)
		go domus.ListenForDeviceUpdates(devices, errs, done)

		switch outputFormat {
		case "text":
			fmt.Println("Waiting for events, Ctrl+C to exit")
		case "json":
			fmt.Println("[")
		}

		for {
			select {
			case d := <-devices:
				output(outputFormat, godomus.DeviceUpdate(d))
				if outputFormat == "json" {
					fmt.Println(",")
				}
			case err := <-errs:
				log.Fatal(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listenCmd)
	listenCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "text", "Output format: text, json or yaml")
}
