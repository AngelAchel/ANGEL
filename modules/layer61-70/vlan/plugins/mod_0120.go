package vlan

import (
    "time"
)

type vlan0120 struct{}

func Newvlan0120() *vlan0120 {
    return &vlan0120{}
}

func (e *vlan0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "vlan:done")
    return results, nil
}

func (e *vlan0120) Name() string { return "vlan0120" }
func (e *vlan0120) Timestamp() time.Time { return time.Now() }
