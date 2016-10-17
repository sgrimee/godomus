package godomus

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type UserKey string

type User struct {
	Key        UserKey    `json:"user_key"`
	Nickname   string     `json:"nickname"`
	PictureKey PictureKey `json:"picture_key"`
}

// GetUsers returns the list of users for a given site
func (d *Domus) GetUsers(site SiteKey) ([]User, error) {
	resp, err := d.Get("/Mobile/GetUsers", map[string]string{
		"site_key": string(site),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		List []User `json:"user"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return []User(body.List), nil
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
