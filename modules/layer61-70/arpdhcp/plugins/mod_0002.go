package arpdhcp

import (
    "time"
)

type arpdhcp0002 struct{}

func Newarpdhcp0002() *arpdhcp0002 {
    return &arpdhcp0002{}
}

func (e *arpdhcp0002) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0002) Name() string { return "arpdhcp0002" }
func (e *arpdhcp0002) Timestamp() time.Time { return time.Now() }
