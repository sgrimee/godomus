package godomus

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type TargetKey string

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
			if d.Debug {
				log.Println("Refreshing session")
			}
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

// ExecuteAction runs the action on property of device
func (d *Domus) ExecuteAction(action ActionClassId, property PropClassId, tk TargetKey) error {
	queries := map[string]string{
		"target_key":   string(tk),
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
