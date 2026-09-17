package arpdhcp

import (
    "time"
)

type arpdhcp0176 struct{}

func Newarpdhcp0176() *arpdhcp0176 {
    return &arpdhcp0176{}
}

func (e *arpdhcp0176) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0176) Name() string { return "arpdhcp0176" }
func (e *arpdhcp0176) Timestamp() time.Time { return time.Now() }
