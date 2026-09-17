package arpdhcp

import (
    "time"
)

type arpdhcp0014 struct{}

func Newarpdhcp0014() *arpdhcp0014 {
    return &arpdhcp0014{}
}

func (e *arpdhcp0014) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0014) Name() string { return "arpdhcp0014" }
func (e *arpdhcp0014) Timestamp() time.Time { return time.Now() }
