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
	case godomus.Sites:
		for _, e := range obj.(godomus.Sites) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Label)
		}
	case godomus.Users:
		for _, e := range obj.(godomus.Users) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Nickname)
		}
	case godomus.Rooms:
		for _, e := range obj.(godomus.Rooms) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Label)
		}
	case godomus.Scenarios:
		for _, e := range obj.(godomus.Scenarios) {
			fmt.Printf("| %5d | %s\n", e.Key.Num(), e.Label)
		}
	case godomus.Devices:
		for _, e := range obj.(godomus.Devices) {
			fmt.Printf("| %5d | %18s | %5d | %s\n", e.RoomKey.Num(), e.RoomLabel, e.Key.Num(), e.Label)
		}
	case godomus.Device:
		e := obj.(godomus.Device)
		fmt.Printf("| %18s | %5d | %s\n", e.RoomLabel, e.Key.Num(), e.Label)
	case godomus.Categories:
		for _, e := range obj.(godomus.Categories) {
			fmt.Printf("| %16s | %28s | %2d devices\n", e.CatClsId, e.Label, e.DevicesCount)
		}
	}
}
