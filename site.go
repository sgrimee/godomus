package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

type SiteKey string
type Sites []Site

type Site struct {
	Key        SiteKey `json:"site_key"`
	Label      string  `json:"label"`
	PictureKey string  `json:"picture_key"`
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
