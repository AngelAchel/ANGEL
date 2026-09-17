package arpdhcp

import (
    "time"
)

type arpdhcp0047 struct{}

func Newarpdhcp0047() *arpdhcp0047 {
    return &arpdhcp0047{}
}

func (e *arpdhcp0047) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0047) Name() string { return "arpdhcp0047" }
func (e *arpdhcp0047) Timestamp() time.Time { return time.Now() }
