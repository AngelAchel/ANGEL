package arpdhcp

import (
    "time"
)

type arpdhcp0183 struct{}

func Newarpdhcp0183() *arpdhcp0183 {
    return &arpdhcp0183{}
}

func (e *arpdhcp0183) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0183) Name() string { return "arpdhcp0183" }
func (e *arpdhcp0183) Timestamp() time.Time { return time.Now() }
