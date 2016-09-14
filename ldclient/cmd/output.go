package cmd

import (
	"encoding/json"
	"fmt"
	"os"

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
		fmt.Printf("Unknown output format: %s\n", format)
	}
}

func printJson(obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(-1)
	}
	fmt.Println(string(j))
}

func printYaml(obj interface{}) {
	y, err := yaml.Marshal(obj)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(string(y))
}

func printText(obj interface{}) {
	switch t := obj.(type) {
	default:
		fmt.Printf("printText does not support type %T\n", t)
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
	}
}
