package arpdhcp

import (
    "time"
)

type arpdhcp0154 struct{}

func Newarpdhcp0154() *arpdhcp0154 {
    return &arpdhcp0154{}
}

func (e *arpdhcp0154) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0154) Name() string { return "arpdhcp0154" }
func (e *arpdhcp0154) Timestamp() time.Time { return time.Now() }
