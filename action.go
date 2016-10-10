package godomus

import "errors"

func (d *Domus) On(dev Device) error {
	return d.binaryAction(dev, ActionClassIdOn)
}

func (d *Domus) Off(dev Device) error {
	return d.binaryAction(dev, ActionClassIdOff)
}

func (d *Domus) Toggle(dev Device) error {
	return d.binaryAction(dev, ActionClassIdToggle)
}

func (d *Domus) binaryAction(dev Device, action ActionClassId) error {
	// is it a supported action
	switch action {
	case ActionClassIdOn,
		ActionClassIdOff,
		ActionClassIdToggle:
	default:
		return errors.New("Invalid action class id")
	}

	// is it a supported device
	for _, devAction := range dev.Actions.Action {
		property := devAction.PropClsId
		switch property {
		case PropClassIdBinarySwitch,
			PropClassIdDimmerSwitch:
			return d.ExecuteAction(action, property, dev.Key)
		}
	}
	return errors.New("Invalid property class id")
}
