package godomus

import (
	"fmt"
	"net/url"
	"strings"
)

const sessionKeyLen = 40

type SessionKey string

type Domus struct {
	siteKey    SiteKey
	userKey    UserKey
	password   string
	apiUrl     *url.URL
	socketAddr string
	sessionKey SessionKey
	Debug      bool
}

type LoginInfos struct {
	SessionKey SessionKey                    `json:"session_key"`
	Rooms      struct{ Room []Room }         `json:"rooms"`
	Scenarios  struct{ Scenario []Scenario } `json:"scenarios"`
	//Devices    struct{ Device []Device }     `json:"devices"`
	// Bookmarks  interface{}                   `json:"bookmarks"`
	// Groups     interface{}                   `json:"groups"`
}

// New returns a new Domus object
// socketPort is the port of the LD event socket. If 0, event handling is disabled.
func New(apiUrl string, socketPort int) (*Domus, error) {
	d := new(Domus)
	api, err := url.Parse(fmt.Sprintf("%s/DomoBox/rs", apiUrl))
	if err != nil {
		return nil, err
	}
	d.apiUrl = api
	if socketPort != 0 {
		host := strings.Split(api.Host, ":")[0] // take host from api url without port
		d.socketAddr = fmt.Sprintf("%s:%d", host, socketPort)
	}
	return d, nil
}
