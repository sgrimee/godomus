package godomus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (d Domus) GetSites() ([]Site, error) {
	resp, err := http.Get(fmt.Sprintf("%s/Mobile/GetSites", d.Api))
	defer resp.Body.Close()
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
	return body.List, nil
}
