package arpdhcp

import (
    "time"
)

type arpdhcp0143 struct{}

func Newarpdhcp0143() *arpdhcp0143 {
    return &arpdhcp0143{}
}

func (e *arpdhcp0143) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0143) Name() string { return "arpdhcp0143" }
func (e *arpdhcp0143) Timestamp() time.Time { return time.Now() }
