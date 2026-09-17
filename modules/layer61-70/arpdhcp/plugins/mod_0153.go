package arpdhcp

import (
    "time"
)

type arpdhcp0153 struct{}

func Newarpdhcp0153() *arpdhcp0153 {
    return &arpdhcp0153{}
}

func (e *arpdhcp0153) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0153) Name() string { return "arpdhcp0153" }
func (e *arpdhcp0153) Timestamp() time.Time { return time.Now() }
