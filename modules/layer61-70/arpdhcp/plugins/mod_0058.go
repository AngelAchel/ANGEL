package arpdhcp

import (
    "time"
)

type arpdhcp0058 struct{}

func Newarpdhcp0058() *arpdhcp0058 {
    return &arpdhcp0058{}
}

func (e *arpdhcp0058) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0058) Name() string { return "arpdhcp0058" }
func (e *arpdhcp0058) Timestamp() time.Time { return time.Now() }
