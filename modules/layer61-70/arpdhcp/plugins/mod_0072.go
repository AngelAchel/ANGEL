package arpdhcp

import (
    "time"
)

type arpdhcp0072 struct{}

func Newarpdhcp0072() *arpdhcp0072 {
    return &arpdhcp0072{}
}

func (e *arpdhcp0072) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0072) Name() string { return "arpdhcp0072" }
func (e *arpdhcp0072) Timestamp() time.Time { return time.Now() }
