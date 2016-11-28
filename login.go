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

// Login saves the session key using credentials from config
func (d *Domus) Login() error {
	resp, err := d.Get("/Mobile/Login", map[string]string{
		"site_key": string(d.config.SiteKey),
		"user_key": string(d.config.UserKey),
		"password": d.config.Password,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var n int
	body := make([]byte, sessionKeyLen+10)
	if n, err = resp.Body.Read(body); n <= 0 {
		if err == io.EOF {
			return errors.New("Login: invalid credentials (EOF received)")
		}
		return err
	}
	if n != sessionKeyLen {
		return fmt.Errorf("Login: session key should be 40 bytes, is %d: %s", n, body)
	}

	d.setSessionKey(SessionKey(body[:sessionKeyLen]))
	if d.Debug {
		fmt.Printf("sessionKey: %s\n", d.SessionKey)
	}
	return nil
}

// LoginInfos returns a struct with misc infos
func (d *Domus) LoginInfos() (LoginInfos, error) {
	var infos LoginInfos
	resp, err := d.Get("/Mobile/LoginInfos", map[string]string{
		"site_key": string(d.config.SiteKey),
		"user_key": string(d.config.UserKey),
		"password": d.config.Password,
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
