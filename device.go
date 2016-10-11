package godomus

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const ActionClassIdOn = ActionClassId("CLSID-ACTION-ON")
const ActionClassIdOff = ActionClassId("CLSID-ACTION-OFF")
const ActionClassIdUp = ActionClassId("CLSID-ACTION-UP")
const ActionClassIdDown = ActionClassId("CLSID-ACTION-DOWN")
const ActionClassIdToggle = ActionClassId("CLSID-ACTION-TOGGLE")
const PropClassIdBinarySwitch = PropClassId("CLSID-DEVC-PROP-TOR-SW")
const PropClassIdDimmerSwitch = PropClassId("CLSID-DEVC-PROP-DIMMER-SW")
const PropClassIdMotorUpDown = PropClassId("CLSID-DEVC-PROP-MOTOR-UD")

type DeviceKey TargetKey
type Devices []Device
type DeviceClassId string
type ActionClassId string
type PropClassId string
type StateClassId string

type State struct {
	ClsId  StateClassId            `json:"state_clsid"`
	Type   string                  `json:"type"`
	Label  string                  `json:"label"`
	Values struct{ Value []Value } `json:"values"`
}

type Action struct {
	PropClsId PropClassId   `json:"prop_clsid"`
	ClsId     ActionClassId `json:"action_clsid"`
}

type Device struct {
	Key       DeviceKey                 `json:"device_key"`
	DevClsId  DeviceClassId             `json:"device_clsid"`
	RoomLabel string                    `json:"room_label"`
	CatClsId  CategoryClassId           `json:"category_clsid"`
	Label     string                    `json:"label"`
	Resume    string                    `json:"resume"`
	States    struct{ State []State }   `json:"states"`
	Actions   struct{ Action []Action } `json:"actions"`
	RoomKey   RoomKey
	server    *Domus
}

type Value struct {
	Index       string `json:"index"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Label       string `json:"label"`
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
		return errors.New("Invalid action class id")
	}

	// is it a supported device
	for _, devAction := range dev.Actions.Action {
		property := devAction.PropClsId
		switch property {
		case PropClassIdBinarySwitch,
			PropClassIdDimmerSwitch:
			return dev.server.ExecuteAction(action, property, TargetKey(dev.Key))
		}
	}
	return errors.New("Invalid property class id")
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
		return errors.New("Invalid action class id")
	}

	// is it a supported device
	for _, devAction := range dev.Actions.Action {
		property := devAction.PropClsId
		switch property {
		case PropClassIdMotorUpDown:
			return dev.server.ExecuteAction(action, property, TargetKey(dev.Key))
		}
	}
	return errors.New("Invalid property class id")
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
