package godomus

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"net"
	"strings"
)

type EventResult struct {
	Message string
	Error   error
}

// ListenForEvents blocks while listening for events on the LD socket
// If max is > 0, the function returns after max valid events
func (d *Domus) ListenForEvents(max int) error {
	if d.socketAddr == "" {
		return errors.New("Socket listenting is not active.")
	}
	conn, err := net.Dial("tcp", d.socketAddr)
	if err != nil {
		return err
	}
	connbuf := bufio.NewReader(conn)
	counter := 0
	for {
		bytes, err := connbuf.ReadBytes(0)
		if (len(bytes) == 1) && (bytes[0] == 0) {
			continue // ignore regular byte 0 events
		}
		if len(bytes) != (int(bytes[0]) + 2) {
			return fmt.Errorf("Invalid length for frame: %s", bytes)
		}
		if err != nil {
			return err
		}
		// remove leading byte (length) and trailing (0)
		msg := bytes[1 : len(bytes)-1]
		em, err := ParseMsg(msg)
		if err != nil {
			return err
		}
		fem := FilterEvent(em)
		if fem != nil {
			// ch <- fem
		}
		counter++
		if (max > 0) && (counter >= max) {
			return nil
		}
	}
}

type EventMsg struct {
	XMLName   xml.Name
	ClsId     string `xml:"clsid,attr"`
	DeviceKey string `xml:"device_key,attr"`
	Raw       string
	Error     error
}

func ParseMsg(msg []byte) (*EventMsg, error) {
	em := EventMsg{}
	err := xml.Unmarshal(msg, &em)
	if err != nil {
		return nil, err
	}
	em.Raw = string(msg)
	return &em, err
}

// FilterEvent returns nil for events that not device state updates
// Returns the Event itself otherwise
func FilterEvent(e *EventMsg) *EventMsg {
	if strings.Contains(e.ClsId, "-CONSO") {
		return nil
	}
	if e.DeviceKey != "" {
		return e
	}
	fmt.Printf("  Ignoring clsid: %s\n", e.ClsId)
	return nil
}
