package godomus

import (
	"testing"
)

func TestNew(t *testing.T) {
	domus := New("myself", "abc123", "localhost", 8080)
	if domus == nil {
		t.Error("Could not create domus object")
	}
}
