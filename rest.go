package godomus

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Get makes a GET call to resource after encoding params given in queries
// If no error, it returns an unclosed http.Response
func (d *Domus) Get(resource string, queries map[string]string) (*http.Response, error) {
	api, err := d.apiUrl.Parse(d.apiUrl.Path + resource)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	for k, v := range queries {
		params.Add(k, v)
	}
	api.RawQuery = params.Encode()
	if d.Debug {
		fmt.Println(api)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req := &http.Request{
		URL: api,
		Header: http.Header{
			"Accept": {"application/json;charset=UTF-8"},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSites returns the list of sites managed by the server
func (d *Domus) GetSites() (Sites, error) {
	resp, err := d.Get("/Mobile/GetSites", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		List []Site `json:"site"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return Sites(body.List), nil
}

// GetUsers returns the list of users for a given site
func (d *Domus) GetUsers(site SiteKey) (Users, error) {
	resp, err := d.Get("/Mobile/GetUsers", map[string]string{
		"site_key": string(site),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		List []User `json:"user"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return Users(body.List), nil
}

// Login saves and returns the session key for a user
// The site, user and session keys are saved for later use by other methods
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
		return "", err
	}
	if n != sessionKeyLen {
		return "", fmt.Errorf("Session key should be 40 bytes, is %d: %s", n, body)
	}
	d.sessionKey = SessionKey(body[:sessionKeyLen])
	d.siteKey = sk
	d.userKey = uk
	return d.sessionKey, nil
}

// LoginInfos returns a struct with a session key and additional infos
// The site, user and session keys are saved for later use by other methods
func (d *Domus) LoginInfos(sk SiteKey, uk UserKey, password string) (*LoginInfos, error) {
	resp, err := d.Get("/Mobile/LoginInfos", map[string]string{
		"site_key": string(sk),
		"user_key": string(uk),
		"password": password,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var infos LoginInfos

	if err := json.NewDecoder(resp.Body).Decode(&infos); err != nil {
		return nil, err
	}
	d.siteKey = sk
	d.userKey = uk
	d.sessionKey = infos.SessionKey
	return &infos, nil
}

// DevicesInRoom returns the list of devices in the given roomId
// If class is "" then all devices are returned, otherwise only devices of that class
func (d *Domus) DevicesInRoom(rk RoomKey, class CategoryClassId) (Devices, error) {
	queries := map[string]string{
		"session_key": string(d.sessionKey),
		"room_key":    string(rk),
	}
	if class != "" {
		queries["category_clsid"] = string(class)
	}
	resp, err := d.Get("/Mobile/GetDevices", queries)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		Devices []Device `json:"device"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return Devices(body.Devices), nil
}
