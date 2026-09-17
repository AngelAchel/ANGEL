package vlan

import (
    "time"
)

type vlan0080 struct{}

func Newvlan0080() *vlan0080 {
    return &vlan0080{}
}

func (e *vlan0080) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "vlan:done")
    return results, nil
}

func (e *vlan0080) Name() string { return "vlan0080" }
func (e *vlan0080) Timestamp() time.Time { return time.Now() }
