package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

type GroupKey string
type Groups []Group

// The API for groups returns a bogus key for the device label
type DeviceInGroup struct {
	Device
	Label string `json:"device_label"`
}

type Group struct {
	Key           GroupKey                  `json:"group_key"`
	Label         string                    `json:"label"`
	Resume        string                    `json:"resume"`
	Family        string                    `json:"family"`
	Actions       struct{ Action []Action } `json:"actions"`
	States        struct{ State []State }   `json:"states"`
	DevicesGroups struct {
		DevicesGroup []DeviceInGroup `json:"devices_group"`
	} `json:"devices_groups"`
}

// NewGroupKey returns a GroupKey from a group number as integer
func NewGroupKey(num int) GroupKey {
	return GroupKey(fmt.Sprintf("GRPE_%035d", num))
}

// (GroupKey) Num returns the group number as integer
func (uk GroupKey) Num() int {
	ns := strings.TrimPrefix(string(uk), "GRPE_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}
