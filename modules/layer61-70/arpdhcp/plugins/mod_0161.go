package arpdhcp

import (
    "time"
)

type arpdhcp0161 struct{}

func Newarpdhcp0161() *arpdhcp0161 {
    return &arpdhcp0161{}
}

func (e *arpdhcp0161) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0161) Name() string { return "arpdhcp0161" }
func (e *arpdhcp0161) Timestamp() time.Time { return time.Now() }
