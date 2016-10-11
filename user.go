package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

type UserKey string
type Users []User

type User struct {
	Key        UserKey `json:"user_key"`
	Nickname   string  `json:"nickname"`
	PictureKey string  `json:"picture_key"`
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
