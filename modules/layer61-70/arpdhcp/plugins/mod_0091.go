package arpdhcp

import (
    "time"
)

type arpdhcp0091 struct{}

func Newarpdhcp0091() *arpdhcp0091 {
    return &arpdhcp0091{}
}

func (e *arpdhcp0091) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0091) Name() string { return "arpdhcp0091" }
func (e *arpdhcp0091) Timestamp() time.Time { return time.Now() }
