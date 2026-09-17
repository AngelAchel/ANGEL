package arpdhcp

import (
    "time"
)

type arpdhcp0127 struct{}

func Newarpdhcp0127() *arpdhcp0127 {
    return &arpdhcp0127{}
}

func (e *arpdhcp0127) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0127) Name() string { return "arpdhcp0127" }
func (e *arpdhcp0127) Timestamp() time.Time { return time.Now() }
