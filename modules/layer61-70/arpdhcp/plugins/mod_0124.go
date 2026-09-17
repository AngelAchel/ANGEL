package arpdhcp

import (
    "time"
)

type arpdhcp0124 struct{}

func Newarpdhcp0124() *arpdhcp0124 {
    return &arpdhcp0124{}
}

func (e *arpdhcp0124) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0124) Name() string { return "arpdhcp0124" }
func (e *arpdhcp0124) Timestamp() time.Time { return time.Now() }
