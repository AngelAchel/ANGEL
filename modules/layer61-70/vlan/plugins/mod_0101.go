package vlan

import (
    "time"
)

type vlan0101 struct{}

func Newvlan0101() *vlan0101 {
    return &vlan0101{}
}

func (e *vlan0101) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "vlan:done")
    return results, nil
}

func (e *vlan0101) Name() string { return "vlan0101" }
func (e *vlan0101) Timestamp() time.Time { return time.Now() }
