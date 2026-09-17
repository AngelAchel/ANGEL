package arpdhcp

import (
    "time"
)

type arpdhcp0136 struct{}

func Newarpdhcp0136() *arpdhcp0136 {
    return &arpdhcp0136{}
}

func (e *arpdhcp0136) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0136) Name() string { return "arpdhcp0136" }
func (e *arpdhcp0136) Timestamp() time.Time { return time.Now() }
