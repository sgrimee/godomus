package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

const sessionKeyLen = 40

type SessionKey string
type TargetKey string
type SiteKey string
type UserKey string
type RoomKey string

type Sites []Site
type Users []User
type Rooms []Room
type Devices []Device
type Categories []Category

type CategoryClassId string
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

type Value struct {
	Index       string `json:"index"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Label       string `json:"label"`
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

// (RoomKey) Num returns the room number as integer
func (uk RoomKey) Num() int {
	ns := strings.TrimPrefix(string(uk), "ROOM_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}

// RoomsFromInfos extracts rooms from an info object
// This only keeps rooms with a valid RoomKey, removing GROUPS, CAMERAS, etc
func RoomsFromInfos(infos LoginInfos) Rooms {
	var filtered Rooms
	for _, r := range infos.Rooms.Room {
		if strings.HasPrefix(string(r.Key), "ROOM_") {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

// ScenariosFromInfos extracts scenarions from an info object
func ScenariosFromInfos(infos LoginInfos) Scenarios {
	return Scenarios(infos.Scenarios.Scenario)
}
