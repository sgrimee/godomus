package cmd

import (
	"log"

	"github.com/sgrimee/godomus"
)

var domus *godomus.Domus

func init() {
	d, err := godomus.New(getConfig().DomusConfig)
	if err != nil {
		log.Fatal(err)
	}
	domus = d
}
