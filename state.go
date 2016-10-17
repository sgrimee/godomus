package godomus

import "encoding/json"

type StateClassId string

type State struct {
	ClsId  StateClassId `json:"state_clsid"`
	Type   string       `json:"type"`
	Label  string       `json:"label"`
	Values []Value
}

// avoid recursion in UnmarshalJSON
type state State

func (s *State) UnmarshalJSON(data []byte) error {
	tmp := struct {
		state
		Values struct {
			Value []Value `json:"value"`
		} `json:"values"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = State(tmp.state)
	s.Values = tmp.Values.Value
	return nil
}
