package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sgrimee/godomus"

	yaml "gopkg.in/yaml.v2"
)

// output prints the given object to the writer in the desired format
// format can be 'text', 'json' or 'yaml'
func output(format string, obj interface{}) {
	switch format {
	case "yaml":
		printYaml(obj)
	case "json":
		printJson(obj)
	case "text":
		printText(obj)
	default:
		log.Fatalf("Unknown output format: %s\n", format)
	}
}

func printJson(obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}

func printYaml(obj interface{}) {
	y, err := yaml.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(y))
}

func printText(obj interface{}) {
	switch t := obj.(type) {
	default:
		log.Fatalf("printText does not support type %T\n", t)
	case []godomus.Site:
		for _, e := range obj.([]godomus.Site) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Label)
		}
	case []godomus.User:
		for _, e := range obj.([]godomus.User) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Nickname)
		}
	case []godomus.Room:
		for _, e := range obj.([]godomus.Room) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Label)
		}
	case []godomus.Scenario:
		for _, e := range obj.([]godomus.Scenario) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Label)
		}
	case []godomus.Device:
		for _, e := range obj.([]godomus.Device) {
			fmt.Printf("| %18s | %5d | %35s | %s\n",
				e.RoomLabel, e.Key.Num(), e.Label, e.Resume)
		}
	case godomus.Device:
		e := obj.(godomus.Device)
		fmt.Printf("| %18s | %5d | %s\n", e.RoomLabel, e.Key.Num(), e.Label)
	case []godomus.Category:
		for _, e := range obj.([]godomus.Category) {
			fmt.Printf("| %16s | %28s | %2d devices\n", e.CatClsId, e.Label, e.DevicesCount)
		}
	case []godomus.Group:
		for _, e := range obj.([]godomus.Group) {
			fmt.Printf("| %5d | %28s | %s\n", e.Key.Num(), e.Label, e.Resume)
		}
	case godomus.Group:
		e := obj.(godomus.Group)
		fmt.Printf("| %5d | %18s | %s\n", e.Key.Num(), e.Label, e.Resume)
		for _, d := range e.Devices {
			fmt.Printf("    %25s: %5d - %s\n", d.RoomLabel, d.Key.Num(), d.Label)
		}
	}
}
