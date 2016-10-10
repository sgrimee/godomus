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

// ListenForEvents gets device update events from the LD event socket
func (d *Domus) ListenForEvents(events chan<- EventMsg, errs chan<- error) {
	if d.socketAddr == "" {
		errs <- errors.New("Event listening is not active, please provide a port number.")
		close(events)
		return
	}
	conn, err := net.Dial("tcp", d.socketAddr)
	if err != nil {
		errs <- err
		close(events)
		return
	}
	connbuf := bufio.NewReader(conn)
	for {
		bytes, err := connbuf.ReadBytes(0)
		if (len(bytes) == 1) && (bytes[0] == 0) {
			continue // ignore regular byte 0 events
		}
		if len(bytes) != (int(bytes[0]) + 2) {
			errs <- fmt.Errorf("Invalid length for frame: %s", bytes)
			continue
		}
		if err != nil {
			errs <- err
			continue
		}
		// remove leading byte (length) and trailing (0)
		msg := bytes[1 : len(bytes)-1]
		em, err := ParseMsg(msg)
		if err != nil {
			errs <- err
			continue
		}
		if IsDeviceUpdate(em) {
			events <- em
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

func ParseMsg(msg []byte) (EventMsg, error) {
	em := EventMsg{}
	em.Raw = string(msg)
	err := xml.Unmarshal(msg, &em)
	return em, err
}

// IsDeviceUpdate returns true only if the event is a device state update
func IsDeviceUpdate(e EventMsg) bool {
	if strings.Contains(e.ClsId, "-CONSO") {
		return false
	}
	if e.DeviceKey != "" {
		return true
	}
	//fmt.Printf("  Ignoring clsid: %s\n", e.ClsId)
	return false
}
