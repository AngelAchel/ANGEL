package arpdhcp

import (
    "time"
)

type arpdhcp0196 struct{}

func Newarpdhcp0196() *arpdhcp0196 {
    return &arpdhcp0196{}
}

func (e *arpdhcp0196) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0196) Name() string { return "arpdhcp0196" }
func (e *arpdhcp0196) Timestamp() time.Time { return time.Now() }
