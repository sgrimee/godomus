package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

type RoomKey string

type Room struct {
	Key        RoomKey    `json:"room_key"`
	Label      string     `json:"label"`
	PictureKey PictureKey `json:"picture_key"`
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

// IsValidRoom returns true only for valid RoomKey
// thus removing GROUPS, CAMERAS, etc
func isValidRoom(r Room) bool {
	return strings.HasPrefix(string(r.Key), "ROOM_")
}

// RoomsFromInfos extracts rooms from an info object
// This only keeps rooms with a
func RoomsFromInfos(infos LoginInfos) []Room {
	var filtered []Room
	for _, r := range infos.Rooms {
		if isValidRoom(r) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
