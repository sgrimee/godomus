package godomus

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"net"
	"strings"
)

type EventClassId string

type EventResult struct {
	Message string
	Error   error
}

type EventMsg struct {
	XMLName   xml.Name
	ClsId     EventClassId `xml:"clsid,attr"`
	DeviceKey DeviceKey    `xml:"device_key,attr"`
	Raw       string
	Error     error
}

type DeviceUpdate Device

// ListenForEvents gets events from the LD event socket
func (d *Domus) ListenForEvents(events chan<- EventMsg, errs chan<- error, done <-chan struct{}) {
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
		select {
		case <-done:
			close(events)
			close(errs)
			return
		default:
			readOneEvent(connbuf, events, errs)
		}
	}
}

func readOneEvent(connbuf *bufio.Reader, events chan<- EventMsg, errs chan<- error) {
	bytes, err := connbuf.ReadBytes(0)
	if err != nil {
		errs <- err
		return
	}
	if (len(bytes) == 1) && (bytes[0] == 0) {
		return // ignore regular byte 0 events
	}
	if len(bytes) != (int(bytes[0]) + 2) {
		errs <- fmt.Errorf("Invalid length for frame: %s", bytes)
		return
	}
	// remove leading byte (length) and trailing (0)
	msg := bytes[1 : len(bytes)-1]
	em, err := parseMsg(msg)
	if err != nil {
		errs <- err
		return
	}
	events <- em
}

// ListenForDeviceUpdates listens for server events that represent a device update
// and sends the updated device status on the channel
func (d *Domus) ListenForDeviceUpdates(devices chan<- Device, errs chan<- error, done <-chan struct{}) {
	events := make(chan EventMsg)
	go d.ListenForEvents(events, errs, done)
	for {
		select {
		case <-done:
			close(devices)
			// errs will be closed by the ListenForEvents goroutine
			return
		case ev := <-events:
			if isDeviceUpdate(ev) {
				dev, err := d.GetDeviceState(ev.DeviceKey)
				if err != nil {
					errs <- err
					continue
				}
				devices <- dev
			}
		}
	}
}

func parseMsg(msg []byte) (EventMsg, error) {
	em := EventMsg{}
	em.Raw = string(msg)
	err := xml.Unmarshal(msg, &em)
	return em, err
}

// isDeviceUpdate returns true only if the event is a device state update
func isDeviceUpdate(e EventMsg) bool {
	if strings.Contains(string(e.ClsId), "-CONSO") {
		return false
	}
	if !strings.Contains(string(e.ClsId), "-STATE-") {
		return false
	}
	if e.DeviceKey == "" {
		return false
	}
	return true
}
