package vlan

import (
    "time"
)

type vlan0100 struct{}

func Newvlan0100() *vlan0100 {
    return &vlan0100{}
}

func (e *vlan0100) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "vlan:done")
    return results, nil
}

func (e *vlan0100) Name() string { return "vlan0100" }
func (e *vlan0100) Timestamp() time.Time { return time.Now() }
