package godomus

import (
	"fmt"
	"strconv"
	"strings"
)

type ScenarioKey TargetKey

type Scenario struct {
	Key    ScenarioKey `json:"scenario_key"`
	Label  string      `json:"label"`
	Resume string      `json:"resume"`
}

const ActionClassIdRun = ActionClassId("CLSID-ACTION-RUN")

func (d *Domus) RunScenario(sk ScenarioKey) error {
	return d.ExecuteAction(ActionClassIdRun, PropClassId("0"), TargetKey(sk))
}

// NewScenarioKey returns a Scenario from a scenario number as integer
func NewScenarioKey(num int) ScenarioKey {
	return ScenarioKey(fmt.Sprintf("SCNR_%035d", num))
}

// (ScenarioKey) Num returns the scenario number as integer
func (uk ScenarioKey) Num() int {
	ns := strings.TrimPrefix(string(uk), "SCNR_")
	num, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}
	return num
}
