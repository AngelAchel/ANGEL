package arpdhcp

import (
    "time"
)

type arpdhcp0063 struct{}

func Newarpdhcp0063() *arpdhcp0063 {
    return &arpdhcp0063{}
}

func (e *arpdhcp0063) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0063) Name() string { return "arpdhcp0063" }
func (e *arpdhcp0063) Timestamp() time.Time { return time.Now() }
