package godomus

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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

// GetWithSession makes a Get call that requires a session key
// If ensures the session key is set
// If the session has expired it tries to refresh it
// A login must have been performed before using GetWithSession
func (d *Domus) GetWithSession(resource string, queries map[string]string) (*http.Response, error) {
	if d.sessionKey == "" {
		return nil, errors.New("DevicesInRoom: session key is not set, login first")
	}
	q := map[string]string{
		"session_key": string(d.sessionKey),
	}
	for k, v := range queries {
		q[k] = v
	}
	resp, err := d.Get(resource, q)
	if resp.StatusCode != 200 {
		e := jsonError(resp)
		if e.Error() == "INVALID_SESSION" {
			// make a second attempt after logging in again to refresh sessionKey
			log.Println("Refreshing session")
			key, err := d.Login(d.siteKey, d.userKey, d.password)
			if err != nil {
				return nil, err
			}
			q["session_key"] = string(key)
			resp, err = d.Get(resource, q)
		} else {
			return nil, e
		}
	}
	return resp, err
}

// jsonError tries to get an error code from the Json in the body
func jsonError(resp *http.Response) error {
	var errBody struct {
		Message string
		Code    string
	}
	if err := json.NewDecoder(resp.Body).Decode(&errBody); err != nil {
		return err
	}
	if errBody.Code != "" {
		return errors.New(errBody.Code)
	} else {
		return errors.New(resp.Status)
	}
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
	d.sessionKey = SessionKey(body[:sessionKeyLen])
	if d.Debug {
		fmt.Printf("sessionKey: %s\n", d.sessionKey)
	}
	d.siteKey = sk
	d.userKey = uk
	d.password = password
	return d.sessionKey, nil
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
	d.sessionKey = infos.SessionKey
	d.siteKey = sk
	d.userKey = uk
	d.password = password
	return infos, nil
}

// DevicesInRoom returns the list of devices in the given roomId
// If class is "" then all devices are returned, otherwise only devices of that class
func (d *Domus) DevicesInRoom(rk RoomKey, class CategoryClassId) (Devices, error) {
	queries := map[string]string{
		"room_key": string(rk),
	}
	if class != "" {
		queries["category_clsid"] = string(class)
	}
	resp, err := d.GetWithSession("/Mobile/GetDevices", queries)
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

// CategoriesInRoom returns the list of categories in the given roomId
func (d *Domus) CategoriesInRoom(rk RoomKey) (Categories, error) {
	queries := map[string]string{
		"room_key": string(rk),
	}
	resp, err := d.GetWithSession("/Mobile/GetCategories", queries)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body struct {
		Categories []Category `json:"category"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return Categories(body.Categories), nil
}

// ExecuteAction runs the action on property of device
func (d *Domus) ExecuteAction(action ActionClassId, property PropClassId, dk DeviceKey) error {
	queries := map[string]string{
		"target_key":   string(dk),
		"prop_clsid":   string(property),
		"action_clsid": string(action),
	}
	resp, err := d.GetWithSession("/Mobile/ExecuteAction", queries)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// GetDeviceState returns all infos on one device
func (d *Domus) GetDeviceState(dk DeviceKey) (Device, error) {
	var dev Device
	queries := map[string]string{
		"device_key": string(dk),
	}
	resp, err := d.GetWithSession("/Mobile/GetDeviceState", queries)
	if err != nil {
		return dev, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&dev); err != nil {
		return dev, err
	}
	dev.server = d
	return dev, nil
}
