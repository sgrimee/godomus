package godomus

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const PropClassIdBinarySwitch = PropClassId("CLSID-DEVC-PROP-TOR-SW")
const PropClassIdDimmerSwitch = PropClassId("CLSID-DEVC-PROP-DIMMER-SW")
const PropClassIdMotorUpDown = PropClassId("CLSID-DEVC-PROP-MOTOR-UD")

type DeviceKey TargetKey
type DeviceClassId string
type PropClassId string

type Device struct {
	Key       DeviceKey       `json:"device_key"`
	DevClsId  DeviceClassId   `json:"device_clsid"`
	RoomLabel string          `json:"room_label"`
	CatClsId  CategoryClassId `json:"category_clsid"`
	Label     string          `json:"label"`
	Resume    string          `json:"resume"`
	States    []State
	Actions   []Action
	RoomKey   *RoomKey
	server    *Domus
}

// GetDeviceState returns all infos on one device
func (d *Domus) GetDeviceState(dk DeviceKey) (Device, error) {
	var dev Device
	queries := map[string]string{
		"device_key": string(dk),
	}
	resp, err := d.GetWithSession("/Mobile/GetDeviceState", queries)
	if err != nil {
		return dev, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&dev); err != nil {
		return dev, err
	}
	dev.server = d
	return dev, nil
}

// DevicesInRoom returns the list of devices in the given roomId
// If class is "" then all devices are returned, otherwise only devices of that class
func (d *Domus) DevicesInRoom(rk RoomKey, class CategoryClassId) ([]Device, error) {
	queries := map[string]string{
		"room_key": string(rk),
	}
	if class != "" {
		queries["category_clsid"] = string(class)
	}
	resp, err := d.GetWithSession("/Mobile/GetDevices", queries)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		Devices []Device `json:"device"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	devices := []Device(body.Devices)
	for i := range devices {
		devices[i].server = d
		devices[i].RoomKey = &rk
	}
	return devices, nil
}

// avoid recursion in UnmarshalJSON
type device Device

// UnmarshalJSON fixes the States and Actions fields to be lists
// It fixes Label if given as "device_label" instead of "label"
func (dev *Device) UnmarshalJSON(data []byte) error {
	tmp := struct {
		device
		DeviceLabel string `json:"device_label"`
		States      struct {
			State []State `json:"state"`
		} `json:"states"`
		Actions struct {
			Action []Action `json:"action"`
		} `json:"actions"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*dev = Device(tmp.device)
	if (tmp.device.Label == "") && (tmp.DeviceLabel != "") {
		dev.Label = tmp.DeviceLabel
	}
	dev.States = tmp.States.State
	dev.Actions = tmp.Actions.Action
	//fmt.Printf("dev: %+v\n", dev)
	return nil
}

func (dev Device) On() error {
	return dev.switchAction(ActionClassIdOn)
}

func (dev Device) Off() error {
	return dev.switchAction(ActionClassIdOff)
}

func (dev Device) Toggle() error {
	return dev.switchAction(ActionClassIdToggle)
}

func (dev Device) switchAction(action ActionClassId) error {
	// is it a supported action
	switch action {
	case ActionClassIdOn,
		ActionClassIdOff,
		ActionClassIdToggle:
	default:
		return errors.New("switchAction: invalid action class id")
	}

	// is it a supported device
	for _, devAction := range dev.Actions {
		property := devAction.PropClsId
		switch property {
		case PropClassIdBinarySwitch,
			PropClassIdDimmerSwitch:
			if dev.server == nil {
				return errors.New("switchAction: device server reference is nil")
			}
			return dev.server.ExecuteAction(action, property, TargetKey(dev.Key))
		}
	}
	return errors.New("switchAction: invalid property class id")
}

func (dev Device) Up() error {
	return dev.motorAction(ActionClassIdUp)
}

func (dev Device) Down() error {
	return dev.motorAction(ActionClassIdDown)
}

func (dev Device) motorAction(action ActionClassId) error {
	// is it a supported action
	switch action {
	case ActionClassIdUp,
		ActionClassIdDown:
	default:
		return errors.New("motorAction: invalid action class id")
	}

	// is it a supported device
	for _, devAction := range dev.Actions {
		property := devAction.PropClsId
		switch property {
		case PropClassIdMotorUpDown:
			if dev.server == nil {
				return errors.New("motorAction: device server reference is nil")
			}
			return dev.server.ExecuteAction(action, property, TargetKey(dev.Key))
		}
	}
	return errors.New("motorAction: invalid property class id")
}

// NewDeviceKey returns a device key from a device number as integer
func NewDeviceKey(num int) DeviceKey {
	return DeviceKey(fmt.Sprintf("DEVC_%035d", num))
}

// (DeviceKey) Num returns the device number as integer
func (dk DeviceKey) Num() int {
	ns := strings.TrimPrefix(string(dk), "DEVC_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}
