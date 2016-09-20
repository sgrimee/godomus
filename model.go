package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

const sessionKeyLen = 40

type SessionKey string

type SiteKey string
type UserKey string
type RoomKey string
type ScenarioKey string
type DeviceKey string

type Sites []Site
type Users []User
type Rooms []Room
type Devices []Device
type Scenarios []Scenario
type Categories []Category

type DeviceClassId string
type CategoryClassId string
type ActionClassId string
type PropClassId string
type StateClassId string

type Site struct {
	Key        SiteKey `json:"site_key"`
	Label      string  `json:"label"`
	PictureKey string  `json:"picture_key"`
}

type User struct {
	Key        UserKey `json:"user_key"`
	Nickname   string  `json:"nickname"`
	PictureKey string  `json:"picture_key"`
}

type Room struct {
	Key        RoomKey `json:"room_key"`
	Label      string  `json:"label"`
	PictureKey string  `json:"picture_key"`
}

type Scenario struct {
	Key    ScenarioKey `json:"scenario_key"`
	Label  string      `json:"label"`
	Resume string      `json:"resume"`
}

type Value struct {
	Index       string `json:"index"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Label       string `json:"label"`
}

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
}

type LoginInfos struct {
	SessionKey SessionKey                    `json:"session_key"`
	Rooms      struct{ Room []Room }         `json:"rooms"`
	Scenarios  struct{ Scenario []Scenario } `json:"scenarios"`
	//Devices    struct{ Device []Device }     `json:"devices"`
	// Bookmarks  interface{}                   `json:"bookmarks"`
	// Groups     interface{}                   `json:"groups"`
}

type Category struct {
	CatClsId     CategoryClassId `json:"category_clsid"`
	Label        string          `json:"label"`
	PictureKey   string          `json:"picture_key"`
	DevicesCount int             `json:"devices_count,string"`
}

// NewSiteKey returns a SiteKey from a site number as integer
func NewSiteKey(num int) SiteKey {
	return SiteKey(fmt.Sprintf("SITE_%035d", num))
}

// (SiteKey) Num returns the site number as integer
func (sk SiteKey) Num() int {
	ns := strings.TrimPrefix(string(sk), "SITE_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}

// NewUserKey returns a UserKey from a user number as integer
func NewUserKey(num int) UserKey {
	return UserKey(fmt.Sprintf("USER_%035d", num))
}

// (UserKey) Num returns the user number as integer
func (uk UserKey) Num() int {
	ns := strings.TrimPrefix(string(uk), "USER_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}

// NewRoomKey returns a RoomKey from a room number as integer
func NewRoomKey(num int) RoomKey {
	return RoomKey(fmt.Sprintf("ROOM_%035d", num))
}

// (RoomKey) Num returns the user number as integer
func (uk RoomKey) Num() int {
	ns := strings.TrimPrefix(string(uk), "ROOM_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}

// NewScenarioKey returns a Scenario from a scenario number as integer
func NewScenarioKey(num int) ScenarioKey {
	return ScenarioKey(fmt.Sprintf("SCNR_%035d", num))
}

// (ScenarioKey) Num returns the user number as integer
func (uk ScenarioKey) Num() int {
	ns := strings.TrimPrefix(string(uk), "SCNR_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}

// NewDeviceKey returns a Scenario from a scenario number as integer
func NewDeviceKey(num int) DeviceKey {
	return DeviceKey(fmt.Sprintf("DEVC_%035d", num))
}

// (DeviceKey) Num returns the user number as integer
func (dk DeviceKey) Num() int {
	ns := strings.TrimPrefix(string(dk), "DEVC_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}

// RoomsFromInfos extracts rooms from an info object
// This only keeps rooms with a valid RoomKey, removing GROUPS, CAMERAS, etc
func RoomsFromInfos(infos *LoginInfos) Rooms {
	var filtered Rooms
	for _, r := range infos.Rooms.Room {
		if strings.HasPrefix(string(r.Key), "ROOM_") {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

// ScenariosFromInfos extracts scenarions from an info object
func ScenariosFromInfos(infos *LoginInfos) Scenarios {
	return Scenarios(infos.Scenarios.Scenario)
}
