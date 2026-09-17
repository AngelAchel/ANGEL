package arpdhcp

import (
    "time"
)

type arpdhcp0036 struct{}

func Newarpdhcp0036() *arpdhcp0036 {
    return &arpdhcp0036{}
}

func (e *arpdhcp0036) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0036) Name() string { return "arpdhcp0036" }
func (e *arpdhcp0036) Timestamp() time.Time { return time.Now() }
