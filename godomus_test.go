package godomus

import (
	"fmt"
	"testing"
)

var (
	domus *Domus
)

func TestNew(t *testing.T) {
	domus = New("myself", "abc123", "http://10.0.1.6:8080")
	if domus == nil {
		t.Error("Could not create domus object")
	}
}

func TestGetSites(t *testing.T) {
	sites, err := domus.GetSites()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sites)
}
