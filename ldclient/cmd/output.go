package cmd

import (
	"log"
	"os"

	"github.com/sgrimee/godomus"
)

// output prints the object in the given format to stdout
func output(format string, obj interface{}) {
	if err := godomus.Output(os.Stdout, format, obj); err != nil {
		log.Fatal(err)
	}
}
