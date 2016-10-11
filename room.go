package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

type RoomKey string
type Rooms []Room

type Room struct {
	Key        RoomKey `json:"room_key"`
	Label      string  `json:"label"`
	PictureKey string  `json:"picture_key"`
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
