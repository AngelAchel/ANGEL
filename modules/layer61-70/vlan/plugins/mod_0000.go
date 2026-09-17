package vlan

import (
    "time"
)

type vlan0000 struct{}

func Newvlan0000() *vlan0000 {
    return &vlan0000{}
}

func (e *vlan0000) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "vlan:done")
    return results, nil
}

func (e *vlan0000) Name() string { return "vlan0000" }
func (e *vlan0000) Timestamp() time.Time { return time.Now() }
