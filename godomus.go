package godomus

import "fmt"

type Domus struct {
	User,
	Password,
	Api,
	HostUrl string
}

func New(user, password, hostUrl string) *Domus {
	d := &Domus{User: user, Password: password, HostUrl: hostUrl}
	d.Api = fmt.Sprintf("%s/DomoBox/rs", hostUrl)
	return d
}
