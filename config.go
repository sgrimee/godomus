package godomus

// config required to connect to a LD server
type Config struct {
	SiteKey    SiteKey
	UserKey    UserKey
	Password   string
	Url        string
	SocketPort int
	Debug      bool
}
