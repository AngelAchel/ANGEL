package vlan

import (
    "time"
)

type vlan0123 struct{}

func Newvlan0123() *vlan0123 {
    return &vlan0123{}
}

func (e *vlan0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "vlan:done")
    return results, nil
}

func (e *vlan0123) Name() string { return "vlan0123" }
func (e *vlan0123) Timestamp() time.Time { return time.Now() }
