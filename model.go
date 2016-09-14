package godomus

import "strings"

type Sites []Site
type Users []User
type Rooms []Room
type Devices []Device
type Scenarios []Scenario

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
