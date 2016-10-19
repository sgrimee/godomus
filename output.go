package godomus

import (
	"encoding/json"
	"fmt"
	"io"

	yaml "gopkg.in/yaml.v2"
)

// output prints the given object to the writer in the desired format
// format can be 'text', 'json' or 'yaml'
func Output(w io.Writer, format string, obj interface{}) error {
	switch format {
	case "yaml":
		return printYaml(w, obj)
	case "json":
		return printJson(w, obj)
	case "text":
		return printText(w, obj)
	default:
		return fmt.Errorf("Unknown output format: %s\n", format)
	}
}

// BUG(sg): piping to jq does not work any more with json stream (event listening)

func printJson(w io.Writer, obj interface{}) error {
	j, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if _, err = io.WriteString(w, fmt.Sprintln(string(j))); err != nil {
		return err
	}

	return nil
}

func printYaml(w io.Writer, obj interface{}) error {
	y, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, fmt.Sprintf("---\n%s\n", string(y)))
	return err
}

func printText(w io.Writer, obj interface{}) (err error) {
	switch t := obj.(type) {
	default:
		return fmt.Errorf("printText does not support type %T\n", t)
	case string:
		s := obj.(string)
		if _, err = io.WriteString(w, s); err != nil {
			return err
		}
	case []Site:
		for _, e := range obj.([]Site) {
			if _, err = io.WriteString(w, fmt.Sprintf("| %5d | %s\n",
				e.Key.Num(), e.Label)); err != nil {
				return err
			}
		}
	case []User:
		for _, e := range obj.([]User) {
			if _, err = io.WriteString(w, fmt.Sprintf("| %5d | %s\n",
				e.Key.Num(), e.Nickname)); err != nil {
				return err
			}
		}
	case []Room:
		for _, e := range obj.([]Room) {
			if _, err = io.WriteString(w, fmt.Sprintf("| %5d | %s\n",
				e.Key.Num(), e.Label)); err != nil {
				return err
			}
		}
	case []Scenario:
		for _, e := range obj.([]Scenario) {
			if _, err = io.WriteString(w, fmt.Sprintf("| %5d | %s\n",
				e.Key.Num(), e.Label)); err != nil {
				return err
			}
		}
	case []Device:
		for _, e := range obj.([]Device) {
			if _, err = io.WriteString(w, fmt.Sprintf("| %18s | %5d | %35s | %s\n",
				e.RoomLabel, e.Key.Num(), e.Label, e.Resume)); err != nil {
				return err
			}
		}
	case Device:
		e := obj.(Device)
		if _, err = io.WriteString(w, fmt.Sprintf("| %18s | %5d | %s\n",
			e.RoomLabel, e.Key.Num(), e.Label)); err != nil {
			return err
		}
	case DeviceUpdate:
		d := Device(obj.(DeviceUpdate))
		val := d.States[0].Values[0]
		if _, err = io.WriteString(w, fmt.Sprintf("%20s: %5d - %25s - %s %s %s\n",
			d.RoomLabel, d.Key.Num(), d.Label, val.Label, val.Value, val.Unit)); err != nil {
			return err
		}
	case []Category:
		for _, e := range obj.([]Category) {
			if _, err = io.WriteString(w, fmt.Sprintf("| %16s | %28s | %2d devices\n",
				e.CatClsId, e.Label, e.DevicesCount)); err != nil {
				return err
			}
		}
	case []Group:
		for _, e := range obj.([]Group) {
			if _, err = io.WriteString(w, fmt.Sprintf("| %5d | %28s | %s\n",
				e.Key.Num(), e.Label, e.Resume)); err != nil {
				return err
			}
		}
	case Group:
		e := obj.(Group)
		if _, err = io.WriteString(w, fmt.Sprintf("| %5d | %18s | %s\n",
			e.Key.Num(), e.Label, e.Resume)); err != nil {
			return err
		}
		for _, d := range e.Devices {
			if _, err = io.WriteString(w, fmt.Sprintf("    %25s: %5d - %s\n",
				d.RoomLabel, d.Key.Num(), d.Label)); err != nil {
				return err
			}
		}
	}
	return err
}
