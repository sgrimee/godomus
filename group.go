package godomus

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type GroupKey string

type Group struct {
	Key     GroupKey `json:"group_key"`
	Label   string   `json:"label"`
	Resume  string   `json:"resume"`
	Family  string   `json:"family"`
	Devices []Device
	Actions []Action
	States  []State
}

// GetGroups returns all groups
func (d *Domus) GetGroups() ([]Group, error) {
	queries := map[string]string{}
	resp, err := d.GetWithSession("/Mobile/GetGroups", queries)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body struct {
		Groups []Group `json:"group"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return []Group(body.Groups), nil
}

// GetGroup returns a single group
func (d *Domus) GetGroup(gk GroupKey) (Group, error) {
	var group Group
	queries := map[string]string{
		"group_key": string(gk),
	}
	resp, err := d.GetWithSession("/Mobile/GetGroup", queries)
	if err != nil {
		return group, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return group, err
	}
	return group, nil
}

// avoid recursion in UnmarshalJSON
type group Group

// UnmarshalJSON fixes the Devices, Actions and States fields to be lists
func (g *Group) UnmarshalJSON(data []byte) error {
	tmp := struct {
		group
		DevicesGroups struct {
			DevicesGroup []Device `json:"devices_group"`
		} `json:"devices_groups"`
		Actions struct{ Action []Action }
		States  struct{ State []State }
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = Group(tmp.group)
	g.Devices = tmp.DevicesGroups.DevicesGroup
	g.Actions = tmp.Actions.Action
	g.States = tmp.States.State
	return nil
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
