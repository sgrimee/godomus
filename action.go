package godomus

const ActionClassIdOn = ActionClassId("CLSID-ACTION-ON")
const ActionClassIdOff = ActionClassId("CLSID-ACTION-OFF")
const ActionClassIdUp = ActionClassId("CLSID-ACTION-UP")
const ActionClassIdDown = ActionClassId("CLSID-ACTION-DOWN")
const ActionClassIdToggle = ActionClassId("CLSID-ACTION-TOGGLE")

type ActionClassId string

type Action struct {
	PropClsId PropClassId   `json:"prop_clsid"`
	ClsId     ActionClassId `json:"action_clsid"`
}
