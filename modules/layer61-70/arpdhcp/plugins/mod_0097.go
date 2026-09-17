package arpdhcp

import (
    "time"
)

type arpdhcp0097 struct{}

func Newarpdhcp0097() *arpdhcp0097 {
    return &arpdhcp0097{}
}

func (e *arpdhcp0097) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0097) Name() string { return "arpdhcp0097" }
func (e *arpdhcp0097) Timestamp() time.Time { return time.Now() }
