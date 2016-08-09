package godomus

type Domus struct {
	User,
	Password,
	Host string
	Port int
}

func New(user, password, host string, port int) *Domus {
	return &Domus{User: user, Password: password, Host: host, Port: port}
}
