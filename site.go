package godomus

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type SiteKey string

type Site struct {
	Key        SiteKey    `json:"site_key"`
	Label      string     `json:"label"`
	PictureKey PictureKey `json:"picture_key"`
}

// GetSites returns the list of sites managed by the server
func (d *Domus) GetSites() ([]Site, error) {
	resp, err := d.Get("/Mobile/GetSites", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		List []Site `json:"site"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return []Site(body.List), nil
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
