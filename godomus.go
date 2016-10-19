package godomus

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

const sessionKeyLen = 40

type SessionKey string
type PictureKey string

type Credentials struct {
	sync.Mutex
	siteKey    SiteKey
	userKey    UserKey
	password   string
	sessionKey SessionKey
}

type Domus struct {
	Debug      bool
	creds      Credentials
	apiUrl     *url.URL
	socketAddr string
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

// Make the Domus.creds thread safe because the Login function could be invoked from
// the ListenForEvents goroutine and it updates creds.

func (d *Domus) SiteKey() SiteKey {
	d.creds.Lock()
	defer d.creds.Unlock()
	return d.creds.siteKey
}

func (d *Domus) setSiteKey(sk SiteKey) {
	d.creds.Lock()
	defer d.creds.Unlock()
	d.creds.siteKey = sk
}

func (d *Domus) UserKey() UserKey {
	d.creds.Lock()
	defer d.creds.Unlock()
	return d.creds.userKey
}

func (d *Domus) setUserKey(uk UserKey) {
	d.creds.Lock()
	defer d.creds.Unlock()
	d.creds.userKey = uk
}

func (d *Domus) SessionKey() SessionKey {
	d.creds.Lock()
	defer d.creds.Unlock()
	return d.creds.sessionKey
}

func (d *Domus) setSessionKey(sk SessionKey) {
	d.creds.Lock()
	defer d.creds.Unlock()
	d.creds.sessionKey = sk
}

func (d *Domus) Password() string {
	d.creds.Lock()
	defer d.creds.Unlock()
	return d.creds.password
}

func (d *Domus) setPassword(p string) {
	d.creds.Lock()
	defer d.creds.Unlock()
	d.creds.password = p
}
