package arpdhcp

import (
    "time"
)

type arpdhcp0042 struct{}

func Newarpdhcp0042() *arpdhcp0042 {
    return &arpdhcp0042{}
}

func (e *arpdhcp0042) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0042) Name() string { return "arpdhcp0042" }
func (e *arpdhcp0042) Timestamp() time.Time { return time.Now() }
