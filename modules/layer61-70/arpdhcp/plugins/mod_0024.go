package arpdhcp

import (
    "time"
)

type arpdhcp0024 struct{}

func Newarpdhcp0024() *arpdhcp0024 {
    return &arpdhcp0024{}
}

func (e *arpdhcp0024) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0024) Name() string { return "arpdhcp0024" }
func (e *arpdhcp0024) Timestamp() time.Time { return time.Now() }
