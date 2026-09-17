package arpdhcp

import (
    "time"
)

type arpdhcp0051 struct{}

func Newarpdhcp0051() *arpdhcp0051 {
    return &arpdhcp0051{}
}

func (e *arpdhcp0051) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0051) Name() string { return "arpdhcp0051" }
func (e *arpdhcp0051) Timestamp() time.Time { return time.Now() }
