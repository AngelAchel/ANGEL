package arpdhcp

import (
    "time"
)

type arpdhcp0172 struct{}

func Newarpdhcp0172() *arpdhcp0172 {
    return &arpdhcp0172{}
}

func (e *arpdhcp0172) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0172) Name() string { return "arpdhcp0172" }
func (e *arpdhcp0172) Timestamp() time.Time { return time.Now() }
