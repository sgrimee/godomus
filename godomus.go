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

type Domus struct {
	Debug bool

	config     Config
	apiUrl     *url.URL
	socketAddr string

	sync.Mutex
	sessionKey SessionKey
}

// New returns a new Domus object
// socketPort is the port of the LD event socket. If 0, event handling is disabled.
func New(config Config) (*Domus, error) {
	api, err := url.Parse(fmt.Sprintf("%s/DomoBox/rs", config.Url))
	if err != nil {
		return nil, err
	}

	d := Domus{
		config: config,
		Debug:  config.Debug,
		apiUrl: api,
	}

	if config.SocketPort != 0 {
		host := strings.Split(api.Host, ":")[0] // take host from api url without port
		d.socketAddr = fmt.Sprintf("%s:%d", host, config.SocketPort)
	}

	return &d, nil
}

// Make the sessionKey thread safe because the Login function could be invoked from
// the ListenForEvents goroutine

func (d *Domus) SessionKey() SessionKey {
	d.Lock()
	defer d.Unlock()
	return d.sessionKey
}

func (d *Domus) setSessionKey(sk SessionKey) {
	d.Lock()
	defer d.Unlock()
	d.sessionKey = sk
}
