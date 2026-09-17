package arpdhcp

import (
    "time"
)

type arpdhcp0093 struct{}

func Newarpdhcp0093() *arpdhcp0093 {
    return &arpdhcp0093{}
}

func (e *arpdhcp0093) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0093) Name() string { return "arpdhcp0093" }
func (e *arpdhcp0093) Timestamp() time.Time { return time.Now() }
