package arpdhcp

import (
    "time"
)

type arpdhcp0086 struct{}

func Newarpdhcp0086() *arpdhcp0086 {
    return &arpdhcp0086{}
}

func (e *arpdhcp0086) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0086) Name() string { return "arpdhcp0086" }
func (e *arpdhcp0086) Timestamp() time.Time { return time.Now() }
