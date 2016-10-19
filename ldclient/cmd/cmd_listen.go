package cmd

import (
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
			output("text", "Waiting for events, Ctrl+C to exit")
		case "json":
			output("text", "[")
		}

		for {
			select {
			case d := <-devices:
				output(outputFormat, godomus.DeviceUpdate(d))
				if outputFormat == "json" {
					output("text", ",")
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
