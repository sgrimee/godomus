package godomus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type LoginInfos struct {
	SessionKey SessionKey `json:"session_key"`
	Rooms      []Room
	Scenarios  []Scenario
	// Devices    []]Device
	// Bookmarks  interface{}
	// Groups     Group[]
}

// Login saves and returns the session key for a user
// The site, user, password and session keys are saved for later use by other methods
func (d *Domus) Login(sk SiteKey, uk UserKey, password string) (SessionKey, error) {
	resp, err := d.Get("/Mobile/Login", map[string]string{
		"site_key": string(sk),
		"user_key": string(uk),
		"password": password,
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var n int
	body := make([]byte, sessionKeyLen+10)
	if n, err = resp.Body.Read(body); n <= 0 {
		if err == io.EOF {
			return "", errors.New("Login: invalid credentials (EOF received)")
		}
		return "", err
	}
	if n != sessionKeyLen {
		return "", fmt.Errorf("Login: session key should be 40 bytes, is %d: %s", n, body)
	}

	d.setSessionKey(SessionKey(body[:sessionKeyLen]))
	d.setSiteKey(sk)
	d.setUserKey(uk)
	d.setPassword(password)
	if d.Debug {
		fmt.Printf("sessionKey: %s\n", d.SessionKey)
	}
	return d.SessionKey(), nil
}

// LoginInfos returns a struct with a session key and additional infos
// The site, user, password and session keys are saved for later use by other methods
func (d *Domus) LoginInfos(sk SiteKey, uk UserKey, password string) (LoginInfos, error) {
	var infos LoginInfos
	resp, err := d.Get("/Mobile/LoginInfos", map[string]string{
		"site_key": string(sk),
		"user_key": string(uk),
		"password": password,
	})
	if err != nil {
		return infos, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		return infos, errors.New("LoginInfos: invalid credentials")
	}

	if err := json.NewDecoder(resp.Body).Decode(&infos); err != nil {
		return infos, err
	}
	d.setSessionKey(infos.SessionKey)
	d.setSiteKey(sk)
	d.setUserKey(uk)
	d.setPassword(password)
	return infos, nil
}

// Avoid recursion in UnmarshalJSON
type loginInfos LoginInfos

func (infos *LoginInfos) UnmarshalJSON(data []byte) error {
	tmp := struct {
		loginInfos
		Rooms     struct{ Room []Room }
		Scenarios struct{ Scenario []Scenario }
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*infos = LoginInfos(tmp.loginInfos)
	infos.Rooms = tmp.Rooms.Room
	infos.Scenarios = tmp.Scenarios.Scenario
	return nil
}
