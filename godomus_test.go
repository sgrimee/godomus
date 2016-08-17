package godomus

import (
	"encoding/xml"
	"net/http"
	"testing"

	"github.com/justwatchcom/goat"
)

var (
	domus *Domus
)

func TestNew(t *testing.T) {
	domus = New("myself", "abc123", "localhost", 8080)
	if domus == nil {
		t.Error("Could not create domus object")
	}
}

func TestGetSitesRuntimeSoap(t *testing.T) {
	c := &http.Client{}
	ws := goat.NewWebservice(c, map[string]interface{}{})
	err := ws.AddServices("http://10.0.1.6:8080/DomoBox/Mobile?wsdl")
	if err != nil {
		t.Error(err)
	}
	resp := struct {
		XMLName          xml.Name `xml:"GetSitesResponse"`
		GetSitesResponse struct {
			XMLName xml.Name `xml:"return"`
			Entries []struct {
				XMLName    xml.Name `xml:"site"`
				SiteKey    string   `xml:"site_key"`
				Label      bool     `xml:"label"`
				PictureKey string   `xml:"picture_key"`
			} `xml:"return"`
		}
	}{}
	_ = "breakpoint"
	err = ws.Do("Mobile", "GetSites", &resp, map[string]interface{}{
		"Entries": []string{
			"SiteKey",
			"Label",
			"PictureKey",
		},
	})
	if err != nil {
		t.Error(err)
	}
}

// func TestGetSitesGeneratedSoap(t *testing.T) {
// 	cli := soap.Client{
// 		URL:       "http://10.0.1.6:8080",
// 		Namespace: ldsoap.Namespace,
// 	}
// 	conn := ldsoap.NewService(&cli)
// 	reply, err := conn.GetSites(cli)
// 	if err != nil {
// 		t.Error("Error getting sites: %s", err)
// 	}
// 	t.Logf("Sites: %+v", reply)
// }
