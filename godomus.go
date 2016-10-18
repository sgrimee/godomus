package godomus

import (
	"fmt"
	"net/url"
	"strings"
)

const sessionKeyLen = 40

type SessionKey string
type PictureKey string

type Domus struct {
	Debug      bool
	siteKey    SiteKey
	userKey    UserKey
	password   string
	apiUrl     *url.URL
	socketAddr string
	sessionKey SessionKey
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
