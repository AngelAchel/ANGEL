package arpdhcp

import (
    "time"
)

type arpdhcp0080 struct{}

func Newarpdhcp0080() *arpdhcp0080 {
    return &arpdhcp0080{}
}

func (e *arpdhcp0080) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0080) Name() string { return "arpdhcp0080" }
func (e *arpdhcp0080) Timestamp() time.Time { return time.Now() }
